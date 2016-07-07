package cli

// FIXME use ByName remote calls

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var linkRoleHelp = `
role [roleName] [permissionIds]
Add permissions to a role. 
Examples:

	$ steam link role engineer ViewCluster ViewModel ViewWorkgroup
`

func linkRole(c *context) *cobra.Command {
	cmd := newCmd(c, linkRoleHelp, func(c *context, args []string) {
		if len(args) < 2 {
			log.Fatalln("Invalid usage. See 'steam help link role'.")
		}

		// -- Args --

		roleName := args[0]
		permissions := args[1:len(args)]
		fmt.Println(roleName, permissions)

		// -- Execution --

		role, err := c.remote.GetRoleByName(roleName)
		if err != nil {
			log.Fatalln(err) //FIXME format error
		}

		permissionIds, err := getPermissionIds(c, permissions...)
		if err != nil {
			log.Fatalln(err) //FIXME format error
		}

		c.remote.LinkRoleAndPermissions(role.Id, permissionIds)
		fmt.Println("Role", roleName, "linked to permissions:", permissions)
	})

	return cmd
}
