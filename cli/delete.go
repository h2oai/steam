package cli

import (
	"fmt"
	"log"
	// "strconv"

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
	cmd.AddCommand(deleteEngine(c))
	cmd.AddCommand(deleteModel(c))
	cmd.AddCommand(deleteService(c))
	return cmd
}

var deleteCloudHelp = `
cloud [cloudName]
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

		// FIXME

		// if err := c.remote.DeleteCloud(cloudName); err != nil {
		// 	log.Fatalln(err)
		// }

		fmt.Println("Cloud deleted:", cloudName)
	})

	return cmd
}

var deleteEngineHelp = `
engine [engineName]
Deletes a specified engine from the database.
Examples:

	$ steam delete engine engine1
`

func deleteEngine(c *context) *cobra.Command {
	cmd := newCmd(c, deleteEngineHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Incorrect number of arugments. See 'steam help delete engine'.")
		}

		engineName := args[0]

		// FIXME

		// if err := c.remote.DeleteEngine(engineName); err != nil {
		// 	log.Fatalln(err)
		// }

		fmt.Println("Engine deleted:", engineName)
	})

	return cmd
}

var deleteModelHelp = `
model [modelName]
Deletes a sepcified model from the database.
Examples:
	
	$ steam delete model model3
`

func deleteModel(c *context) *cobra.Command {
	cmd := newCmd(c, deleteModelHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Incorrect number of arguments. See 'steam help delete model'.")
		}

		modelName := args[0]

		// FIXME

		// if err := c.remote.DeleteModel(modelName); err != nil {
		// 	log.Fatalln(err)
		// }

		fmt.Println("Model deleted:", modelName)
	})

	return cmd
}

var deleteServiceHelp = `
service [modelName] [port]
Deletes a specified scoring service deployed on the specified port from the database.
Examples:

	$ steam delete service model3 59876
`

func deleteService(c *context) *cobra.Command {
	cmd := newCmd(c, deleteServiceHelp, func(c *context, args []string) {
		if len(args) != 2 {
			log.Fatalln("Incorrect number of arguments. See 'steam help delete service'.")

		}

		// FIXME

		// modelName := args[0]
		// port, err := strconv.Atoi(args[1])
		// if err != nil {
		// 	log.Fatalln(err)
		// }

		// if err := c.remote.DeleteScoringService(modelName, port); err != nil {
		// 	log.Fatalln(err)
		// }

		// fmt.Println("Service deleted on:", modelName)
	})

	return cmd
}
