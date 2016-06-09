package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var retrieveHelp = `
retrieve
Retrieve a resource of the specified type into steam.
Examples:

	$ steam import model
`

func retrieve(c *context) *cobra.Command {
	cmd := newCmd(c, retrieveHelp, nil)
	cmd.AddCommand(retrieveModel(c))
	return cmd
}

var retrieveModelHelp = `
model [modelName] [cloudName]
Retrieve a model from an H2O cluster into steam
Examples:

	$ steam retrieve model model3 cloud42
`

func retrieveModel(c *context) *cobra.Command {
	cmd := newCmd(c, retrieveModelHelp, func(c *context, args []string) {
		if len(args) != 2 {
			log.Fatalln("Incorrect number of arguments. See 'steam help retrieve model'.")
		}

		modelName := args[0]
		cloudName := args[1]

		if _, err := c.remote.GetModelFromCloud(cloudName, modelName); err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Retireved model:", modelName)
	})

	return cmd
}
