package cli

import "github.com/spf13/cobra"

var deactivateHelp = `
deactivate [entityType]
Deactivate and entity type.
Examples:

	$ steam deactivate identity
`

func deactivate(c *context) *cobra.Command {
	cmd := newCmd(c, deactivateHelp, nil)
	cmd.AddCommand(deactivateIdentity(c))
	return cmd
}
