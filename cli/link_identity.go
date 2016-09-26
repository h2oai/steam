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

var linkIdentityHelp = `
identity [identityName] [role | workgroup] [entityName]
Add a permissions entity group to an identity.
Example:

	$ steam link identity minksy role engineer
`

func linkIdentity(c *context) *cobra.Command {
	cmd := newCmd(c, linkIdentityHelp, func(c *context, args []string) {
		if len(args) != 3 {
			log.Fatalln("Invalid usage. See 'steam help link identity'.")
		}

		// -- Args --

		identityName := args[0]
		entityType := args[1]
		entityName := args[2]

		// -- Execution --

		identity, err := c.remote.GetIdentityByName(identityName)
		if err != nil {
			log.Fatalf("Identity %s not found: %v", identityName, err)
		}

		if entityType == "role" { // CASE: Link identity -> role
			entity, err := c.remote.GetRoleByName(entityName)
			if err != nil {
				log.Fatalf("Role %s not found: %v", entityName, err)
			}

			if err := c.remote.LinkIdentityAndRole(identity.Id, entity.Id); err != nil {
				log.Fatalf("Failed to link identity %s to role %s: %v",
					identityName, entityName, err)
			}
		} else if entityType == "workgroup" { // CASE: Link identity -> workgroup
			entity, err := c.remote.GetWorkgroupByName(entityName)
			if err != nil {
				log.Fatalf("Workgroup %s not found: %v", entityName, err)
			}

			if err := c.remote.LinkIdentityAndWorkgroup(identity.Id, entity.Id); err != nil {
				log.Fatalf("Failed to link identity %s to workgroup %s: %v",
					identityName, entityName, err)
			}
		} else {
			log.Fatalln("Entity type must be either role or workgroup")
		}

		// -- Formatting --

		fmt.Println("Successfully linked identity", identityName, "to", entityName)
	})

	return cmd
}
