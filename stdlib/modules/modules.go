package modules

import "bridgelang.dev/bridge-stdlib/spec"

var moduleExports = map[string]func() []Export{
	"Console":    ConsoleExports,
	"Math":       MathExports,
	"Memory":     MemoryExports,
	"Filesystem": FileSystemExports,
}

// For runtime usage
func All() map[string][]Export {
	out := make(map[string][]Export, len(moduleExports))
	for _, m := range spec.All() {
		if fn, ok := moduleExports[m.Name]; ok {
			out[m.Name] = fn()
		}
	}
	return out
}
