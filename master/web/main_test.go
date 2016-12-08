package web

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/h2oai/steam/lib/fs"
	"github.com/h2oai/steam/master/az"
	"github.com/h2oai/steam/master/data"
)

const (
	superuser = "superuser"
)

var (
	test_compilationServiceAddress string
)

func init() {
	flag.StringVar(&test_compilationServiceAddress, "compilation-service-address", ":8080", "Where to find the compilation service")
}

func TestMain(m *testing.M) {
	flag.Parse()

	os.Exit(m.Run())
}

func testSetup(testType, driver string) (*Service, az.Principal, string) {
	// Determine working directory
	temp, wd := setupWD(testType)
	ds := setupDS(driver, wd)
	pz := setupPz(ds)

	svc := NewService(
		wd,
		ds,
		test_compilationServiceAddress,
		"",
		":9001",
		[2]int{65525, 65535},
		false,
	)

	return svc, pz, temp
}

func setupWD(testType string) (string, string) {
	dir, err := ioutil.TempDir("", testType)
	if err != nil {
		log.Fatalf("Creating temp directory: %+v", err)
	}
	fsDir, err := fs.MkWorkingDirectory(filepath.Join(dir, "var", "master"))
	if err != nil {
		log.Fatalf("Creating working directory: %+v", err)
	}
	return dir, fsDir
}

func setupDS(driver, wd string) *data.Datastore {
	var dbOpts string
	switch driver {
	case "sqlite3":
		dbOpts = filepath.Join(wd, fs.DbDir, "steam.db")
	}

	ds, err := data.NewDatastore(driver, dbOpts)
	if err != nil {
		log.Fatalf("Creating datastore: %+v", err)
	}
	return ds
}

func setupPz(ds *data.Datastore) az.Principal {
	_, err := ds.CreateSuperuser(superuser, superuser)
	if err != nil {
		log.Fatalf("Creating superuser: %+v", err)
	}

	pz, err := ds.Lookup(superuser)
	if err != nil {
		log.Fatalf("Looking up principal", err)
	}

	return pz
}
