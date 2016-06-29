package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var updateRoleHelp = `
role [roleName]
Change or update a role in the database.
Examples:

	$ steam update role engineer --desc="A better engineer" --name=
`

func updateRole(c *context) *cobra.Command {
	var description, name string
	cmd := newCmd(c, updateRoleHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Invalid usage. See 'steam help update role'.")
		}

		// -- Args --

		roleName := args[0]

		// -- Execution --

		role, err := c.remote.GetRoleByName(roleName)
		if err != nil {
			log.Fatalln(err) // TODO
		}
		if name == "" {
			name = role.Name
		}
		if description == "" {
			description = role.Description
		}

		if err := c.remote.UpdateRole(role.Id, name, description); err != nil {
			log.Fatalln(err) // TODO
		}

		// -- Formatting --

		fmt.Println("Successfully updated role:", roleName)
	})
	cmd.Flags().StringVarP(&description, "desc", "d", "", "An updated role description.")
	cmd.Flags().StringVarP(&name, "name", "n", "", "An updated role name.")

	return cmd
}
