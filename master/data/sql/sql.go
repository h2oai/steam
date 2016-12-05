package sql

import (
	"database/sql"
	"flag"

	"github.com/pkg/errors"
	"gopkg.in/doug-martin/goqu.v3"
)

const (
	VERSION = "1.1.0"
)

type metadata map[string]string

type Datastore struct {
	db          *goqu.Database
	ClusterType clusterTypeKeys
	State       stateKeys
	Permission  permissionKeys
}

func init() {
	flag.BoolVar(&DEBUG, "debug", false, "Set to enable debug mode")
	flag.Parse()
}

func NewDatastore(driver, dbOpts string) (*Datastore, error) {
	// Connect to db
	db, err := Open(driver, dbOpts)
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

func Open(driver, dbOpts string) (*goqu.Database, error) {
	// Open connection
	db, err := sql.Open(driver, dbOpts)
	if err != nil {
		return nil, errors.Wrap(err, "opening database connection")
	}
	// Set configurations (eg. use fk constraints)
	switch driver {
	case "sqlite3":
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
		if err := primeClusterTypes(tx, CLUSTER_TYPES...); err != nil {
			return errors.Wrap(err, "initializing cluster types")
		}
		if err := primeStates(tx, STATES...); err != nil {
			return errors.Wrap(err, "initializing states")
		}
		if err := primePermissions(tx, PERMISSIONS...); err != nil {
			return errors.Wrap(err, "initializing permissions")
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

func initDatastore(db *goqu.Database) (*Datastore, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "beginning transaction")
	}

	var (
		clusterTypes []ClusterType
		states       []State
		permissions  []Permission
	)
	if err := tx.Wrap(func() error {
		var err error
		clusterTypes, err = readClusterTypes(tx)
		if err != nil {
			return errors.Wrap(err, "reading cluster types")
		}
		states, err = readStates(tx)
		if err != nil {
			return errors.Wrap(err, "reading states")
		}
		permissions, err = readPermissions(tx)
		if err != nil {
			return errors.Wrap(err, "reading permissions")
		}
		return nil
	}); err != nil {
		return nil, errors.Wrap(err, "committing transaction")
	}

	return &Datastore{
		db:          db,
		ClusterType: newClusterTypeKeys(clusterTypes),
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
