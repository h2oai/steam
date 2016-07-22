package main

import (
	cli "github.com/h2oai/steamY/cli2"
)

var VERSION string
var BUILD_DATE string

func main() {
	cli.Run(VERSION, BUILD_DATE)
}
