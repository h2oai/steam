package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var deactivateIdentityHelp = `
identity [identityName]
Deactivate a user.
Examples:

	$ steam deactivate identity minsky
`

func deactivateIdentity(c *context) *cobra.Command {
	cmd := newCmd(c, deactivateIdentityHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Invalid usage. See 'steam help deactivate identity'.")
		}

		// -- Args --

		identityName := args[0]

		// -- Execution --

		identity, err := c.remote.GetIdentityByName(identityName)
		if err != nil {
			log.Fatalln(err) // TODO
		}

		if err := c.remote.DeactivateIdentity(identity.Id); err != nil {
			log.Fatalln(err) // TODO
		}

		// -- Formatting --

		fmt.Println("Successfully deactivated user", identityName)
	})

	return cmd
}
