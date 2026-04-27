package bridgeparser

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"bridgelang.dev/bridge-stdlib/spec"
	"tinygo.org/x/go-llvm"
)

func (v *Visitor) declareStdlibFunc(moduleName, funcName string) {
	mod, ok := spec.Lookup(moduleName)
	if !ok {
		return
	}
	for _, fn := range mod.Funcs {
		if fn.Name != funcName {
			continue
		}
		qualName := moduleName + "::" + funcName
		if existing := v.Llvmmod.NamedFunction(qualName); !existing.IsNil() {
			return
		}
		params := make([]llvm.Type, len(fn.ParamTypes))
		for i, pt := range fn.ParamTypes {
			params[i] = v.ResolveType(pt)
		}
		var retType llvm.Type
		if fn.ReturnType == "" {
			retType = v.Llvmctx.VoidType()
		} else {
			retType = v.ResolveType(fn.ReturnType)
		}
		llvm.AddFunction(v.Llvmmod, qualName, llvm.FunctionType(retType, params, fn.Variadic))
		return
	}
}

func (v *Visitor) declareStdlibModule(moduleName string) bool {
	mod, ok := spec.Lookup(moduleName)
	if !ok {
		return false
	}
	v.declareMod(mod)
	return true
}

func (v *Visitor) declarePackageModule(name string) {
	bridgePath := os.Getenv("BRIDGEPATH")
	if bridgePath == "" {
		panic(fmt.Sprintf("bridgec: @Import %q: BRIDGEPATH is not set", name))
	}

	pkgDir := filepath.Join(bridgePath, "packages", name)
	entries, err := os.ReadDir(pkgDir)
	if err != nil {
		panic(fmt.Sprintf("bridgec: @Import %q: cannot read package directory %q: %v", name, pkgDir, err))
	}

	version := ""
	for _, e := range entries {
		if e.IsDir() && e.Name() > version {
			version = e.Name()
		}
	}
	if version == "" {
		panic(fmt.Sprintf("bridgec: @Import %q: no versions found in %q", name, pkgDir))
	}

	specPath := filepath.Join(pkgDir, version, "spec.json")
	specBytes, err := os.ReadFile(specPath)
	if err != nil {
		panic(fmt.Sprintf("bridgec: @Import %q: cannot read %q: %v", name, specPath, err))
	}

	var mod spec.Module
	if err := json.Unmarshal(specBytes, &mod); err != nil {
		panic(fmt.Sprintf("bridgec: @Import %q: invalid spec.json: %v", name, err))
	}
	mod.Name = name

	v.declareMod(mod)
	v.LibDeps[name] = version
}

func (v *Visitor) declareMod(mod spec.Module) {
	for _, obj := range mod.Objects {
		if _, exists := v.structFields[obj.Name]; exists {
			continue
		}
		fieldNames := make([]string, len(obj.Fields))
		fieldTypeStrs := make([]string, len(obj.Fields))
		fieldTypes := make([]llvm.Type, len(obj.Fields))
		for i, f := range obj.Fields {
			fieldNames[i] = f.Name
			fieldTypeStrs[i] = f.Type
			fieldTypes[i] = v.ResolveType(f.Type)
		}
		structType := v.Llvmctx.StructCreateNamed(obj.Name)
		structType.StructSetBody(fieldTypes, false)
		v.Results[obj.Name] = structType
		v.structFields[obj.Name] = fieldNames
		v.structFieldTypes[obj.Name] = fieldTypeStrs

		v.pkgStructs[obj.Name] = mod.Name

		for _, f := range obj.Fields {
			if !strings.HasPrefix(f.Type, "func(") {
				continue
			}
			fnType, ok := v.parseStaticMethodType(f.Type, structType)
			if !ok {
				continue
			}
			// Declare the package's method as a direct import: PackageName::MethodName
			qualName := mod.Name + "::" + f.Name
			if existing := v.Llvmmod.NamedFunction(qualName); existing.IsNil() {
				llvm.AddFunction(v.Llvmmod, qualName, fnType)
			}
			v.methodTypes[obj.Name+"."+f.Name] = fnType
			retStr := methodReturnTypeStr(f.Type)
			if strings.HasSuffix(retStr, "?") {
				base := strings.TrimSuffix(retStr, "?")
				if _, ok := v.Results[base]; ok {
					v.funcReturnStructName[qualName] = base
				}
			}
		}
	}
	for _, fn := range mod.Funcs {
		qualName := mod.Name + "::" + fn.Name
		if existing := v.Llvmmod.NamedFunction(qualName); !existing.IsNil() {
			continue
		}
		params := make([]llvm.Type, len(fn.ParamTypes))
		for i, pt := range fn.ParamTypes {
			params[i] = v.ResolveType(pt)
		}
		var retType llvm.Type
		if fn.ReturnType == "" {
			retType = v.Llvmctx.VoidType()
		} else {
			retType = v.ResolveType(fn.ReturnType)
		}
		if fn.Variadic {
			params = append(params, llvm.PointerType(v.Llvmctx.Int8Type(), 0))
			llvm.AddFunction(v.Llvmmod, qualName, llvm.FunctionType(retType, params, false))
			v.stdlibVariadicFuncs[qualName] = true
		} else {
			llvm.AddFunction(v.Llvmmod, qualName, llvm.FunctionType(retType, params, false))
		}
	}
}

func (v *Visitor) parseStaticMethodType(t string, structType llvm.Type) (llvm.Type, bool) {
	t = strings.TrimPrefix(t, "func")
	if !strings.HasPrefix(t, "(") {
		return llvm.Type{}, false
	}
	close := strings.Index(t, ")")
	if close < 0 {
		return llvm.Type{}, false
	}
	paramStr := t[1:close]
	retStr := strings.TrimSpace(t[close+1:])

	paramTypes := []llvm.Type{llvm.PointerType(v.Llvmctx.Int8Type(), 0)}
	if paramStr != "" {
		for _, p := range strings.Split(paramStr, ",") {
			p = strings.TrimSpace(p)
			if p == "self" {
				continue // already prepended
			}
			paramTypes = append(paramTypes, v.ResolveType(p))
		}
	}

	var retType llvm.Type
	if retStr == "" {
		retType = v.Llvmctx.VoidType()
	} else {
		retType = v.ResolveType(retStr)
	}
	return llvm.FunctionType(retType, paramTypes, false), true
}

func (v *Visitor) parseClosureFuncType(t string, structType llvm.Type) (llvm.Type, bool) {
	t = strings.TrimPrefix(t, "func")
	if !strings.HasPrefix(t, "(") {
		return llvm.Type{}, false
	}
	close := strings.Index(t, ")")
	if close < 0 {
		return llvm.Type{}, false
	}
	paramStr := t[1:close]
	retStr := strings.TrimSpace(t[close+1:])

	i8ptr := llvm.PointerType(v.Llvmctx.Int8Type(), 0)
	// Closure calling convention: first param is env i8*
	paramTypes := []llvm.Type{i8ptr}
	if paramStr != "" {
		for _, p := range strings.Split(paramStr, ",") {
			p = strings.TrimSpace(p)
			if p == "self" {
				paramTypes = append(paramTypes, structType)
			} else {
				paramTypes = append(paramTypes, v.ResolveType(p))
			}
		}
	}

	var retType llvm.Type
	if retStr == "" {
		retType = v.Llvmctx.VoidType()
	} else {
		retType = v.ResolveType(retStr)
	}
	return llvm.FunctionType(retType, paramTypes, false), true
}

func (v *Visitor) parseFuncType(t string) (llvm.Type, bool) {
	t = strings.TrimPrefix(t, "func")
	if !strings.HasPrefix(t, "(") {
		return llvm.Type{}, false
	}
	close := strings.Index(t, ")")
	if close < 0 {
		return llvm.Type{}, false
	}
	paramStr := t[1:close]
	retStr := strings.TrimSpace(t[close+1:])

	var paramTypes []llvm.Type
	if paramStr != "" {
		for _, p := range strings.Split(paramStr, ",") {
			p = strings.TrimSpace(p)
			paramTypes = append(paramTypes, v.ResolveType(p))
		}
	}

	var retType llvm.Type
	if retStr == "" {
		retType = v.Llvmctx.VoidType()
	} else {
		retType = v.ResolveType(retStr)
	}
	return llvm.FunctionType(retType, paramTypes, false), true
}

func methodReturnTypeStr(t string) string {
	t = strings.TrimPrefix(t, "func")
	close := strings.Index(t, ")")
	if close < 0 {
		return ""
	}
	return strings.TrimSpace(t[close+1:])
}
