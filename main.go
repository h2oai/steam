/*
Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>
Everyone is permitted to copy and distribute verbatim copies
of this license document, but changing it is not allowed.
 */

package main


import (
	cli "github.com/h2oai/steamY/cli2"
)

var VERSION string
var BUILD_DATE string

func main() {
	cli.Run(VERSION, BUILD_DATE)
}
