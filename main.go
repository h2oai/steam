package main

import (
	"github.com/h2oai/steamY/cli"
)

var VERSION string
var BUILD_DATE string

func main() {
	cli.Run(VERSION, BUILD_DATE)
}
