package cli

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var importModelHelp = `
model [clusterId] [modelName]
Import a model from an H2O cluster into steam
Examples:

	$ steam import model 42 model3
`

func importModel(c *context) *cobra.Command {
	cmd := newCmd(c, importModelHelp, func(c *context, args []string) {
		if len(args) != 2 {
			log.Fatalln("Incorrect number of arguments. See 'steam help import model'.")
		}

		clusterId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatalf("Incorrect value for clusterId: %s: %v", args[0], err)
		}
		modelName := args[1]

		if _, err := c.remote.ImportModelFromCluster(clusterId, modelName); err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Retireved model:", modelName)
	})

	return cmd
}
