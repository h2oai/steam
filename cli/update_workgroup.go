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

	"github.com/spf13/cobra"
)

var updateWorkgroupHelp = `
workgroup [workgroupName]
Change or update a workgroup in the database.
Examples:

	$ steam update workgroup production --desc="A deploy workgroup" --name="deploy"
`

func updateWorkgroup(c *context) *cobra.Command {
	var description, name string
	cmd := newCmd(c, updateWorkgroupHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Invalid usage. See 'steam help update workgroup'.")
		}

		// -- Args --

		workgroupName := args[0]

		// -- Execution --

		workgroup, err := c.remote.GetWorkgroupByName(workgroupName)
		if err != nil {
			log.Fatalln(err) // TODO
		}
		if name == "" {
			name = workgroup.Name
		}
		if description == "" {
			description = workgroup.Description
		}

		if err := c.remote.UpdateWorkgroup(workgroup.Id, name, description); err != nil {
			log.Fatalln(err) // TODO
		}

		// -- Formatting --

		fmt.Println("Successfully updated workgroup:", workgroupName)
	})
	cmd.Flags().StringVarP(&description, "desc", "d", "", "An updated workgroup description.")
	cmd.Flags().StringVarP(&name, "name", "n", "", "An updated workgroup name.")

	return cmd
}
