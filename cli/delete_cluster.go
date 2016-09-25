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

var deleteClusterHelp = `
cluster [clusterId]
Deletes a specified cluster from the database.
Examples:
	
	$ steam delete cluster 42
`

func deleteCluster(c *context) *cobra.Command {
	cmd := newCmd(c, deleteClusterHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Missing clusterName. See 'steam help delete cluster'.")
		}

		clusterId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatalf("Invalid usage of clusterId %q: expecting integer: %v", args[0], err)
		}

		if err := c.remote.DeleteCluster(clusterId); err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Cluster deleted:", clusterId)
	})

	return cmd
}
