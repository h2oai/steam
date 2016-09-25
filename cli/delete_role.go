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

var deleteRoleHelp = `
role [roleName]
Deletes a role from the database.
Examples:

	$ steam delete role engineer
`

func deleteRole(c *context) *cobra.Command {
	cmd := newCmd(c, deleteRoleHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Invalid usage. See 'steam help delete role'.")
		}

		// -- Args --

		roleName := args[0]

		// -- Execution --

		role, err := c.remote.GetRoleByName(roleName)
		if err != nil {
			log.Fatalln(err) // FIXME
		}

		if err := c.remote.DeleteRole(role.Id); err != nil {
			log.Fatalln(err) // FIXME
		}

		// -- Formatting --

		fmt.Println("Successfully deleted role:", roleName)
	})

	return cmd
}
