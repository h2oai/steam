package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var deleteHelp = `
delete [resource-type]
Deletes the specified resource from the database.
Examples:

    $ steam delete cloud
`

func delete(c *context) *cobra.Command {
	cmd := newCmd(c, deleteHelp, nil)
	cmd.AddCommand(deleteCloud(c))
	return cmd
}

var deleteCloudHelp = `
cloud [modelName]
Deletes a specified cloud from the database.
Examples:
	
	$ steam delete cloud cloud42
`

func deleteCloud(c *context) *cobra.Command {
	cmd := newCmd(c, deleteCloudHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Missing cloudName. See 'steam help delete cloud'.")
		}

		cloudName := args[0]

		err := c.remote.DeleteModel(cloudName)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Model deleted:", cloudName)
	})

	return cmd
}
