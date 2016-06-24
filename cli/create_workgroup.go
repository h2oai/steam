package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var createWorkgroupHelp = `
workgroup [workgroupName] 
Creates a user permissions workgroup.
Exampes:

	$ steam create workgroup production --desc="The production group"
`

func createWorkgroup(c *context) *cobra.Command {
	var description string

	cmd := newCmd(c, createWorkgroupHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Incorrect usage. See 'steam help create workgroup'.")
		}

		workgroupName := args[0]

		workgroupId, err := c.remote.CreateWorkgroup(workgroupName, description)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Created workgroup", workgroupName, "ID:", workgroupId)
	})
	cmd.Flags().StringVarP(&description, "desc", "d", "", "A description for this workgroup")

	return cmd
}
