package cli

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteModelHelp = `
model [modelId]
Deletes a sepcified model from the database.
Examples:
	
	$ steam delete model 3
`

func deleteModel(c *context) *cobra.Command {
	cmd := newCmd(c, deleteModelHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Incorrect number of arguments. See 'steam help delete model'.")
		}

		modelId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatalf("Invalid usage of modelId %s: Integer value required: %v", args[0], err)
		}

		if err := c.remote.DeleteModel(modelId); err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Model deleted:", modelId)
	})

	return cmd
}
