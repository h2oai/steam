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

var deactivateIdentityHelp = `
identity [identityName]
Deactivate a user.
Examples:

	$ steam deactivate identity minsky
`

func deactivateIdentity(c *context) *cobra.Command {
	cmd := newCmd(c, deactivateIdentityHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Invalid usage. See 'steam help deactivate identity'.")
		}

		// -- Args --

		identityName := args[0]

		// -- Execution --

		identity, err := c.remote.GetIdentityByName(identityName)
		if err != nil {
			log.Fatalln(err) // TODO
		}

		if err := c.remote.DeactivateIdentity(identity.Id); err != nil {
			log.Fatalln(err) // TODO
		}

		// -- Formatting --

		fmt.Println("Successfully deactivated user", identityName)
	})

	return cmd
}
