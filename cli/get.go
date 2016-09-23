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

import "github.com/spf13/cobra"

var getHelp = `
get [resource-type]
List or view resources of the specified type.
Examples:

    $ steam get clusters
`

func get(c *context) *cobra.Command {
	cmd := newCmd(c, getHelp, nil)
	cmd.AddCommand(getClusters(c))
	cmd.AddCommand(getCluster(c))
	cmd.AddCommand(getEngines(c))
	cmd.AddCommand(getEngine(c))
	cmd.AddCommand(getModels(c))
	cmd.AddCommand(getModel(c))
	cmd.AddCommand(getServices(c))
	cmd.AddCommand(getService(c))
	cmd.AddCommand(getIdentities(c))
	cmd.AddCommand(getIdentity(c))
	cmd.AddCommand(getRoles(c))
	cmd.AddCommand(getRole(c))
	cmd.AddCommand(getWorkgroups(c))
	cmd.AddCommand(getWorkgroup(c))
	cmd.AddCommand(getPermissions(c))
	cmd.AddCommand(getEntities(c))
	cmd.AddCommand(getHistory(c))
	return cmd
}

/* getClusterTypes
 *
 * Support function to map TypeIds to cluster types (i.e. yarn, external, etc.)
 */
func getClusterTypes(c *context) (map[int64]string, error) {
	raw, err := c.remote.GetSupportedClusterTypes()
	if err != nil {
		return nil, err //FIXME format error
	}

	ret := make(map[int64]string)
	for _, ct := range raw {
		ret[ct.Id] = ct.Name
	}

	return ret, nil
}

/* identityStatus
 *
 * Support function to return active or inactive from bool
 */
func identityStatus(isActive bool) string {
	if isActive {
		return "Active"
	}

	return "Inactive"
}
