package bridgelib

import (
	"context"
	"fmt"
	"sync"

	wasmtime "github.com/bytecodealliance/wasmtime-go/v28"
)

type BridgeContext struct {
	MemoryCeiling int64
	context.Context
}

type Instance struct {
	store    *wasmtime.Store
	instance *wasmtime.Instance
	memory   *wasmtime.Memory
	mu       sync.Mutex
	ctx      BridgeContext
}

type Result struct {
	Values []interface{}
	Err    error
}

func New(wasmBytes []byte, ctx BridgeContext, deps []LibDep) (*Instance, error) {
	cfg := wasmtime.NewConfig()
	cfg.SetWasmMemory64(true)

	engine := wasmtime.NewEngineWithConfig(cfg)
	store := wasmtime.NewStore(engine)

	module, err := wasmtime.NewModule(engine, wasmBytes)
	if err != nil {
		return nil, fmt.Errorf("bridgelib: jit: %w", err)
	}

	const initialPages = 16
	memType := wasmtime.NewMemoryType64(initialPages, (ctx.MemoryCeiling != 0), uint64(ctx.MemoryCeiling), false)
	mem, err := wasmtime.NewMemory(store, memType)
	if err != nil {
		return nil, fmt.Errorf("bridgelib: memory: %w", err)
	}

	stackPtr, err := wasmtime.NewGlobal(store, wasmtime.NewGlobalType(wasmtime.NewValType(wasmtime.KindI64), true), wasmtime.ValI64(initialPages*65536))
	if err != nil {
		return nil, fmt.Errorf("bridgelib: stack pointer global: %w", err)
	}

	sm := &sharedMem{store: store}

	linker := wasmtime.NewLinker(engine)
	if err := linker.DefineWasi(); err != nil {
		return nil, fmt.Errorf("bridgelib: linker: %w", err)
	}
	if err := linker.Define(store, "env", "__linear_memory", mem); err != nil {
		return nil, fmt.Errorf("bridgelib: linker define memory: %w", err)
	}
	if err := linker.Define(store, "env", "__stack_pointer", stackPtr); err != nil {
		return nil, fmt.Errorf("bridgelib: linker define stack pointer: %w", err)
	}

	alloc := &vmAllocator{
		cursor: uint64(initialPages) * 65536 * 2,
		sizes:  make(map[uint64]uint64),
		sm:     sm,
	}
	sm.alloc = alloc

	mallocType := wasmtime.NewFuncType(
		[]*wasmtime.ValType{wasmtime.NewValType(wasmtime.KindI64)},
		[]*wasmtime.ValType{wasmtime.NewValType(wasmtime.KindI64)},
	)
	if err := linker.FuncNew("env", "malloc", mallocType,
		func(caller *wasmtime.Caller, args []wasmtime.Val) ([]wasmtime.Val, *wasmtime.Trap) {
			ptr := alloc.Alloc(uint64(args[0].I64()))
			return []wasmtime.Val{wasmtime.ValI64(int64(ptr))}, nil
		},
	); err != nil {
		return nil, fmt.Errorf("bridgelib: linker define malloc: %w", err)
	}
	freeType := wasmtime.NewFuncType(
		[]*wasmtime.ValType{wasmtime.NewValType(wasmtime.KindI64)},
		nil,
	)
	if err := linker.FuncNew("env", "free", freeType,
		func(caller *wasmtime.Caller, args []wasmtime.Val) ([]wasmtime.Val, *wasmtime.Trap) {
			alloc.Free(uint64(args[0].I64()))
			return nil, nil
		},
	); err != nil {
		return nil, fmt.Errorf("bridgelib: linker define free: %w", err)
	}
	strcmpType := wasmtime.NewFuncType(
		[]*wasmtime.ValType{wasmtime.NewValType(wasmtime.KindI64), wasmtime.NewValType(wasmtime.KindI64)},
		[]*wasmtime.ValType{wasmtime.NewValType(wasmtime.KindI32)},
	)
	if err := linker.FuncNew("env", "strcmp", strcmpType,
		func(caller *wasmtime.Caller, args []wasmtime.Val) ([]wasmtime.Val, *wasmtime.Trap) {
			mem := sm.mem.UnsafeData(store)
			a, b := uint64(args[0].I64()), uint64(args[1].I64())
			var result int32
			for {
				ca, cb := mem[a], mem[b]
				if ca != cb {
					result = int32(ca) - int32(cb)
					break
				}
				if ca == 0 {
					break
				}
				a++
				b++
			}
			return []wasmtime.Val{wasmtime.ValI32(result)}, nil
		},
	); err != nil {
		return nil, fmt.Errorf("bridgelib: linker define strcmp: %w", err)
	}

	for _, imp := range module.Imports() {
		if imp.Module() == "env" && *imp.Name() == "__indirect_function_table" {
			tableType := imp.Type().TableType()

			table, err := wasmtime.NewTable(store, tableType, wasmtime.ValFuncref(nil))
			if err != nil {
				return nil, fmt.Errorf("bridgelib: indirect function table: %w", err)
			}

			const minTableSize = 1024
			if cur := table.Size(store); cur < minTableSize {
				if _, err := table.Grow(store, minTableSize-cur, wasmtime.ValFuncref(nil)); err != nil {
					return nil, fmt.Errorf("bridgelib: grow indirect function table: %w", err)
				}
			}

			if err := linker.Define(store, "env", "__indirect_function_table", table); err != nil {
				return nil, fmt.Errorf("bridgelib: linker define indirect function table: %w", err)
			}
		}
	}

	if err := registerStdlib(linker, sm); err != nil {
		return nil, fmt.Errorf("bridgelib: register stdlib: %w", err)
	}

	if err := registerLibDeps(engine, store, linker, sm, deps); err != nil {
		return nil, fmt.Errorf("bridgelib: register lib deps: %w", err)
	}

	inst, err := linker.Instantiate(store, module)
	if err != nil {
		return nil, fmt.Errorf("bridgelib: instantiate: %w", err)
	}

	if exp := inst.GetExport(store, "memory"); exp != nil {
		if expMem := exp.Memory(); expMem != nil {
			mem = expMem
		}
	}
	sm.mem = mem

	return &Instance{
		store:    store,
		instance: inst,
		memory:   mem,
		ctx:      ctx,
	}, nil
}

func (ins *Instance) Close() {
	ins.store.Close()
}

func (ins *Instance) ReadMemory(addr, length uint32) ([]byte, error) {
	if ins.memory == nil {
		return nil, fmt.Errorf("bridgelib: module has no exported memory")
	}
	data := ins.memory.UnsafeData(ins.store)
	end := uint64(addr) + uint64(length)
	if end > uint64(len(data)) {
		return nil, fmt.Errorf("bridgelib: read out of bounds (addr=%d len=%d memsize=%d)", addr, length, len(data))
	}
	out := make([]byte, length)
	copy(out, data[addr:end])
	return out, nil
}

func (ins *Instance) WriteMemory(addr uint32, p []byte) error {
	if ins.memory == nil {
		return fmt.Errorf("bridgelib: module has no exported memory")
	}
	data := ins.memory.UnsafeData(ins.store)
	end := uint64(addr) + uint64(len(p))
	if end > uint64(len(data)) {
		return fmt.Errorf("bridgelib: write out of bounds (addr=%d len=%d memsize=%d)", addr, len(p), len(data))
	}
	copy(data[addr:end], p)
	return nil
}

func (ins *Instance) MemorySize() (uint32, error) {
	if ins.memory == nil {
		return 0, fmt.Errorf("bridgelib: module has no exported memory")
	}
	pages := ins.memory.Size(ins.store)
	return uint32(pages), nil
}

func (ins *Instance) Call(name string, args ...interface{}) ([]interface{}, error) {
	ins.mu.Lock()
	defer ins.mu.Unlock()
	return ins.call(name, args...)
}

func (ins *Instance) CallAsync(ctx context.Context, name string, args ...interface{}) <-chan Result {
	ch := make(chan Result, 1)
	go func() {
		if err := ctx.Err(); err != nil {
			ch <- Result{Err: err}
			return
		}
		ins.mu.Lock()
		defer ins.mu.Unlock()
		vals, err := ins.call(name, args...)
		ch <- Result{Values: vals, Err: err}
	}()
	return ch
}

func (ins *Instance) call(name string, args ...interface{}) ([]interface{}, error) {
	fn := ins.instance.GetFunc(ins.store, name)
	if fn == nil {
		return nil, fmt.Errorf("bridgelib: function %q not found", name)
	}
	result, err := fn.Call(ins.store, args...)
	if err != nil {
		return nil, fmt.Errorf("bridgelib: call %q: %w", name, err)
	}
	if result == nil {
		return nil, nil
	}
	return []interface{}{result}, nil
}

func (ins *Instance) FuncParamCount(name string) (int, error) {
	fn := ins.instance.GetFunc(ins.store, name)
	if fn == nil {
		return 0, fmt.Errorf("bridgelib: function %q not found", name)
	}
	return len(fn.Type(ins.store).Params()), nil
}

func (ins *Instance) WriteStrings(strs []string) (arrayPtr int64, count int64, err error) {
	if ins.memory == nil {
		return 0, 0, fmt.Errorf("bridgelib: module has no exported memory")
	}
	data := ins.memory.UnsafeData(ins.store)
	memSize := uint64(len(data))

	cursor := memSize

	ptrs := make([]uint64, len(strs))
	for i, s := range strs {
		b := append([]byte(s), 0)
		need := cursor + uint64(len(b))
		if need > memSize {
			pages := (need-memSize)/65536 + 1
			if _, growErr := ins.memory.Grow(ins.store, pages); growErr != nil {
				return 0, 0, fmt.Errorf("bridgelib: grow memory: %w", growErr)
			}
			data = ins.memory.UnsafeData(ins.store)
			memSize = uint64(len(data))
		}
		copy(data[cursor:], b)
		ptrs[i] = cursor
		cursor += uint64(len(b))
	}

	// Slice header layout: [len: i64][ptr0: i64][ptr1: i64]...
	arraySize := uint64(1+len(ptrs)) * 8
	need := cursor + arraySize
	if need > memSize {
		pages := (need-memSize)/65536 + 1
		if _, growErr := ins.memory.Grow(ins.store, pages); growErr != nil {
			return 0, 0, fmt.Errorf("bridgelib: grow memory for array: %w", growErr)
		}
		data = ins.memory.UnsafeData(ins.store)
	}
	headerBase := cursor
	n := uint64(len(strs))
	data[headerBase+0] = byte(n)
	data[headerBase+1] = byte(n >> 8)
	data[headerBase+2] = byte(n >> 16)
	data[headerBase+3] = byte(n >> 24)
	data[headerBase+4] = byte(n >> 32)
	data[headerBase+5] = byte(n >> 40)
	data[headerBase+6] = byte(n >> 48)
	data[headerBase+7] = byte(n >> 56)
	for i, p := range ptrs {
		off := headerBase + 8 + uint64(i)*8
		data[off+0] = byte(p)
		data[off+1] = byte(p >> 8)
		data[off+2] = byte(p >> 16)
		data[off+3] = byte(p >> 24)
		data[off+4] = byte(p >> 32)
		data[off+5] = byte(p >> 40)
		data[off+6] = byte(p >> 48)
		data[off+7] = byte(p >> 56)
	}

	return int64(headerBase), int64(len(strs)), nil
}
