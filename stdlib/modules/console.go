package modules

import (
	"context"
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
		},
	})
}

// ConsoleModule provides Console::* stdlib functions for Bridge programs.
// Bridge usage: @Import(Console)
//   Console::Printf(format string, ...) int32
//   Console::Print(str string)
//   Console::Read() string

const ConsoleName = "Console"

func ConsoleExports() []Export {
	return []Export{
		{Name: "Printf", Fn: consolePrintf, Params: []ValueType{ValueTypeI64, ValueTypeI64}, Results: []ValueType{ValueTypeI32}},
		{Name: "Print", Fn: consolePrint, Params: []ValueType{ValueTypeI64}, Results: nil},
		{Name: "Read", Fn: consoleReadLine, Params: nil, Results: []ValueType{ValueTypeI64}},
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

func consoleReadLine(ctx context.Context, m RawMemory, stack []uint64) {
	var line string
	fmt.Scanln(&line)
	ptr := writeString(m, line)
	stack[0] = ptr
}
