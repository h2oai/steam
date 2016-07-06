package cli

import "github.com/spf13/cobra"

var registerHelp = `
register [resource-type]
Register an external resource.
Examples:

	$ steam register cluster
`

func register(c *context) *cobra.Command {
	cmd := newCmd(c, registerHelp, nil)
	cmd.AddCommand(registerCluster(c))
	return cmd
}
