package cli

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteEngineHelp = `
engine [engineid]
Deletes a specified engine from the database.
Examples:

	$ steam delete engine 2
`

func deleteEngine(c *context) *cobra.Command {
	cmd := newCmd(c, deleteEngineHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Incorrect number of arugments. See 'steam help delete engine'.")
		}

		engineId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatalf("Invalid usage of engineId %s: Integer value required: %v", args[0], err)
		}

		if err := c.remote.DeleteEngine(engineId); err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Engine deleted:", engineId)
	})

	return cmd
}
