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

	"github.com/h2oai/steam/srv/web"
	"github.com/spf13/cobra"
)

var getPermissionsHelp = `
permissions
Get permissions and their corresponding codes.
Examples:

	$ steam get permissions --role=2
`

func getPermissions(c *context) *cobra.Command {
	var roleName, identityName string
	cmd := newCmd(c, getPermissionsHelp, func(c *context, args []string) {

		// -- Execution --

		var ps []*web.Permission

		if roleName != "" && identityName != "" {
			log.Fatalln("Cannot view by role and identity.")

			// Role set
		} else if roleName != "" {
			role, err := c.remote.GetRoleByName(roleName)
			if err != nil {
				log.Fatalln(err)
			}

			ps, err = c.remote.GetPermissionsForRole(role.Id)
			if err != nil {
				log.Fatalln(err)
			}

			// Identity set
		} else if identityName != "" {
			identity, err := c.remote.GetIdentityByName(identityName)
			if err != nil {
				log.Fatalln(err)
			}

			ps, err = c.remote.GetPermissionsForIdentity(identity.Id)
			if err != nil {
				log.Fatalln(err)
			}

			// Niether set
		} else {
			var err error
			ps, err = c.remote.GetSupportedPermissions()
			if err != nil {
				log.Fatalln(err)
			}
		}

		// -- Formatting --

		lines := make([]string, len(ps))
		for i, p := range ps {
			lines[i] = fmt.Sprintf("%d\t%s\t%s", p.Id, p.Description, p.Code)
		}

		c.printt("ID\tDESCRIPTION\tCODE", lines)
	})

	cmd.Flags().StringVarP(&roleName, "role", "r", "", "View permissions for a role.")
	cmd.Flags().StringVarP(&identityName, "identity", "i", "", "View permissions for an identity.")
	return cmd
}
