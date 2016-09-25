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

var getIdentitiesHelp = `
identities
List all identities.
Examples:

	$ steam get identities --workgroup=production
`

func getIdentities(c *context) *cobra.Command {
	var workgroup, role string

	cmd := newCmd(c, getIdentitiesHelp, func(c *context, args []string) {

		// -- Execution --

		var ids []*web.Identity

		// Both workgroup and role set
		if workgroup != "" && role != "" {
			log.Fatalln("Cannot use both workgroup and role at the same time.")

			// Workgroup set
		} else if workgroup != "" {
			wg, err := c.remote.GetWorkgroupByName(workgroup)
			if err != nil {
				log.Fatalln(err)
			}
			ids, err = c.remote.GetIdentitiesForWorkgroup(wg.Id)
			if err != nil {
				log.Fatalln(err)
			}

			// Role set
		} else if role != "" {
			rl, err := c.remote.GetRoleByName(role)
			if err != nil {
				log.Fatalln(err)
			}

			ids, err = c.remote.GetIdentitiesForRole(rl.Id)
			if err != nil {
				log.Fatalln(err)
			}

			// Neither set
		} else {
			var err error
			ids, err = c.remote.GetIdentities(0, 10000)
			if err != nil {
				log.Fatalln(err)
			}
		}

		// -- Formatting --

		lines := make([]string, len(ids))
		for i, id := range ids {
			lines[i] = fmt.Sprintf("%s\t%d\t%s\t%s",
				id.Name,
				id.Id,
				fmtAgo(id.LastLogin),
				fmtAgo(id.Created))
		}

		c.printt("NAME\tID\tLAST LOGIN\tAGE", lines)
	})

	cmd.Flags().StringVarP(&workgroup, "workgroup", "w", "", "Search for users by workgroup.")
	cmd.Flags().StringVarP(&role, "role", "r", "", "Search for users by role.")

	return cmd
}
