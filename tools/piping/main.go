package main

import (
	"fmt"
	"github.com/h2oai/steamY/srv/web/api"
	"github.com/serenize/snaker"
	"io/ioutil"
	"strings"
)

func main() {
	ix, err := Define("Service", &api.Service{})
	if err != nil {
		panic(err)
	}

	generate(ix, "srv/web/api/go.template", "srv/web/service.go", map[string]interface{}{
		"lower": lower,
		"snake": snaker.CamelToSnake,
	})
	generate(ix, "srv/web/api/typescript.template", "gui/src/proxy.ts", map[string]interface{}{
		"lower":   lower,
		"snake":   snaker.CamelToSnake,
		"js_type": jsTypeOf,
	})
	generate(ix, "srv/web/api/python.template", "python/steam.py", map[string]interface{}{
		"lower": lower,
		"snake": snaker.CamelToSnake,
	})

	methodGroups := toCLIMethodGroups(ix)
	generate(methodGroups, "srv/web/api/cli.template", "cli2/cli.go", map[string]interface{}{
		"lower": lower,
		"upper": upper,
		"snake": snaker.CamelToSnake,
		"flag":  toFlagName,
	})
}

func generate(ix interface{}, input, output string, funcMap map[string]interface{}) {
	fmt.Println(input, "-->", output)

	tmpl, err := ioutil.ReadFile(input)
	if err != nil {
		panic(err)
	}

	code, err := Generate(ix, string(tmpl), funcMap)
	if err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(output, code, 0644); err != nil {
		panic(err)
	}
}

func lower(s string) string {
	switch len(s) {
	case 0:
		return ""
	case 1:
		return strings.ToLower(s)
	default:
		return strings.ToLower(string(s[0])) + s[1:]
	}
}

func upper(s string) string {
	switch len(s) {
	case 0:
		return ""
	case 1:
		return strings.ToUpper(s)
	default:
		return strings.ToUpper(string(s[0])) + s[1:]
	}
}

func jsTypeOf(t string) string {
	switch t {
	case "bool":
		return "boolean"
	case
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"uintptr",
		"float32",
		"float64":
		return "number"
	default:
		return t
	}
}
