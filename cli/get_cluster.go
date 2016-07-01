package cli

import (
	"fmt"
	"log"
	"strconv"

	"github.com/h2oai/steamY/master/data"
	"github.com/spf13/cobra"
)

var getClusterHelp = `
cluster [clusterId]
View detailed cluster information.
Examples:

	$ steam get cluster
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
		var lines, info, yarn []string

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

		// FIXME: formmating

		c.printt("\t"+cluster.Name, lines)
	})

	return cmd
}
