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
			log.Fatalf("Invalid usage of clusterId %s: Integer value required: %v", args[0], err)
		}

		if err := c.remote.DeleteCluster(clusterId); err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Cluster deleted:", clusterId)
	})

	return cmd
}
