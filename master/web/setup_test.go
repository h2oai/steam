package web

import (
	"github.com/h2oai/steamY/lib/fs"
	"github.com/h2oai/steamY/master/az"
	"github.com/h2oai/steamY/master/data"
	"github.com/h2oai/steamY/srv/web"
	"os"
	"path"
	"path/filepath"
	"testing"
)

const superuser = "superuser"

type test struct {
	t   *testing.T
	svc web.Service
	dir az.Directory
	su  az.Principal
}

func newTest(t *testing.T) *test {
	dbOpts := driverDBOpts{
		"steam",
		"steam",
		"disable",
		superuser,
		superuser,
	}

	// Truncate database tables

	if err := data.Destroy(dbOpts.Name, dbOpts.Username, dbOpts.SSLMode); err != nil {
		t.Fatalf("Failed truncating database: %s", err)
	}

	// Determine current directory

	wd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		t.Fatalf("Failed determining current directory: %s", err)
	}

	// Create service instance

	opts := driverOpts{
		path.Join(wd, fs.VarDir, "master"),
		":9001",
		":8080",
		"",
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
	permissions, err := t.svc.GetSupportedPermissions(t.su)
	t.nil(err)

	permissionMap := make(map[string]int64)
	for _, permission := range permissions {
		permissionMap[permission.Code] = permission.Id
	}

	return permissionMap
}

func buildEntityTypeMap(t *test) map[string]int64 {
	entityTypes, err := t.svc.GetSupportedEntityTypes(t.su)
	t.nil(err)

	entityTypeMap := make(map[string]int64)
	for _, entityType := range entityTypes {
		entityTypeMap[entityType.Name] = entityType.Id
	}
	return entityTypeMap
}
