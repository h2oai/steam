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

	"github.com/h2oai/steam/master/data"
	"github.com/spf13/cobra"
)

var getClusterHelp = `
cluster [clusterId]
View detailed cluster information.
Examples:

	$ steam get cluster 42
`

func getCluster(c *context) *cobra.Command {
	cmd := newCmd(c, getClusterHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Invalid usage. See 'steam help get cluster'.")
		}

		// -- Args --

		clusterId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatalln("Invalid usage for clusterId %s: expecting int: %v", args[0], err)
		}

		clusterTypes, err := getClusterTypes(c)
		if err != nil {
			log.Fatalln(err) //FIXME format error
		}

		// -- Execution --

		// create lines
		var (
			lines, info, yarn, models []string
			modelCt                   int
		)

		// base cluster info
		cluster, err := c.remote.GetCluster(clusterId)
		if err != nil {
			log.Fatalln(err)
		}

		if clusterTypes[cluster.TypeId] == data.ClusterYarn {
			clusterYarn, err := c.remote.GetYarnCluster(clusterId)
			if err != nil {
				log.Fatalln(err) //FIXME format error
			}

			yarn = []string{
				fmt.Sprintf("YARN ID:\t%s", clusterYarn.ApplicationId),
				fmt.Sprintf("H2O ENGINE:\t%d", clusterYarn.EngineId),
			}
		}

		if cluster.State == data.StartedState {
			clusterInfo, err := c.remote.GetClusterStatus(clusterId)
			if err != nil {
				log.Fatalln(err) //FIXME format error
			}

			info = []string{
				fmt.Sprintf("STATE:\t%s", clusterInfo.Status),
				fmt.Sprintf("H2O VERSION:\t%s", clusterInfo.Version),
				fmt.Sprintf("MEMORY:\t%s", clusterInfo.MaxMemory),
				fmt.Sprintf("TOTAL CPUS:\t%d", clusterInfo.TotalCpuCount),
			}

			clusterModels, err := c.remote.GetClusterModels(clusterId)
			if err != nil {
				log.Fatalln(err) //FIXME format error
			}

			modelCt = len(clusterModels)
			models = make([]string, modelCt)
			for i, clusterModel := range clusterModels {
				models[i] = fmt.Sprintf("%s\t%s\t%s",
					clusterModel.Name,
					clusterModel.Algorithm,
					fmtAgo(clusterModel.CreatedAt),
				)
			}
		} else {
			info = []string{fmt.Sprintf("STATE:\t%s", cluster.State)}
		}

		lines = []string{
			fmt.Sprintf("TYPE:\t%s", clusterTypes[cluster.TypeId]),
		}
		lines = append(lines, yarn...)
		lines = append(lines, info...)
		lines = append(lines,
			fmt.Sprintf("ID:\t%d", cluster.Id),
			fmt.Sprintf("AGE:\t%s", fmtAgo(cluster.CreatedAt)),
		)

		// -- Formatting --

		c.printt("\t"+cluster.Name, lines)
		fmt.Println("Models in cluster:", modelCt)
		if modelCt > 0 {
			c.printt("NAME\tALGO\tAGE", models)
		}
	})

	return cmd
}
