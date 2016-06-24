package cli

import "github.com/spf13/cobra"

var linkHelp = `
link [permission-type]
Add authentication permissions.
Examples:

	$ steam link permission
`

func link(c *context) *cobra.Command {
	cmd := newCmd(c, linkHelp, nil)
	cmd.AddCommand(linkRole(c))
	cmd.AddCommand(linkIdentity(c))
	return cmd
}
