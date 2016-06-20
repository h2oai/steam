package cli

import (
	// "fmt"
	"log"

	"github.com/spf13/cobra"
)

var importHelp = `
import
Import a resource of the specified type into steam.
Examples:

	$ steam import model
`

func import_(c *context) *cobra.Command {
	cmd := newCmd(c, importHelp, nil)
	cmd.AddCommand(importModel(c))
	return cmd
}

var importModelHelp = `
model [modelName] [cloudName]
Import a model from an H2O cluster into steam
Examples:

	$ steam import model model3 cloud42
`

func importModel(c *context) *cobra.Command {
	cmd := newCmd(c, importModelHelp, func(c *context, args []string) {
		if len(args) != 2 {
			log.Fatalln("Incorrect number of arguments. See 'steam help import model'.")
		}

		// FIXME

		// modelName := args[0]
		// cloudName := args[1]

		// if _, err := c.remote.GetModelFromCloud(cloudName, modelName); err != nil {
		// 	log.Fatalln(err)
		// }

		// fmt.Println("Retireved model:", modelName)
	})

	return cmd
}
