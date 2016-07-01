package cli

import (
	"fmt"
	"log"

	"github.com/h2oai/steamY/srv/web"
	"github.com/spf13/cobra"
)

var getIdentitiesHelp = `
identities
List all identities.
Examples:

	$ steam get identities --workgroup=production
`

func getIdentities(c *context) *cobra.Command {
	var workgroup, role string

	cmd := newCmd(c, getIdentitiesHelp, func(c *context, args []string) {

		// -- Execution --

		var ids []*web.Identity

		// Both workgroup and role set
		if workgroup != "" && role != "" {
			log.Fatalln("Cannot use both workgroup and role at the same time.")

			// Workgroup set
		} else if workgroup != "" {
			wg, err := c.remote.GetWorkgroupByName(workgroup)
			if err != nil {
				log.Fatalln(err)
			}
			ids, err = c.remote.GetIdentitiesForWorkgroup(wg.Id)
			if err != nil {
				log.Fatalln(err)
			}

			// Role set
		} else if role != "" {
			rl, err := c.remote.GetRoleByName(role)
			if err != nil {
				log.Fatalln(err)
			}

			ids, err = c.remote.GetIdentititesForRole(rl.Id)
			if err != nil {
				log.Fatalln(err)
			}

			// Neither set
		} else {
			var err error
			ids, err = c.remote.GetIdentities(0, 10000)
			if err != nil {
				log.Fatalln(err)
			}
		}

		// -- Formatting --

		lines := make([]string, len(ids))
		for i, id := range ids {
			lines[i] = fmt.Sprintf("%s\t%d\t%s\t%s",
				id.Name,
				id.Id,
				fmtAgo(id.LastLogin),
				fmtAgo(id.Created))
		}

		c.printt("NAME\tID\tLAST LOGIN\tAGE", lines)
	})

	cmd.Flags().StringVarP(&workgroup, "workgroup", "w", "", "Search for users by workgroup.")
	cmd.Flags().StringVarP(&role, "role", "r", "", "Search for users by role.")

	return cmd
}
