package cli

import "github.com/spf13/cobra"

var createHelp = `
create [resource-type]
Creates an instance of the specified resource.
Examples:

	$ steam create identity
`

func create(c *context) *cobra.Command {
	cmd := newCmd(c, createHelp, nil)
	cmd.AddCommand(createIdentity(c))
	cmd.AddCommand(createRole(c))
	cmd.AddCommand(createWorkgroup(c))
	return cmd
}
