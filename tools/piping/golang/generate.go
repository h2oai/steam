package golang

import (
	"strings"

	"github.com/h2oai/steamY/tools/piping/parser"
	"github.com/serenize/snaker"
)

func genParam(p *parser.Param) string {
	return p.Name + " " + p.Type
}

func genMember(p *parser.Param) string {
	return p.Name + " " + p.Type + " `json:\"" + snaker.CamelToSnake(p.Name) + "\"`"
}

func genStruct(s *parser.Struct) string {
	c := "type " + s.Name + " struct {\n"
	for _, p := range s.Params {
		c += "\t" + genMember(p) + "\n"
	}
	return c + "}\n\n"
}

func genFunc(f *parser.Func) string {
	c := "\t" + f.Name + "(pz az.Principal,"
	l := len(f.Params) - 1
	for i, p := range f.Params {
		c += genParam(p)
		if i < l {
			c += ", "
		}
	}
	c += ") "
	if f.Return != nil {
		c += "(" + f.Return.Type + ", error)\n"
	} else {
		c += "error\n"
	}
	return c
}
func capitalize(s string) string {
	switch len(s) {
	case 0:
		return ""
	case 1:
		return strings.ToUpper(s)
	default:
		return strings.ToUpper(string(s[0])) + s[1:]
	}
}

func genRequest(f *parser.Func) *parser.Struct {
	fs := make([]*parser.Param, len(f.Params))
	for i, p := range f.Params {
		fs[i] = &parser.Param{capitalize(p.Name), p.Type}
	}
	return &parser.Struct{capitalize(f.Name) + "In", fs}
}

func genResponse(f *parser.Func) *parser.Struct {
	var fs []*parser.Param
	if f.Return != nil {
		fs = []*parser.Param{&parser.Param{capitalize(f.Return.Name), f.Return.Type}}
	} else {
		fs = make([]*parser.Param, 0)
	}
	return &parser.Struct{capitalize(f.Name) + "Out", fs}
}

func genClientDefs() string {
	return `type Remote struct {
	Proc Proc
}

type Proc interface {
	Call(name string, in, out interface{}) error
}

`
}

func genServerDefs(i *parser.Interface) string {
	c := "type Impl struct {\n"
	c += "\tService " + i.Name + "\n\tAz az.Az\n}\n\n"
	return c
}

func genClientStub(f *parser.Func) string {
	c := "func (this *Remote) " + f.Name + "("
	args := make([]string, len(f.Params))
	for i, p := range f.Params {
		args[i] = genParam(p)
	}
	c += strings.Join(args, ", ") + ") ("
	if f.Return != nil {
		c += f.Return.Type + ", "
	}
	c += " error) {\n"
	c += "\tin := " + f.Name + "In{"
	params := make([]string, len(f.Params))
	for i, p := range f.Params {
		params[i] = p.Name
	}
	c += strings.Join(params, ", ") + "}\n"
	c += "\tvar out " + f.Name + "Out\n"
	c += "\terr := this.Proc.Call(\"" + f.Name + "\", &in, &out)\n"
	c += "\tif err != nil {\n"
	c += "\t\treturn "
	if f.Return != nil {
		c += defaultOf(f.Return.Type) + ", "
	}
	c += "err\n\t}\n"
	c += "\treturn "
	if f.Return != nil {
		c += "out." + capitalize(f.Return.Name) + ", "
	}
	c += "nil\n}\n\n"
	return c
}

func genServerStub(f *parser.Func) string {
	c := "func (this *Impl) " + f.Name + "(r *http.Request, in *" + f.Name + "In, out *" + f.Name + "Out) error {\n\t"

	c += `
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	`
	c += "log.Println(pz, \"Called " + f.Name + "\")\n\n"

	if f.Return != nil {
		c += "it, "
	}
	c += "err := this.Service." + f.Name + "(pz,"
	if len(f.Params) != 0 {
		params := make([]string, len(f.Params))
		for i, p := range f.Params {
			params[i] = "in." + capitalize(p.Name)
		}
		c += strings.Join(params, ", ")
	}
	c += ")\n\tif err != nil {\n"
	c += "\t\tlog.Printf(\"%s Failed to " + f.Name + ": %v\", pz, err)\n"
	c += "\t\treturn err\n\t}\n"
	if f.Return != nil {
		c += "\tout." + capitalize(f.Return.Name) + " = it\n"
	}
	c += "\treturn nil\n}\n\n"

	return c
}

func defaultOf(t string) string {
	switch t {
	case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "float32", "float64", "int", "uint":
		return "0"
	case "string":
		return "\"\""
	case "bool":
		return "false"
	default:
		return "nil"
	}
}

func Generate(i *parser.Interface) string {
	c := "// ----------------------------------\n"
	c += "// --- Generated with go:generate ---\n"
	c += "// ---        DO NOT EDIT         ---\n"
	c += "// ----------------------------------\n\n"
	c += "package " + i.Package + "\n\n"

	if len(i.Imports) > 0 {
		c += "import(\n"
		for _, a := range i.Imports {
			c += "\t\"" + a + "\"\n"
		}
		c += ")\n\n"
	}

	if len(i.Aliases) > 0 {
		c += "// --- Aliases ---\n\n"
		for _, a := range i.Aliases {
			c += "type " + a.Name + " " + a.Type + "\n"
		}

		c += "\n"
	}

	if len(i.Consts) > 0 {
		c += "// --- Consts ---\n\n"
		for _, a := range i.Consts {
			c += "const (\n"
			for i, l := range a.Values {
				c += "\t" + a.Prefix + l
				if i == 0 {
					c += " " + a.Type
				}
				c += " = \"" + l + "\"\n"
			}
			c += ")\n\n"
		}
	}

	if len(i.Structs) > 0 {
		c += "// --- Types ---\n\n"
		for _, s := range i.Structs {
			c += genStruct(s)
		}
	}

	if len(i.Funcs) > 0 {
		c += "// --- Interfaces ---\n\n"

		c += `
		type Az interface {
			Identify(r *http.Request) (az.Principal, error)
		}

		`

		c += "type " + i.Name + " interface {\n"
		for _, f := range i.Funcs {
			c += genFunc(f)
		}
		c += "}\n\n"

		c += "// --- Messages ---\n\n"

		for _, f := range i.Funcs {
			c += genStruct(genRequest(f))
			c += genStruct(genResponse(f))
		}

		c += "// --- Client Stub ---\n\n"

		c += genClientDefs()
		for _, f := range i.Funcs {
			c += genClientStub(f)
		}

		c += "// --- Server Stub ---\n\n"

		c += genServerDefs(i)
		for _, f := range i.Funcs {
			c += genServerStub(f)
		}
	}

	return c
}
