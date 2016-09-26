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

	"github.com/spf13/cobra"
)

var linkHelp = `
link [permission-type]
Add authentication permissions.
Examples:

	$ steam link permission
`

func link(c *context) *cobra.Command {
	cmd := newCmd(c, linkHelp, nil)
	cmd.AddCommand(linkRole(c))
	cmd.AddCommand(linkIdentity(c))
	return cmd
}

/* getPermissionIds
 *
 * Support function to return ids from codes
 */
func getPermissionIds(c *context, codes ...string) ([]int64, error) {
	permissions, err := c.remote.GetSupportedPermissions()
	if err != nil {
		return nil, err //FIXME format error
	}

	permissionsMap := make(map[string]int64)
	for _, p := range permissions {
		permissionsMap[p.Code] = p.Id
	}

	ids := make([]int64, len(codes))
	for i, c := range codes {
		var ok bool
		ids[i], ok = permissionsMap[c]
		if !ok {
			return nil, fmt.Errorf("Failed getting permissions: cannot locate permission with corresponding code: %s", c)
		}
	}

	return ids, nil
}
