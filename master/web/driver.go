package web

import (
	"github.com/h2oai/steamY/lib/fs"
	"github.com/h2oai/steamY/master/az"
	"github.com/h2oai/steamY/master/data"
	web "github.com/h2oai/steamY/srv/web"
	"log"
)

type testOpts struct {
	WorkingDirectory          string
	ClusterProxyAddress       string
	CompilationServiceAddress string
	ScoringServiceHost        string
	Yarn                      testYarnOpts
	DB                        testDBOpts
}

type testDBOpts struct {
	Name              string
	Username          string
	SSLMode           string
	SuperuserName     string
	SuperuserPassword string
}

type testYarnOpts struct {
	KerberosEnabled bool
	Username        string
	Keytab          string
}

func newService(opts testOpts) (web.Service, az.Directory, error) {
	wd, err := fs.MkWorkingDirectory(opts.WorkingDirectory)
	if err != nil {
		log.Fatalln(err)
	}

	ds, err := data.Init(
		opts.DB.Name,
		opts.DB.Username,
		opts.DB.SSLMode,
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
		opts.Yarn.KerberosEnabled,
		opts.Yarn.Username,
		opts.Yarn.Keytab,
	), ds, nil
}
