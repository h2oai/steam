package main

/*
		    GNU AFFERO GENERAL PUBLIC LICENSE
                       Version 3, 19 November 2007

 Copyright (C) 2007 Free Software Foundation, Inc. <http://fsf.org/>
 Everyone is permitted to copy and distribute verbatim copies
 of this license document, but changing it is not allowed.
 */

import (
	cli "github.com/h2oai/steamY/cli2"
)

var VERSION string
var BUILD_DATE string

func main() {
	cli.Run(VERSION, BUILD_DATE)
}
