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

var importModelHelp = `
model [clusterId] [modelName]
Import a model from an H2O cluster into steam
Examples:

	$ steam import model 42 model3
`

func importModel(c *context) *cobra.Command {
	var (
		projectId int64
	)
	cmd := newCmd(c, importModelHelp, func(c *context, args []string) {
		if len(args) != 2 {
			log.Fatalln("Incorrect number of arguments. See 'steam help import model'.")
		}

		clusterId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatalf("Incorrect value for clusterId: %s: %v", args[0], err)
		}
		modelName := args[1]

		if _, err := c.remote.ImportModelFromCluster(clusterId, projectId, modelName); err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Retireved model:", modelName)
	})
	cmd.Flags().Int64Var(&projectId, "project-id", projectId, "ID of project to import this model into")

	return cmd
}
