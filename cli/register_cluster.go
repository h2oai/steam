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
