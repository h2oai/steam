package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var createIdentityHelp = `
identity [username] [password]
Creates a user.
Examples:

	$ steam create minsky m1n5kypassword
`

func createIdentity(c *context) *cobra.Command {
	cmd := newCmd(c, createIdentityHelp, func(c *context, args []string) {
		if len(args) != 2 {
			log.Fatalln("Incorrect usage. See 'steam help create identity'.")
		}

		username := args[0]
		password := args[1]

		identityId, err := c.remote.CreateIdentity(username, password)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Created user", username, "ID:", identityId)
	})

	return cmd
}
