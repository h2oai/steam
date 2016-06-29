package cli

import "github.com/spf13/cobra"

var getHelp = `
get [resource-type]
List or view resources of the specified type.
Examples:

    $ steam get clusters
`

func get(c *context) *cobra.Command {
	cmd := newCmd(c, getHelp, nil)
	cmd.AddCommand(getClusters(c))
	cmd.AddCommand(getEngines(c))
	cmd.AddCommand(getModels(c))
	cmd.AddCommand(getServices(c))
	cmd.AddCommand(getIdentities(c))
	cmd.AddCommand(getRoles(c))
	cmd.AddCommand(getWorkgroups(c))
	cmd.AddCommand(getPermissions(c))
	cmd.AddCommand(getEntities(c))
	cmd.AddCommand(getHistory(c))
	return cmd
}
