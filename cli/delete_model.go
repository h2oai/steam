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
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteModelHelp = `
model [modelId]
Deletes a sepcified model from the database.
Examples:
	
	$ steam delete model 3
`

func deleteModel(c *context) *cobra.Command {
	cmd := newCmd(c, deleteModelHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Incorrect number of arguments. See 'steam help delete model'.")
		}

		modelId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatalf("Invalid usage of modelId %q: expecting integer: %v", args[0], err)
		}

		if err := c.remote.DeleteModel(modelId); err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Model deleted:", modelId)
	})

	return cmd
}
