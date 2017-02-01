/*
  Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"text/template"

	"github.com/serenize/snaker"
)

type Interface struct {
	Facade  *Facade
	Structs map[string]*Struct
}

type Facade struct {
	Name    string
	Imports []*Struct
	Methods []*Method
}

type Method struct {
	Name    string
	Inputs  []*Field
	Outputs []*Field
	Help    string
}

type Struct struct {
	Name    string
	HasJSON bool
	Fields  []*Field
}

type Field struct {
	typ          reflect.Type
	Name         string
	Type         string
	IsArray      bool
	IsStruct     bool
	DefaultValue string
	Format       string
	Struct       *Struct
	Help         string
	LogOmit      bool
}

func Define(name string, instance interface{}) (*Interface, error) {
	dict := make(map[string]*Struct)
	err := collate(dict, instance)
	if err != nil {
		return nil, err
	}
	return toInterface(name, dict)
}

func Generate(i interface{}, tmpl string, funcMap map[string]interface{}) ([]byte, error) {
	t, err := template.New("test").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, i)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func collate(dict map[string]*Struct, instance interface{}) error {
	s, err := toStruct(instance)
	if err != nil {
		return err
	}

	dict[s.Name] = s

	for _, f := range s.Fields {
		if !f.IsStruct {
			continue
		}
		if _, ok := dict[f.Type]; ok {
			continue
		}
		v := reflect.New(f.typ)                  // new instance
		p := v.Elem().Interface()                // pointer to new instance
		if err := collate(dict, p); err != nil { // recurse
			return err
		}
	}
	return nil
}

func toStruct(s interface{}) (*Struct, error) {
	hasJSON := false
	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("%v is not a struct", t.Kind())
	}

	fields := make([]*Field, t.NumField())
	for i := 0; i < len(fields); i++ {
		f := t.Field(i)
		if f.Anonymous {
			continue
		}

		ft := f.Type
		isArray := ft.Kind() == reflect.Slice
		isStruct := ft.Kind() == reflect.Struct

		switch ft.Kind() {
		case
			reflect.Slice,
			reflect.Struct,
			reflect.Bool,
			reflect.Int,
			reflect.Uint,
			reflect.Int64,
			reflect.Float32,
			reflect.Float64,
			reflect.String:
			// ok: supported
		default:
			return nil, fmt.Errorf("Unsupported type: %s.%s", t.Name(), f.Name)
		}

		if isArray {
			ft = ft.Elem()
			isStruct = ft.Kind() == reflect.Struct
		}

		help := f.Tag.Get("help")
		if len(help) == 0 {
			help = "No description available"
		}
		var logOmit bool
		logTag := f.Tag.Get("log")
		if len(logTag) != 0 {
			// logOmit = strings.Contains(log, "-")
			logOmit = true
		}

		if strings.HasPrefix(f.Name, "JSON") {
			hasJSON = true
		}

		fields[i] = &Field{
			ft,
			f.Name,
			ft.Name(),
			isArray,
			isStruct,
			defaultValueOf(ft.Name(), isArray, isStruct),
			formatOf(ft.Name(), isArray, isStruct),
			nil,
			help,
			logOmit,
		}
	}
	return &Struct{t.Name(), hasJSON, fields}, nil
}

func defaultValueOf(t string, isArray, isStruct bool) string {
	if isArray || isStruct {
		return "nil"
	}
	switch t {
	case "bool":
		return "false"
	case "string":
		return "\"\""
	default:
		return "0"
	}
}

func formatOf(t string, isArray, isStruct bool) string {
	if isArray || isStruct {
		return "%+v"
	} else {
		return "%v"
	}
}

func toInterface(facadeName string, dict map[string]*Struct) (*Interface, error) {
	facade, ok := dict[facadeName]
	if !ok {
		return nil, fmt.Errorf("Could not find facade definition %s", facadeName)
	}

	imports := make([]*Struct, 0)
	methods := make([]*Method, 0)
	structs := make(map[string]*Struct)

	for _, m := range facade.Fields {
		method, ok := dict[m.Type]
		if !ok {
			return nil, fmt.Errorf("Could not find method definition %s", m.Type)
		}

		inputs := make([]*Field, 0)
		outputs := make([]*Field, 0)

		in := true
		for _, f := range method.Fields {
			if f.Name == "_" { // "_" acts as a separator between input and output parameters
				in = false
				continue
			}
			if in {
				if f.IsStruct {
					if _, ok := structs[f.Type]; !ok {
						s, ok := dict[f.Type]
						if !ok {
							return nil, fmt.Errorf("Could not find parameter definition %s", m.Type)
						}
						imports = append(imports, s)
					}
				}
				inputs = append(inputs, f)
			} else {
				outputs = append(outputs, f)
			}
			if f.IsStruct {
				if _, ok := structs[f.Type]; !ok {
					s, ok := dict[f.Type]
					if !ok {
						return nil, fmt.Errorf("Could not find parameter definition %s", m.Type)
					}
					structs[f.Type] = s
				}
			}
		}

		methods = append(methods, &Method{m.Name, inputs, outputs, m.Help})
	}

	// Store a reference to the actual struct in the field, for downstream pretty-printing.
	for _, m := range methods {
		for _, i := range m.Inputs {
			if i.IsStruct && i.Struct == nil {
				i.Struct = structs[i.Type]
			}
		}
		for _, o := range m.Outputs {
			if o.IsStruct && o.Struct == nil {
				o.Struct = structs[o.Type]
			}
		}
	}

	ks := make([]string, 0)
	for k, _ := range structs {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	types := make(map[string]*Struct, len(ks))
	for _, n := range ks {
		types[structs[n].Name] = structs[n]
	}

	return &Interface{
		&Facade{facade.Name, imports, methods},
		types,
	}, nil
}

func multiMarshal(typ, outputName, fieldName string) string {
	return fmt.Sprintf(`aux := make([]%s, len(out.%s))
for i, val := range out.%s {
	aux[i] = *val
	aux[i].%s = "JSON DATA OMITTED..."
}

res, merr := json.Marshal(aux)`, typ, outputName, outputName, fieldName)
}

func singlMarshal(output, field string) string {
	return fmt.Sprintf(`aux := *out.%s
aux.%s = "JSON DATA OMITTED..."

res, merr := json.Marshal(aux)`, output, field)
}

func cleanJSON(outputs []*Field) string {
	// Search outputs
	for _, output := range outputs {
		// Only structs contain JSON fields
		if output.IsStruct {
			// All JSON fields are prepended with the text `JSON`
			for _, field := range output.Struct.Fields {
				if strings.HasPrefix(field.Name, "JSON") {
					// Handle arrays differently than values
					if output.IsArray {
						return multiMarshal(output.Type, output.Name, field.Name)
					} else {
						return singlMarshal(output.Name, field.Name)
					}
				}
			}
		}
	}

	// If all else fails, return default output
	return `res, merr := json.Marshal(out)`
}

func createMaps(field *Field, prefix string) string {
	if field.IsStruct && field.Struct != nil {
		if field.IsArray {
			return mapArrayFromStruct(field.Struct, field.Name, prefix)
		}
		return mapFromStruct(field.Struct, fmt.Sprintf("%s.%s", prefix, field.Name))
	}
	return fmt.Sprintf("%s.%s", prefix, field.Name)
}

var structMap = `map[string]interface{}{
	{{- range .Fields}}{{- if not .LogOmit}}
	"{{snake .Name}}": {{createMaps . "%s"}},
	{{- end}}{{- end}}
}`

func mapFromStruct(s *Struct, prefix string) string {
	tmpl := template.Must(
		template.New(
			"mapper",
		).Funcs(template.FuncMap{
			"snake":      snaker.CamelToSnake,
			"createMaps": createMaps,
		}).Parse(
			fmt.Sprintf(structMap, prefix),
		))

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return buf.String()
}

var structArrayMap = `make([]map[string]interface{}, len(%s.%s))
	for i, v := range %s.%s {
		tmp := make(map[string]interface{})
		{{- range .Fields}}{{- if not .LogOmit}}
		tmp["{{snake .Name}}"] = v.{{.Name}}
		{{- end}}{{- end}}
		json%s["%s"].([]map[string]interface{})[i] = tmp
	}
`

func mapArrayFromStruct(s *Struct, name, prefix string) string {
	tmpl := template.Must(
		template.New(
			"arraymapper",
		).Funcs(template.FuncMap{
			"snake":      snaker.CamelToSnake,
			"createMaps": createMaps,
		}).Parse(
			fmt.Sprintf(structArrayMap, prefix, name, prefix, name, strings.Title(prefix), snaker.CamelToSnake(name)),
		))

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return buf.String()
}
