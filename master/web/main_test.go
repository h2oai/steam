package web

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/h2oai/steam/lib/fs"
	"github.com/h2oai/steam/master/az"
	"github.com/h2oai/steam/master/data"
)

const (
	admin = "admin"
)

var (
	test_compilationServiceAddress string
	test_compilationService        bool
	test_h2oAddress                string
	test_h2o                       bool
)

var cluster_url, compilation_service_url string

func TestMain(m *testing.M) {
	flag.StringVar(&test_compilationServiceAddress, "compilation-service-address", ":8080", "Where to find the compilation service")
	flag.BoolVar(&test_compilationService, "test-compilation-service", true, "Set to false to skip compilation service tests")
	flag.StringVar(&test_h2oAddress, "h2o-address", "localhost:54321", "Where to locate the H2O instance")
	flag.BoolVar(&test_h2o, "test-h2o", true, "Set to false to skip H2O tests")

	flag.Parse()
	cluster_url = (&url.URL{Scheme: "http", Host: test_h2oAddress}).String()
	compilation_service_url = (&url.URL{Scheme: "http", Host: test_compilationServiceAddress}).String()
	if test_h2o && !pingExternal(cluster_url) {
		log.Fatalf("unable to reach h2o at %s", cluster_url)
	}
	if test_compilationService && !pingExternal(compilation_service_url) {
		log.Fatalf("unable to reach compilation service at %s", compilation_service_url)
	}

	os.Exit(m.Run())
}

func testSetup(testType, driver string) (*Service, az.Principal, string) {
	// Determine working directory
	temp, wd := setupWD(testType)
	ds := setupDS(driver, wd)
	pz := setupPz(ds)

	svc := NewService(
		"test",
		wd,
		ds,
		nil,
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
	dbOpts := data.DBOpts{
		Driver:    driver,
		AdminName: admin,
		AdminPass: "adminadmin",
	}
	switch driver {
	case "sqlite3":
		dbOpts.Path = filepath.Join(wd, fs.DbDir, "steam.db")
	}

	ds, err := data.NewDatastore(dbOpts, false)
	if err != nil {
		log.Fatalf("Creating datastore: %+v", err)
	}
	return ds
}

func setupPz(ds *data.Datastore) az.Principal {
	pz, err := ds.Lookup(admin)
	if err != nil {
		log.Fatalf("Looking up principal", err)
	}

	return pz
}

func pingExternal(u string) bool {
	_, err := http.Get(u)
	if err != nil {
		return false
	}
	return true
}
