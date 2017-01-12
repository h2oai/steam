package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

// Struct for walking through top level nodes
type top struct {
	pName   string
	genDecl *decl
}

func NewTop() *top {
	return &top{
		genDecl: &decl{tables: make([]Table, 0, 8)},
	}
}

func (t *top) Visit(n ast.Node) ast.Visitor {
	switch x := n.(type) {
	case *ast.File: // Top Level file: continue with t
		t.pName = x.Name.Name
		return t
	case *ast.GenDecl: // Top Level GenDecl
		if x.Tok == token.TYPE { // Top Level struct: continue with genDecl
			return t.genDecl
		}
	}

	return nil
}

// Struct for walking through ast.GenDecl
type decl struct {
	tables []Table
}

func (d *decl) Visit(n ast.Node) ast.Visitor {
	switch x := n.(type) {
	case *ast.TypeSpec:
		if y, ok := x.Type.(*ast.StructType); ok {
			// log.Println(x.Name.Name)
			var t Table
			t.Name = x.Name.Name
			t.IsExported = ast.IsExported(x.Name.Name)
			ast.Walk(&t, y) // Start new Walk to generate columns

			d.tables = append(d.tables, t)
		}
	}

	return nil
}

func parseFiles(filePaths ...string) (Package, error) {

	t := NewTop()
	for _, filePath := range filePaths {
		fset := token.NewFileSet() // Positions are relative to fset
		f, err := parser.ParseFile(fset, filePath, nil, 0)
		if err != nil {
			return Package{}, err
		}

		// ast.Print(fset, f)
		ast.Walk(t, f)
	}

	return Package{Name: t.pName, Tables: t.genDecl.tables}, nil
}

func (t *Table) Visit(n ast.Node) ast.Visitor {
	switch x := n.(type) {
	case *ast.StructType: // Struct Top level: continue
		return t
	case *ast.FieldList: // Struct Fields: continue
		t.Cols = make([]Col, 0, x.NumFields())
		return t
	case *ast.Field:
		cols := make([]Col, len(x.Names))
		for i, field := range x.Names {
			var c Col
			c.Name = strings.TrimSpace(field.Name)
			var isPk bool
			c.TblName, c.Default, c.IsArg, isPk = parseTag(x.Tag)
			if c.TblName == "" {
				c.TblName = c.Name
			}
			if isPk {
				t.Key = c.TblName
			}
			c.Type = parseType(x.Type)

			cols[i] = c
		}

		t.Cols = append(t.Cols, cols...)

	case nil:
	default:
		panic(fmt.Sprintf("unhandled type: %T", x))
	}

	return nil
}

func parseTag(tag *ast.BasicLit) (string, string, bool, bool) {
	var (
		name, def   string
		isArg, isPk bool
	)

	if tag != nil {
		vals := splitTags(tag.Value)
		// log.Println(vals)
		for _, val := range vals {
			s := strings.TrimSpace(val)
			switch {
			case strings.HasPrefix(s, "def="):
				def = strings.TrimPrefix(s, "def=")
			case s == "arg":
				isArg = true
			case s == "pk":
				isPk = true
			case !strings.Contains(s, "="):
				name = s
			}
		}
	}

	return name, def, isArg, isPk
}

func splitTags(tags string) []string {
	x := strings.Index(tags, "db:")
	cont := strings.Split(strings.Trim(tags[x+3:], "`"), ":")
	if len(cont) > 1 {
		cont = []string{strings.TrimRightFunc(cont[0], func(r rune) bool { return r != '"' })}
	}
	vals := cont[0]

	return strings.Split(vals[1:len(vals)-1], ",")
}

func parseType(typ ast.Expr) string {
	switch x := typ.(type) {
	case *ast.Ident:
		return strings.TrimSpace(x.Name)
	case *ast.SelectorExpr:
		if y, ok := x.X.(*ast.Ident); ok {
			return strings.TrimSpace(fmt.Sprintf("%s.%s", y.Name, x.Sel.Name))
		} else {
			panic(fmt.Sprintf("unhandled selector: %s", y.Name))
		}

	default:
		panic(fmt.Sprintf("unhandled field type: %T", x))
	}

	return ""
}
