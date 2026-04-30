//go:build darwin || linux || freebsd

package modules

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include <stdint.h>
// #include <stdlib.h>
//
// typedef int64_t (*ffi_fn_i64_0)(void);
// typedef int64_t (*ffi_fn_i64_1)(int64_t);
// typedef int64_t (*ffi_fn_i64_2)(int64_t,int64_t);
// typedef int64_t (*ffi_fn_i64_3)(int64_t,int64_t,int64_t);
// typedef int64_t (*ffi_fn_i64_4)(int64_t,int64_t,int64_t,int64_t);
// typedef int64_t (*ffi_fn_i64_5)(int64_t,int64_t,int64_t,int64_t,int64_t);
// typedef int64_t (*ffi_fn_i64_6)(int64_t,int64_t,int64_t,int64_t,int64_t,int64_t);
// typedef int64_t (*ffi_fn_i64_7)(int64_t,int64_t,int64_t,int64_t,int64_t,int64_t,int64_t);
// typedef int64_t (*ffi_fn_i64_8)(int64_t,int64_t,int64_t,int64_t,int64_t,int64_t,int64_t,int64_t);
//
// static int64_t bridge_ffi_call_i64(void* fn, int64_t* args, int n) {
//     switch (n) {
//     case 0: return ((ffi_fn_i64_0)fn)();
//     case 1: return ((ffi_fn_i64_1)fn)(args[0]);
//     case 2: return ((ffi_fn_i64_2)fn)(args[0],args[1]);
//     case 3: return ((ffi_fn_i64_3)fn)(args[0],args[1],args[2]);
//     case 4: return ((ffi_fn_i64_4)fn)(args[0],args[1],args[2],args[3]);
//     case 5: return ((ffi_fn_i64_5)fn)(args[0],args[1],args[2],args[3],args[4]);
//     case 6: return ((ffi_fn_i64_6)fn)(args[0],args[1],args[2],args[3],args[4],args[5]);
//     case 7: return ((ffi_fn_i64_7)fn)(args[0],args[1],args[2],args[3],args[4],args[5],args[6]);
//     case 8: return ((ffi_fn_i64_8)fn)(args[0],args[1],args[2],args[3],args[4],args[5],args[6],args[7]);
//     default: return -1;
//     }
// }
import "C"

import (
	"context"
	"sync"
	"sync/atomic"
	"unsafe"

	"bridgelang.dev/bridge-stdlib/spec"
)

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

var (
	ffiHandles sync.Map
	ffiCounter int64
)

func nextFFIHandle() int64 {
	return atomic.AddInt64(&ffiCounter, 1)
}

// (ID < 0) = lib handles, (ID > 0) = fn handles.
func ffiLoad(_ context.Context, m RawMemory, stack []uint64) {
	path := readString(m, stack[0])
	cs := C.CString(path)
	defer C.free(unsafe.Pointer(cs))
	handle := C.dlopen(cs, C.RTLD_NOW)
	if handle == nil {
		stack[0] = 0
		return
	}
	id := nextFFIHandle()
	ffiHandles.Store(id, unsafe.Pointer(handle))
	stack[0] = uint64(id)
}

func ffiLookup(_ context.Context, m RawMemory, stack []uint64) {
	libID := int64(stack[0])
	sym := readString(m, stack[1])

	raw, ok := ffiHandles.Load(libID)
	if !ok {
		stack[0] = 0
		return
	}
	lib := raw.(unsafe.Pointer)

	cs := C.CString(sym)
	defer C.free(unsafe.Pointer(cs))
	fn := C.dlsym(lib, cs)
	if fn == nil {
		stack[0] = 0
		return
	}
	id := nextFFIHandle()
	ffiHandles.Store(id, unsafe.Pointer(fn))
	stack[0] = uint64(id)
}

func ffiCall(_ context.Context, m RawMemory, stack []uint64) {
	fnID := int64(stack[0])
	slicePtr := uint32(stack[1])

	raw, ok := ffiHandles.Load(fnID)
	if !ok {
		stack[0] = ^uint64(0) // -1 in two's complement
		return
	}
	fn := raw.(unsafe.Pointer)

	lenBytes, ok := m.Read(slicePtr, 8)
	if !ok || len(lenBytes) < 8 {
		stack[0] = ^uint64(0)
		return
	}
	argsLen := int64(lenBytes[0]) | int64(lenBytes[1])<<8 | int64(lenBytes[2])<<16 | int64(lenBytes[3])<<24 |
		int64(lenBytes[4])<<32 | int64(lenBytes[5])<<40 | int64(lenBytes[6])<<48 | int64(lenBytes[7])<<56

	nargs := int(argsLen)
	if nargs > 8 {
		nargs = 8
	}

	dataPtr := slicePtr + 8
	cargs := make([]C.int64_t, nargs)
	for i := 0; i < nargs; i++ {
		b, ok := m.Read(dataPtr+uint32(i)*8, 8)
		if !ok || len(b) < 8 {
			break
		}
		v := int64(b[0]) | int64(b[1])<<8 | int64(b[2])<<16 | int64(b[3])<<24 |
			int64(b[4])<<32 | int64(b[5])<<40 | int64(b[6])<<48 | int64(b[7])<<56
		cargs[i] = C.int64_t(v)
	}

	var cargsPtr *C.int64_t
	if nargs > 0 {
		cargsPtr = &cargs[0]
	}

	result := C.bridge_ffi_call_i64(fn, cargsPtr, C.int(nargs))
	stack[0] = uint64(int64(result))
}

func ffiClose(_ context.Context, _ RawMemory, stack []uint64) {
	id := int64(stack[0])
	if raw, ok := ffiHandles.LoadAndDelete(id); ok {
		C.dlclose(raw.(unsafe.Pointer))
	}
}
