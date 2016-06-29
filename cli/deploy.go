package cli

import "github.com/spf13/cobra"

var deployHelp = `
deploy [resource-type]
Deploy a resource of the specified type.
Examples:

	$ steam deploy engine
`

func deploy(c *context) *cobra.Command {
	cmd := newCmd(c, deployHelp, nil)
	cmd.AddCommand(deployEngine(c))
	return cmd
}
