package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/h2oai/steam/tools/piping/parser"
	"io/ioutil"
	"strings"
	"unicode"
)

var header = `
// ----------------------------------
// --- Generated with go:generate ---
// ---        DO NOT EDIT         ---
// ----------------------------------

package db

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"time"
)

`

var codeTemplate = `
func (ds *DS) HasStruct(id string) (bool, error) {
	has := false
	err := ds.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Struct"))
		if b == nil {
			return fmt.Errorf("Bucket Struct does not exist.")
		}

		v := b.Get([]byte(id))
		if v != nil {
			has = true
		}
		return nil
	})
	return has, err
}

func (ds *DS) HasStructs(ids []string) (bool, error) {
	has := true
	err := ds.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Struct"))
		if b == nil {
			return fmt.Errorf("Bucket Struct does not exist.")
		}

		for _, id := range ids {
			v := b.Get([]byte(id))
			if v == nil {
				has = false
				return nil
			}
		}
		return nil
	})
	return has, err
}

func readStructs(tx *bolt.Tx, ids []string) ([]*Struct, error) {
	objs := make([]*Struct, len(ids))
	b := tx.Bucket([]byte("Struct"))
	if b == nil {
		return nil, fmt.Errorf("Bucket Struct does not exist.")
	}

	for i, id := range ids {
		v := b.Get([]byte(id))
		if v == nil {
			continue
		}

		o, err := DecodeStruct(v)
		if err != nil {
			return nil, err
		}

		objs[i] = o
	}

	return objs, nil
}

func (ds *DS) ReadStructs(ids []string) ([]*Struct, error) {
	var objs []*Struct
	err := ds.db.View(func(tx *bolt.Tx) error {
		os, err := readStructs(tx, ids)
		if err != nil {
			return err
		}
		objs = os
		return nil
	})
	return objs, err
}

func (ds *DS) ReadStruct(id string) (*Struct, error) {
	var obj *Struct
	err := ds.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Struct"))
		if b == nil {
			return fmt.Errorf("Bucket Struct does not exist.")
		}

		v := b.Get([]byte(id))
		if v == nil {
			return nil
		}
		
		o, err := DecodeStruct(v)
		if err != nil {
			return err
		}

		obj = o

		return nil
	})
	return obj, err
}

func (ds *DS) CreateStruct(o *Struct) error {
	return ds.writeStruct(o, true)
}

func (ds *DS) UpdateStruct(o *Struct) error {
	return ds.writeStruct(o, false)
}

func (ds *DS) writeStruct(o *Struct, create bool) error {
	err := ds.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Struct"))
		if b == nil {
			return fmt.Errorf("Bucket Struct does not exist.")
		}

		o.ModifiedAt = time.Now().UTC().Unix()

		v, err := EncodeStruct(o)
		if err != nil {
			return err
		}

		err = b.Put([]byte(o.ID), v)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("Struct write failed: %v", err)
	}

	return nil
}

func (ds *DS) DeleteStruct(id string) error {
	err := ds.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Struct"))
		if b == nil {
			return fmt.Errorf("Bucket Struct does not exist.")
		}
		return b.Delete([]byte(id))
	})
	if err != nil {
		return fmt.Errorf("Struct delete failed: %v", err)
	}
	return nil
}

func (ds *DS) ListStruct() ([]*Struct, error) {
	var objs []*Struct
	err := ds.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Struct"))
		if b == nil {
			return fmt.Errorf("Bucket Struct does not exist.")
		}

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			o, err := DecodeStruct(v)
			if err != nil {
				return err
			}
			objs = append(objs, o)
		}
		return nil
	})
	return objs, err
}

func EncodeStruct(o *Struct) ([]byte, error) {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(o)
	if err != nil {
		return nil, fmt.Errorf("Struct encode failed: %v", err)
	}
	return b.Bytes(), nil
}

func DecodeStruct(b []byte) (*Struct, error) {
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	var o Struct
	err := dec.Decode(&o)
	if err != nil {
		return nil, fmt.Errorf("Struct decode failed: %v", err)
	}
	return &o, nil
}

func (ds *DS) PrintStruct(b []byte) (string, error) {
	o, err := DecodeStruct(b)
	if err != nil {
		return "", err
	}

	js, err := json.MarshalIndent(o, "", " ")
	if err != nil {
		return "", err
	}

	return string(js), nil
}

`

func genStruct(s *parser.Struct) string {
	c := "type " + s.Name + " struct {\n\t*Record\n"
	for _, p := range s.Params {
		c += "\t" + p.Name + " " + p.Type + "\n"
	}
	return c + "}\n\n"
}

func camelCase(s string) string {
	a := []rune(s)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}

func genCtor(s *parser.Struct) string {
	c := "func New" + s.Name + "(id string,"
	for _, p := range s.Params {
		c += " " + camelCase(p.Name) + " " + p.Type + ","
	}
	c += ") *" + s.Name + " {\nreturn &" + s.Name + "{\n"
	c += `&Record{
			id,
			time.Now().UTC().Unix(),
			0,
		},
		`
	for _, p := range s.Params {
		c += camelCase(p.Name) + ",\n"
	}
	return c + "}\n}\n\n"
}

func generate(i *parser.Interface) string {
	const find = "Struct"

	schemas := make([]string, len(i.Structs))
	for i, s := range i.Structs {
		schemas[i] = s.Name
	}

	var buf bytes.Buffer
	buf.WriteString(header)
	buf.WriteString("func (ds *DS) Init() {\n")
	for _, s := range schemas {
		buf.WriteString("\t" + `Printers["` + s + `"] = ds.Print` + s + "\n")
	}
	buf.WriteString("}\n")

	buf.WriteString("var Buckets = []string{\n")
	for _, s := range schemas {
		buf.WriteString("\t" + `"` + s + `",` + "\n")
	}
	buf.WriteString("}\n")

	for _, s := range i.Structs {
		buf.WriteString(genStruct(s))
		buf.WriteString(genCtor(s))
	}

	for _, s := range schemas {
		buf.WriteString(strings.Replace(codeTemplate, find, s, -1))
	}
	return buf.String()
}

func main() {
	// Usage: icing -idl schema.pipe -go widget.go
	file := flag.String("idl", "service.pipe", "Path to IDL file")
	goDest := flag.String("go", "", "Output file name for Go")

	flag.Parse()

	idl, err := ioutil.ReadFile(*file)
	if err != nil {
		panic(fmt.Sprintf("Error opening IDL: %s: %s", *file, err))
	}

	i, err := parser.Parse(string(idl))
	if err != nil {
		panic(fmt.Sprintf("Error parsing IDL: %s", err))
	}

	fmt.Println()

	if *goDest != "" {
		if err = ioutil.WriteFile(*goDest, []byte(generate(i)), 0644); err != nil {
			panic(fmt.Sprintf("Error writing Go output: %s: %s", *goDest, err))
		}

		fmt.Println("Go schema definition created:", *goDest)
	}

}
