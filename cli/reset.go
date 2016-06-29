package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var resetHelp = `
reset
Reset Steam client configuration.
Examples:

    $ steam reset
`

func reset(c *context) *cobra.Command {
	cmd := newCmd(c, resetHelp, func(c *context, args []string) {
		if err := c.resetConfig(); err != nil {
			log.Fatalln("Failed configuration reset:", err)
			return
		}
		fmt.Println("Configuration reset successfully. Use 'steam login <server-address>' to re-authenticate to steam")
	})
	return cmd
}
