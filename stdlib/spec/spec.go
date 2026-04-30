package spec

type Func struct {
	Name       string   `json:"name"`
	ParamTypes []string `json:"params"`
	ReturnType string   `json:"return"`
	Variadic   bool     `json:"variadic,omitempty"`
}

type ObjectField struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Object struct {
	Name   string        `json:"name"`
	Fields []ObjectField `json:"fields"`
}

type Module struct {
	Name    string   `json:"name"`
	Funcs   []Func   `json:"funcs"`
	Objects []Object `json:"objects,omitempty"`
}

var registry []Module

func Register(m Module) {
	registry = append(registry, m)
}

func All() []Module {
	return registry
}

func Lookup(name string) (Module, bool) {
	for _, m := range registry {
		if m.Name == name {
			return m, true
		}
	}
	return Module{}, false
}
