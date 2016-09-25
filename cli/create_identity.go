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

var createIdentityHelp = `
identity [username] [password]
Creates a user.
Examples:

	$ steam create minsky m1n5kypassword
`

func createIdentity(c *context) *cobra.Command {
	cmd := newCmd(c, createIdentityHelp, func(c *context, args []string) {
		if len(args) != 2 {
			log.Fatalln("Incorrect usage. See 'steam help create identity'.")
		}

		username := args[0]
		password := args[1]

		identityId, err := c.remote.CreateIdentity(username, password)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Created user", username, "ID:", identityId)
	})

	return cmd
}
