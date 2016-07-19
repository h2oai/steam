package cli

import "github.com/spf13/cobra"

var unregisterHelp = `
unregister [resourceType]
Unregister an external resource.
Examples:

	$ steam unregister cluster
`

func unregister(c *context) *cobra.Command {
	cmd := newCmd(c, unregisterHelp, nil)
	cmd.AddCommand(unregisterCluster(c))
	return cmd
}
