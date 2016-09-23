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

var createRoleHelp = `
role [roleName] 
Creates a user permissions role.
Examples:

	$ steam create role engineer --desc="a default engineer role"
`

func createRole(c *context) *cobra.Command {
	var description string

	cmd := newCmd(c, createRoleHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Incorrect usage. See 'steam help create role'.")
		}

		roleName := args[0]

		roleId, err := c.remote.CreateRole(roleName, description)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Created role", roleName, "ID:", roleId)
	})
	cmd.Flags().StringVarP(&description, "desc", "d", "", "A description for this role")

	return cmd
}
