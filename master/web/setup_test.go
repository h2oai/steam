package web

import (
	"fmt"
	"github.com/h2oai/steamY/lib/fs"
	"github.com/h2oai/steamY/master/az"
	"github.com/h2oai/steamY/master/data"
	web "github.com/h2oai/steamY/srv/web"
	"os"
	"path"
	"path/filepath"
	"testing"
)

const superuser = "superuser"

func setup() (web.Service, az.Directory, error) {
	dbOpts := driverDBOpts{
		"steam",
		"steam",
		"disable",
		superuser,
		superuser,
	}

	// Truncate database tables

	if err := data.Destroy(dbOpts.Name, dbOpts.Username, dbOpts.SSLMode); err != nil {
		return nil, nil, fmt.Errorf("Failed truncating database: %s", err)
	}

	// Determine current directory

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return nil, nil, fmt.Errorf("Failed determining current directory: %s", err)
	}

	//

	opts := driverOpts{
		path.Join(dir, fs.VarDir, "master"),
		":9001",
		":8080",
		"",
		driverYarnOpts{false, "", ""},
		dbOpts,
	}
	return newService(opts)
}

func check(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
