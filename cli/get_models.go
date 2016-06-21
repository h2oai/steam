package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var getModelsHelp = `
models
List all models.
Examples:

	$ steam get models
`

func getModels(c *context) *cobra.Command {
	cmd := newCmd(c, getModelsHelp, func(c *context, args []string) {

		// FIXME

		ms, err := c.remote.GetModels(0, 10000)
		if err != nil {
			log.Fatalln(err)
		}

		lines := make([]string, len(ms))
		for i, m := range ms {
			lines[i] = fmt.Sprintf("%s\t%d\t%s\t%s\t%s\t%s\t", m.Name, m.Id, m.Algorithm, m.DatasetName, m.ResponseColumnName, fmtAgo(m.CreatedAt))
		}
		c.printt("NAME\tALGO\tDATASET\tTARGET\tAGE\t", lines)
	})

	return cmd
}
