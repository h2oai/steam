package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var getEnginesHelp = `
engines
List all engines.
Examples:

	$ steam get engines
`

func getEngines(c *context) *cobra.Command {
	cmd := newCmd(c, getEnginesHelp, func(c *context, args []string) {
		es, err := c.remote.GetEngines()
		if err != nil {
			log.Fatalln(err)
		}

		lines := make([]string, len(es))
		for i, e := range es {
			lines[i] = fmt.Sprintf("%s\t%d\t%s", e.Name, e.Id, fmtAgo(e.CreatedAt))
		}
		c.printt("NAME\tID\tAGE", lines)
	})

	return cmd
}
