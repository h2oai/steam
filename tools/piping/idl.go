package main

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"text/template"
)

type Interface struct {
	Facade  *Facade
	Structs []*Struct
}

type Facade struct {
	Name    string
	Methods []*Method
}

type Method struct {
	Name    string
	Inputs  []*Field
	Outputs []*Field
	Help    string
}

type Struct struct {
	Name   string
	Fields []*Field
}

type Field struct {
	typ          reflect.Type
	Name         string
	Type         string
	IsArray      bool
	IsStruct     bool
	DefaultValue string
	Help         string
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

		fields[i] = &Field{
			ft,
			f.Name,
			ft.Name(),
			isArray,
			isStruct,
			defaultValueOf(ft.Name(), isArray, isStruct),
			help,
		}
	}
	return &Struct{t.Name(), fields}, nil
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

func toInterface(facadeName string, dict map[string]*Struct) (*Interface, error) {
	facade, ok := dict[facadeName]
	if !ok {
		return nil, fmt.Errorf("Could not find facade definition %s", facadeName)
	}

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
		for _, p := range method.Fields {
			if p.Name == "_" { // "_" acts as a separator between input and output parameters
				in = false
				continue
			}
			if in {
				inputs = append(inputs, p)
			} else {
				outputs = append(outputs, p)
			}
			if p.IsStruct {
				if _, ok := structs[p.Type]; !ok {
					s, ok := dict[p.Type]
					if !ok {
						return nil, fmt.Errorf("Could not find parameter definition %s", m.Type)
					}
					structs[p.Type] = s
				}
			}
		}

		methods = append(methods, &Method{m.Name, inputs, outputs, m.Help})
	}

	ks := make([]string, 0)
	for k, _ := range structs {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	types := make([]*Struct, len(ks))
	for i, n := range ks {
		types[i] = structs[n]
	}

	return &Interface{
		&Facade{facade.Name, methods},
		types,
	}, nil
}
