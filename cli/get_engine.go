package cli

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var getEngineHelp = `
engine [engineIds]
View detailed engine information.
Examples:

	$ steam get engine 2
`

func getEngine(c *context) *cobra.Command {
	cmd := newCmd(c, getEngineHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Invalid usage. See 'steam help get engine'.")
		}

		engineId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatalf("Invalid usage for engineId %s: expecting int: %v", args[0], err)
		}

		engine, err := c.remote.GetEngine(engineId)
		if err != nil {
			log.Fatalln(err)
		}

		base := []string{
			fmt.Sprintf("ID:\t%d", engine.Id),
			fmt.Sprintf("AGE:\t%s", fmtAgo(engine.CreatedAt)),
		}

		c.printt("\t"+engine.Name, base)
	})

	return cmd
}
