package cli

import "github.com/spf13/cobra"

var stopHelp = `
stop [resource-type]
Stop the specified resource.
Examples:

    $ steam stop cluster
`

func stop(c *context) *cobra.Command {
	cmd := newCmd(c, stopHelp, nil)
	cmd.AddCommand(stopCluster(c))
	cmd.AddCommand(stopService(c))
	return cmd
}
