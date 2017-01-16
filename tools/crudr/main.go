package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"

	"github.com/serenize/snaker"
)

type Package struct {
	Name   string
	Tables []Table
}

type Table struct {
	Name       string
	Key        string
	Cols       []Col
	IsExported bool
}

type Col struct {
	Name    string
	TblName string
	Type    string
	Default string
	IsArg   bool
}

func main() {
	flag.Parse()

	files := flag.Args()

	filePaths := make([]string, len(files))
	for i, file := range files {
		var err error
		filePaths[i], err = filepath.Abs(file)
		if err != nil {
			panic(err)
		}

	}

	pkg, err := parseFiles(filePaths...)
	if err != nil {
		panic(err)
	}

	funcMap := template.FuncMap{
		"strLen":       strLen,
		"toArgs":       toArgs,
		"title":        strings.Title,
		"lowerFirst":   lowerFirst,
		"toFieldVar":   toFieldVar,
		"toPluralName": toPluralName,
		"camelToSnake": snaker.CamelToSnake,
	}

	tmpl, err := template.New("crud").Funcs(funcMap).Parse(src)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(buf, "main", pkg); err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile("crud.go", buf.Bytes(), 0644); err != nil {
		panic(err)
	}
}

func isTitle(str string) bool {
	r, _ := utf8.DecodeRuneInString(str)
	return unicode.IsUpper(r)
}

func lowerFirst(str string) string {
	r := []rune(str)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}

func strLen(str string) []struct{} {
	return make([]struct{}, len(str))
}

func toArgs(cols []Col) string {
	var args, curType string
	for _, col := range cols {
		if col.IsArg {
			switch {
			case strings.HasPrefix(col.Type, "sql.Null"):
				col.Type = lowerFirst(col.Type[8:])
			case col.Type == "pq.NullTime":
				col.Type = "time.Time"
			}
			change := curType != col.Type
			if change && args != "" {
				args += " " + curType
			}
			curType = col.Type

			if args != "" {
				args += ", "
			}
			args += toFieldVar(col.Name)
		}
	}
	if args != "" {
		args += " " + curType + ", "
	}

	return args
}

func toFieldVar(str string) string {
	val := lowerFirst(str)

	switch val {
	case "type":
		return "typ"
	}

	return val
}

func toPluralName(str string) string {
	if strings.HasSuffix(str, "y") {
		return strings.TrimRight(str, "y") + "ie"
	}

	return str
}
