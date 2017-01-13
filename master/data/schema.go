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

	"github.com/pkg/errors"
)

func createSQLiteDB(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "beginning transaction")
	}
	defer tx.Rollback()

	for tbl, query := range schema {
		_, err := tx.Exec(query)
		if err != nil {
			return errors.Wrapf(err, "initializing %s", tbl)
		}
	}

	return errors.Wrap(tx.Commit(), "committing transaction")
}

var schema = map[string]string{
	"authentication":      createTableAuthentication,
	"binomial_model":      createTableBinomialModel,
	"cluster":             createTableCluster,
	"cluster_type":        createTableClusterType,
	"cluster_yarn_detail": createTableClusterYarnDetail,
	"engine":              createTableEngine,
	"entity_type":         createTableEntityType,
	"identity":            createTableIdentity,
	"identity_role":       createTableIdentityRole,
	"identity_workgroup":  createTableIdentityWorkgroup,
	"history":             createTableHistory,
	"label":               createTableLabel,
	"meta":                createTableMeta,
	"model":               createTableModel,
	"multinomial_model":   createTableMultinomialModel,
	"permission":          createTablePermission,
	"privilege":           createTablePrivilege,
	"project":             createTableProject,
	"regression_model":    createTableRegressionModel,
	"role":                createTableRole,
	"role_permission":     createTableRolePermission,
	"service":             createTableService,
	"state":               createTableState,
	"workgroup":           createTableWorkgroup,
}

var createTableAuthentication = `
CREATE TABLE authentication (
    id integer PRIMARY KEY AUTOINCREMENT,
    key text NOT NULL UNIQUE,
    value texts NOT NULL,
    enabled boolean UNIQUE
)
`

var createTableBinomialModel = `
CREATE TABLE binomial_model (
    model_id integer NOT NULL,
    mse double precision,
    r_squared double precision,
    logloss double precision,
    auc double precision,
    gini double precision, 

    PRIMARY KEY (model_id),
    FOREIGN KEY (model_id) REFERENCES model(id) ON DELETE CASCADE
)
`

var createTableCluster = `
CREATE TABLE cluster (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL,
    context_path text,
    type_id integer NOT NULL,
    detail_id integer,
    address text UNIQUE,
    token text,
    state text NOT NULL,
    created datetime NOT NULL,

    CONSTRAINT type_id FOREIGN KEY (type_id) REFERENCES cluster_type(id),
    CONSTRAINT detail_id FOREIGN KEY (detail_id) REFERENCES cluster_yarn_detail(id),
    CONSTRAINT state FOREIGN KEY (state) REFERENCES state(name)
)
`

var createTableClusterType = `
CREATE TABLE cluster_type (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL UNIQUE
)
`

var createTableClusterYarnDetail = `
CREATE TABLE cluster_yarn_detail (
    id integer PRIMARY KEY AUTOINCREMENT,
    engine_id integer NOT NULL,
    size integer NOT NULL,
    application_id text NOT NULL,
    memory text NOT NULL,
    output_dir text NOT NULL,

    FOREIGN KEY (engine_id) REFERENCES engine(id)
)
`

var createTableEngine = `
CREATE TABLE engine (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL UNIQUE,
    location text NOT NULL,
    created datetime NOT NULL
)
`

var createTableEntityType = `
CREATE TABLE entity_type (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL UNIQUE
)
`

var createTableHistory = `
CREATE TABLE history (
    id integer PRIMARY KEY AUTOINCREMENT,
    action text NOT NULL,
    identity_id integer NOT NULL,
    entity_type_id integer NOT NULL,
    entity_id integer NOT NULL,
    description text,
    created datetime NOT NULL,

    FOREIGN KEY (entity_type_id) REFERENCES entity_type(id),
    FOREIGN KEY (identity_id) REFERENCES identity(id)
)
`

var createTableIdentity = `
CREATE TABLE identity (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL UNIQUE,
    auth_type test NOT NULL,
    password text,
    workgroup_id integer,
    is_active boolean NOT NULL,
    last_login integer with time zone,
    created datetime NOT NULL
)
`

var createTableIdentityRole = `
CREATE TABLE identity_role (
    identity_id integer NOT NULL,
    role_id integer NOT NULL,

    FOREIGN KEY (identity_id) REFERENCES identity(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES role(id) ON DELETE CASCADE,
    PRIMARY KEY (identity_id, role_id)
)
`

var createTableIdentityWorkgroup = `
CREATE TABLE identity_workgroup (
    identity_id integer NOT NULL,
    workgroup_id integer NOT NULL,

    PRIMARY KEY (identity_id, workgroup_id),
    FOREIGN KEY (identity_id) REFERENCES identity(id) ON DELETE CASCADE,
    FOREIGN KEY (workgroup_id) REFERENCES workgroup(id) ON DELETE CASCADE
)
`

var createTableLabel = `
CREATE TABLE label (
    id integer PRIMARY KEY AUTOINCREMENT,
    project_id integer NOT NULL,
    model_id integer,
    name text NOT NULL,
    description text NOT NULL,
    created datetime NOT NULL,

    FOREIGN KEY (model_id) REFERENCES model(id) ON DELETE SET NULL,
    FOREIGN KEY (project_id) REFERENCES project(id) ON DELETE CASCADE
)
`

var createTableMeta = `
CREATE TABLE meta (
    id integer NOT NULL,
    key text NOT NULL UNIQUE,
    value text NOT NULL,

    PRIMARY KEY (id)
)
`

var createTableModel = `
CREATE TABLE model (
    id integer PRIMARY KEY AUTOINCREMENT,
    project_id integer NOT NULL,
    name text NOT NULL,
    cluster_id integer,
    cluster_name text,
    model_key text NOT NULL,
    algorithm text NOT NULL,
    model_category text NOT NULL,
    dataset_name text,
    response_column_name text NOT NULL,
    logical_name text,
    location text,
    model_object_type text,
    max_run_time integer,
    schema text NOT NULL,
    schema_version text NOT NULL,
    created datetime NOT NULL,

    FOREIGN KEY (project_id) REFERENCES project(id) ON DELETE CASCADE,
    FOREIGN KEY (cluster_id) REFERENCES cluster(id) ON DELETE SET NULL

)
`

// FOREIGN KEY (training_dataset_id) REFERENCES dataset(id),
// FOREIGN KEY (validation_dataset_id) REFERENCES dataset(id),

var createTableMultinomialModel = `
CREATE TABLE multinomial_model (
    model_id integer NOT NULL,
    mse double precision NOT NULL,
    r_squared double precision NOT NULL,
    logloss double precision NOT NULL,

    PRIMARY KEY (model_id),
    FOREIGN KEY (model_id) REFERENCES model(id) ON DELETE CASCADE
)
`

var createTablePermission = `
CREATE TABLE permission (
    id integer PRIMARY KEY AUTOINCREMENT,
    code text NOT NULL UNIQUE,
    description text NOT NULL
)
`

var createTablePrivilege = `
CREATE TABLE privilege (
    privilege_type text NOT NULL,
    identity_id integer NOT NULL,
    workgroup_id integer NOT NULL,
    entity_type_id integer NOT NULL,
    entity_id integer NOT NULL,

    PRIMARY KEY (identity_id, privilege_type, workgroup_id, entity_type_id, entity_id),
    FOREIGN KEY (identity_id) REFERENCES identity(id) ON DELETE CASCADE,
    FOREIGN KEY (entity_type_id) REFERENCES entity_type(id),
    FOREIGN KEY (workgroup_id) REFERENCES workgroup(id) ON DELETE CASCADE
)
`

var createTableProject = `
CREATE TABLE project (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL UNIQUE,
    description text NOT NULL,
    model_category text NOT NULL,
    created datetime NOT NULL  
)
`

var createTableRegressionModel = `
CREATE TABLE regression_model (
    model_id integer NOT NULL,
    mse double precision,
    r_squared double precision,
    mean_residual_deviance double precision,

    PRIMARY KEY (model_id),
    FOREIGN KEY (model_id) REFERENCES model(id) ON DELETE CASCADE
)
`

var createTableRole = `
CREATE TABLE role (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL UNIQUE,
    description text,
    created datetime NOT NULL  
)
`

var createTableRolePermission = `
CREATE TABLE role_permission (
    role_id integer NOT NULL,
    permission_id integer NOT NULL,

    PRIMARY KEY (role_id, permission_id),
    FOREIGN KEY (permission_id) REFERENCES permission(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES role(id) ON DELETE CASCADE
)
`

var createTableService = `
CREATE TABLE service (
    id integer PRIMARY KEY AUTOINCREMENT,
    project_id integer NOT NULL,
    model_id integer NOT NULL,
    name text NOT NULL,
    host text,
    port integer,
    process_id integer,
    state job_state NOT NULL,
    created datetime NOT NULL,

    FOREIGN KEY (project_id) REFERENCES project(id)
    FOREIGN KEY (model_id) REFERENCES model(id)
)
`

var createTableState = `
CREATE TABLE state (
    id integer NOT NULL,
    name text NOT NULL UNIQUE,

    PRIMARY KEY (id)
)
`

var createTableWorkgroup = `
CREATE TABLE workgroup (
    id integer PRIMARY KEY AUTOINCREMENT,
    type workgroup_type NOT NULL,
    name text NOT NULL UNIQUE,
    description text,
    created datetime NOT NULL 
)
`
