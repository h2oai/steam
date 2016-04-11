package main

import (
	"flag"
	"fmt"
	"github.com/h2oai/steamY/tools/piping/golang"
	"github.com/h2oai/steamY/tools/piping/parser"
	"github.com/h2oai/steamY/tools/piping/typescript"
	"io/ioutil"
)

func main() {

	// Usage: piping -idl widget.pipe -go widget.go -ts widget.ts

	file := flag.String("idl", "service.pipe", "Path to IDL file")
	goDest := flag.String("go", "", "Output file name for Go")
	tsDest := flag.String("ts", "", "Output file name for Typescript")

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
		if err = ioutil.WriteFile(*goDest, []byte(golang.Generate(i)), 0644); err != nil {
			panic(fmt.Sprintf("Error writing Go output: %s: %s", *goDest, err))
		}

		fmt.Println("Go service definition created:", *goDest)
	}

	if *tsDest != "" {
		if err = ioutil.WriteFile(*tsDest, []byte(typescript.Generate(i)), 0644); err != nil {
			panic(fmt.Sprintf("Error writing Typescript output: %s: %s", *tsDest, err))
		}

		fmt.Println("Typescript service definition created:", *tsDest)
	}

}
