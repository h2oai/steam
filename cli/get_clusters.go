package cli

import (
	"fmt"
	"log"

	"github.com/h2oai/steamY/master/data"
	"github.com/spf13/cobra"
)

var getClustersHelp = `
clusters
List all clusters.
Examples

	$ steam get clusters
`

func getClusters(c *context) *cobra.Command {
	var details bool

	cmd := newCmd(c, getClustersHelp, func(c *context, args []string) {
		cs, err := c.remote.GetClusters(0, 10000)
		if err != nil {
			log.Fatalln(err)
		}

		clusterTypes, err := getClusterTypes(c)
		if err != nil {
			log.Fatalln(err) //FIXME format error
		}

		lines := make([]string, len(cs))
		for i, cl := range cs {
			if !details {
				lines[i] = fmt.Sprintf("%s\t%d\t%s\t%s\t%s\t%s", cl.Name, cl.Id,
					cl.Address, cl.State, clusterTypes[cl.TypeId], fmtAgo(cl.CreatedAt))
			} else {
				if (cl.State != data.StoppedState) || (cl.State != data.DisconnectedState) {
					info, err := c.remote.GetClusterStatus(cl.Id)
					if err != nil {
						log.Fatalln(err)
					}

					lines[i] = fmt.Sprintf("%s\t%d\t%s\t%s\t%d\t%s\t%s\t%s", cl.Name,
						cl.Id, cl.Address, info.MaxMemory, info.TotalCpuCount,
						info.Status, clusterTypes[cl.TypeId], fmtAgo(cl.CreatedAt))
				} else {
					lines[i] = fmt.Sprintf("%s\t%d\t%s\tNA\tNA\t%s\t%s\t%s", cl.Name,
						cl.Id, cl.Address, cl.State, clusterTypes[cl.TypeId],
						fmtAgo(cl.CreatedAt))
				}
			}
		}

		var header string
		if !details {
			header = "NAME\tID\tADDRESS\tSTATE\tTYPE\tAGE"
		} else {
			header = "NAME\tID\tADDRESS\tMEMORY\tCPUs\tSTATE\tTYPE\tAGE"
		}

		c.printt(header, lines)
	})

	cmd.Flags().BoolVarP(&details, "details", "d", false, "Detailed cluster information")

	return cmd
}
