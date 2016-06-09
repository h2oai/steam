package cli

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var stopHelp = `
stop [resource-type]
Stop the specified resource.
Examples:

    $ steam stop cloud
`

func stop(c *context) *cobra.Command {
	cmd := newCmd(c, stopHelp, nil)
	cmd.AddCommand(stopCloud(c))
	cmd.AddCommand(stopService(c))
	return cmd
}

var stopCloudHelp = `
cloud [cloudName]
Stop a cloud.
Examples:

    $ steam stop cloud cloud42
`

func stopCloud(c *context) *cobra.Command {
	var (
		kerberos, force  bool
		username, keytab string
	)

	cmd := newCmd(c, stopCloudHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Missing cloudName. See 'steam help stop cloud'.")
		}

		name := args[0]
		// --- add additional args here ---

		if err := c.remote.StopCloud(name); err != nil {
			log.Fatalln(err)
		}
		// if err := yarn.StopCloud(kerberos, name, id, job, username, keytab); err != nil {
		// 	log.Fatalln(err)
		// }

	})

	cmd.Flags().BoolVar(&kerberos, "kerberos", true, "Set false on systems with no kerberos authentication.")
	cmd.Flags().StringVar(&username, "username", "", "The valid kerberos username.")
	cmd.Flags().StringVar(&keytab, "keytab", "", "The name of the keytab file to use")
	cmd.Flags().BoolVar(&force, "force", false, "Force-kill all H2O instances in the cloud")

	return cmd
}

var stopServiceHelp = `
service [modelName] [port]
Stop a scoring service.
Examples:

    $ steam stop service model3 
`

func stopService(c *context) *cobra.Command {
	cmd := newCmd(c, stopServiceHelp, func(c *context, args []string) {
		if len(args) != 2 {
			log.Fatalln("Invalid number of arguments. See 'steam help stop service'.")
		}

		modelName := args[0]
		port, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatalln(err)
		}

		if err := c.remote.StopScoringService(modelName, port); err != nil {
			log.Fatalln(err)
		}

		log.Println("Service stopped on:", modelName)
	})

	return cmd
}
