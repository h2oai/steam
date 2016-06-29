package cli

import "github.com/spf13/cobra"

var unlinkHelp = `
unlink [entityType]
Remove authentication permissions.
Examples:

	$ steam unlink identity minksy role engineer
`

func unlink(c *context) *cobra.Command {
	cmd := newCmd(c, unlinkHelp, nil)
	cmd.AddCommand(unlinkIdentity(c))
	return cmd
}
