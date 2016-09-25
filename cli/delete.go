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

var deleteHelp = `
delete [resource-type]
Deletes the specified resource from the database.
Examples:

    $ steam delete cluster
`

func delete_(c *context) *cobra.Command {
	cmd := newCmd(c, deleteHelp, nil)
	cmd.AddCommand(deleteCluster(c))
	cmd.AddCommand(deleteEngine(c))
	cmd.AddCommand(deleteModel(c))
	cmd.AddCommand(deleteService(c))
	cmd.AddCommand(deleteRole(c))
	cmd.AddCommand(deleteWorkgroup(c))
	return cmd
}
