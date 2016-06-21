package cli

import "github.com/spf13/cobra"

var startHelp = `
start [resource-type]
Start a new resource.
Examples:

    $ steam start cloud
`

func start(c *context) *cobra.Command {
	cmd := newCmd(c, startHelp, nil)
	cmd.AddCommand(startCluster(c))
	cmd.AddCommand(startService(c))
	return cmd
}
