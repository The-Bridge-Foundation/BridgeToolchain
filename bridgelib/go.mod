module bridgelang.dev/bridgelib

go 1.25.1

require (
	bridgelang.dev/bridge-stdlib v0.0.0
	github.com/bytecodealliance/wasmtime-go/v28 v28.0.0
)

replace bridgelang.dev/bridge-stdlib v0.0.0 => ../stdlib
