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
	"strconv"

	"github.com/spf13/cobra"
)

var stopClusterHelp = `
cluster [clusterId]
Stop a cluster.
Examples:

    $ steam stop cluster 42
`

func stopCluster(c *context) *cobra.Command {
	var (
		kerberos, force  bool
		username, keytab string
	)

	cmd := newCmd(c, stopClusterHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Missing clusterId. See 'steam help stop cluster'.")
		}

		clusterId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatalf("Incorrect usage of clusterId %s: Should be an integer value", args[0])
		}
		// --- add additional args here ---

		if err := c.remote.StopYarnCluster(clusterId); err != nil {
			log.Fatalln(err)
		}

	})

	cmd.Flags().BoolVar(&kerberos, "kerberos", true, "Set false on systems with no kerberos authentication.")
	cmd.Flags().StringVar(&username, "username", "", "The valid kerberos username.")
	cmd.Flags().StringVar(&keytab, "keytab", "", "The name of the keytab file to use")
	cmd.Flags().BoolVar(&force, "force", false, "Force-kill all H2O instances in the cluster")

	return cmd
}
