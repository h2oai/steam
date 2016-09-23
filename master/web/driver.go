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

package web

import (
	"log"

	"github.com/h2oai/steam/lib/fs"
	"github.com/h2oai/steam/master/az"
	"github.com/h2oai/steam/master/data"
	web "github.com/h2oai/steam/srv/web"
)

type driverOpts struct {
	WorkingDirectory          string
	ClusterProxyAddress       string
	CompilationServiceAddress string
	ScoringServiceHost        string
	ScoringServicePorts       [2]int
	Yarn                      driverYarnOpts
	DB                        driverDBOpts
}

type driverDBOpts struct {
	DBPath            string
	SuperuserName     string
	SuperuserPassword string
}

type driverYarnOpts struct {
	KerberosEnabled bool
	Username        string
	Keytab          string
}

func newService(opts driverOpts) (web.Service, az.Directory, error) {
	wd, err := fs.MkWorkingDirectory(opts.WorkingDirectory)
	if err != nil {
		log.Fatalln(err)
	}

	ds, err := data.Create(
		opts.DB.DBPath,
		opts.DB.SuperuserName,
		opts.DB.SuperuserPassword,
	)

	if err != nil {
		log.Fatalln(err)
	}

	return NewService(
		wd,
		ds,
		opts.CompilationServiceAddress,
		opts.ScoringServiceHost,
		opts.ClusterProxyAddress,
		opts.ScoringServicePorts,
		opts.Yarn.KerberosEnabled,
		opts.Yarn.Username,
		opts.Yarn.Keytab,
	), ds, nil
}
