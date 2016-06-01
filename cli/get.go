package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var getHelp = `
get [resource-type]
List or view resources of the specified type.
Examples:

    $ steam get clouds
`

func get(c *context) *cobra.Command {
	cmd := newCmd(c, getHelp, nil)
	cmd.AddCommand(getClouds(c))
	return cmd
}

var getCloudsHelp = `
clouds
List all clouds.
Examples:

	$ steam get clouds
`

func getClouds(c *context) *cobra.Command {
	var details bool
	cmd := newCmd(c, getCloudsHelp, func(c *context, args []string) {
		cs, err := c.remote.GetClouds()
		if err != nil {
			log.Fatalln(err)
		}

		lines := make([]string, len(cs))
		if details {
			for i, cl := range cs {
				info, err := c.remote.GetCloudStatus(cl.Name)
				if err != nil {
					log.Fatalln(err)
				}
				lines[i] = fmt.Sprintf("%s\t%s\t%s\t%d\t%s", info.Name, info.EngineName, info.Memory, info.Size, info.State)
			}
			c.printt("NAME\tENGINE\tMEMORY\tSIZE\tSTATE", lines)
		} else {
			lines := make([]string, len(cs))
			for i, cl := range cs {
				lines[i] = fmt.Sprintf("%s\t%s\t%s", cl.Name, cl.EngineName, cl.State)
			}
			c.printt("NAME\tENGINE\tSTATE", lines)
		}
	})

	cmd.Flags().BoolVarP(&details, "details", "d", false, "Detailed cluster information")

	return cmd
}

//FIXME: getCloud requires storage of all nodes in cluster
