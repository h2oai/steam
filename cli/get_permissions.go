package cli

import (
	"fmt"
	"log"

	"github.com/h2oai/steamY/srv/web"
	"github.com/spf13/cobra"
)

var getPermissionsHelp = `
permissions
Get permissions and their corresponding codes.
Examples:

	$ steam get permissions --role=2
`

func getPermissions(c *context) *cobra.Command {
	var roleName, identityName string
	cmd := newCmd(c, getPermissionsHelp, func(c *context, args []string) {

		// -- Execution --

		var ps []*web.Permission

		if roleName != "" && identityName != "" {
			log.Fatalln("Cannot view by role and identity.")

			// Role set
		} else if roleName != "" {
			role, err := c.remote.GetRoleByName(roleName)
			if err != nil {
				log.Fatalln(err)
			}

			ps, err = c.remote.GetPermissionsForRole(role.Id)
			if err != nil {
				log.Fatalln(err)
			}

			// Identity set
		} else if identityName != "" {
			identity, err := c.remote.GetIdentityByName(identityName)
			if err != nil {
				log.Fatalln(err)
			}

			ps, err = c.remote.GetPermissionsForIdentity(identity.Id)
			if err != nil {
				log.Fatalln(err)
			}

			// Niether set
		} else {
			var err error
			ps, err = c.remote.GetSupportedPermissions()
			if err != nil {
				log.Fatalln(err)
			}
		}

		// -- Formatting --

		lines := make([]string, len(ps))
		for i, p := range ps {
			lines[i] = fmt.Sprintf("%d\t%s\t%s", p.Id, p.Description, p.Code)
		}

		c.printt("ID\tDESCRIPTION\tCODE", lines)
	})

	cmd.Flags().StringVarP(&roleName, "role", "r", "", "View permissions for a role.")
	cmd.Flags().StringVarP(&identityName, "identity", "i", "", "View permissions for an identity.")
	return cmd
}
