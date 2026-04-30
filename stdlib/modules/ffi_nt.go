//go:build !darwin && !linux && !freebsd

package modules

import (
	"context"

	"bridgelang.dev/bridge-stdlib/spec"
)

// TODO: Implement for windows

func init() {
	spec.Register(spec.Module{
		Name: "FFI",
		Funcs: []spec.Func{
			{Name: "Load", ParamTypes: []string{"string"}, ReturnType: "any"},
			{Name: "Lookup", ParamTypes: []string{"any", "string"}, ReturnType: "any"},
			{Name: "Call", ParamTypes: []string{"any", "int64[]"}, ReturnType: "int64"},
			{Name: "Close", ParamTypes: []string{"any"}, ReturnType: ""},
		},
	})
}

const FFIName = "FFI"

func FFIExports() []Export {
	return []Export{
		{Name: "Load", Fn: ffiLoad, Params: []ValueType{ValueTypeI64}, Results: []ValueType{ValueTypeI64}},
		{Name: "Lookup", Fn: ffiLookup, Params: []ValueType{ValueTypeI64, ValueTypeI64}, Results: []ValueType{ValueTypeI64}},
		{Name: "Call", Fn: ffiCall, Params: []ValueType{ValueTypeI64, ValueTypeI64}, Results: []ValueType{ValueTypeI64}},
		{Name: "Close", Fn: ffiClose, Params: []ValueType{ValueTypeI64}, Results: nil},
	}
}

func ffiLoad(_ context.Context, m RawMemory, stack []uint64)   { stack[0] = 0 }
func ffiLookup(_ context.Context, m RawMemory, stack []uint64) { stack[0] = 0 }
func ffiCall(_ context.Context, m RawMemory, stack []uint64)   { stack[0] = uint64(int64(-1)) }
func ffiClose(_ context.Context, _ RawMemory, stack []uint64)  {}
