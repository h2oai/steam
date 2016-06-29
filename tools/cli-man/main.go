package main

import (
	"github.com/h2oai/steamY/cli"
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

	header := &cobra.GenManHeader{
		Title:   "STEAM",
		Section: "1",
		Manual:  "Steam CLI Manual",
		Source:  "(c) 2016 H2O.ai",
	}

	steam.GenManTree(header, p)
}
