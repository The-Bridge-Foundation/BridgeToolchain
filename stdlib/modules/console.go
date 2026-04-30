package modules

import (
	"context"
	"encoding/binary"
	"fmt"
	"math"

	"bridgelang.dev/bridge-stdlib/spec"
)

func init() {
	spec.Register(spec.Module{
		Name: "Console",
		Funcs: []spec.Func{
			{Name: "Printf", ParamTypes: []string{"string"}, ReturnType: "int32", Variadic: true},
			{Name: "Print", ParamTypes: []string{"string"}, ReturnType: ""},
			{Name: "Read", ParamTypes: nil, ReturnType: "string"},
			{Name: "Scanf", ParamTypes: []string{"string"}, ReturnType: "int32", Variadic: true},
		},
	})
}

// ConsoleModule provides Console::* stdlib functions for Bridge programs.
// Bridge usage: @Import(Console)
//   Console::Printf(format string, ...) int32
//   Console::Print(str string)
//   Console::Read() string
//	 Console::Scanf(str string, ...) int32

const ConsoleName = "Console"

func ConsoleExports() []Export {
	return []Export{
		{Name: "Printf", Fn: consolePrintf, Params: []ValueType{ValueTypeI64, ValueTypeI64}, Results: []ValueType{ValueTypeI32}},
		{Name: "Print", Fn: consolePrint, Params: []ValueType{ValueTypeI64}, Results: nil},
		{Name: "Read", Fn: consoleReadLine, Params: nil, Results: []ValueType{ValueTypeI64}},
		{Name: "Scanf", Fn: consoleScanf, Params: []ValueType{ValueTypeI64, ValueTypeI64}, Results: []ValueType{ValueTypeI32}},
	}
}

func consolePrintf(ctx context.Context, m RawMemory, stack []uint64) {
	fmtStr := readString(m, stack[0])
	verbs := parseFmtVerbs(fmtStr)
	vaListPtr := stack[1]
	args := make([]any, len(verbs))
	for i, verb := range verbs {
		val := uint64(readInt64(m, vaListPtr+uint64(i)*8))
		switch verb {
		case 's':
			args[i] = readString(m, val)
		case 'f', 'F', 'e', 'E', 'g', 'G':
			args[i] = math.Float64frombits(val)
		default:
			args[i] = int64(val)
		}
	}
	n, _ := fmt.Printf(fmtStr, args...)
	stack[0] = uint64(n)
}

func parseFmtVerbs(s string) []byte {
	var verbs []byte
	for i := 0; i < len(s); i++ {
		if s[i] != '%' {
			continue
		}
		i++
		// skip flags, width, precision
		for i < len(s) && (s[i] == '-' || s[i] == '+' || s[i] == ' ' || s[i] == '#' || s[i] == '0' || (s[i] >= '1' && s[i] <= '9') || s[i] == '.' || s[i] == '*') {
			i++
		}
		if i < len(s) && s[i] != '%' {
			verbs = append(verbs, s[i])
		}
	}
	return verbs
}

func consolePrint(ctx context.Context, m RawMemory, stack []uint64) {
	str := readString(m, stack[0])
	fmt.Println(str)
}

func consoleScanf(ctx context.Context, m RawMemory, stack []uint64) {
	fmtStr := readString(m, stack[0])
	verbs := parseFmtVerbs(fmtStr)
	vaListPtr := stack[1]

	ptrs := make([]uint64, len(verbs))
	for i := range ptrs {
		ptrs[i] = uint64(readInt64(m, vaListPtr+uint64(i)*8))
	}

	args := make([]any, len(verbs))
	scanTargets := make([]any, len(verbs))
	for i, verb := range verbs {
		switch verb {
		case 's':
			s := new(string)
			args[i] = s
			scanTargets[i] = s
		case 'f', 'F', 'e', 'E', 'g', 'G':
			f := new(float64)
			args[i] = f
			scanTargets[i] = f
		default:
			n := new(int64)
			args[i] = n
			scanTargets[i] = n
		}
	}

	n, _ := fmt.Scanf(fmtStr, args...)

	for i, verb := range verbs {
		if i >= n {
			break
		}
		dst := ptrs[i]
		switch verb {
		case 's':
			ptr := writeString(m, *scanTargets[i].(*string))
			var buf [8]byte
			binary.LittleEndian.PutUint64(buf[:], ptr)
			m.Write(uint32(dst), buf[:])
		case 'f', 'F', 'e', 'E', 'g', 'G':
			bits := math.Float64bits(*scanTargets[i].(*float64))
			var buf [8]byte
			binary.LittleEndian.PutUint64(buf[:], bits)
			m.Write(uint32(dst), buf[:])
		default:
			val := uint64(*scanTargets[i].(*int64))
			var buf [8]byte
			binary.LittleEndian.PutUint64(buf[:], val)
			m.Write(uint32(dst), buf[:])
		}
	}
	stack[0] = uint64(n)
}

func consoleReadLine(ctx context.Context, m RawMemory, stack []uint64) {
	var line string
	fmt.Scanln(&line)
	ptr := writeString(m, line)
	stack[0] = ptr
}
