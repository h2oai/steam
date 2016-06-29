package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var getServicesHelp = `
services
List all services.
Examples:

	$ steam get services
`

func getServices(c *context) *cobra.Command {
	cmd := newCmd(c, getServicesHelp, func(c *context, args []string) {

		ss, err := c.remote.GetScoringServices(0, 1000)
		if err != nil {
			log.Fatalln(err)
		}

		lines := make([]string, len(ss))
		for i, s := range ss {
			lines[i] = fmt.Sprintf("%d\t%s:%d\t%s\t%s", s.Id, s.Address, s.Port, s.State, fmtAgo(s.CreatedAt))
		}

		c.printt("ID\tAddress\tSTATE\tAGE", lines)
	})

	return cmd
}
