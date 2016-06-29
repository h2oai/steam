package cli

import (
	"fmt"
	"log"

	"github.com/h2oai/steamY/srv/web"
	"github.com/spf13/cobra"
)

var getRolesHelp = `
roles
List permission roles.
Exampes:

	$ steam get roles --identity=minksy
`

func getRoles(c *context) *cobra.Command {
	var identityName string
	cmd := newCmd(c, getRolesHelp, func(c *context, args []string) {
		var rs []*web.Role

		// -- Execution --

		if identityName == "" {
			var err error
			rs, err = c.remote.GetRoles(0, 10000)
			if err != nil {
				log.Fatalln(err)
			}

			// Identity set
		} else {
			identity, err := c.remote.GetIdentityByName(identityName)
			if err != nil {
				log.Fatalln(err)
			}

			rs, err = c.remote.GetRolesForIdentity(identity.Id)
			if err != nil {
				log.Fatalln(err)
			}
		}

		// -- Formatting --

		lines := make([]string, len(rs))
		for i, r := range rs {
			lines[i] = fmt.Sprintf("%s\t%d\t%s\t%s",
				r.Name,
				r.Id,
				r.Description,
				fmtAgo(r.Created))
		}

		c.printt("NAME\tID\tDESCRIPTION\tAGE", lines)
	})
	cmd.Flags().StringVarP(&identityName, "identity", "i", "", "Search by identity.")
	return cmd
}
