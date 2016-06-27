package cli

import "github.com/spf13/cobra"

var serveHelp = `
serve [agent-type]
Lauch a new service.
Examples:

    $ steam serve master
`

func serve(c *context) *cobra.Command {
	cmd := newCmd(c, serveHelp, nil)
	cmd.AddCommand(serveMaster(c))
	return cmd
}
