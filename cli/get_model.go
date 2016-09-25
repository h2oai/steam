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

var getModelHelp = `
model [modelId]
View detailed model information.
Examples:

	$ steam get model 3
`

func getModel(c *context) *cobra.Command {
	cmd := newCmd(c, getModelHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Invalid usage. See 'steam help get model'.")
		}

		modelId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatalf("Invalid usage for modelId %s: expecting int: %v", args[0], err)
		}

		model, err := c.remote.GetModel(modelId)
		if err != nil {
			log.Fatalln(err)
		}

		services, err := c.remote.GetScoringServicesForModel(modelId, 0, 10000)
		if err != nil {
			log.Fatalln(err) //FIXME format error
		}

		base := []string{
			fmt.Sprintf("ALGO:\t%s", model.Algorithm),
			fmt.Sprintf("DATASET:\t%s", model.DatasetName),
			fmt.Sprintf("TARGET:\t%s", model.ResponseColumnName),
			fmt.Sprintf("CLUSTER:\t%s", model.ClusterName),
			fmt.Sprintf("ID:\t%d", model.Id),
			fmt.Sprintf("AGE:\t%s", fmtAgo(model.CreatedAt)),
		}
		c.printt("\t"+model.Name, base)

		fmt.Println("Scoring Services:", len(services))
		if len(services) > 0 {
			ss := make([]string, len(services))
			for i, service := range services {
				ss[i] = fmt.Sprintf("%d\t%d\t%s\t%s", service.Id, service.Port, service.State, fmtAgo(service.CreatedAt))
			}

			c.printt("ID\tPORT\tSTATE\tAGE", ss)
		}
	})

	return cmd
}
