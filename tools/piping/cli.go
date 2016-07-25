package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type CLIMethodGroup struct {
	Func     string
	Verb     string
	Help     string
	Variants []*CLIMethodVariant
}

type CLIMethodVariant struct {
	Verb    string
	Noun    string
	Help    string
	Methods []*CLIMethod
	Inputs  []*Field
	HasFlag bool
}

type CLIMethod struct {
	Verb    string
	Noun    string
	Switch  string
	Help    string
	Flag    string
	HasFlag bool
	Args    []string
	Method  *Method
}

var camelCaseTokenizer = regexp.MustCompile("[A-Z][a-z0-9]+")

func tokenizeCamelCase(name string) []string {
	return camelCaseTokenizer.FindAllString(name, -1)
}

func toCLIMethod(method *Method) *CLIMethod {
	tokens := tokenizeCamelCase(method.Name)
	var verb, noun, switch_, flag string
	var hasFlag bool
	switch len(tokens) {
	case 0:
		panic(fmt.Sprintf("Could not tokenize method name %s", method.Name))
	case 1:
		verb = tokens[0]
	case 2:
		verb = tokens[0]
		noun = tokens[1]
	default:
		verb = tokens[0]
		noun = tokens[1]
		switch_ = strings.Join(tokens[2:], "")
		flag = strings.Join(tokens[2:], "-")
		hasFlag = true
	}
	args := make([]string, 0)
	for _, input := range method.Inputs {
		args = append(args, toFlagName(input.Name))
	}
	return &CLIMethod{
		strings.ToLower(verb),
		strings.ToLower(noun),
		switch_,
		method.Help,
		strings.ToLower(flag),
		hasFlag,
		args,
		method,
	}
}

func sanitize(name string) string {
	switch name {
	case "import", "delete":
		return name + "_"
	default:
		return name
	}
}

func collateInputs(methods []*CLIMethod) []*Field {
	dict := make(map[string]*Field)
	for _, method := range methods {
		for _, field := range method.Method.Inputs {
			f, ok := dict[field.Name]
			if ok && f.Type != field.Type {
				panic(fmt.Sprintf("Field type mismatch: %s != %s", f.Type, field.Type))
			}
			dict[field.Name] = field
		}
	}

	inputs := make([]*Field, 0)
	for _, v := range dict {
		inputs = append(inputs, v)
	}

	return inputs
}

func groupMethods(methods []*CLIMethod) []*CLIMethodGroup {
	dict := make(map[string]map[string][]*CLIMethod)
	for _, m := range methods {
		dict[m.Verb] = make(map[string][]*CLIMethod)
	}
	for _, m := range methods {
		dict[m.Verb][m.Noun] = make([]*CLIMethod, 0)
	}
	for _, m := range methods {
		dict[m.Verb][m.Noun] = append(dict[m.Verb][m.Noun], m)
	}

	groupNames := make([]string, 0)
	for verb, _ := range dict {
		groupNames = append(groupNames, verb)
	}
	sort.Strings(groupNames)

	groups := make([]*CLIMethodGroup, len(groupNames))
	for i, verb := range groupNames {
		d := dict[verb]

		methodNames := make([]string, 0)
		for noun, _ := range d {
			methodNames = append(methodNames, noun)
		}
		sort.Strings(methodNames)

		variants := make([]*CLIMethodVariant, len(methodNames))
		for j, noun := range methodNames {
			methods := d[noun]
			hasFlag := false
			for _, m := range methods {
				if m.HasFlag {
					hasFlag = true
					break
				}
			}
			variants[j] = &CLIMethodVariant{
				verb,
				noun,
				upper(verb) + " " + upper(noun),
				methods,
				collateInputs(methods),
				hasFlag,
			}
		}
		groups[i] = &CLIMethodGroup{
			sanitize(verb),
			verb,
			upper(verb) + " entities",
			variants,
		}
	}

	return groups
}

func toFlagName(name string) string {
	tokens := tokenizeCamelCase(name)
	return strings.ToLower(strings.Join(tokens, "-"))
}

func toCLIMethodGroups(ix *Interface) []*CLIMethodGroup {
	f := ix.Facade
	methods := make([]*CLIMethod, 0)
	for _, m := range f.Methods {
		portable := true

		for _, input := range m.Inputs {
			if input.IsArray || input.IsStruct {
				portable = false
				break
			}
		}

		if portable {
			methods = append(methods, toCLIMethod(m))
		} else {
			fmt.Println("Skipping method (unportable): ", m.Name)
		}
	}
	return groupMethods(methods)
}
