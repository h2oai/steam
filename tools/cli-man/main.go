/*
  Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

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

	header := &cobra.GenManHeader{
		Title:   "STEAM",
		Section: "1",
		Manual:  "Steam CLI Manual",
		Source:  "(c) 2016 H2O.ai",
	}

	steam.GenManTree(header, p)
}
