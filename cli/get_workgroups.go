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

// FIXME use ByName remote calls

import (
	"fmt"
	"log"

	"github.com/h2oai/steam/srv/web"
	"github.com/spf13/cobra"
)

var getWorkgroupsHelp = `
workgroups
List permission workgroups.
Examples:

	$ steam get workgroups --identity=2
`

func getWorkgroups(c *context) *cobra.Command {
	var identityName string
	cmd := newCmd(c, getWorkgroupsHelp, func(c *context, args []string) {

		// -- Execution --

		var ws []*web.Workgroup

		if identityName == "" {
			var err error
			ws, err = c.remote.GetWorkgroups(0, 10000)
			if err != nil {
				log.Fatalln(err)
			}

			// Identity set
		} else {
			identity, err := c.remote.GetIdentityByName(identityName)
			if err != nil {
				log.Fatalln(err)
			}

			ws, err = c.remote.GetWorkgroupsForIdentity(identity.Id)
			if err != nil {
				log.Fatalln(err)
			}
		}

		// -- Formatting --

		lines := make([]string, len(ws))
		for i, w := range ws {
			lines[i] = fmt.Sprintf("%s\t%d\t%s\t%s",
				w.Name,
				w.Id,
				w.Description,
				fmtAgo(w.Created))
		}

		c.printt("NAME\tID\tDESCRIPTION\tAGE", lines)
	})
	cmd.Flags().StringVarP(&identityName, "identity", "i", "", "Search by identity.")

	return cmd
}
