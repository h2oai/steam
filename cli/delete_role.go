package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var deleteRoleHelp = `
role [roleName]
Deletes a role from the database.
Examples:

	$ steam delete role engineer
`

func deleteRole(c *context) *cobra.Command {
	cmd := newCmd(c, deleteRoleHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Invalid usage. See 'steam help delete role'.")
		}

		// -- Args --

		roleName := args[0]

		// -- Execution --

		role, err := c.remote.GetRoleByName(roleName)
		if err != nil {
			log.Fatalln(err) // FIXME
		}

		if err := c.remote.DeleteRole(role.Id); err != nil {
			log.Fatalln(err) // FIXME
		}

		// -- Formatting --

		fmt.Println("Successfully deleted role:", roleName)
	})

	return cmd
}
