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

var getServicesHelp = `
services
List all services.
Examples:

	$ steam get services
`

func getServices(c *context) *cobra.Command {
	cmd := newCmd(c, getServicesHelp, func(c *context, args []string) {

		ss, err := c.remote.GetScoringServices(0, 1000)
		if err != nil {
			log.Fatalln(err)
		}

		lines := make([]string, len(ss))
		for i, s := range ss {
			lines[i] = fmt.Sprintf("%d\t%s:%d\t%s\t%s", s.Id, s.Address, s.Port, s.State, fmtAgo(s.CreatedAt))
		}

		c.printt("ID\tAddress\tSTATE\tAGE", lines)
	})

	return cmd
}
