package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var getEntitiesHelp = `
entities
View the supported entity types.
Examples:

	$ steam get entities
`

func getEntities(c *context) *cobra.Command {
	cmd := newCmd(c, getEntitiesHelp, func(c *context, args []string) {

		// -- Execution

		es, err := c.remote.GetSupportedEntityTypes()
		if err != nil {
			log.Fatalln(err)
		}

		// -- Formatting --

		lines := make([]string, len(es))
		for i, e := range es {
			lines[i] = fmt.Sprintf("%s\t%d", e.Name, e.Id)
		}

		c.printt("NAME\tID", lines)
	})

	return cmd
}
