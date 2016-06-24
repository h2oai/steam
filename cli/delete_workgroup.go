package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var deleteWorkgroupHelp = `
workgroup [workgroupName]
Deletes a workgroup from the database.
Examples:

	$ steam delete workgroup production
`

func deleteWorkgroup(c *context) *cobra.Command {
	cmd := newCmd(c, deleteWorkgroupHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Invalid usage. See 'steam help delete workgroup'.")
		}

		// -- Args --

		workgroupName := args[0]

		// -- Execution --

		workgroup, err := c.remote.GetWorkgroupByName(workgroupName)
		if err != nil {
			log.Fatalln(err) // TODO
		}

		if err := c.remote.DeleteWorkgroup(workgroup.Id); err != nil {
			log.Fatalln(err) // TODO
		}

		// -- Formatting --

		fmt.Println("Successfully deleted workgroup:", workgroupName)
	})

	return cmd
}
