package bridgeparser

import (
	"crypto/sha256"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	parser "bridgelang.dev/bridgec-libs/antlr"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/google/uuid"
	"tinygo.org/x/go-llvm"
)

type Visitor struct {
	*parser.BaseBridgeVisitor
	Results       map[string]any
	PrintAst      bool
	depth         int
	lastFunc      string
	lastRetType   llvm.Type
	returnEmitted bool

	LibDeps map[string]string

	vars                 map[string]llvm.Value
	globalVars           map[string]llvm.Value // top-level globals, survive function scope resets
	varTypes             map[string]string     // identifier: typeName
	symlinkPointeeTypes  map[string]llvm.Type  // pointerName: pointeeType
	selfType             string
	structFields         map[string][]string
	structFieldTypes     map[string][]string
	methodTypes          map[string]llvm.Type
	methodIsClosure      map[string]bool
	pkgStructs           map[string]string // structName → packageName
	funcReturnStructName map[string]string // qualFuncName → struct name for String? returns
	varargFuncs          map[string]bool
	stdlibVariadicFuncs  map[string]bool
	breakTarget          llvm.BasicBlock
	continueTarget       llvm.BasicBlock
	escapeFlag           llvm.Value

	closureCaptures        map[string]llvm.Value
	closureVarFuncTypes    map[string]llvm.Type
	pendingClosureFuncType llvm.Type
	pendingClosureInnerFn  llvm.Value
	pendingJournaled       bool

	Llvmctx llvm.Context
	Llvmmod llvm.Module
	Builder llvm.Builder
}

func GetIndent(n int) string {
	indent := ""
	for i := 0; i < n; i++ {
		indent += " "
	}
	return indent
}

func (v *Visitor) MakeVisitor() {
	v.Results = make(map[string]any)
	v.vars = make(map[string]llvm.Value)
	v.globalVars = make(map[string]llvm.Value)
	v.LibDeps = make(map[string]string)
	v.varTypes = make(map[string]string)
	v.symlinkPointeeTypes = make(map[string]llvm.Type)
	v.structFields = make(map[string][]string)
	v.structFieldTypes = make(map[string][]string)
	v.methodTypes = make(map[string]llvm.Type)
	v.methodIsClosure = make(map[string]bool)
	v.pkgStructs = make(map[string]string)
	v.funcReturnStructName = make(map[string]string)
	v.varargFuncs = make(map[string]bool)
	v.stdlibVariadicFuncs = make(map[string]bool)
	v.closureCaptures = make(map[string]llvm.Value)
	v.closureVarFuncTypes = make(map[string]llvm.Type)
}

func (v *Visitor) ResolveType(t string) llvm.Type {
	if idx := strings.Index(t, "["); idx != -1 {
		inner := v.ResolveType(t[:idx])
		closeIdx := strings.Index(t[idx:], "]")
		lenStr := ""
		if closeIdx > 1 {
			lenStr = t[idx+1 : idx+closeIdx]
		}
		length := 0
		if lenStr != "" {
			fmt.Sscanf(lenStr, "%d", &length)
		}
		if length == 0 {
			return llvm.PointerType(v.Llvmctx.Int64Type(), 0)
		}
		return llvm.ArrayType(inner, length)
	}
	if strings.HasSuffix(t, "?") {
		return llvm.PointerType(v.ResolveType(strings.TrimSuffix(t, "?")), 0)
	}
	if strings.HasPrefix(t, "func") { // Closures
		return llvm.PointerType(v.Llvmctx.Int8Type(), 0)
	}
	switch {
	case t == "bool":
		return v.Llvmctx.Int1Type()
	case t == "byte":
		return v.Llvmctx.Int8Type()
	case t == "int8" || t == "uint8":
		return v.Llvmctx.Int8Type()
	case t == "int16" || t == "uint16":
		return v.Llvmctx.Int16Type()
	case t == "int32" || t == "uint32":
		return v.Llvmctx.Int32Type()
	case t == "int" || t == "uint" || t == "int64" || t == "uint64":
		return v.Llvmctx.Int64Type()
	case t == "f32" || t == "float":
		return v.Llvmctx.FloatType()
	case t == "f64":
		return v.Llvmctx.DoubleType()
	case t == "string" || t == "bstr":
		return llvm.PointerType(v.Llvmctx.Int8Type(), 0)
	case t == "obj" || t == "any" || t == "symlink":
		return llvm.PointerType(v.Llvmctx.Int8Type(), 0)
	case t == "self":
		if v.selfType != "" {
			if st, ok := v.Results[v.selfType].(llvm.Type); ok {
				return st
			}
		}
		return llvm.PointerType(v.Llvmctx.Int8Type(), 0)
	default:
		if st, ok := v.Results[t]; ok {
			if llvmT, ok := st.(llvm.Type); ok {
				return llvmT
			}
		}
		return llvm.PointerType(v.Llvmctx.Int8Type(), 0)
	}
}

func (v *Visitor) PrintNode(s string) {
	if !v.PrintAst {
		return
	}
	if len(s) > 60 {
		s = s[:60]
		s += "..."
	}
	fmt.Printf("%s%s\n", GetIndent(v.depth), s)
}

func (v *Visitor) blockHasTerminator() bool {
	bb := v.Builder.GetInsertBlock()
	if bb.IsNil() {
		return true
	}
	last := bb.LastInstruction()
	if last.IsNil() {
		return false
	}
	op := last.InstructionOpcode()
	return op == llvm.Ret || op == llvm.Br || op == llvm.Switch || op == llvm.Unreachable
}

func (v *Visitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *Visitor) VisitStart(ctx *parser.StartContext) any {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitExternalDecl(ctx *parser.ExternalDeclContext) any {
	if ctx.FuncDecl() != nil {
		ret := v.Visit(ctx.FuncDecl())
		fun := ret.(llvm.Value)
		fun.SetFunctionCallConv(llvm.CallConv(llvm.CCallConv))
		return ret
	}
	if ctx.Annotation() != nil {
		return v.Visit(ctx.Annotation())
	}
	if ctx.Decl() != nil {
		return v.Visit(ctx.Decl())
	}
	if ctx.ClassDecl() != nil {
		return v.Visit(ctx.ClassDecl())
	}
	return nil
}

func (v *Visitor) VisitFuncDecl(ctx *parser.FuncDeclContext) any {
	v.PrintNode(fmt.Sprintf("Function Declaration: %s", ctx.GetText()))
	var functype llvm.Type
	var argTypes []llvm.Type
	var localid string
	if ctx.Arguments() != nil {
		argsCtx := ctx.Arguments().(*parser.ArgumentsContext)
		if argsCtx.ELLIPSIS() != nil {
			argTypes = append(argTypes, llvm.PointerType(llvm.PointerType(v.Llvmctx.Int8Type(), 0), 0))
		} else {
			for _, a := range ctx.Arguments().GetChildren() {
				arg, ok := a.(*parser.ArgumentContext)
				if !ok {
					continue
				}
				argTypes = append(argTypes, v.ResolveType(arg.TypeKey().GetText()))
			}
		}
	}
	if ctx.TypeKey() != nil {
		functype = llvm.FunctionType(v.ResolveType(ctx.TypeKey().GetText()), argTypes, false)
	} else {
		functype = llvm.FunctionType(v.Llvmctx.VoidType(), argTypes, false)
	}
	argsText := ""
	if ctx.Arguments() != nil {
		argsText = ctx.Arguments().GetText()
	}
	if ctx.Identifier() != nil {
		name := ctx.Identifier().GetText()
		signature := name + " " + argsText
		h := sha256.Sum256([]byte(signature))
		localid = fmt.Sprintf("%x", h[:])
		fn := llvm.AddFunction(v.Llvmmod, name, functype)
		if name == "Main" || v.pendingJournaled {
			attr := v.Llvmctx.CreateStringAttribute("wasm-export-name", name)
			fn.AddFunctionAttr(attr)
		}
		v.pendingJournaled = false
		v.Results[localid] = fn
		if ctx.Arguments() != nil && ctx.Arguments().(*parser.ArgumentsContext).ELLIPSIS() != nil {
			v.varargFuncs[name] = true
		}
	} else {
		if len(v.vars) > 0 && ctx.Body() != nil {
			bodyIds := collectBodyIdentifiers(ctx.Body())
			var captureNames []string
			for name := range bodyIds {
				if _, ok := v.vars[name]; ok {
					captureNames = append(captureNames, name)
				}
			}
			sort.Strings(captureNames)
			if len(captureNames) > 0 {
				return v.compileClosure(ctx, captureNames, functype)
			}
		}
		uid := strings.ReplaceAll(uuid.New().String(), "-", "")
		signature := uid + " " + argsText
		h := sha256.Sum256([]byte(signature))
		localid = fmt.Sprintf("%x", h[:])
		v.Results[localid] = llvm.AddFunction(v.Llvmmod, uid, functype)
	}
	savedLastFunc := v.lastFunc
	savedLastRetType := v.lastRetType
	savedReturnEmitted := v.returnEmitted
	savedVars := v.vars
	savedVarTypes := v.varTypes
	savedSymlinkPointeeTypes := v.symlinkPointeeTypes
	savedSelfType := v.selfType
	savedBreakTarget := v.breakTarget
	savedContinueTarget := v.continueTarget

	v.vars = make(map[string]llvm.Value)
	v.varTypes = make(map[string]string)
	v.symlinkPointeeTypes = make(map[string]llvm.Type)
	v.lastRetType = functype.ReturnType()
	v.lastFunc = localid
	fn := v.Results[localid]

	if ctx.Arguments() != nil {
		argsCtx := ctx.Arguments().(*parser.ArgumentsContext)
		if argsCtx.ELLIPSIS() != nil {
			v.varTypes[argsCtx.Identifier().GetText()] = "symlink[]"
		} else {
			for _, a := range ctx.Arguments().GetChildren() {
				arg, ok := a.(*parser.ArgumentContext)
				if !ok {
					continue
				}
				v.varTypes[arg.Identifier().GetText()] = arg.TypeKey().GetText()
			}
		}
	}

	if ctx.Arguments() != nil {
		argsCtx := ctx.Arguments().(*parser.ArgumentsContext)
		llvmFn := fn.(llvm.Value)
		if argsCtx.ELLIPSIS() != nil {
			name := argsCtx.Identifier().GetText()
			llvmFn.Param(0).SetName(name)
		} else {
			i := 0
			for _, a := range ctx.Arguments().GetChildren() {
				arg, ok := a.(*parser.ArgumentContext)
				if !ok {
					continue
				}
				name := arg.Identifier().GetText()
				param := llvmFn.Param(i)
				param.SetName(name)
				i++
			}
		}
	}

	if ctx.Body() != nil {
		v.Visit(ctx.Body())
	}
	v.lastFunc = savedLastFunc
	v.lastRetType = savedLastRetType
	v.returnEmitted = savedReturnEmitted
	v.vars = savedVars
	v.varTypes = savedVarTypes
	v.symlinkPointeeTypes = savedSymlinkPointeeTypes
	v.selfType = savedSelfType
	v.breakTarget = savedBreakTarget
	v.continueTarget = savedContinueTarget
	return fn
}

func (v *Visitor) VisitClassDecl(ctx *parser.ClassDeclContext) any {
	name := ctx.Identifier().GetText()
	v.PrintNode(fmt.Sprintf("Object Declaration: %s", name))
	var fieldTypes []llvm.Type
	var fieldNames []string
	var fieldTypeStrs []string
	for _, arg := range ctx.ClassBody().(*parser.ClassBodyContext).AllArgument() {
		argCtx := arg.(*parser.ArgumentContext)
		fieldNames = append(fieldNames, argCtx.Identifier().GetText())
		typeStr := argCtx.TypeKey().GetText()
		fieldTypeStrs = append(fieldTypeStrs, typeStr)
		fieldTypes = append(fieldTypes, v.ResolveType(typeStr))
	}
	structType := v.Llvmctx.StructCreateNamed(name)
	structType.StructSetBody(fieldTypes, false)
	v.Results[name] = structType
	v.structFields[name] = fieldNames
	v.structFieldTypes[name] = fieldTypeStrs
	return structType
}

func (v *Visitor) VisitClassBody(ctx *parser.ClassBodyContext) any {
	v.PrintNode(fmt.Sprintf("Object Body: %s", ctx.GetText()))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitReturnStatement(ctx *parser.ReturnStatementContext) any {
	v.PrintNode(fmt.Sprintf("Return: %s", ctx.GetText()))
	v.returnEmitted = true
	if v.lastRetType == v.Llvmctx.VoidType() {
		return v.Builder.CreateRetVoid()
	}
	if ctx.Value() != nil {
		val := v.Visit(ctx.Value())
		if lval, ok := val.(llvm.Value); ok {
			lval = v.coerce(lval, v.lastRetType)
			return v.Builder.CreateRet(lval)
		}
	}
	if ctx.Statement() != nil {
		val := v.Visit(ctx.Statement())
		if lval, ok := val.(llvm.Value); ok {
			lval = v.coerce(lval, v.lastRetType)
			return v.Builder.CreateRet(lval)
		}
	}
	return v.Builder.CreateRet(llvm.Undef(v.lastRetType))
}

func (v *Visitor) VisitAnnotation(ctx *parser.AnnotationContext) any {
	v.PrintNode(fmt.Sprintf("Annotation: %s", ctx.GetText()))
	found := false

	if ctx.Identifier().GetText() == "Journaled" {
		v.pendingJournaled = true
		found = true
	}

	if ctx.Identifier().GetText() == "Import" && ctx.Identifiers() != nil {
		found = true
		if ids, ok := ctx.Identifiers().(*parser.IdentifiersContext); ok {
			for _, id := range ids.AllIdentifier() {
				name := id.GetText()
				if !v.declareStdlibModule(name) {
					v.declarePackageModule(name)
				}
			}
		}
	}

	if !found {
		// TODO: search for annotation extensions in $BRIDGEPATH/annotations
		// if external
		//	...load etc. etc.
		// else
		fmt.Fprintf(os.Stderr, "Warning: unknown annotation %s\n", ctx.Identifier().GetText())
	}

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitIdentifiers(ctx *parser.IdentifiersContext) any {
	identifiers := ctx.AllIdentifier()
	for _, id := range identifiers {
		v.PrintNode(fmt.Sprintf("ID: %s", id.GetText()))
	}
	return identifiers
}

func (v *Visitor) VisitArguments(ctx *parser.ArgumentsContext) any {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitArgument(ctx *parser.ArgumentContext) any {
	v.PrintNode(fmt.Sprintf("Argument: %s (%s)", ctx.Identifier().GetText(), ctx.TypeKey().GetText()))
	return v.VisitChildren(ctx)
}

func (v *Visitor) emitStatements(stmts []parser.IStatementContext) {
	for _, stmt := range stmts {
		if v.blockHasTerminator() {
			break
		}
		v.Visit(stmt)
		if v.returnEmitted {
			break
		}
	}
}

func (v *Visitor) VisitBody(ctx *parser.BodyContext) any {
	savedLastFunc := v.lastFunc
	savedLastRetType := v.lastRetType
	savedReturnEmitted := v.returnEmitted
	savedInsertBlock := v.Builder.GetInsertBlock()

	thisFunc := v.Results[v.lastFunc].(llvm.Value)
	entry := llvm.AddBasicBlock(thisFunc, "entry")
	v.Builder.SetInsertPointAtEnd(entry)
	v.returnEmitted = false
	savedEscapeFlag := v.escapeFlag
	i1 := v.Llvmctx.Int1Type()
	v.escapeFlag = v.Builder.CreateAlloca(i1, "escape.flag")
	v.Builder.CreateStore(llvm.ConstInt(i1, 0, false), v.escapeFlag)
	for i := 0; i < thisFunc.ParamsCount(); i++ {
		param := thisFunc.Param(i)
		name := param.Name()
		if name != "" {
			alloca := v.Builder.CreateAlloca(param.Type(), name)
			v.Builder.CreateStore(param, alloca)
			v.vars[name] = alloca
		}
	}
	v.emitStatements(ctx.AllStatement())
	if !v.returnEmitted && !v.blockHasTerminator() {
		if v.lastRetType == v.Llvmctx.VoidType() {
			v.Builder.CreateRetVoid()
		} else {
			v.Builder.CreateRet(llvm.Undef(v.lastRetType))
		}
	}

	v.lastFunc = savedLastFunc
	v.lastRetType = savedLastRetType
	v.returnEmitted = savedReturnEmitted
	v.escapeFlag = savedEscapeFlag
	if savedInsertBlock.IsNil() {
		v.Builder.ClearInsertionPoint()
	} else {
		v.Builder.SetInsertPointAtEnd(savedInsertBlock)
	}
	return entry
}

func (v *Visitor) VisitStatement(ctx *parser.StatementContext) any {
	switch {
	case ctx.ReturnStatement() != nil:
		return v.Visit(ctx.ReturnStatement())
	case ctx.FuncCall() != nil:
		return v.Visit(ctx.FuncCall())
	case ctx.Assignment() != nil:
		return v.Visit(ctx.Assignment())
	case ctx.ExternalDecl() != nil:
		return v.Visit(ctx.ExternalDecl())
	case ctx.LoopStatement() != nil:
		return v.Visit(ctx.LoopStatement())
	case ctx.IfStatement() != nil:
		return v.Visit(ctx.IfStatement())
	case ctx.OperationWithReturn() != nil:
		return v.Visit(ctx.OperationWithReturn())
	case ctx.MemberAccess() != nil:
		return v.Visit(ctx.MemberAccess())
	case ctx.BreakStatement() != nil:
		return v.Visit(ctx.BreakStatement())
	case ctx.SwitchStatement() != nil:
		return v.Visit(ctx.SwitchStatement())
	case ctx.DeleteOp() != nil:
		return v.Visit(ctx.DeleteOp())
	}
	return nil
}

func (v *Visitor) VisitMemberAccess(ctx *parser.MemberAccessContext) any {
	return v.evalMemberAccess(ctx)
}

func (v *Visitor) VisitNamespaceAccess(ctx *parser.NamespaceAccessContext) any {
	var member string
	if fc := ctx.FuncCall(); fc != nil {
		if conc, ok := fc.(*parser.FuncCallContext); ok {
			member = conc.Identifier().GetText()
		}
	} else {
		member = ctx.Identifier(1).GetText()
	}
	v.PrintNode(fmt.Sprintf("Access to namespace %s: %s", ctx.Identifier(0), member))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFuncCall(ctx *parser.FuncCallContext) any {
	var funcName string
	if ctx.NamespaceAccess() == nil {
		funcName = ctx.Identifier().GetText()
	} else {
		ns := ctx.NamespaceAccess().(*parser.NamespaceAccessContext)
		namespace := ns.Identifier(0).GetText()
		var fnName string
		if fc := ns.FuncCall(); fc != nil {
			// ANTLR matched the rhs as a nested funcCall; extract just the name
			fnName = fc.(*parser.FuncCallContext).Identifier().GetText()
		} else {
			fnName = ns.Identifier(1).GetText()
		}
		funcName = namespace + "::" + fnName
	}
	callArgsText := ""
	if ctx.CallArgs() != nil {
		callArgsText = ctx.CallArgs().GetText()
	}
	v.PrintNode(fmt.Sprintf("Function Call: %s(%s)", funcName, callArgsText))

	fn := v.Llvmmod.NamedFunction(funcName)
	if fn.IsNil() {
		// Check if it's a closure stored in a local variable
		if alloca, ok := v.vars[funcName]; ok {
			return v.emitClosureCall(funcName, alloca, ctx)
		}
		fmt.Fprintf(os.Stderr, "Error: unknown function %q\n", funcName)
		os.Exit(-1)
		return nil
	}

	fnType := fn.GlobalValueType()
	var args []llvm.Value
	if ctx.CallArgs() != nil {
		if v.stdlibVariadicFuncs[funcName] {
			paramTypes := fnType.ParamTypes()
			fixedCount := len(paramTypes) - 1
			callVals := ctx.CallArgs().(*parser.CallArgsContext).AllValue()
			for i := 0; i < fixedCount && i < len(callVals); i++ {
				raw := v.Visit(callVals[i])
				lv, ok := raw.(llvm.Value)
				if !ok {
					continue
				}
				args = append(args, v.coerce(lv, paramTypes[i]))
			}
			vaVals := callVals[fixedCount:]
			i64t := v.Llvmctx.Int64Type()
			bufType := llvm.ArrayType(i64t, len(vaVals))
			bufAlloca := v.Builder.CreateAlloca(bufType, "")
			for i, valCtx := range vaVals {
				raw := v.Visit(valCtx)
				lv, ok := raw.(llvm.Value)
				if !ok {
					continue
				}
				var slot llvm.Value
				switch lv.Type().TypeKind() {
				case llvm.FloatTypeKind:
					slot = v.Builder.CreateBitCast(v.Builder.CreateFPExt(lv, v.Llvmctx.DoubleType(), ""), i64t, "")
				case llvm.DoubleTypeKind:
					slot = v.Builder.CreateBitCast(lv, i64t, "")
				case llvm.PointerTypeKind:
					slot = v.Builder.CreatePtrToInt(lv, i64t, "")
				default:
					slot = v.Builder.CreateIntCast(lv, i64t, "")
				}
				gep := v.Builder.CreateInBoundsGEP(bufType, bufAlloca, []llvm.Value{
					llvm.ConstInt(i64t, 0, false),
					llvm.ConstInt(i64t, uint64(i), false),
				}, "")
				v.Builder.CreateStore(slot, gep)
			}
			bufPtr := v.Builder.CreateInBoundsGEP(bufType, bufAlloca, []llvm.Value{
				llvm.ConstInt(i64t, 0, false),
				llvm.ConstInt(i64t, 0, false),
			}, "")
			args = append(args, bufPtr)
		} else if v.varargFuncs[funcName] {
			callVals := ctx.CallArgs().(*parser.CallArgsContext).AllValue()
			i8ptr := llvm.PointerType(v.Llvmctx.Int8Type(), 0)
			arrType := llvm.ArrayType(i8ptr, len(callVals))
			arrAlloca := v.Builder.CreateAlloca(arrType, "varargs")
			for i, valCtx := range callVals {
				raw := v.Visit(valCtx)
				lv, ok := raw.(llvm.Value)
				if !ok {
					continue
				}
				lv = v.coerce(lv, i8ptr)
				gep := v.Builder.CreateInBoundsGEP(arrType, arrAlloca,
					[]llvm.Value{
						llvm.ConstInt(v.Llvmctx.Int64Type(), 0, false),
						llvm.ConstInt(v.Llvmctx.Int64Type(), uint64(i), false),
					}, "")
				v.Builder.CreateStore(lv, gep)
			}
			arrayPtr := v.Builder.CreateInBoundsGEP(arrType, arrAlloca,
				[]llvm.Value{
					llvm.ConstInt(v.Llvmctx.Int64Type(), 0, false),
					llvm.ConstInt(v.Llvmctx.Int64Type(), 0, false),
				}, "")
			args = append(args, arrayPtr)
		} else {
			paramTypes := fnType.ParamTypes()
			isVararg := fnType.IsFunctionVarArg()
			for i, valCtx := range ctx.CallArgs().(*parser.CallArgsContext).AllValue() {
				raw := v.Visit(valCtx)
				if lv, ok := raw.(llvm.Value); ok {
					if i < len(paramTypes) {
						lv = v.coerce(lv, paramTypes[i])
					} else if isVararg {
						if lv.Type().TypeKind() == llvm.FloatTypeKind {
							lv = v.Builder.CreateFPExt(lv, v.Llvmctx.DoubleType(), "")
						}
					}
					args = append(args, lv)
				}
			}
		}
	}
	return v.Builder.CreateCall(fnType, fn, args, "")
}

func (v *Visitor) VisitCallArgs(ctx *parser.CallArgsContext) any {
	for _, val := range ctx.AllValue() {
		v.PrintNode(fmt.Sprintf("Call argument: %s", val.GetText()))
	}
	return ctx.AllValue()
}

func (v *Visitor) VisitDecl(ctx *parser.DeclContext) any {
	if ctx.Declarator().GetStart().GetTokenType() == parser.BridgeLexerLET {
		if ctx.AssignmentOp() == nil {
			fmt.Fprintf(os.Stderr, "invalid operation:%v \n\nUse of let without known value", ctx.GetPayload())
			os.Exit(-1)
		}
		if ctx.TypeKey() != nil {
			fmt.Fprintf(os.Stderr, "invalid operation:%s \n\nUse of let with constant type", ctx.GetText())
			os.Exit(-1)
		}
	}

	typekey := "auto"
	if ctx.TypeKey() != nil {
		typekey = ctx.TypeKey().GetText()
	}
	v.PrintNode(fmt.Sprintf("Declaration: %s %s type:%s", ctx.Declarator().GetText(), ctx.Identifiers().GetText(), typekey))

	// Evaluate initializer if present
	var initVal llvm.Value
	hasInit := false
	if ctx.AssignmentOp() != nil {
		raw := v.Visit(ctx.AssignmentOp().(*parser.AssignmentOpContext).Value())
		if lv, ok := raw.(llvm.Value); ok {
			initVal = lv
			hasInit = true
		}
	}

	pendingFuncType := v.pendingClosureFuncType
	v.pendingClosureFuncType = llvm.Type{}

	var llvmType llvm.Type
	if typekey != "auto" {
		llvmType = v.ResolveType(typekey)
	} else if hasInit {
		llvmType = initVal.Type()
	} else {
		llvmType = v.Llvmctx.Int64Type()
	}

	// For symlink declarations, resolve the RHS identifier to its alloca so we
	// store the address (not the loaded value) and record the pointee type
	var symlinkPointee llvm.Type
	if typekey == "symlink" && hasInit {
		valCtx := ctx.AssignmentOp().(*parser.AssignmentOpContext).Value().(*parser.ValueContext)
		if valCtx.Identifier() != nil {
			ident := valCtx.Identifier().GetText()
			if targetAlloca, ok := v.vars[ident]; ok {
				initVal = targetAlloca
				symlinkPointee = targetAlloca.AllocatedType()
			}
		} else if !initVal.IsAAllocaInst().IsNil() {
			symlinkPointee = initVal.AllocatedType()
		}
	}

	isTopLevel := v.Builder.GetInsertBlock().IsNil()

	for _, id := range ctx.Identifiers().(*parser.IdentifiersContext).AllIdentifier() {
		name := id.GetText()
		var alloca llvm.Value
		if isTopLevel {
			alloca = llvm.AddGlobal(v.Llvmmod, llvmType, name)
			if hasInit && initVal.IsConstant() {
				alloca.SetInitializer(initVal)
			} else {
				alloca.SetInitializer(llvm.ConstNull(llvmType))
			}
			v.globalVars[name] = alloca
		} else {
			alloca = v.Builder.CreateAlloca(llvmType, name)
		}
		v.vars[name] = alloca
		if typekey != "auto" {
			v.varTypes[name] = typekey
		} else if hasInit {
			if sn := v.structNameOf(llvmType); sn != "" {
				v.varTypes[name] = sn
			}
		}
		if !pendingFuncType.IsNil() {
			v.closureVarFuncTypes[name] = pendingFuncType
		}
		if !symlinkPointee.IsNil() {
			v.symlinkPointeeTypes[name] = symlinkPointee
		}
		if !isTopLevel {
			if hasInit {
				val := v.coerce(initVal, llvmType)
				v.Builder.CreateStore(val, alloca)
			} else if strings.HasSuffix(typekey, "?") {
				v.Builder.CreateStore(llvm.ConstNull(llvmType), alloca)
			}
		}
	}
	return nil
}

func (v *Visitor) VisitDeclarator(ctx *parser.DeclaratorContext) any {
	return ctx.GetText()
}

func (v *Visitor) VisitAssignmentOp(ctx *parser.AssignmentOpContext) any {
	v.PrintNode(fmt.Sprintf("Default assignment operation: %s %s", ctx.AssignmentOperator().GetText(), ctx.Value().GetText()))
	return v.Visit(ctx.Value())
}

func (v *Visitor) VisitValue(ctx *parser.ValueContext) any {
	v.PrintNode(fmt.Sprintf("Value: %s", ctx.GetText()))

	// Parenthesized value or condition
	if ctx.LPAREN() != nil {
		if ctx.Condition() != nil {
			return v.Visit(ctx.Condition())
		}
		if len(ctx.AllValue()) == 1 {
			return v.Visit(ctx.Value(0))
		}
	}

	// Binary operation: value (op value)+
	// ANTLR builds a right-associative tree regardless of operator precedence,
	// so we flatten the entire tree into a token list first, then apply
	// shunting-yard to emit IR with correct precedence
	if len(ctx.AllValue()) >= 2 {
		var valCtxs []parser.IValueContext
		var opStrs []string
		v.flattenBinOpTree(ctx, &valCtxs, &opStrs)
		operands := make([]llvm.Value, len(valCtxs))
		for i, vc := range valCtxs {
			lv, ok := v.Visit(vc).(llvm.Value)
			if !ok {
				return nil
			}
			operands[i] = lv
		}
		return v.emitBinOpPrec(operands, opStrs)
	}

	if ctx.MemberAccess() != nil {
		return v.evalMemberAccess(ctx.MemberAccess().(*parser.MemberAccessContext))
	}

	if ctx.FuncCall() != nil {
		return v.Visit(ctx.FuncCall())
	}

	// Lambdas
	if ctx.FuncDecl() != nil {
		return v.Visit(ctx.FuncDecl())
	}

	if ctx.OperationWithReturn() != nil {
		return v.Visit(ctx.OperationWithReturn())
	}

	if ctx.TRUE() != nil {
		return llvm.ConstInt(v.Llvmctx.Int1Type(), 1, false)
	}
	if ctx.FALSE() != nil {
		return llvm.ConstInt(v.Llvmctx.Int1Type(), 0, false)
	}

	if ctx.NULL() != nil {
		return llvm.ConstNull(llvm.PointerType(v.Llvmctx.Int8Type(), 0))
	}

	if ctx.FloatingPointLiteral() != nil {
		text := ctx.FloatingPointLiteral().GetText()
		text = strings.TrimRight(text, "fFdD")
		f, err := strconv.ParseFloat(text, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: cannot parse float literal %q: %v\n", text, err)
			return llvm.ConstFloat(v.Llvmctx.DoubleType(), 0)
		}
		return llvm.ConstFloat(v.Llvmctx.DoubleType(), f)
	}

	if ctx.IntegerLiteral() != nil {
		text := ctx.IntegerLiteral().GetText()
		text = strings.TrimRight(text, "lL")
		n, err := strconv.ParseInt(text, 0, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: cannot parse integer literal %q: %v\n", text, err)
			return llvm.ConstInt(v.Llvmctx.Int64Type(), 0, false)
		}
		return llvm.ConstInt(v.Llvmctx.Int64Type(), uint64(n), true)
	}

	if ctx.StringLiteral() != nil {
		raw := ctx.StringLiteral().GetText()
		// Strip surrounding quotes and unescape basic sequences
		inner := raw[1 : len(raw)-1]
		inner = strings.ReplaceAll(inner, `\"`, `"`)
		inner = strings.ReplaceAll(inner, `\\`, `\`)
		inner = strings.ReplaceAll(inner, `\n`, "\n")
		inner = strings.ReplaceAll(inner, `\t`, "\t")
		inner = strings.ReplaceAll(inner, `\r`, "\r")
		if v.Builder.GetInsertBlock().IsNil() {
			strConst := llvm.ConstString(inner, true)
			g := llvm.AddGlobal(v.Llvmmod, strConst.Type(), "")
			g.SetInitializer(strConst)
			g.SetGlobalConstant(true)
			g.SetLinkage(llvm.InternalLinkage)
			i8ptr := llvm.PointerType(v.Llvmctx.Int8Type(), 0)
			return llvm.ConstBitCast(g, i8ptr)
		}
		return v.Builder.CreateGlobalStringPtr(inner, "")
	}

	if ctx.ADDR() != nil {
		name := ctx.Identifier().GetText()
		if alloca, ok := v.vars[name]; ok {
			if strings.HasSuffix(v.varTypes[name], "?") {
				// Return the stored address (pointer value), not the slot's own address
				return v.Builder.CreateLoad(alloca.AllocatedType(), alloca, name)
			}
			return alloca
		}
	}

	if ctx.Identifier() != nil {
		name := ctx.Identifier().GetText()
		// Closure capture: load through heap box pointer
		if heapPtr, captured := v.closureCaptures[name]; captured {
			return v.Builder.CreateLoad(v.Llvmctx.Int64Type(), heapPtr, name)
		}
		if _, ok := v.vars[name]; !ok {
			if gv, ok := v.globalVars[name]; ok {
				v.vars[name] = gv
				v.varTypes[name] = v.varTypes[name] // preserve if already set
			}
		}
		if alloca, ok := v.vars[name]; ok {
			var elemType llvm.Type
			if !alloca.IsAGlobalVariable().IsNil() {
				elemType = alloca.GlobalValueType()
			} else {
				elemType = alloca.AllocatedType()
			}
			loaded := v.Builder.CreateLoad(elemType, alloca, name)
			// Auto-dereference symlink typed variables
			if v.varTypes[name] == "symlink" {
				if pointeeType, ok := v.symlinkPointeeTypes[name]; ok && !pointeeType.IsNil() {
					return v.Builder.CreateLoad(pointeeType, loaded, name+".deref")
				}
			}
			if strings.HasSuffix(v.varTypes[name], "?") {
				thisFunc := v.Results[v.lastFunc].(llvm.Value)

				derefBlock := llvm.AddBasicBlock(thisFunc, name+".deref")
				nullBlock := llvm.AddBasicBlock(thisFunc, name+".null")
				mergeBlock := llvm.AddBasicBlock(thisFunc, name+".merge")

				pointeeType := v.ResolveType(strings.TrimSuffix(v.varTypes[name], "?"))
				isNull := v.Builder.CreateICmp(llvm.IntEQ, loaded, llvm.ConstNull(loaded.Type()), "")

				v.Builder.CreateCondBr(isNull, nullBlock, derefBlock)
				v.Builder.SetInsertPointAtEnd(derefBlock)

				val := v.Builder.CreateLoad(pointeeType, loaded, name+".val")
				derefBlock = v.Builder.GetInsertBlock()

				v.Builder.CreateBr(mergeBlock)
				v.Builder.SetInsertPointAtEnd(nullBlock)
				v.Builder.CreateBr(mergeBlock)
				v.Builder.SetInsertPointAtEnd(mergeBlock)

				phi := v.Builder.CreatePHI(pointeeType, name+".phi")
				phi.AddIncoming([]llvm.Value{val, llvm.ConstInt(pointeeType, 0, false)}, []llvm.BasicBlock{derefBlock, nullBlock})
				return phi
			}
			return loaded
		}
		// Check if it's a function (used as a value)
		fn := v.Llvmmod.NamedFunction(name)
		if !fn.IsNil() {
			return fn
		}
		// Check if it's a named function parameter already set in Results
		for _, res := range v.Results {
			if lv, ok := res.(llvm.Value); ok && lv.Name() == name {
				return lv
			}
		}
		fmt.Fprintf(os.Stderr, "Error: unknown identifier %q\n", name)
		os.Exit(-1)
		return nil
	}

	if ctx.ArrayMemberRef() != nil {
		return v.VisitArrayMemberRef(ctx.ArrayMemberRef().(*parser.ArrayMemberRefContext))
	}

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitArrayMemberRef(ctx *parser.ArrayMemberRefContext) any {
	var arrayPtr llvm.Value
	var elemTypeKey string
	isSlice := false

	if ctx.Identifier() != nil {
		name := ctx.Identifier().GetText()
		if typeKey, ok := v.varTypes[name]; ok {
			if idx := strings.Index(typeKey, "["); idx != -1 {
				elemTypeKey = typeKey[:idx]
				closeIdx := strings.Index(typeKey[idx:], "]")
				if closeIdx > 0 {
					lenStr := typeKey[idx+1 : idx+closeIdx]
					isSlice = lenStr == ""
				}
			}
		}
		if alloca, ok := v.vars[name]; ok {
			arrayPtr = v.Builder.CreateLoad(alloca.AllocatedType(), alloca, name)
		} else if fn, ok := v.Results[v.lastFunc].(llvm.Value); ok {
			for i := 0; i < fn.ParamsCount(); i++ {
				p := fn.Param(i)
				if p.Name() == name {
					arrayPtr = p
					break
				}
			}
		}
	} else if ctx.MemberAccess() != nil {
		if res, ok := v.evalMemberAccess(ctx.MemberAccess().(*parser.MemberAccessContext)).(llvm.Value); ok {
			arrayPtr = res
		}
	}

	if arrayPtr.IsNil() {
		return nil
	}

	accessCtx := ctx.ArrayAccess().(*parser.ArrayAccessContext)
	i64 := v.Llvmctx.Int64Type()
	var idxVal llvm.Value
	if accessCtx.IntegerLiteral() != nil {
		idx := int64(0)
		fmt.Sscanf(accessCtx.IntegerLiteral().GetText(), "%d", &idx)
		idxVal = llvm.ConstInt(i64, uint64(idx), false)
	} else if accessCtx.Identifier() != nil {
		name := accessCtx.Identifier().GetText()
		if alloca, ok := v.vars[name]; ok {
			idxVal = v.Builder.CreateLoad(i64, alloca, "")
		} else {
			idxVal = llvm.ConstInt(i64, 0, false)
		}
	} else {
		idxVal = llvm.ConstInt(i64, 0, false)
	}

	var elemType llvm.Type
	if elemTypeKey != "" {
		elemType = v.ResolveType(elemTypeKey)
	} else {
		elemType = llvm.PointerType(v.Llvmctx.Int8Type(), 0)
	}

	if isSlice {
		one := llvm.ConstInt(i64, 1, false)
		offsetIdx := v.Builder.CreateAdd(idxVal, one, "")
		ptr := v.Builder.CreateGEP(i64, arrayPtr, []llvm.Value{offsetIdx}, "")
		raw := v.Builder.CreateLoad(i64, ptr, "")
		return v.coerce(raw, elemType)
	}

	ptr := v.Builder.CreateGEP(elemType, arrayPtr, []llvm.Value{idxVal}, "")
	return v.Builder.CreateLoad(elemType, ptr, "")
}

func opPrec(op string) int {
	switch op {
	case "||":
		return 1
	case "&&":
		return 2
	case "|":
		return 3
	case "&":
		return 4
	case "<<", ">>":
		return 5
	case "+", "-":
		return 6
	case "*", "/", "%":
		return 7
	}
	return 0
}

func (v *Visitor) flattenBinOpTree(ctx parser.IValueContext, vals *[]parser.IValueContext, ops *[]string) {
	vc, ok := ctx.(*parser.ValueContext)
	if !ok || len(vc.AllValue()) < 2 {
		*vals = append(*vals, ctx)
		return
	}
	allVals := vc.AllValue()
	allOps := vc.AllOp()
	v.flattenBinOpTree(allVals[0], vals, ops)
	for i, op := range allOps {
		*ops = append(*ops, op.GetText())
		v.flattenBinOpTree(allVals[i+1], vals, ops)
	}
}

func (v *Visitor) emitBinOpPrec(operands []llvm.Value, operators []string) llvm.Value {
	var valStack []llvm.Value
	var opStack []string
	applyTop := func() {
		op := opStack[len(opStack)-1]
		opStack = opStack[:len(opStack)-1]
		b := valStack[len(valStack)-1]
		a := valStack[len(valStack)-2]
		valStack = valStack[:len(valStack)-2]
		valStack = append(valStack, v.emitBinOp(op, a, b))
	}
	valStack = append(valStack, operands[0])
	for i, op := range operators {
		for len(opStack) > 0 && opPrec(opStack[len(opStack)-1]) >= opPrec(op) {
			applyTop()
		}
		opStack = append(opStack, op)
		valStack = append(valStack, operands[i+1])
	}
	for len(opStack) > 0 {
		applyTop()
	}
	return valStack[0]
}

func (v *Visitor) emitBinOp(op string, lhs, rhs llvm.Value) llvm.Value {
	i64 := v.Llvmctx.Int64Type()
	lhsIsPtr := lhs.Type().TypeKind() == llvm.PointerTypeKind
	if lhsIsPtr && (op == "+" || op == "-") {
		lhsInt := v.Builder.CreatePtrToInt(lhs, i64, "")
		rhsInt := v.coerce(rhs, i64)
		var result llvm.Value
		if op == "+" {
			result = v.Builder.CreateAdd(lhsInt, rhsInt, "")
		} else {
			result = v.Builder.CreateSub(lhsInt, rhsInt, "")
		}
		return v.Builder.CreateIntToPtr(result, lhs.Type(), "")
	}
	isFloat := lhs.Type().TypeKind() == llvm.FloatTypeKind || lhs.Type().TypeKind() == llvm.DoubleTypeKind
	// Coerce rhs to lhs type if possible
	rhs = v.coerce(rhs, lhs.Type())
	switch op {
	case "+":
		if isFloat {
			return v.Builder.CreateFAdd(lhs, rhs, "")
		}
		return v.Builder.CreateAdd(lhs, rhs, "")
	case "-":
		if isFloat {
			return v.Builder.CreateFSub(lhs, rhs, "")
		}
		return v.Builder.CreateSub(lhs, rhs, "")
	case "*":
		if isFloat {
			return v.Builder.CreateFMul(lhs, rhs, "")
		}
		return v.Builder.CreateMul(lhs, rhs, "")
	case "/":
		if isFloat {
			return v.Builder.CreateFDiv(lhs, rhs, "")
		}
		return v.Builder.CreateSDiv(lhs, rhs, "")
	case "%":
		return v.Builder.CreateSRem(lhs, rhs, "")
	case "&":
		return v.Builder.CreateAnd(lhs, rhs, "")
	case "|":
		return v.Builder.CreateOr(lhs, rhs, "")
	case "&&":
		return v.Builder.CreateAnd(lhs, rhs, "")
	case "||":
		return v.Builder.CreateOr(lhs, rhs, "")
	case "<<":
		return v.Builder.CreateShl(lhs, rhs, "")
	case ">>":
		return v.Builder.CreateAShr(lhs, rhs, "")
	default:
		fmt.Fprintf(os.Stderr, "Warning: unhandled binary operator %q\n", op)
		return lhs
	}
}

func (v *Visitor) coerce(val llvm.Value, target llvm.Type) llvm.Value {
	if val.Type() == target {
		return val
	}
	srcKind := val.Type().TypeKind()
	dstKind := target.TypeKind()

	switch {
	case srcKind == llvm.IntegerTypeKind && dstKind == llvm.IntegerTypeKind:
		srcBits := val.Type().IntTypeWidth()
		dstBits := target.IntTypeWidth()
		if srcBits < dstBits {
			return v.Builder.CreateSExt(val, target, "")
		}
		if srcBits > dstBits {
			return v.Builder.CreateTrunc(val, target, "")
		}
	case srcKind == llvm.IntegerTypeKind && (dstKind == llvm.FloatTypeKind || dstKind == llvm.DoubleTypeKind):
		return v.Builder.CreateSIToFP(val, target, "")
	case (srcKind == llvm.FloatTypeKind || srcKind == llvm.DoubleTypeKind) && dstKind == llvm.IntegerTypeKind:
		return v.Builder.CreateFPToSI(val, target, "")
	case srcKind == llvm.FloatTypeKind && dstKind == llvm.DoubleTypeKind:
		return v.Builder.CreateFPExt(val, target, "")
	case srcKind == llvm.DoubleTypeKind && dstKind == llvm.FloatTypeKind:
		return v.Builder.CreateFPTrunc(val, target, "")
	case srcKind == llvm.PointerTypeKind && dstKind == llvm.PointerTypeKind:
		return v.Builder.CreateBitCast(val, target, "")
	case srcKind == llvm.IntegerTypeKind && dstKind == llvm.PointerTypeKind:
		return v.Builder.CreateIntToPtr(val, target, "")
	case srcKind == llvm.PointerTypeKind && dstKind == llvm.IntegerTypeKind:
		return v.Builder.CreatePtrToInt(val, target, "")
	}
	return val
}

func (v *Visitor) resolveVarType(varName string) (llvm.Type, string) {
	typeName := v.varTypes[varName]
	if typeName == "self" {
		typeName = v.selfType
	}
	st, ok := v.Results[typeName].(llvm.Type)
	if !ok {
		return llvm.Type{}, ""
	}
	return st, typeName
}

func (v *Visitor) structNameOf(t llvm.Type) string {
	if t.TypeKind() != llvm.StructTypeKind {
		return ""
	}
	if name := t.StructName(); name != "" {
		return name
	}
	// if StructName() fails match by type identity
	for name := range v.structFields {
		if st, ok := v.Results[name].(llvm.Type); ok && st == t {
			return name
		}
	}
	return ""
}

func (v *Visitor) fieldIndexOf(structName, fieldName string) int {
	for i, f := range v.structFields[structName] {
		if f == fieldName {
			return i
		}
	}
	return -1
}

func (v *Visitor) allocaFor(val llvm.Value, hint llvm.Value) llvm.Value {
	if !hint.IsNil() {
		return hint
	}
	tmp := v.Builder.CreateAlloca(val.Type(), "")
	v.Builder.CreateStore(val, tmp)
	return tmp
}

func (v *Visitor) emitMethodCallOnValue(baseVal llvm.Value, baseAlloca llvm.Value, fc *parser.FuncCallContext) llvm.Value {
	structType := baseVal.Type()
	structName := v.structNameOf(structType)
	if structName == "" {
		return llvm.Value{}
	}
	if st, ok := v.Results[structName].(llvm.Type); ok {
		structType = st
	}

	fieldName := fc.Identifier().GetText()
	idx := v.fieldIndexOf(structName, fieldName)
	if idx < 0 {
		fmt.Fprintf(os.Stderr, "Error: unknown method %q on %q\n", fieldName, structName)
		os.Exit(-1)
	}

	if pkgName, isPkg := v.pkgStructs[structName]; isPkg {
		qualName := pkgName + "::" + fieldName
		fn := v.Llvmmod.NamedFunction(qualName)
		if fn.IsNil() {
			fmt.Fprintf(os.Stderr, "Error: package %q has no method %q\n", pkgName, fieldName)
			os.Exit(-1)
		}
		fnType := fn.GlobalValueType()
		paramTypes := fnType.ParamTypes()
		ptr := v.allocaFor(baseVal, baseAlloca)
		args := []llvm.Value{ptr}
		if fc.CallArgs() != nil {
			for i, valCtx := range fc.CallArgs().(*parser.CallArgsContext).AllValue() {
				raw := v.Visit(valCtx)
				lv, ok := raw.(llvm.Value)
				if !ok {
					continue
				}
				pi := i + 1
				if pi < len(paramTypes) {
					lv = v.coerce(lv, paramTypes[pi])
				}
				args = append(args, lv)
			}
		}
		return v.Builder.CreateCall(fnType, fn, args, "")
	}

	ptr := v.allocaFor(baseVal, baseAlloca)
	fieldPtr := v.Builder.CreateStructGEP(structType, ptr, idx, "")
	fieldType := structType.StructElementTypes()[idx]
	fnRaw := v.Builder.CreateLoad(fieldType, fieldPtr, "")

	methodKey := structName + "." + fieldName
	fnType, hasFnType := v.methodTypes[methodKey]
	if !hasFnType {
		fmt.Fprintf(os.Stderr, "Warning: no method type recorded for %q\n", methodKey)
		os.Exit(-1)
	}

	if v.methodIsClosure[methodKey] {
		i64 := v.Llvmctx.Int64Type()
		i8ptr := llvm.PointerType(v.Llvmctx.Int8Type(), 0)
		closureI64Ptr := v.Builder.CreateBitCast(fnRaw, llvm.PointerType(i64, 0), "")
		fnPtrI64 := v.Builder.CreateLoad(i64, closureI64Ptr, "fn.i64")
		envSlot := v.Builder.CreateInBoundsGEP(i64, closureI64Ptr, []llvm.Value{llvm.ConstInt(i64, 1, false)}, "")
		envI64 := v.Builder.CreateLoad(i64, envSlot, "env.i64")
		envPtr := v.Builder.CreateIntToPtr(envI64, i8ptr, "env.ptr")
		fnPtr := v.Builder.CreateIntToPtr(fnPtrI64, llvm.PointerType(fnType, 0), "fn.ptr")
		paramTypes := fnType.ParamTypes()
		args := []llvm.Value{envPtr}
		if len(paramTypes) > 1 {
			selfParamType := paramTypes[1]
			var selfArg llvm.Value
			if selfParamType.TypeKind() == llvm.StructTypeKind {
				selfArg = v.Builder.CreateLoad(selfParamType, ptr, "")
			} else {
				selfArg = v.Builder.CreateBitCast(ptr, selfParamType, "")
			}
			args = append(args, selfArg)
		}
		if fc.CallArgs() != nil {
			for i, valCtx := range fc.CallArgs().(*parser.CallArgsContext).AllValue() {
				raw := v.Visit(valCtx)
				lv, ok := raw.(llvm.Value)
				if !ok {
					continue
				}
				pi := i + 2
				if pi < len(paramTypes) {
					lv = v.coerce(lv, paramTypes[pi])
				}
				args = append(args, lv)
			}
		}
		return v.Builder.CreateCall(fnType, fnPtr, args, "")
	}

	fnPtrType := llvm.PointerType(fnType, 0)
	fnPtr := v.Builder.CreateBitCast(fnRaw, fnPtrType, "")
	paramTypes := fnType.ParamTypes()
	var selfArg llvm.Value
	if len(paramTypes) > 0 {
		selfParamType := paramTypes[0]
		if selfParamType.TypeKind() == llvm.StructTypeKind {
			selfArg = v.Builder.CreateLoad(selfParamType, ptr, "")
		} else {
			selfArg = v.Builder.CreateBitCast(ptr, selfParamType, "")
		}
	}
	var args []llvm.Value
	if !selfArg.IsNil() {
		args = append(args, selfArg)
	}
	if fc.CallArgs() != nil {
		for i, valCtx := range fc.CallArgs().(*parser.CallArgsContext).AllValue() {
			raw := v.Visit(valCtx)
			lv, ok := raw.(llvm.Value)
			if !ok {
				continue
			}
			pi := i + 1
			if pi < len(paramTypes) {
				lv = v.coerce(lv, paramTypes[pi])
			}
			args = append(args, lv)
		}
	}
	return v.Builder.CreateCall(fnType, fnPtr, args, "")
}

func (v *Visitor) evalMemberAccess(ma *parser.MemberAccessContext) any {
	if ma.NamespaceAccess() != nil {
		return v.VisitChildren(ma)
	}
	funcCalls := ma.AllFuncCall()
	idents := ma.AllIdentifier()
	dots := len(ma.AllDOT())

	var baseVal llvm.Value
	var baseAlloca llvm.Value
	var memberFuncCall *parser.FuncCallContext
	var memberIdent string

	switch {
	case len(idents) == 1 && len(funcCalls) >= 2 && dots >= 2 && len(funcCalls) == dots:
		name := idents[0].GetText()
		alloca, ok := v.vars[name]
		if !ok {
			fmt.Fprintf(os.Stderr, "Error: unknown identifier %q\n", name)
			os.Exit(-1)
		}
		st, sn := v.resolveVarType(name)
		if st.IsNil() {
			fmt.Fprintf(os.Stderr, "Error: %q is not a struct type (type=%q)\n", name, sn)
			os.Exit(-1)
		}
		cur := v.Builder.CreateLoad(st, alloca, name)
		curAlloca := alloca
		curStructName := sn
		for _, fc := range funcCalls {
			result := v.emitMethodCallOnValue(cur, curAlloca, fc.(*parser.FuncCallContext))
			if result.IsNil() {
				return nil
			}
			fieldName := fc.(*parser.FuncCallContext).Identifier().GetText()
			pkgName := v.pkgStructs[curStructName]
			retStructName := v.funcReturnStructName[pkgName+"::"+fieldName]
			if retStructName != "" {
				if retStructType, ok := v.Results[retStructName].(llvm.Type); ok {
					cur = v.Builder.CreateLoad(retStructType, result, "")
					curAlloca = result
					curStructName = retStructName
					continue
				}
			}
			cur = result
			curAlloca = llvm.Value{}
			curStructName = v.structNameOf(result.Type())
		}
		if cur.Type().TypeKind() == llvm.StructTypeKind {
			v.Builder.CreateStore(cur, alloca)
		}
		return cur

	case len(idents) == 0 && len(funcCalls) == 2 && dots == 1:
		baseFc := funcCalls[0].(*parser.FuncCallContext)
		raw := v.Visit(baseFc)
		lv, ok := raw.(llvm.Value)
		if !ok {
			return nil
		}
		baseVal = lv
		memberFuncCall = funcCalls[1].(*parser.FuncCallContext)

	case len(idents) == 1 && len(funcCalls) == 1 && dots == 1:
		name := idents[0].GetText()
		alloca, ok := v.vars[name]
		if !ok {
			fmt.Fprintf(os.Stderr, "Error: unknown identifier %q\n", name)
			os.Exit(-1)
		}
		baseAlloca = alloca
		st, sn := v.resolveVarType(name)
		if st.IsNil() {
			fmt.Fprintf(os.Stderr, "Error: %q is not a struct type (type=%q)\n", name, sn)
			os.Exit(-1)
		}
		baseVal = v.Builder.CreateLoad(st, alloca, name)
		memberFuncCall = funcCalls[0].(*parser.FuncCallContext)

	case len(idents) == 2 && len(funcCalls) == 0 && dots == 1:
		name := idents[0].GetText()
		alloca, ok := v.vars[name]
		if !ok {
			fmt.Fprintf(os.Stderr, "Error: unknown identifier %q\n", name)
			os.Exit(-1)
		}
		baseAlloca = alloca
		st, sn := v.resolveVarType(name)
		if st.IsNil() {
			fmt.Fprintf(os.Stderr, "Error: %q is not a struct type (type=%q)\n", name, sn)
			os.Exit(-1)
		}
		baseVal = v.Builder.CreateLoad(st, alloca, name)
		memberIdent = idents[1].GetText()

	default:
		return v.VisitChildren(ma)
	}

	structType := baseVal.Type()
	structName := v.structNameOf(structType)
	if structName == "" {
		if len(idents) > 0 {
			_, structName = v.resolveVarType(idents[0].GetText())
		} else if len(funcCalls) >= 1 {
			baseFc := funcCalls[0].(*parser.FuncCallContext)
			baseFnName := baseFc.Identifier().GetText()
			if baseFc.NamespaceAccess() != nil {
				ns := baseFc.NamespaceAccess().(*parser.NamespaceAccessContext)
				nsName := ns.Identifier(0).GetText()
				var fnName string
				if fc := ns.FuncCall(); fc != nil {
					fnName = fc.(*parser.FuncCallContext).Identifier().GetText()
				} else {
					fnName = ns.Identifier(1).GetText()
				}
				baseFnName = nsName + "::" + fnName
			}
			if fn := v.Llvmmod.NamedFunction(baseFnName); !fn.IsNil() {
				structName = v.structNameOf(fn.GlobalValueType().ReturnType())
			}
		}
		if st, ok := v.Results[structName].(llvm.Type); ok {
			structType = st
		}
	}

	if memberIdent != "" {
		idx := v.fieldIndexOf(structName, memberIdent)
		if idx < 0 {
			fmt.Fprintf(os.Stderr, "Error: unknown field %q on %q\n", memberIdent, structName)
			os.Exit(-1)
		}
		ptr := v.allocaFor(baseVal, baseAlloca)
		fieldPtr := v.Builder.CreateStructGEP(structType, ptr, idx, "")
		fieldType := structType.StructElementTypes()[idx]
		return v.Builder.CreateLoad(fieldType, fieldPtr, memberIdent)
	}

	fieldName := memberFuncCall.Identifier().GetText()
	idx := v.fieldIndexOf(structName, fieldName)
	if idx < 0 {
		fmt.Fprintf(os.Stderr, "Error: unknown method %q on %q\n", fieldName, structName)
		os.Exit(-1)
	}

	if pkgName, isPkg := v.pkgStructs[structName]; isPkg {
		qualName := pkgName + "::" + fieldName
		fn := v.Llvmmod.NamedFunction(qualName)
		if fn.IsNil() {
			fmt.Fprintf(os.Stderr, "Error: package %q has no method %q\n", pkgName, fieldName)
			os.Exit(-1)
		}
		fnType := fn.GlobalValueType()
		paramTypes := fnType.ParamTypes()
		ptr := v.allocaFor(baseVal, baseAlloca)
		args := []llvm.Value{ptr}
		if memberFuncCall.CallArgs() != nil {
			for i, valCtx := range memberFuncCall.CallArgs().(*parser.CallArgsContext).AllValue() {
				raw := v.Visit(valCtx)
				lv, ok := raw.(llvm.Value)
				if !ok {
					continue
				}
				pi := i + 1
				if pi < len(paramTypes) {
					lv = v.coerce(lv, paramTypes[pi])
				}
				args = append(args, lv)
			}
		}
		return v.Builder.CreateCall(fnType, fn, args, "")
	}

	ptr := v.allocaFor(baseVal, baseAlloca)
	fieldPtr := v.Builder.CreateStructGEP(structType, ptr, idx, "")
	fieldType := structType.StructElementTypes()[idx]
	fnRaw := v.Builder.CreateLoad(fieldType, fieldPtr, "")

	methodKey := structName + "." + fieldName
	fnType, hasFnType := v.methodTypes[methodKey]
	if !hasFnType {
		fmt.Fprintf(os.Stderr, "Warning: no method type recorded for %q\n", methodKey)
		os.Exit(-1)
	}

	if v.methodIsClosure[methodKey] {
		// fnRaw is an i8* closure pointer: [fn_ptr: i64, env: i64]
		i64 := v.Llvmctx.Int64Type()
		i8ptr := llvm.PointerType(v.Llvmctx.Int8Type(), 0)
		closureI64Ptr := v.Builder.CreateBitCast(fnRaw, llvm.PointerType(i64, 0), "")
		fnPtrI64 := v.Builder.CreateLoad(i64, closureI64Ptr, "fn.i64")
		envSlot := v.Builder.CreateInBoundsGEP(i64, closureI64Ptr, []llvm.Value{llvm.ConstInt(i64, 1, false)}, "")
		envI64 := v.Builder.CreateLoad(i64, envSlot, "env.i64")
		envPtr := v.Builder.CreateIntToPtr(envI64, i8ptr, "env.ptr")

		fnPtr := v.Builder.CreateIntToPtr(fnPtrI64, llvm.PointerType(fnType, 0), "fn.ptr")

		// closure fn signature: (env i8*, self, ...args)
		paramTypes := fnType.ParamTypes()
		args := []llvm.Value{envPtr}
		if len(paramTypes) > 1 {
			selfParamType := paramTypes[1]
			var selfArg llvm.Value
			if selfParamType.TypeKind() == llvm.StructTypeKind {
				selfArg = v.Builder.CreateLoad(selfParamType, ptr, "")
			} else {
				selfArg = v.Builder.CreateBitCast(ptr, selfParamType, "")
			}
			args = append(args, selfArg)
		}
		if memberFuncCall.CallArgs() != nil {
			for i, valCtx := range memberFuncCall.CallArgs().(*parser.CallArgsContext).AllValue() {
				raw := v.Visit(valCtx)
				lv, ok := raw.(llvm.Value)
				if !ok {
					continue
				}
				pi := i + 2 // +2: skip env and self
				if pi < len(paramTypes) {
					lv = v.coerce(lv, paramTypes[pi])
				}
				args = append(args, lv)
			}
		}
		return v.Builder.CreateCall(fnType, fnPtr, args, "")
	}

	fnPtrType := llvm.PointerType(fnType, 0)
	fnPtr := v.Builder.CreateBitCast(fnRaw, fnPtrType, "")

	paramTypes := fnType.ParamTypes()
	var selfArg llvm.Value
	if len(paramTypes) > 0 {
		selfParamType := paramTypes[0]
		if selfParamType.TypeKind() == llvm.StructTypeKind {
			selfArg = v.Builder.CreateLoad(selfParamType, ptr, "")
		} else {
			selfArg = v.Builder.CreateBitCast(ptr, selfParamType, "")
		}
	}
	var args []llvm.Value
	if !selfArg.IsNil() {
		args = append(args, selfArg)
	}
	if memberFuncCall.CallArgs() != nil {
		for i, valCtx := range memberFuncCall.CallArgs().(*parser.CallArgsContext).AllValue() {
			raw := v.Visit(valCtx)
			lv, ok := raw.(llvm.Value)
			if !ok {
				continue
			}
			pi := i + 1 // offset by 1 for self
			if pi < len(paramTypes) {
				lv = v.coerce(lv, paramTypes[pi])
			}
			args = append(args, lv)
		}
	}
	return v.Builder.CreateCall(fnType, fnPtr, args, "")
}

func (v *Visitor) emitMemberAssign(ma *parser.MemberAccessContext, aop *parser.AssignmentOpContext) any {
	idents := ma.AllIdentifier()
	if len(idents) < 2 || ma.NamespaceAccess() != nil || len(ma.AllFuncCall()) != 0 {
		fmt.Println("Error: unsupported member assignment form")
		os.Exit(-1)
	}
	baseName := idents[0].GetText()
	fieldName := idents[len(idents)-1].GetText()
	alloca, ok := v.vars[baseName]
	if !ok {
		fmt.Fprintf(os.Stderr, "Error: undefined variable %q in member assignment\n", baseName)
		os.Exit(-1)
	}
	structName := v.varTypes[baseName]
	if structName == "self" {
		structName = v.selfType
	}
	structType, stOk := v.Results[structName].(llvm.Type)
	if !stOk || structType.IsNil() {
		fmt.Fprintf(os.Stderr, "Error: %q is not a struct type (type=%q)\n", baseName, structName)
		os.Exit(-1)
	}
	idx := v.fieldIndexOf(structName, fieldName)
	if idx < 0 {
		fmt.Fprintf(os.Stderr, "Error: unknown field %q on %q\n", fieldName, structName)
		os.Exit(-1)
	}
	v.selfType = structName
	raw := v.Visit(aop.Value())
	pendingFuncType := v.pendingClosureFuncType
	v.pendingClosureFuncType = llvm.Type{}
	v.selfType = ""
	rhs, ok := raw.(llvm.Value)
	if !ok {
		return nil
	}
	fieldPtr := v.Builder.CreateStructGEP(structType, alloca, idx, "")
	fieldElemType := structType.StructElementTypes()[idx]
	// Check if rhs is a function value by inspecting its pointee type
	rhsGlobalType := rhs.GlobalValueType()
	if !rhsGlobalType.IsNil() && rhsGlobalType.TypeKind() == llvm.FunctionTypeKind {
		v.methodTypes[structName+"."+fieldName] = rhsGlobalType
		origFn := rhs
		if v.Llvmmod.NamedFunction(fieldName).IsNil() {
			origParams := rhsGlobalType.ParamTypes()
			i8ptr := llvm.PointerType(v.Llvmctx.Int8Type(), 0)
			if len(origParams) > 0 && origParams[0].TypeKind() == llvm.StructTypeKind {
				selfStructType := origParams[0]
				wrapperParams := []llvm.Type{i8ptr}
				for _, p := range origParams[1:] {
					wrapperParams = append(wrapperParams, p)
				}
				retType := rhsGlobalType.ReturnType()
				wrapperRetType := retType
				if retType == selfStructType || retType.TypeKind() == llvm.PointerTypeKind {
					wrapperRetType = i8ptr
				}
				wrapperFnType := llvm.FunctionType(wrapperRetType, wrapperParams, false)
				wrapper := llvm.AddFunction(v.Llvmmod, fieldName, wrapperFnType)
				wrapper.AddFunctionAttr(v.Llvmctx.CreateStringAttribute("wasm-export-name", fieldName))
				savedBlock := v.Builder.GetInsertBlock()
				entry := v.Llvmctx.AddBasicBlock(wrapper, "entry")
				v.Builder.SetInsertPointAtEnd(entry)
				selfPtr := wrapper.Param(0)
				selfVal := v.Builder.CreateLoad(selfStructType, selfPtr, "self")
				callArgs := []llvm.Value{selfVal}
				for i := 1; i < wrapper.ParamsCount(); i++ {
					callArgs = append(callArgs, wrapper.Param(i))
				}
				result := v.Builder.CreateCall(rhsGlobalType, origFn, callArgs, "")
				if retType.TypeKind() == llvm.VoidTypeKind {
					v.Builder.CreateRetVoid()
				} else if retType == selfStructType {
					v.Builder.CreateStore(result, selfPtr)
					v.Builder.CreateRet(selfPtr)
				} else if retType.TypeKind() == llvm.PointerTypeKind {
					modified := v.Builder.CreateLoad(selfStructType, result, "modified")
					v.Builder.CreateStore(modified, selfPtr)
					v.Builder.CreateRet(selfPtr)
				} else {
					v.Builder.CreateRet(result)
				}
				v.Builder.SetInsertPointAtEnd(savedBlock)
			} else {
				wrapper := llvm.AddFunction(v.Llvmmod, fieldName, rhsGlobalType)
				wrapper.AddFunctionAttr(v.Llvmctx.CreateStringAttribute("wasm-export-name", fieldName))
				savedBlock := v.Builder.GetInsertBlock()
				entry := v.Llvmctx.AddBasicBlock(wrapper, "entry")
				v.Builder.SetInsertPointAtEnd(entry)
				callArgs := make([]llvm.Value, wrapper.ParamsCount())
				for i := 0; i < wrapper.ParamsCount(); i++ {
					callArgs[i] = wrapper.Param(i)
				}
				result := v.Builder.CreateCall(rhsGlobalType, origFn, callArgs, "")
				if rhsGlobalType.ReturnType().TypeKind() == llvm.VoidTypeKind {
					v.Builder.CreateRetVoid()
				} else {
					v.Builder.CreateRet(result)
				}
				v.Builder.SetInsertPointAtEnd(savedBlock)
			}
		}
		if rhs.Type() != fieldElemType {
			rhs = v.Builder.CreateBitCast(rhs, fieldElemType, "")
		}
	} else if !pendingFuncType.IsNil() {
		// RHS was a closure - record its function type for later method calls
		v.methodTypes[structName+"."+fieldName] = pendingFuncType
		v.methodIsClosure[structName+"."+fieldName] = true
		rhs = v.coerce(rhs, fieldElemType)

		innerFn := v.pendingClosureInnerFn
		if !innerFn.IsNil() && v.Llvmmod.NamedFunction(fieldName).IsNil() {
			closureParams := pendingFuncType.ParamTypes() // (env i8*, self structT, ...args)
			// Wrapper receives self as a pointer so mutations inside the closure
			i8ptr := llvm.PointerType(v.Llvmctx.Int8Type(), 0)
			selfStructType := closureParams[1] // the actual struct type
			selfPtrType := llvm.PointerType(v.Llvmctx.Int8Type(), 0)
			wrapperParams := []llvm.Type{selfPtrType}
			for _, p := range closureParams[2:] { // remaining args after self
				wrapperParams = append(wrapperParams, p)
			}
			retType := pendingFuncType.ReturnType()
			wrapperRetType := retType
			if retType == selfStructType || retType.TypeKind() == llvm.PointerTypeKind {
				wrapperRetType = selfPtrType
			}
			wrapperFnType := llvm.FunctionType(wrapperRetType, wrapperParams, false)
			wrapper := llvm.AddFunction(v.Llvmmod, fieldName, wrapperFnType)
			wrapper.AddFunctionAttr(v.Llvmctx.CreateStringAttribute("wasm-export-name", fieldName))

			savedBlock := v.Builder.GetInsertBlock()
			entry := v.Llvmctx.AddBasicBlock(wrapper, "entry")
			v.Builder.SetInsertPointAtEnd(entry)

			selfPtr := wrapper.Param(0)
			selfVal := v.Builder.CreateLoad(selfStructType, selfPtr, "self")
			callArgs := []llvm.Value{llvm.ConstNull(i8ptr), selfVal}
			for i := 1; i < wrapper.ParamsCount(); i++ {
				callArgs = append(callArgs, wrapper.Param(i))
			}
			result := v.Builder.CreateCall(pendingFuncType, innerFn, callArgs, "")
			if retType.TypeKind() == llvm.VoidTypeKind {
				v.Builder.CreateRetVoid()
			} else if retType == selfStructType {
				v.Builder.CreateStore(result, selfPtr)
				v.Builder.CreateRet(selfPtr)
			} else if retType.TypeKind() == llvm.PointerTypeKind {
				modified := v.Builder.CreateLoad(selfStructType, result, "modified")
				v.Builder.CreateStore(modified, selfPtr)
				v.Builder.CreateRet(selfPtr)
			} else {
				v.Builder.CreateRet(result)
			}

			v.Builder.SetInsertPointAtEnd(savedBlock)
		}
		v.pendingClosureInnerFn = llvm.Value{}
	} else {
		rhs = v.coerce(rhs, fieldElemType)
	}
	return v.Builder.CreateStore(rhs, fieldPtr)
}

func (v *Visitor) VisitOp(ctx *parser.OpContext) any {
	v.PrintNode(fmt.Sprintf("Operator: %s", ctx.GetText()))
	return ctx.GetText()
}

func (v *Visitor) VisitTypeKey(ctx *parser.TypeKeyContext) any {
	return ctx.GetText()
}

func (v *Visitor) VisitAssignment(ctx *parser.AssignmentContext) any {
	v.PrintNode(fmt.Sprintf("Assignment operation: %s", ctx.GetText()))

	if ctx.Identifier() == nil {
		if ctx.MemberAccess() != nil {
			return v.emitMemberAssign(ctx.MemberAccess().(*parser.MemberAccessContext),
				ctx.AssignmentOp().(*parser.AssignmentOpContext))
		}
		return nil
	}
	name := ctx.Identifier().GetText()

	// Closure capture: store/modify through heap box pointer
	if heapPtr, captured := v.closureCaptures[name]; captured {
		raw := v.Visit(ctx.AssignmentOp().(*parser.AssignmentOpContext).Value())
		rhs, ok := raw.(llvm.Value)
		if !ok {
			return nil
		}
		i64 := v.Llvmctx.Int64Type()
		op := ctx.AssignmentOp().(*parser.AssignmentOpContext).AssignmentOperator().GetText()
		rhs = v.coerce(rhs, i64)
		if op == "=" {
			return v.Builder.CreateStore(rhs, heapPtr)
		}
		loaded := v.Builder.CreateLoad(i64, heapPtr, "")
		var storeVal llvm.Value
		switch op {
		case "+=":
			storeVal = v.emitBinOp("+", loaded, rhs)
		case "-=":
			storeVal = v.emitBinOp("-", loaded, rhs)
		case "*=":
			storeVal = v.emitBinOp("*", loaded, rhs)
		case "/=":
			storeVal = v.emitBinOp("/", loaded, rhs)
		default:
			storeVal = rhs
		}
		return v.Builder.CreateStore(storeVal, heapPtr)
	}

	alloca, exists := v.vars[name]
	if !exists {
		fmt.Fprintf(os.Stderr, "Warning: assignment to undeclared variable %q\n", name)
		return nil
	}

	raw := v.Visit(ctx.AssignmentOp().(*parser.AssignmentOpContext).Value())
	rhs, ok := raw.(llvm.Value)
	if !ok {
		return nil
	}

	elemType := alloca.AllocatedType()
	op := ctx.AssignmentOp().(*parser.AssignmentOpContext).AssignmentOperator().GetText()

	var storeVal llvm.Value
	if op == "=" {
		if v.varTypes[name] == "symlink" {
			ptr := v.Builder.CreateLoad(elemType, alloca, "")
			if pointeeType, ok := v.symlinkPointeeTypes[name]; ok && !pointeeType.IsNil() {
				return v.Builder.CreateStore(v.coerce(rhs, pointeeType), ptr)
			}
			return v.Builder.CreateStore(rhs, ptr)
		}
		if strings.HasSuffix(v.varTypes[name], "?") {
			pointeeType := v.ResolveType(strings.TrimSuffix(v.varTypes[name], "?"))
			curPtr := v.Builder.CreateLoad(elemType, alloca, "")
			isNull := v.Builder.CreateICmp(llvm.IntEQ, curPtr, llvm.ConstNull(elemType), "")
			thisFunc := v.Results[v.lastFunc].(llvm.Value)
			allocBlock := llvm.AddBasicBlock(thisFunc, name+".alloc")
			storeBlock := llvm.AddBasicBlock(thisFunc, name+".store")
			v.Builder.CreateCondBr(isNull, allocBlock, storeBlock)
			v.Builder.SetInsertPointAtEnd(allocBlock)
			malloc := v.getOrDeclareMalloc()
			mallocType := llvm.FunctionType(llvm.PointerType(v.Llvmctx.Int8Type(), 0), []llvm.Type{v.Llvmctx.Int64Type()}, false)
			size := llvm.ConstInt(v.Llvmctx.Int64Type(), uint64(pointeeType.IntTypeWidth()/8), false)
			newPtr := v.Builder.CreateCall(mallocType, malloc, []llvm.Value{size}, name+".ptr")
			v.Builder.CreateStore(newPtr, alloca)
			v.Builder.CreateBr(storeBlock)
			v.Builder.SetInsertPointAtEnd(storeBlock)
			ptr := v.Builder.CreateLoad(elemType, alloca, "")
			return v.Builder.CreateStore(v.coerce(rhs, pointeeType), ptr)
		}
		storeVal = v.coerce(rhs, elemType)
	} else if op == "$=" {
		valCtx := ctx.AssignmentOp().(*parser.AssignmentOpContext).Value()
		if valCtx.(*parser.ValueContext).Identifier() != nil {
			ident := valCtx.(*parser.ValueContext).Identifier().GetText()
			if targetAlloca, ok := v.vars[ident]; ok {
				storeVal = v.coerce(targetAlloca, elemType)
				// Update the symlink pointee type to match the new target
				v.symlinkPointeeTypes[name] = targetAlloca.AllocatedType()
			} else {
				storeVal = v.coerce(rhs, elemType)
			}
		} else {
			storeVal = v.coerce(rhs, elemType)
		}
	} else {
		loaded := v.Builder.CreateLoad(elemType, alloca, "")
		rhs = v.coerce(rhs, elemType)
		switch op {
		case "+=":
			storeVal = v.emitBinOp("+", loaded, rhs)
		case "-=":
			storeVal = v.emitBinOp("-", loaded, rhs)
		case "*=":
			storeVal = v.emitBinOp("*", loaded, rhs)
		case "/=":
			storeVal = v.emitBinOp("/", loaded, rhs)
		default:
			fmt.Fprintf(os.Stderr, "Warning: unhandled assignment operator %q\n", op)
			storeVal = rhs
		}
	}
	return v.Builder.CreateStore(storeVal, alloca)
}

func (v *Visitor) VisitLoopStatement(ctx *parser.LoopStatementContext) any {
	v.PrintNode(fmt.Sprintf("Loop: %s", ctx.GetText()))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLoopDecl(ctx *parser.LoopDeclContext) any {
	thisFunc := v.Results[v.lastFunc].(llvm.Value)
	bodyBlock := llvm.AddBasicBlock(thisFunc, "loop.body")
	exitBlock := llvm.AddBasicBlock(thisFunc, "loop.exit")

	v.Builder.CreateBr(bodyBlock)
	v.Builder.SetInsertPointAtEnd(bodyBlock)

	savedBreak := v.breakTarget
	savedContinue := v.continueTarget
	v.breakTarget = exitBlock
	v.continueTarget = bodyBlock

	v.emitStatements(ctx.Body().(*parser.BodyContext).AllStatement())

	if !v.blockHasTerminator() {
		v.Builder.CreateBr(bodyBlock)
	}

	v.breakTarget = savedBreak
	v.continueTarget = savedContinue
	v.Builder.SetInsertPointAtEnd(exitBlock)
	v.emitEscapeCheck(savedBreak)
	return nil
}

func (v *Visitor) VisitTimesDecl(ctx *parser.TimesDeclContext) any {
	v.PrintNode(fmt.Sprintf("Times (%s):", ctx.Value().GetText()))
	// Emit a counted loop: for i = 0; i < count; i++
	thisFunc := v.Results[v.lastFunc].(llvm.Value)
	condBlock := llvm.AddBasicBlock(thisFunc, "times.cond")
	bodyBlock := llvm.AddBasicBlock(thisFunc, "times.body")
	incrBlock := llvm.AddBasicBlock(thisFunc, "times.incr")
	exitBlock := llvm.AddBasicBlock(thisFunc, "times.exit")

	i64 := v.Llvmctx.Int64Type()
	counter := v.Builder.CreateAlloca(i64, "times.counter")
	v.Builder.CreateStore(llvm.ConstInt(i64, 0, false), counter)
	v.Builder.CreateBr(condBlock)

	// Condition: counter < count
	v.Builder.SetInsertPointAtEnd(condBlock)
	countVal, _ := v.Visit(ctx.Value()).(llvm.Value)
	if countVal.IsNil() {
		countVal = llvm.ConstInt(i64, 0, false)
	}
	countVal = v.coerce(countVal, i64)
	cur := v.Builder.CreateLoad(i64, counter, "")
	cmp := v.Builder.CreateICmp(llvm.IntSLT, cur, countVal, "")
	v.Builder.CreateCondBr(cmp, bodyBlock, exitBlock)

	// Body
	v.Builder.SetInsertPointAtEnd(bodyBlock)
	savedBreak := v.breakTarget
	savedContinue := v.continueTarget
	v.breakTarget = exitBlock
	v.continueTarget = incrBlock
	v.emitStatements(ctx.Body().(*parser.BodyContext).AllStatement())
	if !v.blockHasTerminator() {
		v.Builder.CreateBr(incrBlock)
	}
	v.breakTarget = savedBreak
	v.continueTarget = savedContinue

	v.Builder.SetInsertPointAtEnd(incrBlock)
	next := v.Builder.CreateAdd(v.Builder.CreateLoad(i64, counter, ""), llvm.ConstInt(i64, 1, false), "")
	v.Builder.CreateStore(next, counter)
	v.Builder.CreateBr(condBlock)

	v.Builder.SetInsertPointAtEnd(exitBlock)
	v.emitEscapeCheck(savedBreak)
	return nil
}

func (v *Visitor) VisitOperationWithReturn(ctx *parser.OperationWithReturnContext) any {
	if ctx.TypeofOperation() != nil {
		return v.Visit(ctx.TypeofOperation())
	}
	if ctx.NegateOperation() != nil {
		return v.Visit(ctx.NegateOperation())
	}
	if ctx.PrefixOperation() != nil {
		return v.Visit(ctx.PrefixOperation())
	}
	if ctx.SuffixOperation() != nil {
		return v.Visit(ctx.SuffixOperation())
	}
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTypeofOperation(ctx *parser.TypeofOperationContext) any {
	v.PrintNode(fmt.Sprintf("Typeof: %s", ctx.Value()))
	valCtx := ctx.Value().(*parser.ValueContext)
	if valCtx.Identifier() != nil {
		name := valCtx.Identifier().GetText()
		if typeName, ok := v.varTypes[name]; ok {
			return v.Builder.CreateGlobalStringPtr(typeName, "")
		}
	}
	// Fallback: evaluate the value and derive type name from LLVM type
	raw := v.Visit(ctx.Value())
	val, ok := raw.(llvm.Value)
	if !ok {
		return v.Builder.CreateGlobalStringPtr("unknown", "")
	}
	var typeName string
	switch val.Type().TypeKind() {
	case llvm.IntegerTypeKind:
		switch val.Type().IntTypeWidth() {
		case 1:
			typeName = "bool"
		case 8:
			typeName = "byte"
		case 32:
			typeName = "int32"
		default:
			typeName = "int"
		}
	case llvm.DoubleTypeKind, llvm.FloatTypeKind:
		typeName = "float"
	case llvm.PointerTypeKind:
		typeName = "ptr"
	default:
		typeName = "unknown"
	}
	return v.Builder.CreateGlobalStringPtr(typeName, "")
}

func (v *Visitor) VisitPrefixOperation(ctx *parser.PrefixOperationContext) any {
	v.PrintNode(fmt.Sprintf("Operation: %s", ctx.GetText()))
	op := ctx.IncDecOP().GetText()
	name := ctx.Identifier().GetText()
	alloca, ok := v.vars[name]
	if !ok {
		return nil
	}
	elemType := alloca.AllocatedType()
	cur := v.Builder.CreateLoad(elemType, alloca, "")
	one := llvm.ConstInt(elemType, 1, false)
	var next llvm.Value
	if op == "++" {
		next = v.Builder.CreateAdd(cur, one, "")
	} else {
		next = v.Builder.CreateSub(cur, one, "")
	}
	v.Builder.CreateStore(next, alloca)
	return next
}

func (v *Visitor) VisitSuffixOperation(ctx *parser.SuffixOperationContext) any {
	v.PrintNode(fmt.Sprintf("Operation: %s", ctx.GetText()))
	op := ctx.IncDecOP().GetText()
	name := ctx.Identifier().GetText()
	alloca, ok := v.vars[name]
	if !ok {
		return nil
	}
	elemType := alloca.AllocatedType()
	cur := v.Builder.CreateLoad(elemType, alloca, "")
	one := llvm.ConstInt(elemType, 1, false)
	var next llvm.Value
	if op == "++" {
		next = v.Builder.CreateAdd(cur, one, "")
	} else {
		next = v.Builder.CreateSub(cur, one, "")
	}
	v.Builder.CreateStore(next, alloca)
	return cur
}

func (v *Visitor) VisitIncDecOP(ctx *parser.IncDecOPContext) any {
	return ctx.GetText()
}

func (v *Visitor) VisitNegateOperation(ctx *parser.NegateOperationContext) any {
	v.PrintNode(fmt.Sprintf("Operation: %s", ctx.GetText()))
	raw := v.Visit(ctx.Value())
	val, ok := raw.(llvm.Value)
	if !ok {
		return nil
	}
	zero := llvm.ConstInt(val.Type(), 0, false)
	return v.Builder.CreateICmp(llvm.IntEQ, val, zero, "")
}

func (v *Visitor) emitCondition(ctx *parser.ConditionContext) llvm.Value {
	vals := ctx.AllValue()
	if len(vals) == 1 {
		raw := v.Visit(vals[0])
		val, ok := raw.(llvm.Value)
		if !ok {
			return llvm.ConstInt(v.Llvmctx.Int1Type(), 0, false)
		}
		if val.Type().TypeKind() == llvm.IntegerTypeKind && val.Type().IntTypeWidth() != 1 {
			zero := llvm.ConstInt(val.Type(), 0, false)
			return v.Builder.CreateICmp(llvm.IntNE, val, zero, "")
		}
		return val
	}
	// value comparator value
	lhs, _ := v.Visit(vals[0]).(llvm.Value)
	rhs, _ := v.Visit(vals[1]).(llvm.Value)
	if lhs.IsNil() || rhs.IsNil() {
		return llvm.ConstInt(v.Llvmctx.Int1Type(), 0, false)
	}
	rhs = v.coerce(rhs, lhs.Type())
	cmpText := ctx.Comparator().GetText()
	isFloat := lhs.Type().TypeKind() == llvm.FloatTypeKind || lhs.Type().TypeKind() == llvm.DoubleTypeKind
	if isFloat {
		var pred llvm.FloatPredicate
		switch cmpText {
		case "==":
			pred = llvm.FloatOEQ
		case "!=":
			pred = llvm.FloatONE
		case "<":
			pred = llvm.FloatOLT
		case "<=":
			pred = llvm.FloatOLE
		case ">":
			pred = llvm.FloatOGT
		case ">=":
			pred = llvm.FloatOGE
		default:
			pred = llvm.FloatOEQ
		}
		return v.Builder.CreateFCmp(pred, lhs, rhs, "")
	}
	var pred llvm.IntPredicate
	switch cmpText {
	case "==":
		pred = llvm.IntEQ
	case "!=":
		pred = llvm.IntNE
	case "<":
		pred = llvm.IntSLT
	case "<=":
		pred = llvm.IntSLE
	case ">":
		pred = llvm.IntSGT
	case ">=":
		pred = llvm.IntSGE
	default:
		pred = llvm.IntEQ
	}
	return v.Builder.CreateICmp(pred, lhs, rhs, "")
}

func (v *Visitor) VisitIfStatement(ctx *parser.IfStatementContext) any {
	v.PrintNode(fmt.Sprintf("If: %s", ctx.Condition().GetText()))

	thisFunc := v.Results[v.lastFunc].(llvm.Value)
	thenBlock := llvm.AddBasicBlock(thisFunc, "if.then")
	mergeBlock := llvm.AddBasicBlock(thisFunc, "if.merge")

	elseifBlocks := ctx.AllElseifBlock()
	elseBlocks := ctx.AllElseBlock()
	hasElse := len(elseifBlocks) > 0 || len(elseBlocks) > 0

	var elseBlock llvm.BasicBlock
	if hasElse {
		elseBlock = llvm.AddBasicBlock(thisFunc, "if.else")
	} else {
		elseBlock = mergeBlock
	}

	cond := v.emitCondition(ctx.Condition().(*parser.ConditionContext))
	v.Builder.CreateCondBr(cond, thenBlock, elseBlock)

	// Then block
	v.Builder.SetInsertPointAtEnd(thenBlock)
	savedReturnEmitted := v.returnEmitted
	v.returnEmitted = false
	v.emitStatements(ctx.Body().(*parser.BodyContext).AllStatement())
	thenReturned := v.returnEmitted
	if !v.blockHasTerminator() {
		v.Builder.CreateBr(mergeBlock)
	}
	v.returnEmitted = savedReturnEmitted

	// Else / elseif chain
	if hasElse {
		v.Builder.SetInsertPointAtEnd(elseBlock)
		for _, elif := range elseifBlocks {
			v.Visit(elif)
		}
		for _, eb := range elseBlocks {
			v.Visit(eb)
		}
		if !v.blockHasTerminator() {
			v.Builder.CreateBr(mergeBlock)
		}
	}

	_ = thenReturned
	v.Builder.SetInsertPointAtEnd(mergeBlock)
	return nil
}

func (v *Visitor) VisitWhileDecl(ctx *parser.WhileDeclContext) any {
	v.PrintNode(fmt.Sprintf("While: %s", ctx.Condition().GetText()))

	thisFunc := v.Results[v.lastFunc].(llvm.Value)
	condBlock := llvm.AddBasicBlock(thisFunc, "while.cond")
	bodyBlock := llvm.AddBasicBlock(thisFunc, "while.body")
	exitBlock := llvm.AddBasicBlock(thisFunc, "while.exit")

	v.Builder.CreateBr(condBlock)

	v.Builder.SetInsertPointAtEnd(condBlock)
	cond := v.emitCondition(ctx.Condition().(*parser.ConditionContext))
	v.Builder.CreateCondBr(cond, bodyBlock, exitBlock)

	v.Builder.SetInsertPointAtEnd(bodyBlock)
	savedBreak := v.breakTarget
	savedContinue := v.continueTarget
	v.breakTarget = exitBlock
	v.continueTarget = condBlock
	v.emitStatements(ctx.Body().(*parser.BodyContext).AllStatement())
	if !v.blockHasTerminator() {
		v.Builder.CreateBr(condBlock)
	}
	v.breakTarget = savedBreak
	v.continueTarget = savedContinue

	v.Builder.SetInsertPointAtEnd(exitBlock)
	v.emitEscapeCheck(savedBreak)
	return nil
}

func (v *Visitor) VisitCondition(ctx *parser.ConditionContext) any {
	v.PrintNode(fmt.Sprintf("Condition: %s", ctx.GetText()))
	return v.emitCondition(ctx)
}

func (v *Visitor) VisitComparator(ctx *parser.ComparatorContext) any {
	v.PrintNode(fmt.Sprintf("Comparator: %s", ctx.GetText()))
	return ctx.GetText()
}

func (v *Visitor) VisitElseifBlock(ctx *parser.ElseifBlockContext) any {
	v.PrintNode(fmt.Sprintf("Elseif: %s", ctx.GetText()))
	return v.Visit(ctx.IfStatement())
}

func (v *Visitor) VisitElseBlock(ctx *parser.ElseBlockContext) any {
	v.PrintNode(fmt.Sprintf("Else: %s", ctx.GetText()))
	v.emitStatements(ctx.Body().(*parser.BodyContext).AllStatement())
	return nil
}

func (v *Visitor) VisitEachDecl(ctx *parser.EachDeclContext) any {
	cursorName := ctx.Identifier(0).GetText()

	var elems []llvm.Value
	if ctx.ArrayLiteral() != nil {
		al := ctx.ArrayLiteral().(*parser.ArrayLiteralContext)
		if am := al.ArrayMembers(); am != nil {
			for _, valCtx := range am.(*parser.ArrayMembersContext).AllValue() {
				if lv, ok := v.Visit(valCtx).(llvm.Value); ok {
					elems = append(elems, lv)
				}
			}
		}
	} else if ctx.MemberAccess() != nil {
		ma := ctx.MemberAccess().(*parser.MemberAccessContext)
		raw := v.evalMemberAccess(ma)
		slicePtr, ok := raw.(llvm.Value)
		if !ok {
			return nil
		}
		i64 := v.Llvmctx.Int64Type()
		elemType := v.Llvmctx.Int8Type()
		idents := ma.AllIdentifier()
		if len(idents) == 2 {
			baseName := idents[0].GetText()
			fieldName := idents[1].GetText()
			structName := v.varTypes[baseName]
			if structName == "self" {
				structName = v.selfType
			}
			if fields, ok := v.structFields[structName]; ok {
				for i, f := range fields {
					if f == fieldName {
						if typeStrs, ok2 := v.structFieldTypes[structName]; ok2 && i < len(typeStrs) {
							ts := typeStrs[i]
							if idx := strings.Index(ts, "["); idx != -1 {
								elemType = v.ResolveType(ts[:idx])
							}
						}
						break
					}
				}
			}
		}
		lenGEP := v.Builder.CreateGEP(i64, slicePtr, []llvm.Value{llvm.ConstInt(i64, 0, false)}, "")
		runtimeLen := v.Builder.CreateLoad(i64, lenGEP, "each.len")
		v.emitSliceEachLoop(ctx, cursorName, slicePtr, runtimeLen, elemType)
		return nil
	} else {
		srcName := ctx.Identifier(1).GetText()
		i64 := v.Llvmctx.Int64Type()
		isDynSlice := false
		var elemTypeStr string
		if tk, ok := v.varTypes[srcName]; ok {
			if idx := strings.Index(tk, "["); idx != -1 {
				closeIdx := strings.Index(tk[idx:], "]")
				if closeIdx > 0 && tk[idx+1:idx+closeIdx] == "" {
					isDynSlice = true
					elemTypeStr = tk[:idx]
				}
			}
		}
		if isDynSlice {
			if alloca, ok := v.vars[srcName]; ok {
				slicePtr := v.Builder.CreateLoad(alloca.AllocatedType(), alloca, "")
				var sliceElemType llvm.Type
				if elemTypeStr != "" {
					sliceElemType = v.ResolveType(elemTypeStr)
				} else {
					sliceElemType = llvm.PointerType(v.Llvmctx.Int8Type(), 0)
				}
				lenGEP := v.Builder.CreateGEP(i64, slicePtr, []llvm.Value{llvm.ConstInt(i64, 0, false)}, "")
				runtimeLen := v.Builder.CreateLoad(i64, lenGEP, "each.len")
				v.emitSliceEachLoop(ctx, cursorName, slicePtr, runtimeLen, sliceElemType)
			}
			return nil
		}
		if alloca, ok := v.vars[srcName]; ok {
			arrType := alloca.AllocatedType()
			if arrType.TypeKind() == llvm.ArrayTypeKind {
				n := arrType.ArrayLength()
				elemType := arrType.ElementType()
				for i := 0; i < n; i++ {
					gep := v.Builder.CreateInBoundsGEP(arrType, alloca,
						[]llvm.Value{
							llvm.ConstInt(i64, 0, false),
							llvm.ConstInt(i64, uint64(i), false),
						}, "")
					elems = append(elems, v.Builder.CreateLoad(elemType, gep, ""))
				}
			}
		}
	}

	src := ""
	if ctx.ArrayLiteral() != nil {
		src = ctx.ArrayLiteral().GetText()
	} else if ctx.Identifier(1) != nil {
		src = ctx.Identifier(1).GetText()
	}
	v.PrintNode(fmt.Sprintf("Each: %s in %s (%d elems)", cursorName, src, len(elems)))

	if len(elems) == 0 {
		return nil
	}

	elemType := elems[0].Type()

	cursorAlloca, hasCursor := v.vars[cursorName]
	if !hasCursor {
		cursorAlloca = v.Builder.CreateAlloca(elemType, cursorName)
		v.vars[cursorName] = cursorAlloca
	}

	thisFunc := v.Results[v.lastFunc].(llvm.Value)
	n := len(elems)

	// Counter alloca: i = 0; i < n; i++
	i64 := v.Llvmctx.Int64Type()
	counter := v.Builder.CreateAlloca(i64, "each.counter")
	v.Builder.CreateStore(llvm.ConstInt(i64, 0, false), counter)

	condBlock := llvm.AddBasicBlock(thisFunc, "each.cond")
	bodyBlock := llvm.AddBasicBlock(thisFunc, "each.body")
	incrBlock := llvm.AddBasicBlock(thisFunc, "each.incr")
	exitBlock := llvm.AddBasicBlock(thisFunc, "each.exit")

	v.Builder.CreateBr(condBlock)

	// Condition: i < n
	v.Builder.SetInsertPointAtEnd(condBlock)
	cur := v.Builder.CreateLoad(i64, counter, "")
	cmp := v.Builder.CreateICmp(llvm.IntSLT, cur,
		llvm.ConstInt(i64, uint64(n), false), "")
	v.Builder.CreateCondBr(cmp, bodyBlock, exitBlock)

	// Body: switch on counter to load the right element into cursor
	v.Builder.SetInsertPointAtEnd(bodyBlock)
	cur = v.Builder.CreateLoad(i64, counter, "")

	// Build a switch over the counter value; default falls through to incr
	afterSwitch := llvm.AddBasicBlock(thisFunc, "each.after_switch")
	sw := v.Builder.CreateSwitch(cur, afterSwitch, n)
	caseMerge := llvm.AddBasicBlock(thisFunc, "each.case_merge")

	for i, elem := range elems {
		caseBlock := llvm.AddBasicBlock(thisFunc, fmt.Sprintf("each.case%d", i))
		sw.AddCase(llvm.ConstInt(i64, uint64(i), false), caseBlock)
		v.Builder.SetInsertPointAtEnd(caseBlock)
		coerced := v.coerce(elem, elemType)
		v.Builder.CreateStore(coerced, cursorAlloca)
		v.Builder.CreateBr(caseMerge)
	}

	// after_switch (default) also goes to caseMerge
	v.Builder.SetInsertPointAtEnd(afterSwitch)
	v.Builder.CreateBr(caseMerge)

	// Emit body statements with proper break/continue targets
	v.Builder.SetInsertPointAtEnd(caseMerge)
	savedBreak := v.breakTarget
	savedContinue := v.continueTarget
	v.breakTarget = exitBlock
	v.continueTarget = incrBlock
	v.emitStatements(ctx.Body().(*parser.BodyContext).AllStatement())
	if !v.blockHasTerminator() {
		v.Builder.CreateBr(incrBlock)
	}
	v.breakTarget = savedBreak
	v.continueTarget = savedContinue

	// Increment
	v.Builder.SetInsertPointAtEnd(incrBlock)
	next := v.Builder.CreateAdd(v.Builder.CreateLoad(i64, counter, ""),
		llvm.ConstInt(i64, 1, false), "")
	v.Builder.CreateStore(next, counter)
	v.Builder.CreateBr(condBlock)

	v.Builder.SetInsertPointAtEnd(exitBlock)
	v.emitEscapeCheck(savedBreak)

	if !hasCursor {
		delete(v.vars, cursorName)
	}
	return nil
}

func (v *Visitor) emitSliceEachLoop(ctx *parser.EachDeclContext, cursorName string, slicePtr llvm.Value, runtimeLen llvm.Value, elemType llvm.Type) {
	i64 := v.Llvmctx.Int64Type()
	thisFunc := v.Results[v.lastFunc].(llvm.Value)

	cursorAlloca, hasCursor := v.vars[cursorName]
	if !hasCursor {
		cursorAlloca = v.Builder.CreateAlloca(elemType, cursorName)
		v.vars[cursorName] = cursorAlloca
	}

	counter := v.Builder.CreateAlloca(i64, "each.counter")
	v.Builder.CreateStore(llvm.ConstInt(i64, 0, false), counter)

	condBlock := llvm.AddBasicBlock(thisFunc, "each.cond")
	bodyBlock := llvm.AddBasicBlock(thisFunc, "each.body")
	incrBlock := llvm.AddBasicBlock(thisFunc, "each.incr")
	exitBlock := llvm.AddBasicBlock(thisFunc, "each.exit")

	v.Builder.CreateBr(condBlock)

	v.Builder.SetInsertPointAtEnd(condBlock)
	cur := v.Builder.CreateLoad(i64, counter, "")
	cmp := v.Builder.CreateICmp(llvm.IntSLT, cur, runtimeLen, "")
	v.Builder.CreateCondBr(cmp, bodyBlock, exitBlock)

	v.Builder.SetInsertPointAtEnd(bodyBlock)
	cur = v.Builder.CreateLoad(i64, counter, "")
	// +1 to skip the length field at index 0
	elemIdx := v.Builder.CreateAdd(cur, llvm.ConstInt(i64, 1, false), "")
	elemGEP := v.Builder.CreateGEP(i64, slicePtr, []llvm.Value{elemIdx}, "")
	raw := v.Builder.CreateLoad(i64, elemGEP, "")
	coerced := v.coerce(raw, elemType)
	v.Builder.CreateStore(coerced, cursorAlloca)

	savedBreak := v.breakTarget
	savedContinue := v.continueTarget
	v.breakTarget = exitBlock
	v.continueTarget = incrBlock
	v.emitStatements(ctx.Body().(*parser.BodyContext).AllStatement())
	if !v.blockHasTerminator() {
		v.Builder.CreateBr(incrBlock)
	}
	v.breakTarget = savedBreak
	v.continueTarget = savedContinue

	v.Builder.SetInsertPointAtEnd(incrBlock)
	next := v.Builder.CreateAdd(v.Builder.CreateLoad(i64, counter, ""), llvm.ConstInt(i64, 1, false), "")
	v.Builder.CreateStore(next, counter)
	v.Builder.CreateBr(condBlock)

	v.Builder.SetInsertPointAtEnd(exitBlock)
	v.emitEscapeCheck(savedBreak)

	if !hasCursor {
		delete(v.vars, cursorName)
	}
}

func (v *Visitor) VisitArrayLiteral(ctx *parser.ArrayLiteralContext) any {
	v.PrintNode(fmt.Sprintf("Array: %s", ctx.GetText()))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitArrayMembers(ctx *parser.ArrayMembersContext) any {
	return v.VisitChildren(ctx)
}

func (v *Visitor) emitEscapeCheck(outerBreak llvm.BasicBlock) {
	if v.escapeFlag.IsNil() {
		return
	}
	i1 := v.Llvmctx.Int1Type()
	flagVal := v.Builder.CreateLoad(i1, v.escapeFlag, "")
	thisFunc := v.Results[v.lastFunc].(llvm.Value)
	if !outerBreak.IsNil() {
		continueBlock := llvm.AddBasicBlock(thisFunc, "escape.continue")
		v.Builder.CreateCondBr(flagVal, outerBreak, continueBlock)
		v.Builder.SetInsertPointAtEnd(continueBlock)
	} else {
		resetBlock := llvm.AddBasicBlock(thisFunc, "escape.reset")
		continueBlock := llvm.AddBasicBlock(thisFunc, "escape.continue")
		v.Builder.CreateCondBr(flagVal, resetBlock, continueBlock)
		v.Builder.SetInsertPointAtEnd(resetBlock)
		v.Builder.CreateStore(llvm.ConstInt(i1, 0, false), v.escapeFlag)
		v.Builder.CreateBr(continueBlock)
		v.Builder.SetInsertPointAtEnd(continueBlock)
	}
}

func (v *Visitor) getOrDeclareStrcmp() (llvm.Value, llvm.Type) {
	i8ptr := llvm.PointerType(v.Llvmctx.Int8Type(), 0)
	fnType := llvm.FunctionType(v.Llvmctx.Int32Type(), []llvm.Type{i8ptr, i8ptr}, false)
	if fn := v.Llvmmod.NamedFunction("strcmp"); !fn.IsNil() {
		return fn, fnType
	}
	return llvm.AddFunction(v.Llvmmod, "strcmp", fnType), fnType
}

func (v *Visitor) VisitSwitchStatement(ctx *parser.SwitchStatementContext) any {
	v.PrintNode(fmt.Sprintf("Switch: %s", ctx.Value().GetText()))

	switchVal, ok := v.Visit(ctx.Value()).(llvm.Value)
	if !ok {
		return nil
	}

	thisFunc := v.Results[v.lastFunc].(llvm.Value)
	mergeBlock := llvm.AddBasicBlock(thisFunc, "switch.merge")
	cases := ctx.SwitchBody().(*parser.SwitchBodyContext).AllCases()

	savedBreak := v.breakTarget
	v.breakTarget = mergeBlock

	isString := switchVal.Type().TypeKind() == llvm.PointerTypeKind

	if !isString {
		// Integer switch: find the default case first (CreateSwitch needs elseb upfront)
		defaultIntBlock := mergeBlock
		var defaultIntCtx *parser.CasesContext
		for _, c := range cases {
			caseCtx := c.(*parser.CasesContext)
			if caseCtx.DEFAULT() != nil {
				defaultIntBlock = llvm.AddBasicBlock(thisFunc, "switch.default")
				defaultIntCtx = caseCtx
				break
			}
		}

		sw := v.Builder.CreateSwitch(switchVal, defaultIntBlock, len(cases))
		for i, c := range cases {
			caseCtx := c.(*parser.CasesContext)
			if caseCtx.DEFAULT() != nil {
				continue
			}
			caseVal, ok := v.Visit(caseCtx.Value()).(llvm.Value)
			if !ok {
				continue
			}
			if caseVal.Type() != switchVal.Type() {
				caseVal = llvm.ConstTruncOrBitCast(caseVal, switchVal.Type())
			}
			caseBlock := llvm.AddBasicBlock(thisFunc, fmt.Sprintf("switch.case%d", i))
			sw.AddCase(caseVal, caseBlock)
			v.Builder.SetInsertPointAtEnd(caseBlock)
			v.emitStatements(caseCtx.AllStatement())
			if !v.blockHasTerminator() {
				v.Builder.CreateBr(mergeBlock)
			}
		}
		if defaultIntCtx != nil {
			v.Builder.SetInsertPointAtEnd(defaultIntBlock)
			v.emitStatements(defaultIntCtx.AllStatement())
			if !v.blockHasTerminator() {
				v.Builder.CreateBr(mergeBlock)
			}
		}
	} else {
		// String switch: chain of strcmp comparisons
		strcmp, strcmpType := v.getOrDeclareStrcmp()
		zero32 := llvm.ConstInt(v.Llvmctx.Int32Type(), 0, false)

		var nonDefaultCases []*parser.CasesContext
		var defaultCaseCtx *parser.CasesContext
		for _, c := range cases {
			caseCtx := c.(*parser.CasesContext)
			if caseCtx.DEFAULT() != nil {
				defaultCaseCtx = caseCtx
			} else {
				nonDefaultCases = append(nonDefaultCases, caseCtx)
			}
		}

		bodyBlocks := make([]llvm.BasicBlock, len(nonDefaultCases))
		for i := range nonDefaultCases {
			bodyBlocks[i] = llvm.AddBasicBlock(thisFunc, fmt.Sprintf("switch.case%d", i))
		}

		var defaultBlock llvm.BasicBlock
		if defaultCaseCtx != nil {
			defaultBlock = llvm.AddBasicBlock(thisFunc, "switch.default")
		} else {
			defaultBlock = mergeBlock
		}

		for i, caseCtx := range nonDefaultCases {
			caseVal, ok := v.Visit(caseCtx.Value()).(llvm.Value)
			if !ok {
				v.Builder.CreateBr(defaultBlock)
				break
			}
			result := v.Builder.CreateCall(strcmpType, strcmp, []llvm.Value{switchVal, caseVal}, "")
			cmp := v.Builder.CreateICmp(llvm.IntEQ, result, zero32, "")

			var missBlock llvm.BasicBlock
			if i+1 < len(nonDefaultCases) {
				missBlock = llvm.AddBasicBlock(thisFunc, fmt.Sprintf("switch.check%d", i+1))
			} else {
				missBlock = defaultBlock
			}
			v.Builder.CreateCondBr(cmp, bodyBlocks[i], missBlock)

			v.Builder.SetInsertPointAtEnd(bodyBlocks[i])
			v.emitStatements(caseCtx.AllStatement())
			if !v.blockHasTerminator() {
				v.Builder.CreateBr(mergeBlock)
			}

			if i+1 < len(nonDefaultCases) {
				v.Builder.SetInsertPointAtEnd(missBlock)
			}
		}

		if len(nonDefaultCases) == 0 {
			v.Builder.CreateBr(defaultBlock)
		}

		if defaultCaseCtx != nil {
			v.Builder.SetInsertPointAtEnd(defaultBlock)
			v.emitStatements(defaultCaseCtx.AllStatement())
			if !v.blockHasTerminator() {
				v.Builder.CreateBr(mergeBlock)
			}
		}
	}

	v.breakTarget = savedBreak
	v.Builder.SetInsertPointAtEnd(mergeBlock)
	return nil
}

func (v *Visitor) VisitBreakStatement(ctx *parser.BreakStatementContext) any {
	v.PrintNode(fmt.Sprintf("Loop break: %s", ctx.GetText()))
	text := ctx.GetText()
	if text == "continue" && !v.continueTarget.IsNil() {
		v.Builder.CreateBr(v.continueTarget)
	} else if text == "escape" {
		if !v.escapeFlag.IsNil() {
			i1 := v.Llvmctx.Int1Type()
			v.Builder.CreateStore(llvm.ConstInt(i1, 1, false), v.escapeFlag)
		}
		if !v.breakTarget.IsNil() {
			v.Builder.CreateBr(v.breakTarget)
		}
	} else if !v.breakTarget.IsNil() {
		v.Builder.CreateBr(v.breakTarget)
	}
	return nil
}

func (v *Visitor) VisitTerminal(ctx antlr.TerminalNode) any {
	return ctx.GetText()
}

func (v *Visitor) VisitErrorNode(_ antlr.ErrorNode) interface{} { return nil }

func (v *Visitor) VisitChildren(tree antlr.RuleNode) any {
	n := tree.GetChildCount()
	externalDeclarations := make([]any, 1024)
	v.depth++
	for i := 0; i < n; i++ {
		c := tree.GetChild(i)
		val := c.(antlr.ParseTree)
		externalDeclarations[i] = v.Visit(val)
	}
	v.depth--
	return externalDeclarations
}

func collectBodyIdentifiers(tree antlr.ParseTree) map[string]bool {
	result := make(map[string]bool)
	var walk func(antlr.ParseTree)
	walk = func(t antlr.ParseTree) {
		for i := 0; i < t.GetChildCount(); i++ {
			child := t.GetChild(i)
			if term, ok := child.(antlr.TerminalNode); ok {
				if term.GetSymbol().GetTokenType() == parser.BridgeLexerIdentifier {
					result[term.GetText()] = true
				}
			} else if pt, ok := child.(antlr.ParseTree); ok {
				walk(pt)
			}
		}
	}
	walk(tree)
	return result
}

func (v *Visitor) getOrDeclareMalloc() llvm.Value {
	if fn := v.Llvmmod.NamedFunction("malloc"); !fn.IsNil() {
		return fn
	}
	i8ptr := llvm.PointerType(v.Llvmctx.Int8Type(), 0)
	fnType := llvm.FunctionType(i8ptr, []llvm.Type{v.Llvmctx.Int64Type()}, false)
	return llvm.AddFunction(v.Llvmmod, "malloc", fnType)
}

func (v *Visitor) compileClosure(ctx *parser.FuncDeclContext, captureNames []string, innerFunctype llvm.Type) llvm.Value {
	i8ptr := llvm.PointerType(v.Llvmctx.Int8Type(), 0)
	i64 := v.Llvmctx.Int64Type()
	malloc := v.getOrDeclareMalloc()
	mallocType := llvm.FunctionType(i8ptr, []llvm.Type{i64}, false)

	heapBoxes := make(map[string]llvm.Value)
	for _, name := range captureNames {
		alloca, ok := v.vars[name]
		if !ok {
			continue
		}
		elemType := alloca.AllocatedType()
		heapPtr := v.Builder.CreateCall(mallocType, malloc, []llvm.Value{llvm.ConstInt(i64, 8, false)}, "box."+name)
		typedPtr := v.Builder.CreateBitCast(heapPtr, llvm.PointerType(elemType, 0), "")
		curVal := v.Builder.CreateLoad(elemType, alloca, "")
		v.Builder.CreateStore(curVal, typedPtr)
		heapBoxes[name] = heapPtr
	}

	n := len(captureNames)
	envPtr := v.Builder.CreateCall(mallocType, malloc, []llvm.Value{llvm.ConstInt(i64, uint64(8*n), false)}, "env")
	envI64Ptr := v.Builder.CreateBitCast(envPtr, llvm.PointerType(i64, 0), "env.i64")
	for idx, name := range captureNames {
		slot := v.Builder.CreateInBoundsGEP(i64, envI64Ptr, []llvm.Value{llvm.ConstInt(i64, uint64(idx), false)}, "")
		boxI64 := v.Builder.CreatePtrToInt(heapBoxes[name], i64, "")
		v.Builder.CreateStore(boxI64, slot)
	}

	var argTypes []llvm.Type
	if ctx.Arguments() != nil {
		for _, a := range ctx.Arguments().GetChildren() {
			arg, ok := a.(*parser.ArgumentContext)
			if !ok {
				continue
			}
			argTypes = append(argTypes, v.ResolveType(arg.TypeKey().GetText()))
		}
	}
	retType := innerFunctype.ReturnType()
	closureArgTypes := append([]llvm.Type{i8ptr}, argTypes...)
	closureFunctype := llvm.FunctionType(retType, closureArgTypes, false)

	uid := strings.ReplaceAll(uuid.New().String(), "-", "")
	innerFnName := "__closure_" + uid
	innerFn := llvm.AddFunction(v.Llvmmod, innerFnName, closureFunctype)
	innerFn.Param(0).SetName("__env")
	v.pendingClosureInnerFn = innerFn
	if ctx.Arguments() != nil {
		i := 1
		for _, a := range ctx.Arguments().GetChildren() {
			arg, ok := a.(*parser.ArgumentContext)
			if !ok {
				continue
			}
			innerFn.Param(i).SetName(arg.Identifier().GetText())
			i++
		}
	}

	closurePtr := v.Builder.CreateCall(mallocType, malloc, []llvm.Value{llvm.ConstInt(i64, 16, false)}, "closure")
	closureI64Ptr := v.Builder.CreateBitCast(closurePtr, llvm.PointerType(i64, 0), "")
	fnPtrI64 := v.Builder.CreatePtrToInt(innerFn, i64, "")
	v.Builder.CreateStore(fnPtrI64, closureI64Ptr)
	envSlot := v.Builder.CreateInBoundsGEP(i64, closureI64Ptr, []llvm.Value{llvm.ConstInt(i64, 1, false)}, "")
	envI64 := v.Builder.CreatePtrToInt(envPtr, i64, "")
	v.Builder.CreateStore(envI64, envSlot)

	v.pendingClosureFuncType = closureFunctype

	h := sha256.Sum256([]byte(innerFnName))
	localid := fmt.Sprintf("%x", h[:])
	v.Results[localid] = innerFn

	savedLastFunc := v.lastFunc
	savedLastRetType := v.lastRetType
	savedReturnEmitted := v.returnEmitted
	savedVars := v.vars
	savedVarTypes := v.varTypes
	savedSymlinkPointeeTypes := v.symlinkPointeeTypes
	savedSelfType := v.selfType
	savedBreakTarget := v.breakTarget
	savedContinueTarget := v.continueTarget
	savedClosureCaptures := v.closureCaptures
	savedClosureVarFuncTypes := v.closureVarFuncTypes
	savedEscapeFlag := v.escapeFlag
	savedInsertBlock := v.Builder.GetInsertBlock()

	v.lastFunc = localid
	v.lastRetType = retType
	v.returnEmitted = false
	v.vars = make(map[string]llvm.Value)
	v.varTypes = make(map[string]string)
	v.symlinkPointeeTypes = make(map[string]llvm.Type)
	v.closureCaptures = make(map[string]llvm.Value)
	v.closureVarFuncTypes = make(map[string]llvm.Type)

	entry := llvm.AddBasicBlock(innerFn, "entry")
	v.Builder.SetInsertPointAtEnd(entry)

	i1 := v.Llvmctx.Int1Type()
	v.escapeFlag = v.Builder.CreateAlloca(i1, "escape.flag")
	v.Builder.CreateStore(llvm.ConstInt(i1, 0, false), v.escapeFlag)

	envParam := innerFn.Param(0)
	envParamI64Ptr := v.Builder.CreateBitCast(envParam, llvm.PointerType(i64, 0), "env.i64")
	for idx, name := range captureNames {
		slot := v.Builder.CreateInBoundsGEP(i64, envParamI64Ptr, []llvm.Value{llvm.ConstInt(i64, uint64(idx), false)}, "")
		boxI64 := v.Builder.CreateLoad(i64, slot, "cap.box."+name)
		boxPtr := v.Builder.CreateIntToPtr(boxI64, llvm.PointerType(i64, 0), "cap."+name)
		v.closureCaptures[name] = boxPtr
	}

	if ctx.Arguments() != nil {
		i := 1
		for _, a := range ctx.Arguments().GetChildren() {
			arg, ok := a.(*parser.ArgumentContext)
			if !ok {
				continue
			}
			name := arg.Identifier().GetText()
			paramType := v.ResolveType(arg.TypeKey().GetText())
			alloca := v.Builder.CreateAlloca(paramType, name)
			v.Builder.CreateStore(innerFn.Param(i), alloca)
			v.vars[name] = alloca
			v.varTypes[name] = arg.TypeKey().GetText()
			i++
		}
	}

	if ctx.Body() != nil {
		bodyCtx := ctx.Body().(*parser.BodyContext)
		v.emitStatements(bodyCtx.AllStatement())
	}
	if !v.returnEmitted && !v.blockHasTerminator() {
		if retType == v.Llvmctx.VoidType() {
			v.Builder.CreateRetVoid()
		} else {
			v.Builder.CreateRet(llvm.Undef(retType))
		}
	}

	v.lastFunc = savedLastFunc
	v.lastRetType = savedLastRetType
	v.returnEmitted = savedReturnEmitted
	v.vars = savedVars
	v.varTypes = savedVarTypes
	v.symlinkPointeeTypes = savedSymlinkPointeeTypes
	v.selfType = savedSelfType
	v.breakTarget = savedBreakTarget
	v.continueTarget = savedContinueTarget
	v.closureCaptures = savedClosureCaptures
	v.closureVarFuncTypes = savedClosureVarFuncTypes
	v.escapeFlag = savedEscapeFlag

	if savedInsertBlock.IsNil() {
		v.Builder.ClearInsertionPoint()
	} else {
		v.Builder.SetInsertPointAtEnd(savedInsertBlock)
	}

	return closurePtr
}

func (v *Visitor) emitClosureCall(varName string, alloca llvm.Value, ctx *parser.FuncCallContext) llvm.Value {
	i64 := v.Llvmctx.Int64Type()
	i8ptr := llvm.PointerType(v.Llvmctx.Int8Type(), 0)

	closurePtr := v.Builder.CreateLoad(alloca.AllocatedType(), alloca, varName)
	closureI64Ptr := v.Builder.CreateBitCast(closurePtr, llvm.PointerType(i64, 0), "")

	fnPtrI64 := v.Builder.CreateLoad(i64, closureI64Ptr, "fn.i64")
	envSlot := v.Builder.CreateInBoundsGEP(i64, closureI64Ptr, []llvm.Value{llvm.ConstInt(i64, 1, false)}, "")
	envI64 := v.Builder.CreateLoad(i64, envSlot, "env.i64")
	envPtr := v.Builder.CreateIntToPtr(envI64, i8ptr, "env.ptr")

	closureFunctype, hasFuncType := v.closureVarFuncTypes[varName]
	if !hasFuncType {
		closureFunctype = llvm.FunctionType(v.Llvmctx.VoidType(), []llvm.Type{i8ptr}, false)
	}

	fnPtrType := llvm.PointerType(closureFunctype, 0)
	fnPtr := v.Builder.CreateIntToPtr(fnPtrI64, fnPtrType, "fn.ptr")

	args := []llvm.Value{envPtr}
	if ctx.CallArgs() != nil {
		paramTypes := closureFunctype.ParamTypes()
		for i, valCtx := range ctx.CallArgs().(*parser.CallArgsContext).AllValue() {
			raw := v.Visit(valCtx)
			if lv, ok := raw.(llvm.Value); ok {
				paramIdx := i + 1 // +1 for env
				if paramIdx < len(paramTypes) {
					lv = v.coerce(lv, paramTypes[paramIdx])
				}
				args = append(args, lv)
			}
		}
	}

	return v.Builder.CreateCall(closureFunctype, fnPtr, args, "")
}
