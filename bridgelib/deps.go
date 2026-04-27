package bridgelib

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	wasmtime "github.com/bytecodealliance/wasmtime-go/v28"
)

type LibDep struct {
	Name    string
	Version string
	Wasm    []byte
}

func LoadDependencies(importsTable string) ([]LibDep, error) {
	if importsTable == "" {
		return nil, nil
	}

	var table map[string]string
	if err := json.Unmarshal([]byte(importsTable), &table); err != nil {
		return nil, fmt.Errorf("bridgelib: parse imports table: %w", err)
	}

	if len(table) == 0 {
		return nil, nil
	}

	bridgePath := os.Getenv("BRIDGEPATH")
	if bridgePath == "" {
		bridgePath = os.Getenv("HOME") + "/.bridge"
	}

	var deps []LibDep
	for name, version := range table {
		if name == "" || version == "" {
			continue
		}
		bbPath := filepath.Join(bridgePath, "packages", name, version, "module.bb")
		raw, err := os.ReadFile(bbPath)
		if err != nil {
			return nil, fmt.Errorf("bridgelib: load package %q v%s: %w", name, version, err)
		}
		bbf, err := Deserialize(raw)
		if err != nil {
			return nil, fmt.Errorf("bridgelib: deserialize package %q v%s: %w", name, version, err)
		}
		wasm, err := bbf.InflateBytecode()
		if err != nil {
			return nil, fmt.Errorf("bridgelib: inflate package %q v%s: %w", name, version, err)
		}
		deps = append(deps, LibDep{Name: name, Version: version, Wasm: wasm})
	}
	return deps, nil
}

func registerLibDeps(engine *wasmtime.Engine, store *wasmtime.Store, linker *wasmtime.Linker, sm *sharedMem, deps []LibDep) error {
	for _, dep := range deps {
		libMod, err := wasmtime.NewModule(engine, dep.Wasm)
		if err != nil {
			return fmt.Errorf("package %q: compile: %w", dep.Name, err)
		}

		libLinker := wasmtime.NewLinker(engine)
		if err := libLinker.DefineWasi(); err != nil {
			return fmt.Errorf("package %q: linker wasi: %w", dep.Name, err)
		}
		// Share the same memory, stack pointer, malloc/free, and stdlib as the main module
		for _, itemName := range []string{"__linear_memory", "__stack_pointer", "malloc", "free"} {
			ext := linker.Get(store, "env", itemName)
			if ext == nil {
				continue
			}
			if err := libLinker.Define(store, "env", itemName, ext); err != nil {
				return fmt.Errorf("package %q: linker define %q: %w", dep.Name, itemName, err)
			}
		}
		if err := registerStdlib(libLinker, sm); err != nil {
			return fmt.Errorf("package %q: register stdlib: %w", dep.Name, err)
		}

		for _, imp := range libMod.Imports() {
			if imp.Module() == "env" && imp.Name() != nil && *imp.Name() == "__indirect_function_table" {
				if shared := linker.Get(store, "env", "__indirect_function_table"); shared != nil {
					if err := libLinker.Define(store, "env", "__indirect_function_table", shared); err != nil {
						return fmt.Errorf("package %q: define shared indirect table: %w", dep.Name, err)
					}
				} else {
					tableType := imp.Type().TableType()
					table, err := wasmtime.NewTable(store, tableType, wasmtime.ValFuncref(nil))
					if err != nil {
						return fmt.Errorf("package %q: indirect table: %w", dep.Name, err)
					}
					const minTableSize = 1024
					if cur := table.Size(store); cur < minTableSize {
						if _, err := table.Grow(store, minTableSize-cur, wasmtime.ValFuncref(nil)); err != nil {
							return fmt.Errorf("package %q: grow indirect table: %w", dep.Name, err)
						}
					}
					if err := libLinker.Define(store, "env", "__indirect_function_table", table); err != nil {
						return fmt.Errorf("package %q: define indirect table: %w", dep.Name, err)
					}
				}
			}
		}

		libInst, err := libLinker.Instantiate(store, libMod)
		if err != nil {
			return fmt.Errorf("package %q: instantiate: %w", dep.Name, err)
		}

		// Register as "env.{LibName}::{ExportName}"
		libName := dep.Name
		for _, exp := range libMod.Exports() {
			if exp.Type().FuncType() == nil {
				continue
			}
			exportName := exp.Name()
			fn := libInst.GetFunc(store, exportName)
			if fn == nil {
				continue
			}
			qualName := libName + "::" + exportName
			if err := linker.Define(store, "env", qualName, fn); err != nil {
				return fmt.Errorf("package %q: register export %q: %w", dep.Name, exportName, err)
			}
		}
	}
	return nil
}
