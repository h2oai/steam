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
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/h2oai/steam/srv/web/api"
	"github.com/serenize/snaker"
)

func main() {
	ix, err := Define("Service", &api.Service{})
	if err != nil {
		panic(err)
	}

	generate(ix, "srv/web/api/go.template", "srv/web/service.go", map[string]interface{}{
		"lower":      lower,
		"snake":      snaker.CamelToSnake,
		"startswith": strings.HasPrefix,
		"cleanJSON":  cleanJSON,
		"createMaps": createMaps,
	})
	generate(ix, "srv/web/api/typescript.template", "gui/src/Proxy/Proxy.ts", map[string]interface{}{
		"lower":   lower,
		"snake":   snaker.CamelToSnake,
		"js_type": jsTypeOf,
	})
	generate(ix, "srv/web/api/typescript-cli.template", "gui/src/Proxy/CLI.ts", map[string]interface{}{
		"lower":   lower,
		"snake":   snaker.CamelToSnake,
		"js_type": jsTypeOf,
	})
	generate(ix, "srv/web/api/python.template", "python/backend.py", map[string]interface{}{
		"lower": lower,
		"snake": snaker.CamelToSnake,
	})
	generate(ix, "srv/web/api/r.template", "r/backend.r", map[string]interface{}{
		"rType": rType,
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
