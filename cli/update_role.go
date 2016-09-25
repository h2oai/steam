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

var updateRoleHelp = `
role [roleName]
Change or update a role in the database.
Examples:

	$ steam update role engineer --desc="A better engineer" --name="science engineer"
`

func updateRole(c *context) *cobra.Command {
	var description, name string
	cmd := newCmd(c, updateRoleHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Invalid usage. See 'steam help update role'.")
		}

		// -- Args --

		roleName := args[0]

		// -- Execution --

		role, err := c.remote.GetRoleByName(roleName)
		if err != nil {
			log.Fatalln(err) // TODO
		}
		if name == "" {
			name = role.Name
		}
		if description == "" {
			description = role.Description
		}

		if err := c.remote.UpdateRole(role.Id, name, description); err != nil {
			log.Fatalln(err) // TODO
		}

		// -- Formatting --

		fmt.Println("Successfully updated role:", roleName)
	})
	cmd.Flags().StringVarP(&description, "desc", "d", "", "An updated role description.")
	cmd.Flags().StringVarP(&name, "name", "n", "", "An updated role name.")

	return cmd
}
