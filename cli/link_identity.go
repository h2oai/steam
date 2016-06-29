package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var linkIdentityHelp = `
identity [identityName] [role | workgroup] [entityName]
Add an permissions entity group to an identity.
Example:

	$ steam link identity minksy role engineer
`

func linkIdentity(c *context) *cobra.Command {
	cmd := newCmd(c, linkIdentityHelp, func(c *context, args []string) {
		if len(args) != 3 {
			log.Fatalln("Invalid usage. See 'steam help link identity'.")
		}

		// -- Args --

		identityName := args[0]
		entityType := args[1]
		entityName := args[2]

		// -- Execution --

		identity, err := c.remote.GetIdentityByName(identityName)
		if err != nil {
			log.Fatalf("Identity %s not found: %v", identityName, err)
		}

		if entityType == "role" { // CASE: Link identity -> role
			entity, err := c.remote.GetRoleByName(entityName)
			if err != nil {
				log.Fatalf("Role %s not found: %v", entityName, err)
			}

			if err := c.remote.LinkIdentityAndRole(identity.Id, entity.Id); err != nil {
				log.Fatalf("Failed to link identity %s to role %s: %v",
					identityName, entityName, err)
			}
		} else if entityType == "workgroup" { // CASE: Link identity -> workgroup
			entity, err := c.remote.GetWorkgroupByName(entityName)
			if err != nil {
				log.Fatalf("Workgroup %s not found: %v", entityName, err)
			}

			if err := c.remote.LinkIdentityAndWorkgroup(identity.Id, entity.Id); err != nil {
				log.Fatalf("Failed to link identity %s to workgroup %s: %v",
					identityName, entityName, err)
			}
		} else {
			log.Fatalln("Entity type must be either role or workgroup")
		}

		// -- Formatting --

		fmt.Println("Successfully linked identity", identityName, "to", entityName)
	})

	return cmd
}
