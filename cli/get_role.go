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

var getRoleHelp = `
role [roleName]
View detailed role information.
Examples:

	$ steam get role engineer
`

func getRole(c *context) *cobra.Command {
	cmd := newCmd(c, getRoleHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Invalid usage. See 'steam help get role'.")
		}

		// -- Args --

		roleName := args[0]

		// -- Execution --

		role, err := c.remote.GetRoleByName(roleName)
		if err != nil {
			log.Fatalln(err) //FIXME format error
		}

		perms, err := c.remote.GetPermissionsForRole(role.Id)
		if err != nil {
			log.Fatalln(err) //FIXME format error
		}

		idents, err := c.remote.GetIdentitiesForRole(role.Id)
		if err != nil {
			log.Fatalln(err) //FIXME format error
		}

		// -- Formatting --

		base := []string{
			fmt.Sprintf("DESCRIPTION:\t%s", role.Description),
			fmt.Sprintf("ID:\t%d", role.Id),
			fmt.Sprintf("AGE:\t%s", fmtAgo(role.Created)),
		}
		c.printt("\t"+role.Name, base)

		fmt.Println("IDENTITES:", len(idents))
		if len(idents) > 0 {
			ids := make([]string, len(idents))
			for i, id := range idents {
				ids[i] = fmt.Sprintf("%s\t%s\t%s",
					id.Name,
					identityStatus(id.IsActive),
					fmtAgo(id.LastLogin),
				)
			}
			c.printt("NAME\tSTATUS\tLAST LOGIN", ids)
		}

		fmt.Println("PERMISSIONS")
		for _, p := range perms {
			fmt.Println(p.Description)
		}

	})

	return cmd
}
