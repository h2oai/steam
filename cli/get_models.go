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

	"github.com/spf13/cobra"
)

var getModelsHelp = `
models
List all models.
Examples:

	$ steam get models
`

func getModels(c *context) *cobra.Command {
	var projectId int64
	cmd := newCmd(c, getModelsHelp, func(c *context, args []string) {

		// FIXME

		ms, err := c.remote.GetModels(projectId, 0, 10000)
		if err != nil {
			log.Fatalln(err)
		}

		lines := make([]string, len(ms))
		for i, m := range ms {
			lines[i] = fmt.Sprintf("%s\t%d\t%s\t%s\t%s\t%s ", m.Name, m.Id, m.Algorithm, m.DatasetName, m.ResponseColumnName, fmtAgo(m.CreatedAt))
		}
		c.printt("NAME\tALGO\tDATASET\tTARGET\tAGE", lines)
	})
	cmd.Flags().Int64Var(&projectId, "project-id", projectId, "Project ID")
	return cmd
}
