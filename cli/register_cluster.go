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

var registerClusterHelp = `
cluster [clusterAddress]
Connect to a running h2o instance.
Examples:

	$ steam register cluster localhost:54321
`

func registerCluster(c *context) *cobra.Command {
	cmd := newCmd(c, registerClusterHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Invalid usage. See 'steam help register cluster'.")
		}

		// -- Args --

		clusterAddress := args[0]

		// -- Execution --

		clusterId, err := c.remote.RegisterCluster(clusterAddress)
		if err != nil {
			log.Fatalln(err) //FIXME format error
		}

		// -- Formatting --

		fmt.Printf("Successfully connected to cluster %d at address %s\n", clusterId, clusterAddress)
	})

	return cmd
}
