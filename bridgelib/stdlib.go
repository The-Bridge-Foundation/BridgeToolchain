package bridgelib

import (
	"context"
	"math"
	"sort"
	"sync"

	"bridgelang.dev/bridge-stdlib/modules"
	wasmtime "github.com/bytecodealliance/wasmtime-go/v28"
)

type sharedMem struct {
	mem   *wasmtime.Memory
	store *wasmtime.Store
	alloc modules.Allocator
}

type vmAllocator struct {
	mu     sync.Mutex
	sm     *sharedMem
	cursor uint64
	sizes  map[uint64]uint64
	free   []freeBlock
}

type freeBlock struct {
	addr uint64
	size uint64
}

const allocAlign = 8

func alignUp(n uint64) uint64 { return (n + allocAlign - 1) &^ (allocAlign - 1) }

func (a *vmAllocator) Alloc(size uint64) uint64 {
	if size == 0 {
		size = allocAlign
	}
	size = alignUp(size)

	a.mu.Lock()
	defer a.mu.Unlock()

	best, bestIdx := freeBlock{}, -1
	for i, b := range a.free {
		if b.size >= size && (bestIdx == -1 || b.size < best.size) {
			best, bestIdx = b, i
		}
	}

	var ptr uint64
	if bestIdx != -1 {
		ptr = best.addr
		remaining := best.size - size
		if remaining >= allocAlign*2 {
			a.free[bestIdx] = freeBlock{addr: best.addr + size, size: remaining}
		} else {
			// defragment
			size += remaining
			a.free = append(a.free[:bestIdx], a.free[bestIdx+1:]...)
		}
	} else {
		ptr = a.cursor
		a.cursor += size
		a.grow(a.cursor)
	}

	a.sizes[ptr] = size
	return ptr
}

func (a *vmAllocator) Free(ptr uint64) {
	if ptr == 0 {
		return
	}
	a.mu.Lock()
	defer a.mu.Unlock()

	size, ok := a.sizes[ptr]
	if !ok {
		return
	}
	delete(a.sizes, ptr)

	blk := freeBlock{addr: ptr, size: size}
	idx := sort.Search(len(a.free), func(i int) bool { return a.free[i].addr >= ptr })
	a.free = append(a.free, freeBlock{})
	copy(a.free[idx+1:], a.free[idx:])
	a.free[idx] = blk

	if idx+1 < len(a.free) && a.free[idx].addr+a.free[idx].size == a.free[idx+1].addr {
		a.free[idx].size += a.free[idx+1].size
		a.free = append(a.free[:idx+1], a.free[idx+2:]...)
	}

	if idx > 0 && a.free[idx-1].addr+a.free[idx-1].size == a.free[idx].addr {
		a.free[idx-1].size += a.free[idx].size
		a.free = append(a.free[:idx], a.free[idx+1:]...)
	}
}

func (a *vmAllocator) grow(needBytes uint64) {
	m := a.sm.mem
	if m == nil {
		return
	}
	cur := uint64(len(m.UnsafeData(a.sm.store)))
	if needBytes <= cur {
		return
	}
	pages := (needBytes-cur)/65536 + 1
	m.Grow(a.sm.store, pages)
}

func registerStdlib(linker *wasmtime.Linker, sm *sharedMem) error {
	for _, mod := range []struct {
		ns      string
		exports func() []modules.Export
	}{
		{modules.ConsoleName, modules.ConsoleExports},
		{modules.MathName, modules.MathExports},
		{modules.MemoryName, modules.MemoryExports},
		{modules.FileSystemName, modules.FileSystemExports},
	} {
		for _, exp := range mod.exports() {
			if err := registerExport(linker, mod.ns, exp, sm); err != nil {
				return err
			}
		}
	}
	return nil
}

func registerExport(linker *wasmtime.Linker, ns string, exp modules.Export, sm *sharedMem) error {
	fn := exp.Fn
	return linker.FuncNew("env", ns+"::"+exp.Name,
		wasmtime.NewFuncType(valueTypesToWasmtime(exp.Params), valueTypesToWasmtime(exp.Results)),
		func(caller *wasmtime.Caller, args []wasmtime.Val) ([]wasmtime.Val, *wasmtime.Trap) {
			stackSize := len(exp.Params)
			if len(exp.Results) > stackSize {
				stackSize = len(exp.Results)
			}
			stack := make([]uint64, stackSize)
			for i, v := range args {
				switch exp.Params[i] {
				case modules.ValueTypeF64:
					stack[i] = math.Float64bits(v.F64())
				case modules.ValueTypeF32:
					stack[i] = uint64(math.Float32bits(float32(v.F32())))
				case modules.ValueTypeI32:
					stack[i] = uint64(uint32(v.I32()))
				default:
					stack[i] = uint64(v.I64())
				}
			}
			ctx := context.Background()
			if sm.alloc != nil {
				ctx = modules.WithAllocator(ctx, sm.alloc)
			}
			fn(ctx, &callerRawMem{caller: caller, shared: sm}, stack)
			return stackToVals(stack, exp.Results), nil
		},
	)
}

func valueTypesToWasmtime(types []modules.ValueType) []*wasmtime.ValType {
	out := make([]*wasmtime.ValType, len(types))
	for i, t := range types {
		switch t {
		case modules.ValueTypeI32:
			out[i] = wasmtime.NewValType(wasmtime.KindI32)
		case modules.ValueTypeF32:
			out[i] = wasmtime.NewValType(wasmtime.KindF32)
		case modules.ValueTypeF64:
			out[i] = wasmtime.NewValType(wasmtime.KindF64)
		default:
			out[i] = wasmtime.NewValType(wasmtime.KindI64)
		}
	}
	return out
}

func stackToVals(stack []uint64, types []modules.ValueType) []wasmtime.Val {
	out := make([]wasmtime.Val, len(types))
	for i, t := range types {
		switch t {
		case modules.ValueTypeF64:
			out[i] = wasmtime.ValF64(math.Float64frombits(stack[i]))
		case modules.ValueTypeF32:
			out[i] = wasmtime.ValF32(math.Float32frombits(uint32(stack[i])))
		case modules.ValueTypeI32:
			out[i] = wasmtime.ValI32(int32(stack[i]))
		default:
			out[i] = wasmtime.ValI64(int64(stack[i]))
		}
	}
	return out
}

type callerRawMem struct {
	caller *wasmtime.Caller
	shared *sharedMem
}

func (c *callerRawMem) mem() *wasmtime.Memory {
	if exp := c.caller.GetExport("memory"); exp != nil {
		if m := exp.Memory(); m != nil {
			return m
		}
	}
	if c.shared != nil {
		return c.shared.mem
	}
	return nil
}

func (c *callerRawMem) ReadByte(offset uint32) (byte, bool) {
	m := c.mem()
	if m == nil {
		return 0, false
	}
	d := m.UnsafeData(c.caller)
	if int(offset) >= len(d) {
		return 0, false
	}
	return d[offset], true
}

func (c *callerRawMem) Read(offset, size uint32) ([]byte, bool) {
	m := c.mem()
	if m == nil {
		return nil, false
	}
	d := m.UnsafeData(c.caller)
	end := int(offset) + int(size)
	if end > len(d) {
		return nil, false
	}
	out := make([]byte, size)
	copy(out, d[offset:end])
	return out, true
}

func (c *callerRawMem) Write(offset uint32, v []byte) bool {
	m := c.mem()
	if m == nil {
		return false
	}
	d := m.UnsafeData(c.caller)
	end := int(offset) + len(v)
	if end > len(d) {
		return false
	}
	copy(d[offset:end], v)
	return true
}

func (c *callerRawMem) Size() uint32 {
	m := c.mem()
	if m == nil {
		return 0
	}
	return uint32(len(m.UnsafeData(c.caller)))
}

func (c *callerRawMem) Grow(deltaPages uint32) (uint32, bool) {
	m := c.mem()
	if m == nil {
		return 0, false
	}
	prev, err := m.Grow(c.caller, uint64(deltaPages))
	return uint32(prev), err == nil
}
