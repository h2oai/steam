package cli

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var unregisterClusterHelp = `
cluster [clusterId]
Remove an external cluster from steam.
Examples:

	$ steam unregister cluster 9
`

func unregisterCluster(c *context) *cobra.Command {
	cmd := newCmd(c, unregisterClusterHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Invalid usage. See 'steam help unregister cluster'.")
		}

		// -- Args --

		clusterId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatalf("Invalid usage for clusterId %s: expecting int: %v", args[0], err)
		}

		// -- Execution --

		if err := c.remote.UnregisterCluster(clusterId); err != nil {
			log.Fatalln(err) //FIXME format error
		}

		// -- Formatting --

		fmt.Println("Successfully unregisted cluster %d", clusterId)
	})

	return cmd
}
