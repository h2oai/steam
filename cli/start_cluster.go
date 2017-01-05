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

var startClusterHelp = `
cluster [clusterName] [engineId]
Start a new cluster using the specified H2O package.
Examples:

Start a 4-node H2O 3.2.0.9 cluster

    $ steam start cluster 42 2 --size=4
`

func startCluster(c *context) *cobra.Command {
	var (
		size   int
		memory string
		secure bool
	)

	cmd := newCmd(c, startClusterHelp, func(c *context, args []string) {
		if len(args) != 2 {
			log.Fatalln("Incorrect number of arguments. See 'steam help start cluster'.")
		}

		clusterName := args[0]
		engineId, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			log.Fatalf("Ivalid usage of engineId %q: expecting integer: %v", args[1], err)
		}

		// --- add additional args here ---

		log.Println("Attempting to start cluster...")
		clusterId, err := c.remote.StartYarnCluster(clusterName, engineId, size, memory, secure, "")
		if err != nil {
			log.Fatalln(err)
		}

		cluster, err := c.remote.GetCluster(clusterId)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("Cluster %d started at: %s", clusterId, cluster.Address)

		// if details {
		// 	info, err := c.remote.GetClusterStatus(cluster.Name)
		// 	if err != nil {
		// 		log.Fatalln(err)
		// 	}

		// 	fmt.Printf(
		// 		`
		// Engine: %v
		// Version: %v
		// Total Nodes: %v
		// Total Memory: %v
		// `, info.EngineName, info.EngineVersion, info.Size, info.Memory)
		// }
		// // TODO: name corresponds to id for purpose of stopCluster
	})

	cmd.Flags().IntVarP(&size, "size", "n", 1, "The number of nodes to provision.")
	cmd.Flags().StringVarP(&memory, "memory", "m", "10g", "The max amount of memory to use per node.")

	return cmd
}
