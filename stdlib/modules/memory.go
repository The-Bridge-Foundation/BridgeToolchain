package modules

import (
	"context"

	"bridgelang.dev/bridge-stdlib/spec"
)

func init() {
	spec.Register(spec.Module{
		Name: "Memory",
		Funcs: []spec.Func{
			{Name: "Alloc", ParamTypes: []string{"int64"}, ReturnType: "any"},
			{Name: "Free", ParamTypes: []string{"any"}, ReturnType: ""},
			{Name: "Copy", ParamTypes: []string{"any", "any", "int64"}, ReturnType: ""},
			{Name: "Set", ParamTypes: []string{"any", "byte", "int64"}, ReturnType: ""},
			{Name: "Size", ParamTypes: nil, ReturnType: "int64"},
			{Name: "Len", ParamTypes: []string{"any"}, ReturnType: "int64"},
			{Name: "StrLen", ParamTypes: []string{"string"}, ReturnType: "int64"},
		},
	})
}

// MemoryModule provides Memory::* stdlib functions for Bridge programs.
// Bridge usage: @Import(Memory)
//   Memory::Alloc(size int64) any
//   Memory::Free(ptr any)
//   Memory::Copy(dst any, src any, size int64)
//   Memory::Set(ptr any, val byte, size int64)
//   Memory::Size() int64

const MemoryName = "Memory"

func MemoryExports() []Export {
	ptr := ValueTypeI64
	i64 := ValueTypeI64
	i32 := ValueTypeI32

	return []Export{
		{Name: "Alloc", Fn: memAlloc, Params: []ValueType{i64}, Results: []ValueType{ptr}},
		{Name: "Free", Fn: memFree, Params: []ValueType{ptr}, Results: nil},
		{Name: "Copy", Fn: memCopy, Params: []ValueType{ptr, ptr, i64}, Results: nil},
		{Name: "Set", Fn: memSet, Params: []ValueType{ptr, i32, i64}, Results: nil},
		{Name: "Size", Fn: memSize, Params: nil, Results: []ValueType{i64}},
		{Name: "Len", Fn: memLen, Params: []ValueType{ptr}, Results: []ValueType{i64}},
		{Name: "StrLen", Fn: memStrLen, Params: []ValueType{ptr}, Results: []ValueType{i64}},
	}
}

func memAlloc(ctx context.Context, m RawMemory, stack []uint64) {
	a := AllocatorFromCtx(ctx)
	if a == nil {
		stack[0] = 0
		return
	}
	stack[0] = a.Alloc(stack[0])
}

func memFree(ctx context.Context, m RawMemory, stack []uint64) {
	a := AllocatorFromCtx(ctx)
	if a != nil {
		a.Free(stack[0])
	}
}

func memCopy(ctx context.Context, m RawMemory, stack []uint64) {
	dst := uint32(stack[0])
	src := uint32(stack[1])
	size := uint32(stack[2])
	buf, ok := m.Read(src, size)
	if !ok {
		return
	}
	m.Write(dst, buf)
}

func memSet(ctx context.Context, m RawMemory, stack []uint64) {
	ptr := uint32(stack[0])
	val := byte(stack[1])
	size := uint32(stack[2])
	buf := make([]byte, size)
	for i := range buf {
		buf[i] = val
	}
	m.Write(ptr, buf)
}

func memSize(ctx context.Context, m RawMemory, stack []uint64) {
	stack[0] = uint64(m.Size())
}

func memStrLen(ctx context.Context, m RawMemory, stack []uint64) {
	p := uint32(stack[0])
	var n uint64
	for {
		b, ok := m.ReadByte(p)
		if !ok || b == 0 {
			break
		}
		n++
		p++
	}
	stack[0] = n
}

func memLen(ctx context.Context, m RawMemory, stack []uint64) {
	ptr := uint32(stack[0])
	b, ok := m.Read(ptr, 8)
	if !ok || len(b) < 8 {
		stack[0] = 0
		return
	}
	stack[0] = uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
}
