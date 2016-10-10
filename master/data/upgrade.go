package data

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

func upgradeTo_1_1_0(db *sql.DB) (string, error) {

	order := []string{
		"cluster_type", "engine", "entity_type", "identity", "permission", "project", "role", "workgroup",
		"cluster", "cluster_yarn", "datasource", "history", "identity_role", "identity_workgroup", "privilege", "role_permission",
		"dataset",
		"model",
		"binomial_model", "label", "multinomial_model", "regression_model", "service",
	}

	cols := map[string][]string{
		"binomial_model":     []string{"model_id", "mse", "r_squared", "logloss", "auc", "gini"},
		"cluster":            []string{"id", "name", "type_id", "detail_id", "address", "state", "created"},
		"cluster_type":       []string{"id", "name"},
		"cluster_yarn":       []string{"id", "engine_id", "size", "application_id", "memory", "username", "output_dir"},
		"dataset":            []string{"id", "datasource_id", "name", "description", "frame_name", "response_column_name", "properties", "properties_version", "created"},
		"datasource":         []string{"id", "project_id", "name", "description", "kind", "configuration", "created"},
		"engine":             []string{"id", "name", "location", "created"},
		"entity_type":        []string{"id", "name"},
		"history":            []string{"id", "action", "identity_id", "entity_type_id", "entity_id", "description", "created"},
		"identity":           []string{"id", "name", "password", "workgroup_id", "is_active", "last_login", "created"},
		"identity_role":      []string{"identity_id", "role_id"},
		"identity_workgroup": []string{"identity_id", "workgroup_id"},
		"label":              []string{"id", "project_id", "model_id", "name", "description", "created"},
		"model":              []string{"id", "project_id", "training_dataset_id", "validation_dataset_id", "name", "cluster_name", "model_key", "algorithm", "model_category", "dataset_name", "response_column_name", "logical_name", "location", "max_run_time", "metrics", "metrics_version", "created"},
		"multinomial_model":  []string{"model_id", "mse", "r_squared", "logloss"},
		"permission":         []string{"id", "code", "description"},
		"privilege":          []string{"privilege_type", "workgroup_id", "entity_type_id", "entity_id"},
		"project":            []string{"id", "name", "description", "model_category", "created"},
		"regression_model":   []string{"model_id", "mse", "r_squared", "mean_residual_deviance"},
		"service":            []string{"id", "project_id", "model_id", "name", "address", "port", "process_id", "state", "created"},
		"role":               []string{"id", "name", "description", "created"},
		"role_permission":    []string{"role_id", "permission_id"},
		"workgroup":          []string{"id", "type", "name", "description", "created"},
	}
	qrys := map[string]string{
		"binomial_model": `
	model_id integer NOT NULL,
    mse double precision,
    r_squared double precision,
    logloss double precision,
    auc double precision,
    gini double precision, 

    PRIMARY KEY (model_id),
    FOREIGN KEY (model_id) REFERENCES model(id) ON DELETE CASCADE
		`,

		"cluster": `
	id integer PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL,
    type_id integer NOT NULL,
    detail_id integer NOT NULL,
    address text NOT NULL,
    state job_state NOT NULL,
    created datetime NOT NULL,

    FOREIGN KEY (type_id) REFERENCES cluster_type(id)
    `,

		"cluster_type": `
    id integer PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL UNIQUE
    `,
		"cluster_yarn": `
    id integer PRIMARY KEY AUTOINCREMENT,
    engine_id integer NOT NULL,
    size integer NOT NULL,
    application_id text NOT NULL,
    memory text NOT NULL,
    username text NOT NULL,
    output_dir text NOT NULL,

    FOREIGN KEY (engine_id) REFERENCES engine(id)
    `,

		"dataset": `
    id integer PRIMARY KEY AUTOINCREMENT,
    datasource_id integer NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    frame_name text NOT NULL,
    response_column_name text NOT NULL,
    properties text NOT NULL,
    properties_version text NOT NULL,
    created datetime NOT NULL,

    FOREIGN KEY (datasource_id) REFERENCES datasource(id) ON DELETE CASCADE
    `,

		"datasource": `
    id integer PRIMARY KEY AUTOINCREMENT,
    project_id integer NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    kind text NOT NULL,
    configuration text NOT NULL,
    created datetime NOT NULL,

    FOREIGN KEY (project_id) REFERENCES project(id) ON DELETE CASCADE
	`,

		"engine": `
	id integer PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL,
    location text NOT NULL,
    created datetime NOT NULL
    `,

		"entity_type": `
    id integer PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL UNIQUE
    `,

		"history": `
    id integer PRIMARY KEY AUTOINCREMENT,
    action text NOT NULL,
    identity_id integer NOT NULL,
    entity_type_id integer NOT NULL,
    entity_id integer NOT NULL,
    description text NOT NULL,
    created datetime NOT NULL,

    FOREIGN KEY (entity_type_id) REFERENCES entity_type(id),
    FOREIGN KEY (identity_id) REFERENCES identity(id)
    `,

		"identity": `
    id integer PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL UNIQUE,
    password text NOT NULL,
    workgroup_id integer NOT NULL,
    is_active boolean NOT NULL,
    last_login integer with time zone,
    created datetime NOT NULL
    `,

		"identity_role": `
    identity_id integer NOT NULL,
    role_id integer NOT NULL,

    PRIMARY KEY (identity_id, role_id)
    `,

		"identity_workgroup": `
	identity_id integer NOT NULL,
    workgroup_id integer NOT NULL,

    PRIMARY KEY (identity_id, workgroup_id),
    FOREIGN KEY (identity_id) REFERENCES identity(id) ON DELETE CASCADE,
    FOREIGN KEY (workgroup_id) REFERENCES workgroup(id) ON DELETE CASCADE
    `,

		"label": `
    id integer PRIMARY KEY AUTOINCREMENT,
    project_id integer NOT NULL,
    model_id integer,
    name text NOT NULL,
    description text NOT NULL,
    created datetime NOT NULL,

    FOREIGN KEY (model_id) REFERENCES model(id) ON DELETE SET NULL,
    FOREIGN KEY (project_id) REFERENCES project(id) ON DELETE CASCADE
    `,

		"model": `
    id integer PRIMARY KEY AUTOINCREMENT,
    project_id integer NOT NULL,
    training_dataset_id integer NOT NULL,
    validation_dataset_id integer,
    name text NOT NULL,
    cluster_id integer,
    cluster_name text NOT NULL,
    model_key text NOT NULL,
    algorithm text NOT NULL,
    model_category text NOT NULL,
    dataset_name text NOT NULL,
    response_column_name text NOT NULL,
    logical_name text,
    location text NOT NULL,
    model_object_type text,
    max_run_time integer,
    metrics text NOT NULL,
    metrics_version text NOT NULL,
    created datetime NOT NULL,

    FOREIGN KEY (project_id) REFERENCES project(id),
    FOREIGN KEY (training_dataset_id) REFERENCES dataset(id),
    FOREIGN KEY (validation_dataset_id) REFERENCES dataset(id),
    FOREIGN KEY (cluster_id) REFERENCES cluster(id) ON DELETE SET NULL
    `,

		"multinomial_model": `
    model_id integer NOT NULL,
    mse double precision,
    r_squared double precision,
    logloss double precision,

    PRIMARY KEY (model_id),
    FOREIGN KEY (model_id) REFERENCES model(id) ON DELETE CASCADE
	`,

		"permission": `
    id integer PRIMARY KEY AUTOINCREMENT,
    code text NOT NULL UNIQUE,
    description text NOT NULL
    `,

		"privilege": `
    privilege_type text NOT NULL,
    workgroup_id integer NOT NULL,
    entity_type_id integer NOT NULL,
    entity_id integer NOT NULL,

    PRIMARY KEY (privilege_type, workgroup_id, entity_type_id, entity_id),
    FOREIGN KEY (entity_type_id) REFERENCES entity_type(id),
    FOREIGN KEY (workgroup_id) REFERENCES workgroup(id)
    `,

		"project": `
	id integer PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL,
    description text NOT NULL,
    model_category text NOT NULL,
    created datetime NOT NULL
    `,

		"regression_model": `
    model_id integer NOT NULL,
    mse double precision,
    r_squared double precision,
    mean_residual_deviance double precision,

    PRIMARY KEY (model_id),
    FOREIGN KEY (model_id) REFERENCES model(id) ON DELETE CASCADE
    `,

		"role": `
    id integer PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL UNIQUE,
    description text NOT NULL,
    created datetime NOT NULL  
    `,

		"role_permission": `
    role_id integer NOT NULL,
    permission_id integer NOT NULL,

    PRIMARY KEY (role_id, permission_id),
    FOREIGN KEY (permission_id) REFERENCES permission(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES role(id) ON DELETE CASCADE
    `,

		"service": `
    id integer PRIMARY KEY AUTOINCREMENT,
    project_id integer NOT NULL,
    model_id integer NOT NULL,
    name text NOT NULL,
    address text NOT NULL,
    port integer NOT NULL,
    process_id integer NOT NULL,
    state job_state NOT NULL,
    created datetime NOT NULL,

    FOREIGN KEY (model_id) REFERENCES model(id)
    `,

		"workgroup": `
    id integer PRIMARY KEY AUTOINCREMENT,
    type workgroup_type NOT NULL,
    name text NOT NULL UNIQUE,
    description text NOT NULL,
    created datetime NOT NULL 
    `,
	}

	tx, err := db.Begin()
	if err != nil {
		return "", errors.Wrap(err, "starting transaction")
	}
	defer tx.Rollback()

	for _, table := range order {

		if err := createTable(tx, table, cols[table]...); err != nil {
			return "", errors.Wrapf(err, "initializing table for %s", table)
		}

		ok, err := checkTable(tx, table)
		if err != nil {
			return "", errors.Wrapf(err, "checking table values for %s", table)
		}

		if err := createTemp(tx, table); err != nil {
			return "", errors.Wrapf(err, "creating temp for %s", table)
		}

		if err := createNew(tx, table, qrys[table]); err != nil {
			return "", errors.Wrapf(err, "creating new table for %s", table)
		}

		if ok {
			if err := copyTable(tx, table, cols[table]...); err != nil {
				return "", errors.Wrapf(err, "copying values for %s", table)
			}
		}
	}

	for i := len(order) - 1; i >= 0; i-- {
		if err := dropTemp(tx, order[i]); err != nil {
			return "", errors.Wrapf(err, "dropping temp for %s", order[i])
		}
	}

	if _, err := tx.Exec(`UPDATE meta SET value = $1 WHERE id = 1`, "1.1.0"); err != nil {
		return "", errors.Wrap(err, "updating database version")
	}

	return "1.1.0", errors.Wrap(tx.Commit(), "commiting changes")
}

func createTable(tx *sql.Tx, table string, cols ...string) error {
	var colStr string
	for i, col := range cols {
		if i > 0 {
			colStr += ", "
		}
		colStr += col
	}
	qry := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(%s)`, table, colStr)

	_, err := tx.Exec(qry)
	return err
}

func checkTable(tx *sql.Tx, table string) (bool, error) {
	qry := fmt.Sprintf(`SELECT count(*) FROM %s`, table)
	row := tx.QueryRow(qry)

	var ct int64
	err := row.Scan(&ct)

	if ct > 0 {
		return true, err
	}
	return false, err
}

func createTemp(tx *sql.Tx, table string) error {
	tmp := "temp_" + table
	qry := fmt.Sprintf(`ALTER TABLE %s RENAME TO %s`, table, tmp)

	_, err := tx.Exec(qry)
	return err
}

func createNew(tx *sql.Tx, table, cols string) error {
	qry := fmt.Sprintf(`CREATE TABLE %s (%s)`, table, cols)

	_, err := tx.Exec(qry)
	return err
}

func copyTable(tx *sql.Tx, table string, cols ...string) error {
	tmp := "temp_" + table

	var colStr string
	for i, col := range cols {
		if i > 0 {
			colStr += ", "
		}
		colStr += col
	}
	qry := fmt.Sprintf(`INSERT INTO %s(%s) SELECT %s FROM %s`, table, colStr, colStr, tmp)

	_, err := tx.Exec(qry)
	return err
}

func dropTemp(tx *sql.Tx, table string) error {
	tmp := "temp_" + table
	qry := fmt.Sprintf(`DROP TABLE %s`, tmp)

	_, err := tx.Exec(qry)
	return err
}
