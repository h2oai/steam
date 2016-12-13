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
	"database/sql"
	"flag"
	"log"
	"os"
	"time"

	"github.com/h2oai/steam/master/auth"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"gopkg.in/doug-martin/goqu.v3"
)

const (
	VERSION = "1.1.0"
)

type metadata map[string]string

type Datastore struct {
	db *goqu.Database

	// Internal mapping for printing
	entityTypeMap map[int64]string
	permissionMap map[int64]string

	// Enum References in Database
	ClusterType clusterTypeKeys
	EntityType  entityTypeKeys
	State       stateKeys
	Permission  permissionKeys
}

// --- Enums ---
var (
	States states
)

func init() {
	flag.BoolVar(&DEBUG, "debug", false, "Set to enable debug mode")
	flag.Parse()

	States = initState()
}

func NewDatastore(driver, dbOpts string) (*Datastore, error) {
	// Connect to db
	db, err := open(driver, dbOpts)
	if err != nil {
		return nil, errors.Wrap(err, "connecting to database")
	}

	if primed, err := IsPrimed(db); err != nil {
		return nil, errors.Wrap(err, "checking if database is primed")
	} else if !primed {
		if err := prime(db); err != nil {
			return nil, errors.Wrap(err, "priming database")
		}
	}

	ds, err := initDatastore(db)
	if err != nil {
		return nil, errors.Wrap(err, "initializing datastore")
	}

	return ds, nil
}

func open(driver, dbOpts string) (*goqu.Database, error) {
	// Open connection
	db, err := sql.Open(driver, dbOpts)
	if err != nil {
		return nil, errors.Wrap(err, "opening database connection")
	}
	// Set configurations (eg. use fk constraints)
	switch driver {
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

	return goqu.New(driver, db), nil
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

func primePermissions(tx *goqu.TxDatabase, perms ...struct{ code, desc string }) error {
	for _, perm := range perms {
		if _, err := createPermission(tx, perm.code, perm.desc); err != nil {
			return errors.Wrapf(err, "creating permission %s", perm.code)
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
		return errors.Wrap(err, "reading permissions")
	}); err != nil {
		return nil, errors.Wrap(err, "committing transaction")
	}

	return &Datastore{
		db: db,

		entityTypeMap: toEntityTypeMap(entityTypes),
		permissionMap: toPermissionMap(permissions),

		ClusterType: newClusterTypeKeys(clusterTypes),
		EntityType:  newEntityTypeKeys(entityTypes),
		State:       newStateKeys(states),
		Permission:  newPermissionKeys(permissions),
	}, nil
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

func (ds *Datastore) CreateSuperuser(username, password string) (int64, error) {
	hashPassword, err := auth.HashPassword(password)
	if err != nil {
		return 0, errors.Wrap(err, "hasing password")
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
			WithDefaultIdentityWorkgroup)
		if err != nil {
			return errors.Wrap(err, "creating identity")
		}

		roleId, err := createRole(tx, SuperuserRN)
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
		q := NewQueryConfig(ds, tx, table, tx.From(table).Select(goqu.COUNT("*")))
		for _, option := range options {
			if err := option(q); err != nil {
				return errors.Wrap(err, "setting up query options")
			}
		}
		if DEBUG {
			log.Println(q.dataset.ToSql())
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

type NullTime struct {
	Time  time.Time
	Valid bool
}
