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

var builtins = []Module{
	{
		Name: "Console",
		Funcs: []Func{
			{Name: "Printf", ParamTypes: []string{"string"}, ReturnType: "int32", Variadic: true},
			{Name: "Print", ParamTypes: []string{"string"}, ReturnType: ""},
			{Name: "Read", ParamTypes: nil, ReturnType: "string"},
		},
	},
	{
		Name: "Filesystem",
		Funcs: []Func{
			{Name: "Touch", ParamTypes: []string{"string"}, ReturnType: "string"},
			{Name: "DeleteFile", ParamTypes: []string{"string"}, ReturnType: "string"},
			{Name: "WriteFile", ParamTypes: []string{"string", "byte[]", "int"}, ReturnType: "string"},
			{Name: "ReadFile", ParamTypes: []string{"string"}, ReturnType: "byte[]"},
		},
	},
	{
		Name: "Math",
		Funcs: []Func{
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
	},
	{
		Name: "Memory",
		Funcs: []Func{
			{Name: "Alloc", ParamTypes: []string{"int64"}, ReturnType: "any"},
			{Name: "Free", ParamTypes: []string{"any"}, ReturnType: ""},
			{Name: "Copy", ParamTypes: []string{"any", "any", "int64"}, ReturnType: ""},
			{Name: "Set", ParamTypes: []string{"any", "byte", "int64"}, ReturnType: ""},
			{Name: "Size", ParamTypes: nil, ReturnType: "int64"},
			{Name: "Len", ParamTypes: []string{"any"}, ReturnType: "int64"},
			{Name: "StrLen", ParamTypes: []string{"string"}, ReturnType: "int64"},
		},
	},
}

var dynamic []Module

func Register(m Module) {
	dynamic = append(dynamic, m)
}

func All() []Module {
	return append(builtins, dynamic...)
}

func Lookup(name string) (Module, bool) {
	for _, m := range builtins {
		if m.Name == name {
			return m, true
		}
	}
	for _, m := range dynamic {
		if m.Name == name {
			return m, true
		}
	}
	return Module{}, false
}
