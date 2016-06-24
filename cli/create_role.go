package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var createRoleHelp = `
role [roleName] 
Creates a user permissions role.
Exampes:

	$ steam create role engineer --desc="a default engineer role"
`

func createRole(c *context) *cobra.Command {
	var description string

	cmd := newCmd(c, createRoleHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Incorrect usage. See 'steam help create role'.")
		}

		roleName := args[0]

		roleId, err := c.remote.CreateRole(roleName, description)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Created role", roleName, "ID:", roleId)
	})
	cmd.Flags().StringVarP(&description, "desc", "d", "", "A description for this role")

	return cmd
}
