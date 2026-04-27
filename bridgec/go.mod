module bridgelang.dev/bridgec

go 1.25.1

require (
	bridgelang.dev/bridge-stdlib v0.0.0
	github.com/google/uuid v1.6.0
	tinygo.org/x/go-llvm v0.0.0-20250929104024-00fb4309ddd2
)

require golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect

replace bridgelang.dev/bridge-stdlib v0.0.0 => ../stdlib
