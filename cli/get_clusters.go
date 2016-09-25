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

	"github.com/h2oai/steam/master/data"
	"github.com/spf13/cobra"
)

var getClustersHelp = `
clusters
List all clusters.
Examples

	$ steam get clusters
`

func getClusters(c *context) *cobra.Command {
	var details bool

	cmd := newCmd(c, getClustersHelp, func(c *context, args []string) {
		cs, err := c.remote.GetClusters(0, 10000)
		if err != nil {
			log.Fatalln(err)
		}

		clusterTypes, err := getClusterTypes(c)
		if err != nil {
			log.Fatalln(err) //FIXME format error
		}

		lines := make([]string, len(cs))
		for i, cl := range cs {
			if !details {
				lines[i] = fmt.Sprintf("%s\t%d\t%s\t%s\t%s\t%s", cl.Name, cl.Id,
					cl.Address, cl.State, clusterTypes[cl.TypeId], fmtAgo(cl.CreatedAt))
			} else {
				if (cl.State != data.StoppedState) || (cl.State != data.DisconnectedState) {
					info, err := c.remote.GetClusterStatus(cl.Id)
					if err != nil {
						log.Fatalln(err)
					}

					lines[i] = fmt.Sprintf("%s\t%d\t%s\t%s\t%d\t%s\t%s\t%s", cl.Name,
						cl.Id, cl.Address, info.MaxMemory, info.TotalCpuCount,
						info.Status, clusterTypes[cl.TypeId], fmtAgo(cl.CreatedAt))
				} else {
					lines[i] = fmt.Sprintf("%s\t%d\t%s\tNA\tNA\t%s\t%s\t%s", cl.Name,
						cl.Id, cl.Address, cl.State, clusterTypes[cl.TypeId],
						fmtAgo(cl.CreatedAt))
				}
			}
		}

		var header string
		if !details {
			header = "NAME\tID\tADDRESS\tSTATE\tTYPE\tAGE"
		} else {
			header = "NAME\tID\tADDRESS\tMEMORY\tCPUs\tSTATE\tTYPE\tAGE"
		}

		c.printt(header, lines)
	})

	cmd.Flags().BoolVarP(&details, "details", "d", false, "Detailed cluster information")

	return cmd
}
