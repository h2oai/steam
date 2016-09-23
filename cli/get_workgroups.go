package cli

// FIXME use ByName remote calls

import (
	"fmt"
	"log"

	"github.com/h2oai/steam/srv/web"
	"github.com/spf13/cobra"
)

var getWorkgroupsHelp = `
workgroups
List permission workgroups.
Examples:

	$ steam get workgroups --identity=2
`

func getWorkgroups(c *context) *cobra.Command {
	var identityName string
	cmd := newCmd(c, getWorkgroupsHelp, func(c *context, args []string) {

		// -- Execution --

		var ws []*web.Workgroup

		if identityName == "" {
			var err error
			ws, err = c.remote.GetWorkgroups(0, 10000)
			if err != nil {
				log.Fatalln(err)
			}

			// Identity set
		} else {
			identity, err := c.remote.GetIdentityByName(identityName)
			if err != nil {
				log.Fatalln(err)
			}

			ws, err = c.remote.GetWorkgroupsForIdentity(identity.Id)
			if err != nil {
				log.Fatalln(err)
			}
		}

		// -- Formatting --

		lines := make([]string, len(ws))
		for i, w := range ws {
			lines[i] = fmt.Sprintf("%s\t%d\t%s\t%s",
				w.Name,
				w.Id,
				w.Description,
				fmtAgo(w.Created))
		}

		c.printt("NAME\tID\tDESCRIPTION\tAGE", lines)
	})
	cmd.Flags().StringVarP(&identityName, "identity", "i", "", "Search by identity.")

	return cmd
}
