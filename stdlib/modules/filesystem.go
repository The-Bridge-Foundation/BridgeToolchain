package modules

import (
	"context"
	"encoding/binary"
	"os"

	"bridgelang.dev/bridge-stdlib/spec"
)

func init() {
	spec.Register(spec.Module{
		Name: "Filesystem",
		Funcs: []spec.Func{
			{Name: "Touch", ParamTypes: []string{"string"}, ReturnType: "string"},
			{Name: "DeleteFile", ParamTypes: []string{"string"}, ReturnType: "string"},
			{Name: "WriteFile", ParamTypes: []string{"string", "byte[]", "int"}, ReturnType: "string"},
			{Name: "ReadFile", ParamTypes: []string{"string"}, ReturnType: "byte[]"},
		},
	})
}

// FileSystemModule provides Filesystem::* stdlib functions for Bridge programs.
// Bridge usage: @Import(Filesystem)
//   Filesystem::Touch(path string)
//   Filesystem::DeleteFile(path string)
//   Filesystem::WriteFile(path string, data byte[], len int)
//   Filesystem::ReadFile(path string) byte[]

const FileSystemName = "Filesystem"

func FileSystemExports() []Export {
	ptr := ValueTypeI64
	i64 := ValueTypeI64
	return []Export{
		{Name: "Touch", Fn: filesystemTouch, Params: []ValueType{ptr}, Results: []ValueType{ptr}},
		{Name: "DeleteFile", Fn: filesystemDeleteFile, Params: []ValueType{ptr}, Results: []ValueType{ptr}},
		{Name: "WriteFile", Fn: filesystemWriteFile, Params: []ValueType{ptr, ptr, i64}, Results: []ValueType{ptr}},
		{Name: "ReadFile", Fn: filesystemReadFile, Params: []ValueType{ptr}, Results: []ValueType{ptr}},
	}
}

func errString(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}

func filesystemTouch(ctx context.Context, m RawMemory, stack []uint64) {
	s := readString(m, stack[0])
	err := os.WriteFile(s, []byte{}, 0644)
	stack[0] = writeString(m, errString(err))
}

func filesystemDeleteFile(ctx context.Context, m RawMemory, stack []uint64) {
	path := readString(m, stack[0])
	err := os.Remove(path)
	stack[0] = writeString(m, errString(err))
}

func filesystemWriteFile(ctx context.Context, m RawMemory, stack []uint64) {
	path := readString(m, stack[0])
	dataPtr := uint32(stack[1])
	dataLen := uint32(stack[2])
	data, ok := m.Read(dataPtr, dataLen)
	if !ok {
		stack[0] = writeString(m, "read: out of bounds")
		return
	}
	err := os.WriteFile(path, data, 0644)
	stack[0] = writeString(m, errString(err))
}

func filesystemReadFile(ctx context.Context, m RawMemory, stack []uint64) {
	path := readString(m, stack[0])
	data, err := os.ReadFile(path)
	if err != nil {
		stack[0] = 0
		return
	}
	buf := make([]byte, 8+len(data))
	binary.LittleEndian.PutUint64(buf[:8], uint64(len(data)))
	copy(buf[8:], data)
	base := m.Size()
	pages := (uint32(len(buf)) + 65535) / 65536
	if _, ok := m.Grow(pages); !ok {
		stack[0] = 0
		return
	}
	m.Write(base, buf)
	stack[0] = uint64(base)
}
