package typescript

import (
	"github.com/h2oai/steamY/tools/piping/parser"
	"github.com/serenize/snaker"
	"strings"
)

func cast(t string) string {
	degree := 0
	for {
		if !strings.HasPrefix(t, "[]") {
			break
		}
		t = t[2:]
		degree++
	}
	if strings.HasPrefix(t, "*") {
		t = t[1:]
	}
	switch t {
	case "bool":
		t = "boolean"
	case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "float32", "float64", "int", "uint":
		t = "number"
	}
	for i := 0; i < degree; i++ {
		t = t + "[]"

	}
	return t
}

func pascalCase(s string) string {
	switch len(s) {
	case 0:
		return ""
	case 1:
		return strings.ToUpper(s)
	default:
		return strings.ToUpper(string(s[0])) + s[1:]
	}
}

func camelCase(s string) string {
	switch len(s) {
	case 0:
		return ""
	case 1:
		return strings.ToLower(s)
	default:
		return strings.ToLower(string(s[0])) + s[1:]
	}
}

func genParam(p *parser.Param) string {
	return p.Name + ": " + cast(p.Type)
}

func genMember(p *parser.Param) string {
	return snaker.CamelToSnake(p.Name) + ": " + cast(p.Type)
}

func genStruct(s *parser.Struct, public bool) string {
	c := "\t"
	if public {
		c += "export "
	}
	c += "interface " + s.Name + " {\n"
	for _, p := range s.Params {
		c += "\t\t" + genMember(p) + "\n"
	}
	return c + "\t}\n\n"
}

func genReturn(f *parser.Func) string {
	c := "go: (error: Error"
	if f.Return != nil {
		c += ", " + f.Return.Name + ": " + cast(f.Return.Type)
	}
	c += ") => void"
	return c
}

func genFunc(f *parser.Func) string {
	c := camelCase(f.Name) + ": ("
	for _, p := range f.Params {
		c += genParam(p) + ", "
	}
	c += genReturn(f) + ") => void\n"
	return c
}

func genRequest(f *parser.Func) *parser.Struct {
	fs := make([]*parser.Param, len(f.Params))
	for i, p := range f.Params {
		fs[i] = &parser.Param{pascalCase(p.Name), p.Type}
	}
	return &parser.Struct{pascalCase(f.Name) + "In", fs}
}

func genResponse(f *parser.Func) *parser.Struct {
	var fs []*parser.Param
	if f.Return != nil {
		fs = []*parser.Param{&parser.Param{pascalCase(f.Return.Name), f.Return.Type}}
	} else {
		fs = make([]*parser.Param, 0)
	}
	return &parser.Struct{pascalCase(f.Name) + "Out", fs}
}

func genClientStub(f *parser.Func) string {
	c := "\texport function " + camelCase(f.Name) + "("
	args := make([]string, len(f.Params))
	for i, p := range f.Params {
		args[i] = genParam(p)
	}
	c += strings.Join(args, ", ")
	if len(args) > 0 {
		c += ", "
	}
	c += genReturn(f) + "): void {\n"
	c += "\t\tvar req: " + f.Name + "In = {\n"
	for i, p := range f.Params {
		c += "\t\t\t" + snaker.CamelToSnake(p.Name) + ": " + p.Name
		if i < len(f.Params)-1 {
			c += ","
		}
		c += "\n"
	}
	c += "\t\t}\n"

	c += "\t\tProxy.Call(\"" + f.Name + "\", req, function(error, data) {\n"
	if f.Return != nil {
		c += "\t\t\treturn error ? go(error, null) : go(null, (<" + f.Name + "Out>data)." + snaker.CamelToSnake(f.Return.Name) + ")\n"
	} else {
		c += "\t\t\treturn error ? go(error) : go(null)\n"
	}
	c += "\t\t})\n\n"

	c += "\t}\n"

	return c
}

func Generate(i *parser.Interface) string {
	c := "// ----------------------------------\n"
	c += "// --- Generated with go:generate ---\n"
	c += "// ---        DO NOT EDIT         ---\n"
	c += "// ----------------------------------\n\n"

	c += "module Proxy {\n\n"

	if len(i.Aliases) > 0 {
		c += "\t// --- Aliases ---\n\n"
		for _, a := range i.Aliases {
			c += "\texport type " + a.Name + " = " + cast(a.Type) + "\n"
		}
		c += "\n"
	}

	if len(i.Consts) > 0 {
		c += "\t// --- Consts ---\n\n"
		for _, a := range i.Consts {
			for _, l := range a.Values {
				c += "\texport var " + a.Prefix + l + ": " + a.Type + " = " + "\"" + l + "\"\n"
			}
			c += "\n\n"
		}
	}

	if len(i.Structs) > 0 {
		c += "\t// --- Types ---\n\n"
		for _, s := range i.Structs {
			c += genStruct(s, true)
		}
	}

	if len(i.Funcs) > 0 {
		c += "\t// --- Contract ---\n\n"

		c += "\texport interface " + i.Name + " {\n"
		for _, f := range i.Funcs {
			c += "\t\t" + genFunc(f)
		}
		c += "\t}\n\n"

		c += "\t// --- Messages ---\n\n"

		for _, f := range i.Funcs {
			c += genStruct(genRequest(f), false)
			c += genStruct(genResponse(f), false)
		}

		c += "\t// --- Client Stub ---\n\n"

		for _, f := range i.Funcs {
			c += genClientStub(f)
		}
	}

	c += "}\n\n"

	return c
}
