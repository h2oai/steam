package parser

import (
	"fmt"
	"regexp"
	"strings"
)

type Param struct {
	Name string
	Type string
}

type Const struct {
	Type   string
	Prefix string
	Values []string
}

type Struct struct {
	Name   string
	Params []*Param
}

type Func struct {
	Name   string
	Params []*Param
	Return *Param
}

type Interface struct {
	Package string
	Imports []string
	Name    string
	Aliases []*Param
	Consts  []*Const
	Structs []*Struct
	Funcs   []*Func
}

func parseAttr(line string) (string, error) {
	p := strings.Fields(line)
	if len(p) != 2 {
		return "", fmt.Errorf("Bad attr declaration: [%s]", line)
	}
	return p[1], nil
}

func parseParam(line string) (*Param, error) {
	p := strings.Fields(line)
	if len(p) != 2 {
		return nil, fmt.Errorf("Bad param declaration: [%s]", line)
	}
	return &Param{p[0], p[1]}, nil
}

func parseAlias(line string) (*Param, error) {
	p := strings.Fields(line)
	if len(p) != 3 {
		return nil, fmt.Errorf("Bad alias declaration: [%s]", line)
	}
	return &Param{p[1], p[2]}, nil
}

func parseConst(line string) (*Const, error) {
	fs := strings.Fields(line)
	if len(fs) < 4 {
		return nil, fmt.Errorf("Bad const declaration: [%s]", line)
	}
	vs := make([]string, len(fs)-3)
	for i := 3; i < len(fs); i++ {
		vs[i-3] = fs[i]
	}
	return &Const{fs[1], fs[2], vs}, nil
}

func parseStruct(line string) (*Struct, error) {
	fs := strings.Fields(line)
	if len(fs) != 2 {
		return nil, fmt.Errorf("Bad struct declaration: [%s]", line)
	}
	return &Struct{fs[1], make([]*Param, 0)}, nil
}

var funcRegexp = regexp.MustCompile("^(.+?)\\s*\\((.*?)\\)(.*)$")

func parseFunc(line string) (*Func, error) {
	sm := funcRegexp.FindStringSubmatch(line)
	if len(sm) != 4 {
		return nil, fmt.Errorf("Bad func declaration: [%s]", line)
	}
	n := sm[1]

	rt := strings.TrimSpace(sm[3])

	var r *Param
	var err error

	if len(rt) != 0 {
		r, err = parseParam(strings.TrimSpace(sm[3]))
		if err != nil {
			return nil, err
		}
	}

	parts := strings.Split(sm[2], ",")
	ps := make([]*Param, 0)
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if len(part) > 0 {
			p, err := parseParam(part)
			if err != nil {
				return nil, err
			}
			ps = append(ps, p)
		}
	}
	return &Func{n, ps, r}, nil
}

func Parse(idl string) (*Interface, error) {
	lines := strings.Split(idl, "\n")

	var f *Func
	var s *Struct
	var pkg string
	var svc string
	var err error
	imports := make([]string, 0)
	aliases := make([]*Param, 0)
	consts := make([]*Const, 0)
	funcs := make([]*Func, 0)
	structs := make([]*Struct, 0)

	collect := func() {
		if s != nil {
			structs = append(structs, s)
			s = nil
		}
		if f != nil {
			funcs = append(funcs, f)
			f = nil
		}
	}

	for _, line := range lines {
		t := strings.TrimSpace(line)
		if len(t) == 0 { // empty
			continue
		}
		if strings.HasPrefix(t, "#") { // comment
			continue
		}
		if strings.HasPrefix(line, " ") { // struct field
			if s == nil {
				return nil, fmt.Errorf("Bad declaration: %s", t)
			}
			p, err := parseParam(strings.TrimSpace(t))
			if err != nil {
				return nil, err
			}
			s.Params = append(s.Params, p)
		} else {
			collect()
			if strings.HasPrefix(t, "type") {
				if s, err = parseStruct(t); err != nil {
					return nil, err
				}
			} else if strings.HasPrefix(t, "package") {
				if pkg, err = parseAttr(t); err != nil {
					return nil, err
				}
			} else if strings.HasPrefix(t, "service") {
				if svc, err = parseAttr(t); err != nil {
					return nil, err
				}
			} else if strings.HasPrefix(t, "alias") {
				a, err := parseAlias(t)
				if err != nil {
					return nil, err
				}
				aliases = append(aliases, a)
			} else if strings.HasPrefix(t, "const") {
				cs, err := parseConst(t)
				if err != nil {
					return nil, err
				}
				consts = append(consts, cs)
			} else if strings.HasPrefix(t, "import") {
				imp, err := parseAttr(t)
				if err != nil {
					return nil, err
				}
				imports = append(imports, imp)
			} else {
				if f, err = parseFunc(t); err != nil {
					return nil, err
				}
			}
		}
	}
	collect()

	// if we have funcs, import log and net/http for server stubs
	if len(funcs) > 0 {
		imports = append(imports, "log", "net/http")
	}

	return &Interface{pkg, imports, svc, aliases, consts, structs, funcs}, nil
}
