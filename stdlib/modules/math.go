package modules

import (
	"context"
	"math"

	"bridgelang.dev/bridge-stdlib/spec"
)

func init() {
	spec.Register(spec.Module{
		Name: "Math",
		Funcs: []spec.Func{
			{Name: "Abs", ParamTypes: []string{"f64"}, ReturnType: "f64"},
			{Name: "Sqrt", ParamTypes: []string{"f64"}, ReturnType: "f64"},
			{Name: "Pow", ParamTypes: []string{"f64", "f64"}, ReturnType: "f64"},
			{Name: "Floor", ParamTypes: []string{"f64"}, ReturnType: "f64"},
			{Name: "Ceil", ParamTypes: []string{"f64"}, ReturnType: "f64"},
			{Name: "Round", ParamTypes: []string{"f64"}, ReturnType: "f64"},
			{Name: "Min", ParamTypes: []string{"f64", "f64"}, ReturnType: "f64"},
			{Name: "Max", ParamTypes: []string{"f64", "f64"}, ReturnType: "f64"},
			{Name: "Sin", ParamTypes: []string{"f64"}, ReturnType: "f64"},
			{Name: "Cos", ParamTypes: []string{"f64"}, ReturnType: "f64"},
			{Name: "Tan", ParamTypes: []string{"f64"}, ReturnType: "f64"},
			{Name: "Log", ParamTypes: []string{"f64"}, ReturnType: "f64"},
			{Name: "Log2", ParamTypes: []string{"f64"}, ReturnType: "f64"},
			{Name: "Log10", ParamTypes: []string{"f64"}, ReturnType: "f64"},
			{Name: "Pi", ParamTypes: nil, ReturnType: "f64"},
			{Name: "E", ParamTypes: nil, ReturnType: "f64"},
		},
	})
}

// MathModule provides Math::* stdlib functions for Bridge programs.
// Bridge usage: @Import(Math)
//   Math::Abs(x f64) f64
//   Math::Sqrt(x f64) f64
//   Math::Pow(base f64, exp f64) f64
//   Math::Floor(x f64) f64
//   Math::Ceil(x f64) f64
//   Math::Round(x f64) f64
//   Math::Min(a f64, b f64) f64
//   Math::Max(a f64, b f64) f64
//   Math::Sin(x f64) f64
//   Math::Cos(x f64) f64
//   Math::Tan(x f64) f64
//   Math::Log(x f64) f64
//   Math::Log2(x f64) f64
//   Math::Log10(x f64) f64
//   Math::Pi() f64
//   Math::E() f64

const MathName = "Math"

func MathExports() []Export {
	f64x1 := []ValueType{ValueTypeF64}
	f64x2 := []ValueType{ValueTypeF64, ValueTypeF64}
	retf64 := []ValueType{ValueTypeF64}

	return []Export{
		{Name: "Abs", Fn: mathAbs, Params: f64x1, Results: retf64},
		{Name: "Sqrt", Fn: mathSqrt, Params: f64x1, Results: retf64},
		{Name: "Pow", Fn: mathPow, Params: f64x2, Results: retf64},
		{Name: "Floor", Fn: mathFloor, Params: f64x1, Results: retf64},
		{Name: "Ceil", Fn: mathCeil, Params: f64x1, Results: retf64},
		{Name: "Round", Fn: mathRound, Params: f64x1, Results: retf64},
		{Name: "Min", Fn: mathMin, Params: f64x2, Results: retf64},
		{Name: "Max", Fn: mathMax, Params: f64x2, Results: retf64},
		{Name: "Sin", Fn: mathSin, Params: f64x1, Results: retf64},
		{Name: "Cos", Fn: mathCos, Params: f64x1, Results: retf64},
		{Name: "Tan", Fn: mathTan, Params: f64x1, Results: retf64},
		{Name: "Log", Fn: mathLog, Params: f64x1, Results: retf64},
		{Name: "Log2", Fn: mathLog2, Params: f64x1, Results: retf64},
		{Name: "Log10", Fn: mathLog10, Params: f64x1, Results: retf64},
		{Name: "Pi", Fn: mathPi, Params: nil, Results: retf64},
		{Name: "E", Fn: mathE, Params: nil, Results: retf64},
	}
}

func mathAbs(ctx context.Context, m RawMemory, stack []uint64) {
	stack[0] = EncodeF64(math.Abs(DecodeF64(stack[0])))
}

func mathSqrt(ctx context.Context, m RawMemory, stack []uint64) {
	stack[0] = EncodeF64(math.Sqrt(DecodeF64(stack[0])))
}

func mathPow(ctx context.Context, m RawMemory, stack []uint64) {
	stack[0] = EncodeF64(math.Pow(DecodeF64(stack[0]), DecodeF64(stack[1])))
}

func mathFloor(ctx context.Context, m RawMemory, stack []uint64) {
	stack[0] = EncodeF64(math.Floor(DecodeF64(stack[0])))
}

func mathCeil(ctx context.Context, m RawMemory, stack []uint64) {
	stack[0] = EncodeF64(math.Ceil(DecodeF64(stack[0])))
}

func mathRound(ctx context.Context, m RawMemory, stack []uint64) {
	stack[0] = EncodeF64(math.Round(DecodeF64(stack[0])))
}

func mathMin(ctx context.Context, m RawMemory, stack []uint64) {
	stack[0] = EncodeF64(math.Min(DecodeF64(stack[0]), DecodeF64(stack[1])))
}

func mathMax(ctx context.Context, m RawMemory, stack []uint64) {
	stack[0] = EncodeF64(math.Max(DecodeF64(stack[0]), DecodeF64(stack[1])))
}

func mathSin(ctx context.Context, m RawMemory, stack []uint64) {
	stack[0] = EncodeF64(math.Sin(DecodeF64(stack[0])))
}

func mathCos(ctx context.Context, m RawMemory, stack []uint64) {
	stack[0] = EncodeF64(math.Cos(DecodeF64(stack[0])))
}

func mathTan(ctx context.Context, m RawMemory, stack []uint64) {
	stack[0] = EncodeF64(math.Tan(DecodeF64(stack[0])))
}

func mathLog(ctx context.Context, m RawMemory, stack []uint64) {
	stack[0] = EncodeF64(math.Log(DecodeF64(stack[0])))
}

func mathLog2(ctx context.Context, m RawMemory, stack []uint64) {
	stack[0] = EncodeF64(math.Log2(DecodeF64(stack[0])))
}

func mathLog10(ctx context.Context, m RawMemory, stack []uint64) {
	stack[0] = EncodeF64(math.Log10(DecodeF64(stack[0])))
}

func mathPi(ctx context.Context, m RawMemory, stack []uint64) {
	stack[0] = EncodeF64(math.Pi)
}

func mathE(ctx context.Context, m RawMemory, stack []uint64) {
	stack[0] = EncodeF64(math.E)
}
