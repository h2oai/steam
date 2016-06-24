package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var updateWorkgroupHelp = `
workgroup [workgroupName]
Change or update a workgroup in the database.
Examples:

	$ steam update workgroup engineer --desc="A better engineer" --name=
`

func updateWorkgroup(c *context) *cobra.Command {
	var description, name string
	cmd := newCmd(c, updateWorkgroupHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Invalid usage. See 'steam help update workgroup'.")
		}

		// -- Args --

		workgroupName := args[0]

		// -- Execution --

		workgroup, err := c.remote.GetWorkgroupByName(workgroupName)
		if err != nil {
			log.Fatalln(err) // TODO
		}
		if name == "" {
			name = workgroup.Name
		}
		if description == "" {
			description = workgroup.Description
		}

		if err := c.remote.UpdateWorkgroup(workgroup.Id, name, description); err != nil {
			log.Fatalln(err) // TODO
		}

		// -- Formatting --

		fmt.Println("Successfully updated workgroup:", workgroupName)
	})
	cmd.Flags().StringVarP(&description, "desc", "d", "", "An updated workgroup description.")
	cmd.Flags().StringVarP(&name, "name", "n", "", "An updated workgroup name.")

	return cmd
}
