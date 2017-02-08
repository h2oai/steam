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

package data

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/h2oai/steam/master/auth"

	"time"

	"github.com/fatih/color"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/doug-martin/goqu.v3"
	_ "gopkg.in/doug-martin/goqu.v3/adapters/postgres"
	_ "gopkg.in/doug-martin/goqu.v3/adapters/sqlite3"
	_ "gopkg.in/mattn/go-sqlite3.v1"
)

const (
	VERSION       = "1.1.0"
	standard_user = "standard user"
)

var standard_user_init_permissions = []string{"ManageCluster", "ViewCluster", "ViewEngine", "ViewUser", "ManageKeytab", "ViewKeytab"}

type metadata map[string]string

type (
	Datastore struct {
		db    *goqu.Database
		users *userTable

		// Internal mapping for printing
		EntityTypeMap map[int64]string
		PermissionMap map[int64]permission_map

		// Enum References in Database
		ClusterType clusterTypeKeys
		EntityType  entityTypeKeys
		State       stateKeys
		Permission  permissionKeys
		// Permission checks
		ViewPermission   map[int64]int64
		ManagePermission map[int64]int64
	}

	DBOpts struct {
		Driver string

		// SQLite Flags
		Path string

		// Postgres Flags
		Name              string
		User              string
		Pass              string
		Host              string
		Port              string
		ConnectionTimeout string
		SSLMode           string
		SSLCert           string
		SSLKey            string
		SSLRootCert       string

		// Auth Flags
		AdminName string
		AdminPass string

		Flags uint
	}
)

// --- Flags ---
const (
	Debug uint = 1 << iota
)

// --- Enums ---
var (
	States      states
	EntityTypes entity_types
)

func init() {
	pflag.BoolVarP(&debug, "debug", "d", false, "Set to enable debug mode")

	States.init()
	EntityTypes.init()
}

func NewDatastore(dbOpts DBOpts, forceAdmin bool) (*Datastore, error) {
	if dbOpts.Flags&Debug != 0 {
		debug = true
	}

	// Connect to db
	db, err := open(dbOpts)
	if err != nil {
		return nil, errors.Wrap(err, "connecting to database")
	}

	primed, err := IsPrimed(db)
	if err != nil {
		return nil, errors.Wrap(err, "checking if database is primed")
	} else if !primed {
		if forceAdmin {
			dbOpts.AdminName, dbOpts.AdminPass, err = setAdmin(dbOpts)
			if err != nil {
				return nil, errors.Wrap(err, "setting up local admin")
			}
		} else {
			if dbOpts.AdminName == "" || dbOpts.AdminPass == "" {
				return nil, errors.New("First time Steam launch requires setting" +
					" the local admin: please use './steam --set-admin' or set" +
					" the flags '--admin-name' and '--admin-password'")
			}

			if err := auth.ValidateUsername(dbOpts.AdminName); err != nil {
				return nil, errors.Wrap(err, "validating username")
			}
			if err := auth.ValidatePassword(dbOpts.AdminPass); err != nil {
				return nil, errors.Wrap(err, "validating password")
			}
		}
		if err := prime(db); err != nil {
			return nil, errors.Wrap(err, "priming database")
		}
	} else if forceAdmin {
		dbOpts.AdminName, dbOpts.AdminPass, err = setAdmin(dbOpts)
		if err != nil {
			return nil, errors.Wrap(err, "setting local admin")
		}
	}

	ds, err := initDatastore(db)
	if err != nil {
		return nil, errors.Wrap(err, "initializing datastore")
	}

	if !primed || forceAdmin {
		admins, err := ds.ReadIdentities(ForAdminWorkgroup)
		if err != nil {
			return nil, errors.Wrap(err, "searching for admins")
		}
		if len(admins) > 0 {
			hashPassword, err := auth.HashPassword(dbOpts.AdminPass)
			if err != nil {
				return nil, errors.Wrap(err, "hashing password")
			}
			if err := ds.UpdateIdentity(admins[0].Id,
				WithName(dbOpts.AdminName), WithPassword(hashPassword),
			); err != nil {
				return nil, err
			}
		} else {
			if _, err := ds.createAdmin(dbOpts.AdminName, dbOpts.AdminPass); err != nil {
				return nil, err
			}
		}
	}

	return ds, nil
}

func setAdmin(dbOpts DBOpts) (string, string, error) {
	adminName, adminPass := dbOpts.AdminName, dbOpts.AdminPass
	if strings.TrimSpace(adminName) == "" {
		r := bufio.NewReader(os.Stdin)

		fmt.Print("Steam local admin username: ")

		name, err := r.ReadString('\n')
		if err != nil {
			return "", "", err
		}
		adminName = strings.Trim(name, "\n")
	}

	if err := auth.ValidateUsername(adminName); err != nil {
		return "", "", errors.Wrap(err, "validating username")
	}

	if strings.TrimSpace(adminPass) == "" {
		fmt.Print("Steam local admin password: ")

		passBytes, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Println()
			return "", "", err
		}

		if err := auth.ValidatePassword(string(passBytes)); err != nil {
			fmt.Println()
			return "", "", errors.Wrap(err, "validating password")
		}
		fmt.Print("\nValidate local admin password: ")
		valiBytes, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Println()
			return "", "", err
		}
		if string(valiBytes) != string(passBytes) {
			fmt.Println()
			return "", "", errors.New("password mismatch")
		}
		fmt.Println()
		adminPass = strings.Trim(string(passBytes), "\n")
	}

	return adminName, adminPass, nil
}

func open(opts DBOpts) (*goqu.Database, error) {
	// Set connection opts
	var dbOpts string
	switch opts.Driver {
	case "sqlite3":
		dbOpts = opts.Path
	case "postgres":
		dbOpts = toPostgresOpts(opts)
	default:
		return nil, errors.New("unsupported database")

	}

	// Open connection
	db, err := sql.Open(opts.Driver, dbOpts)
	if err != nil {
		return nil, errors.Wrap(err, "opening database connection")
	}
	// Set configurations (eg. use fk constraints)
	switch opts.Driver {
	case "sqlite3":
		if _, err := os.Stat(dbOpts); os.IsNotExist(err) {
			if err := createSQLiteDB(db); err != nil {
				return nil, errors.Wrap(err, "creating sqlite3 database")
			}
		}
		if _, err := db.Exec(`PRAGMA foreign_keys = ON`); err != nil {
			return nil, errors.Wrap(err, "failed configuring database")
		}
	}
	// Verify connection
	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "failed pinging database")
	}

	return goqu.New(opts.Driver, db), nil
}

func toPostgresOpts(o DBOpts) string {
	s := fmt.Sprintf("dbname=%s", o.Name)

	m := map[string]string{
		"user":            o.User,
		"password":        o.Pass,
		"host":            o.Host,
		"port":            o.Port,
		"connect_timeout": o.ConnectionTimeout,
		"sslmode":         o.SSLMode,
		"sslcert":         o.SSLCert,
		"sslkey":          o.SSLKey,
		"sslrootcert":     o.SSLRootCert,
	}

	for k, v := range m {
		if v != "" {
			s = fmt.Sprintf("%s %s=%s", s, k, v)
		}
	}

	return s
}

func IsPrimed(db *goqu.Database) (bool, error) {
	ct, err := db.From("meta").Count()
	return ct > 0, err
}

func prime(db *goqu.Database) error {
	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "beginning transaction")
	}
	// All priming handled in a single transaction
	err = tx.Wrap(func() error {
		if err := primeMetadata(tx, "version", VERSION); err != nil {
			return errors.Wrap(err, "initializing metadata")
		}
		if err := primeClusterTypes(tx, cluster_types_list...); err != nil {
			return errors.Wrap(err, "initializing cluster types")
		}
		if err := primeStates(tx, states_list...); err != nil {
			return errors.Wrap(err, "initializing states")
		}
		if err := primePermissions(tx, permissions_list...); err != nil {
			return errors.Wrap(err, "initializing permissions")
		}
		if err := primeEntityTypes(tx, entity_types_list...); err != nil {
			return errors.Wrap(err, "initializing entity types")
		}
		return nil
	})
	return errors.Wrap(err, "committing transaction")
}

func initRoles(tx *goqu.TxDatabase, permissions []Permission, permissionsToAdd []Permission) error {
	role, doesRoleExist, err := readRole(tx, ByName(standard_user))
	if err != nil {
		return errors.Wrapf(err, "error reading role %s", role)
	}
	if doesRoleExist {
		if err := deleteRolePermission(tx, ByRoleId(role.Id)); err != nil {
			return errors.Wrapf(err, "error deleting role permission %s", role.Id)
		}
		initRolePermissions(tx, role.Id, permissionsToAdd)

	} else {
		roleId, err := createRole(tx, standard_user)
		if err != nil {
			return errors.Wrapf(err, "error creating role %s", role)
		}
		if err := initRolePermissions(tx, roleId, permissionsToAdd); err != nil {
			return errors.Wrapf(err, "error initialization role permissions")
		}
	}
	return err
}

func initRolePermissions(tx *goqu.TxDatabase, roleId int64, permissions []Permission) error {
	for _, permission := range permissions {
		_, err := createRolePermission(tx, roleId, permission.Id)
		if err != nil {
			return errors.Wrapf(err, "error creating permission %s for role %s", permission.Id, roleId)
		}
	}
	return nil
}

func primeMetadata(tx *goqu.TxDatabase, key, value string) error {
	_, err := tx.From("meta").Insert(goqu.Record{"key": key, "value": value}).Exec()
	return errors.Wrap(err, "executing query")
}

func primeClusterTypes(tx *goqu.TxDatabase, names ...string) error {
	for _, name := range names {
		if _, err := createClusterType(tx, name); err != nil {
			return errors.Wrapf(err, "creating cluster type %s", name)
		}
	}
	return nil
}

func primeStates(tx *goqu.TxDatabase, names ...string) error {
	for _, name := range names {
		if _, err := createState(tx, name); err != nil {
			return errors.Wrapf(err, "creating state %s", name)
		}
	}
	return nil
}

func primePermissions(tx *goqu.TxDatabase, perms ...permission_map) error {
	for _, perm := range perms {
		if _, err := createPermission(tx, perm.Code, perm.Desc); err != nil {
			return errors.Wrapf(err, "creating permission %s", perm.Code)
		}
	}
	return nil
}

func primeEntityTypes(tx *goqu.TxDatabase, entityTyes ...string) error {
	for _, et := range entityTyes {
		if _, err := createEntityType(tx, et); err != nil {
			return errors.Wrapf(err, "creating entity type %s", et)
		}
	}
	return nil
}

func initDatastore(db *goqu.Database) (*Datastore, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "beginning transaction")
	}

	var (
		clusterTypes []clusterType
		entityTypes  []entityType
		states       []state
		permissions  []Permission
	)
	if err := tx.Wrap(func() error {
		var err error
		clusterTypes, err = readClusterTypes(tx)
		if err != nil {
			return errors.Wrap(err, "reading cluster types")
		}
		entityTypes, err = readEntityTypes(tx)
		if err != nil {
			return errors.Wrap(err, "reading entity types")
		}
		states, err = readStates(tx)
		if err != nil {
			return errors.Wrap(err, "reading states")
		}
		permissions, err = readPermissions(tx)
		if err != nil {
			return errors.Wrap(err, "reading permissions")
		}

		if err := initRoles(tx, permissions, getPermissionsSubset(permissions, standard_user_init_permissions)); err != nil {
			return errors.Wrap(err, "initialization roles")
		}
		return nil
	}); err != nil {
		return nil, errors.Wrap(err, "committing transaction")
	}

	return &Datastore{
		db:    db,
		users: newUserTable(2 * time.Hour),

		EntityTypeMap: toEntityTypeMap(entityTypes),
		PermissionMap: toPermissionMap(permissions),

		ClusterType: newClusterTypeKeys(clusterTypes),
		EntityType:  newEntityTypeKeys(entityTypes),
		State:       newStateKeys(states),
		Permission:  newPermissionKeys(permissions),
	}, nil
}
func getPermissionsSubset(permissions []Permission, permissionCodes []string) []Permission {
	var permissionsToAdd []Permission
	for _, permission := range permissionCodes {
		for _, globalPermission := range permissions {
			if globalPermission.Code == permission {
				permissionsToAdd = append(permissionsToAdd, globalPermission)
			}
		}
	}
	return permissionsToAdd
}

func getRows(tx *goqu.TxDatabase, dataset *goqu.Dataset) (*sql.Rows, error) {
	sql, args, err := dataset.Prepared(true).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "writing query")
	}

	rows, err := tx.Query(sql, args...)
	return rows, errors.Wrap(err, "executing query")
}

func getRow(tx *goqu.TxDatabase, dataset *goqu.Dataset) (*sql.Row, error) {
	sql, args, err := dataset.Prepared(true).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "writing query")
	}

	return tx.QueryRow(sql, args...), nil
}

func (ds *Datastore) getRows(dataset *goqu.Dataset) (*sql.Rows, error) {
	sql, args, err := dataset.Prepared(true).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "writing query")
	}

	rows, err := ds.db.Query(sql, args...)
	return rows, errors.Wrap(err, "executing query")
}

func (ds *Datastore) getRow(dataset *goqu.Dataset) (*sql.Row, error) {
	sql, args, err := dataset.Prepared(true).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "writing query")
	}

	return ds.db.QueryRow(sql, args...), nil
}

func (ds *Datastore) createAdmin(username, password string) (int64, error) {
	hashPassword, err := auth.HashPassword(password)
	if err != nil {
		return 0, errors.Wrap(err, "hashing password")
	}

	tx, err := ds.db.Begin()
	if err != nil {
		return 0, errors.Wrap(err, "beginning transaction")
	}

	var id int64
	err = tx.Wrap(func() error {
		var err error
		id, err = createIdentity(tx, username,
			WithPassword(hashPassword),
			WithAdminWorkgroup)
		if err != nil {
			return errors.Wrap(err, "creating identity")
		}

		roleId, err := createRole(tx, AdminRN)
		if err != nil {
			return errors.Wrap(err, "creating role")
		}

		_, err = createIdentityRole(tx, WithIdentityId(id), WithRoleId(roleId))
		return errors.Wrap(err, "linking identity and rold")
	})

	return id, errors.Wrap(err, "committing transaction")
}

func (ds *Datastore) Count(table string, options ...QueryOpt) (int64, error) {
	tx, err := ds.db.Begin()
	if err != nil {
		return 0, errors.Wrap(err, "beginning transaction")
	}

	var ct int64
	err = tx.Wrap(func() error {
		q := NewQueryConfig(ds, tx, "", table, nil)
		q.dataset = q.dataset.Select(goqu.COUNT("*"))
		for _, option := range options {
			if err := option(q); err != nil {
				return errors.Wrap(err, "setting up query options")
			}
		}
		if debug {
			color.Set(color.FgYellow)
			log.Println(q.dataset.ToSql())
			color.Unset()
		}
		// Execute query
		row, err := getRow(tx, q.dataset)
		if err != nil {
			return err
		}
		if err := row.Scan(&ct); err != nil {
			return errors.Wrap(err, "scanning count")
		}
		for _, post := range q.postFunc {
			if err := post(q); err != nil {
				return errors.Wrap(err, "running post functions")
			}
		}

		return nil
	})
	return ct, errors.Wrap(err, "committing transaction")
}
