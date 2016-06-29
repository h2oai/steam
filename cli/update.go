package cli

import "github.com/spf13/cobra"

var updateHelp = `
update [entityType]
Updates an entity in the database.
Examples:

	$ steam update workgroup
`

func update(c *context) *cobra.Command {
	cmd := newCmd(c, updateHelp, nil)
	cmd.AddCommand(updateRole(c))
	cmd.AddCommand(updateWorkgroup(c))
	return cmd
}
