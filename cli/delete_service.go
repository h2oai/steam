/*
  Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

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
