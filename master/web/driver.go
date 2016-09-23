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
