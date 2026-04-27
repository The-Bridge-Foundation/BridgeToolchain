package modules

import (
	"context"
	"math"
)

type ValueType = byte

const (
	ValueTypeI32 ValueType = 0x7F
	ValueTypeI64 ValueType = 0x7E
	ValueTypeF32 ValueType = 0x7D
	ValueTypeF64 ValueType = 0x7C
)

func EncodeF64(v float64) uint64 { return math.Float64bits(v) }
func DecodeF64(v uint64) float64 { return math.Float64frombits(v) }

type RawMemory interface {
	ReadByte(offset uint32) (byte, bool)
	Read(offset, size uint32) ([]byte, bool)
	Write(offset uint32, v []byte) bool
	Size() uint32
	Grow(deltaPages uint32) (uint32, bool)
}

type Export struct {
	Name    string
	Fn      func(ctx context.Context, m RawMemory, stack []uint64)
	Params  []ValueType
	Results []ValueType
}

type Allocator interface {
	Alloc(size uint64) uint64
	Free(ptr uint64)
}

type allocKey struct{}

func WithAllocator(ctx context.Context, a Allocator) context.Context {
	return context.WithValue(ctx, allocKey{}, a)
}

func AllocatorFromCtx(ctx context.Context) Allocator {
	if a, ok := ctx.Value(allocKey{}).(Allocator); ok {
		return a
	}
	return nil
}

func readInt64(m RawMemory, ptr uint64) int64 {
	b, ok := m.Read(uint32(ptr), 8)
	if !ok || len(b) < 8 {
		return 0
	}
	return int64(b[0]) | int64(b[1])<<8 | int64(b[2])<<16 | int64(b[3])<<24 |
		int64(b[4])<<32 | int64(b[5])<<40 | int64(b[6])<<48 | int64(b[7])<<56
}

func readString(m RawMemory, ptr uint64) string {
	p := uint32(ptr)
	var buf []byte
	for {
		b, ok := m.ReadByte(p)
		if !ok || b == 0 {
			break
		}
		buf = append(buf, b)
		p++
	}
	return string(buf)
}

func writeString(m RawMemory, s string) uint64 {
	data := append([]byte(s), 0)
	base := m.Size()
	pages := (uint32(len(data)) + 65535) / 65536
	if _, ok := m.Grow(pages); !ok {
		return 0
	}
	m.Write(base, data)
	return uint64(base)
}
