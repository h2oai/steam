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

var getIdentityHelp = `
identity [username]
View detailed user information.
Examples:

	$ steam get identity minsky
`

func getIdentity(c *context) *cobra.Command {
	cmd := newCmd(c, getIdentityHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Invalid usage. See 'steam help get identity'.")
		}

		// -- Args --

		username := args[0]

		// -- Execution --

		identity, err := c.remote.GetIdentityByName(username)
		if err != nil {
			log.Fatalln(err) //FIXME format error
		}

		workgroups, err := c.remote.GetWorkgroupsForIdentity(identity.Id)
		if err != nil {
			log.Fatalln(err) //FIXME format error
		}

		roles, err := c.remote.GetRolesForIdentity(identity.Id)
		if err != nil {
			log.Fatalln(err) //FIXME format error
		}

		permissions, err := c.remote.GetPermissionsForIdentity(identity.Id)
		if err != nil {
			log.Fatalln(err) //FIXME format error
		}

		// -- Formatting --

		base := []string{
			fmt.Sprintf("STATUS:\t%s", identityStatus(identity.IsActive)),
			fmt.Sprintf("LAST LOGIN:\t%s", fmtAgo(identity.LastLogin)),
			fmt.Sprintf("ID:\t%d", identity.Id),
			fmt.Sprintf("AGE:\t%s", fmtAgo(identity.Created)),
		}

		c.printt("\t"+identity.Name, base)

		ws := make([]string, len(workgroups))
		for i, w := range workgroups {
			ws[i] = fmt.Sprintf("%s\t%s", w.Name, w.Description)
		}
		c.printt("WORKGROUP\tDESCRIPTION", ws)

		rs := make([]string, len(roles))
		for i, r := range roles {
			rs[i] = fmt.Sprintf("%s\t%s", r.Name, r.Description)
		}
		c.printt("ROLE\tDESCRIPTION", rs)

		fmt.Println("PERMISSIONS")
		for _, p := range permissions {
			fmt.Println(p.Description)
		}

	})

	return cmd
}
