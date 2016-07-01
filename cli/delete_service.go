package cli

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteServiceHelp = `
service [serviceId]
Deletes a specified scoring service from the database.
Examples:

	$ steam delete service 76
`

func deleteService(c *context) *cobra.Command {
	cmd := newCmd(c, deleteServiceHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Incorrect number of arguments. See 'steam help delete service'.")

		}

		serviceId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatalf("Invalid usage of serviceId %q: expecting integer: %v", args[0], err)
		}

		if err := c.remote.DeleteScoringService(serviceId); err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Service deleted on:", serviceId)
	})

	return cmd
}
