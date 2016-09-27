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
	"strconv"

	"github.com/spf13/cobra"
)

var getHistoryHelp = `
history [entityType] [entityName | entityId]
View history for a specified entity. Only enity types identity, role, and workgroup can be called by name.
Examples:

	$ steam get history identity minsky
	$ steam get history cluster 9
`

func getHistory(c *context) *cobra.Command {
	cmd := newCmd(c, getHistoryHelp, func(c *context, args []string) {
		if len(args) != 2 {
			log.Fatalln("Invalid usage. See 'steam help get history'.")
		}

		// -- Load command line args --

		entityType := args[0]
		entityName := args[1]

		entityTypeId, err := getEntityId(c, entityType)
		if err != nil {
			log.Fatalln(err)
		}

		// allow passing of identity, role, and workgroup by name
		id, err := strconv.ParseInt(entityName, 10, 64)
		if err != nil {
			switch entityType {
			case "identity":
				identity, err := c.remote.GetIdentityByName(entityName)
				if err != nil {
					log.Fatalln(err) // TODO
				}
				id = identity.Id

			case "role":
				role, err := c.remote.GetRoleByName(entityName)
				if err != nil {
					log.Fatalln(err) // TODO
				}
				id = role.Id

			case "workgroup":
				workgroup, err := c.remote.GetWorkgroupByName(entityName)
				if err != nil {
					log.Fatalln(err) // TODO
				}
				id = workgroup.Id

			default:
				log.Fatalf("Invalid usage for id %s: expecting int: %v", entityName, err)
			}
		}

		// -- Execution --

		hs, err := c.remote.GetEntityHistory(entityTypeId, id, 0, 10000)
		if err != nil {
			log.Fatalln(err) // TODO
		}

		lines := make([]string, len(hs))
		for i, h := range hs {
			lines[i] = fmt.Sprintf("%d\t%s\t%s\t%s", h.IdentityId, h.Action, h.Description, fmtAgo(h.CreatedAt))
		}

		c.printt("USER\tACTION\tDESCRITPION\tTIME", lines)
	})

	return cmd
}

func getEntityId(c *context, entityName string) (int64, error) {
	es, err := c.remote.GetSupportedEntityTypes()
	if err != nil {
		return 0, err
	}

	for _, e := range es {
		if e.Name == entityName {
			return e.Id, nil
		}
	}

	return 0, fmt.Errorf("Failed to locate entity type %s", entityName)
}
