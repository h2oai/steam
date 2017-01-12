package main

import (
	"fmt"
	"testing"
)

const (
	filePath = "examples/tables.go"
)

func TestParse(t *testing.T) {
	pkg, err := parseFiles(filePath)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	fmt.Println(pkg)
}
