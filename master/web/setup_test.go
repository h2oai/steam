package web

import (
	"flag"
	"os"
	"path"
	"path/filepath"
	"runtime/debug"
	"testing"

	"github.com/h2oai/steamY/lib/fs"
	"github.com/h2oai/steamY/master/az"
	"github.com/h2oai/steamY/master/data"
	"github.com/h2oai/steamY/srv/web"
)

const (
	superuser = "superuser"
)

var h2oFrames = []struct {
	name string
}{
	{"bin_hex"},
	{"mul_hex"},
	{"reg_hex"},
}

var h2oModels = []struct {
	name     string
	category string
}{
	{"bin_gbm", "Binomial"},
	{"bin_glm", "Binomial"},
	{"bin_dpl", "Binomial"},
	{"mul_gbm", "Multinomial"},
	{"mul_dpl", "Multinomial"},
	{"reg_gbm", "Regression"},
	{"reg_glm", "Regression"},
	{"reg_dpl", "Regression"},
}

type test struct {
	t   *testing.T
	svc web.Service
	dir az.Directory
	su  az.Principal
}

var clusterAddress, workingDirectory, compilationServiceAddress string

func init() {
	flag.StringVar(&clusterAddress, "cluster-address", "localhost:54321", "Where the h2o cluster can be reached.")
	flag.StringVar(&workingDirectory, "working-directory", "", "Where the var folder will be located.")
	flag.StringVar(&compilationServiceAddress, "compilation-service-address", ":8080", "Where to find the compilation service.")
}

func newTest(t *testing.T) *test {
	dbOpts := driverDBOpts{
		data.Connection{DbName: "steam", User: "steam", SSLMode: "disable"},
		superuser,
		superuser,
	}

	// Truncate database tables

	if err := data.Destroy(dbOpts.Connection); err != nil {
		t.Fatalf("Failed truncating database: %s", err)
	}

	// Determine current directory

	wd, err := filepath.Abs(filepath.Dir(workingDirectory + "/"))
	if err != nil {
		t.Fatalf("Failed determining current directory: %s", err)
	}

	// Delete any remnant models in models directory

	if err := os.RemoveAll(path.Join(wd, "var/master/model")); err != nil {
		t.Fatalf("Failed removing old model directory: %v", err)
	}

	// Create service instance

	opts := driverOpts{
		path.Join(wd, fs.VarDir, "master"),
		":9001",
		compilationServiceAddress,
		"",
		[2]int{1025, 65535},
		driverYarnOpts{false, "", ""},
		dbOpts,
	}
	svc, dir, err := newService(opts)
	if err != nil {
		t.Fatal(err)
	}

	// Lookup superuser

	su, err := dir.Lookup(superuser)
	if err != nil {
		t.Fatal(err)
	}

	return &test{t, svc, dir, su}
}

func (t *test) log(args ...interface{}) {
	t.t.Log(args...)
}

func (t *test) fail(format string, args ...interface{}) {
	debug.PrintStack()
	t.t.Fatalf(format, args...)
}

func (t *test) nil(arg interface{}) {
	if arg != nil {
		t.fail("unexpected: %s", arg)
	}
}

func (t *test) notnil(arg interface{}) {
	if arg == nil {
		t.fail("unexpected nil: %s", arg)
	}
}

func (t *test) ok(condition bool, format string, args ...interface{}) {
	if !condition {
		t.fail(format, args...)
	}
}

func buildPermissionMap(t *test) map[string]int64 {
	permissions, err := t.svc.GetAllPermissions(t.su)
	t.nil(err)

	permissionMap := make(map[string]int64)
	for _, permission := range permissions {
		permissionMap[permission.Code] = permission.Id
	}

	return permissionMap
}

func buildEntityTypeMap(t *test) map[string]int64 {
	entityTypes, err := t.svc.GetAllEntityTypes(t.su)
	t.nil(err)

	entityTypeMap := make(map[string]int64)
	for _, entityType := range entityTypes {
		entityTypeMap[entityType.Name] = entityType.Id
	}
	return entityTypeMap
}
