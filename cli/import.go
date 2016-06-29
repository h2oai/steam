package cli

import "github.com/spf13/cobra"

var importHelp = `
import
Import a resource of the specified type into steam.
Examples:

	$ steam import model
`

func import_(c *context) *cobra.Command {
	cmd := newCmd(c, importHelp, nil)
	cmd.AddCommand(importModel(c))
	return cmd
}
