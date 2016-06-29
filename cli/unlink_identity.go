package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var unlinkIdentityHelp = `
identity [identityName] [role | workgroup] [entityName]
Unlink an identity from a permissions entityt
Examples:

	$ steam unlink identity minksy role engineer
`

func unlinkIdentity(c *context) *cobra.Command {
	cmd := newCmd(c, unlinkIdentityHelp, func(c *context, args []string) {
		if len(args) != 3 {
			log.Fatalln("Invalid usage. See 'steam help unlink identity'.")
		}

		// -- Args --

		identityName := args[0]
		entityType := args[1]
		entityName := args[2]

		// -- Execution --

		identity, err := c.remote.GetIdentityByName(identityName)
		if err != nil {
			log.Fatalln(err) // TODO
		}

		// -- CASE: Role --
		if entityType == "role" {
			role, err := c.remote.GetRoleByName(entityName)
			if err != nil {
				log.Fatalln(err) // TODO
			}

			if err := c.remote.UnlinkIdentityAndRole(identity.Id, role.Id); err != nil {
				log.Fatalln(err) // TODO
			}

			// -- CASE: Workgroup --
		} else if entityType == "workgroup" {
			workgroup, err := c.remote.GetWorkgroupByName(entityName)
			if err != nil {
				log.Fatalln(err) // TODO
			}

			if err := c.remote.UnlinkIdentityAndWorkgroup(identity.Id, workgroup.Id); err != nil {
				log.Fatalln(err) // TODO
			}

			// -- CASE: Niether --
		} else {
			log.Fatalln("Invalid usage. See 'steam help unlink identity'.")
		}

		// -- Formatting --

		fmt.Println("Successfully unlinked identity", identityName, "and", entityType, entityName)
	})

	return cmd
}
