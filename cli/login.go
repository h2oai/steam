package cli

import (
	"log"

	"github.com/spf13/cobra"
)

var loginHelp = `
login [address]
Access a Steam server.
Examples:

	$ steam login steam.local
	$ steam login steam.local:9000
`

func login(c *context) *cobra.Command {
	cmd := newCmd(c, loginHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("*** Missing address. See 'steam help login'.")
		}
		addr := args[0]

		c.config.CurrentHost = addr

		c.saveConfig(c.config)
	})

	return cmd
}
