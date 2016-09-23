package main

import (
	cli "github.com/h2oai/steam/cli2"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

func main() {
	steam := cli.Steam("", "", ioutil.Discard, ioutil.Discard, ioutil.Discard)

	var p string
	if len(os.Args) > 1 {
		p = os.Args[1]
	} else {
		p = "./"
	}

	cobra.GenMarkdownTree(steam, p)
}
