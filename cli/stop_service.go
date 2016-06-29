package cli

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var stopServiceHelp = `
service [serviceId]
Stop a scoring service.
Examples:

    $ steam stop service model3 
`

func stopService(c *context) *cobra.Command {
	cmd := newCmd(c, stopServiceHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Invalid number of arguments. See 'steam help stop service'.")
		}

		serviceId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatalln("Invalid usage of serviceId %s: Integer value required: %v", args[0], err)
		}

		if err := c.remote.StopScoringService(serviceId); err != nil {
			log.Fatalln(err)
		}

		log.Println("Service", serviceId, "stopped.")
	})

	return cmd
}
