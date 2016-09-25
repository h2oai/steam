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

package cli

import (
	"log"
	"path"

	"github.com/h2oai/steam/lib/fs"
	"github.com/spf13/cobra"
)

var deployEngineHelp = `
engine [enginePath]
Deploy an H2O engine to Steam.
Examples:

	$ steam deploy engine path/to/engine
`

func deployEngine(c *context) *cobra.Command {
	cmd := newCmd(c, deployEngineHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Incorrect number of arguments. See 'steam help deploy engine'.")
		}

		enginePath := args[0]

		if err := c.uploadFile(enginePath, fs.KindEngine); err != nil {
			log.Fatalln(err)
		}

		log.Println("Engine deployed:", path.Base(enginePath))
	})

	return cmd
}
