module bridgelang.dev/bridgec

go 1.25.1

require (
	bridgelang.dev/bridge-stdlib v0.0.0
	bridgelang.dev/bridgec-libs/antlr v0.0.0
	bridgelang.dev/bridgelib v0.0.0
	github.com/antlr/antlr4 v0.0.0-20210311221813-5e5b6d35b418
	github.com/google/uuid v1.6.0
	tinygo.org/x/go-llvm v0.0.0-20250929104024-00fb4309ddd2
)

require github.com/bytecodealliance/wasmtime-go/v28 v28.0.0 // indirect

replace bridgelang.dev/bridge-stdlib v0.0.0 => ../stdlib

replace bridgelang.dev/bridgec-libs/antlr v0.0.0 => ../libs/antlr

replace bridgelang.dev/bridgelib v0.0.0 => ../bridgelib
