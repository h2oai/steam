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

var createWorkgroupHelp = `
workgroup [workgroupName] 
Creates a user permissions workgroup.
Examples:

	$ steam create workgroup production --desc="The production group"
`

func createWorkgroup(c *context) *cobra.Command {
	var description string

	cmd := newCmd(c, createWorkgroupHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Incorrect usage. See 'steam help create workgroup'.")
		}

		workgroupName := args[0]

		workgroupId, err := c.remote.CreateWorkgroup(workgroupName, description)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Created workgroup", workgroupName, "ID:", workgroupId)
	})
	cmd.Flags().StringVarP(&description, "desc", "d", "", "A description for this workgroup")

	return cmd
}
