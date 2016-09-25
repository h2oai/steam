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

var getRolesHelp = `
roles
List permission roles.
Exampes:

	$ steam get roles --identity=minksy
`

func getRoles(c *context) *cobra.Command {
	var identityName string
	cmd := newCmd(c, getRolesHelp, func(c *context, args []string) {
		var rs []*web.Role

		// -- Execution --

		if identityName == "" {
			var err error
			rs, err = c.remote.GetRoles(0, 10000)
			if err != nil {
				log.Fatalln(err)
			}

			// Identity set
		} else {
			identity, err := c.remote.GetIdentityByName(identityName)
			if err != nil {
				log.Fatalln(err)
			}

			rs, err = c.remote.GetRolesForIdentity(identity.Id)
			if err != nil {
				log.Fatalln(err)
			}
		}

		// -- Formatting --

		lines := make([]string, len(rs))
		for i, r := range rs {
			lines[i] = fmt.Sprintf("%s\t%d\t%s\t%s",
				r.Name,
				r.Id,
				r.Description,
				fmtAgo(r.Created))
		}

		c.printt("NAME\tID\tDESCRIPTION\tAGE", lines)
	})
	cmd.Flags().StringVarP(&identityName, "identity", "i", "", "Search by identity.")
	return cmd
}
