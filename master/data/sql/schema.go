package sql

var schema = map[string]string{
	"cluster":             createTableCluster,
	"cluster_type":        createTableClusterType,
	"cluster_yarn_detail": createTableClusterYarnDetail,
	"engine":              createTableEngine,
	"identity_workgroup":  createTableIdentityWorkgroup,
	"history":             createTableHistory,
	"meta":                createTableMeta,
	"permission":          createTablePermission,
	"privilege":           createTablePrivilege,
	"state":               createTableState,
}

var createTableCluster = `
CREATE TABLE cluster (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL,
    type_id integer NOT NULL,
    detail_id integer,
    address text UNIQUE,
    state integer NOT NULL,
    created datetime NOT NULL,

    CONSTRAINT type_id FOREIGN KEY (type_id) REFERENCES cluster_type(id),
    CONSTRAINT detail_id FOREIGN KEY (detail_id) REFERENCES cluster_yarn_detail(id),
    CONSTRAINT state FOREIGN KEY (state) REFERENCES state(id)
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
    size integer,
    application_id text,
    memory text,
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

var createEntityType = `
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
    password text NOT NULL,
    workgroup_id integer NOT NULL,
    is_active boolean NOT NULL,
    last_login integer with time zone,
    created datetime NOT NULL
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

var createTableMeta = `
CREATE TABLE meta (
    id integer NOT NULL,
    key text NOT NULL UNIQUE,
    value text NOT NULL,

    PRIMARY KEY (id)
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
    id integer PRIMARY KEY AUTOINCREMENT,
    privilege_type text NOT NULL,
    workgroup_id integer NOT NULL,
    entity_type_id integer NOT NULL,
    entity_id integer NOT NULL,

    UNIQUE (privilege_type, workgroup_id, entity_type_id, entity_id),
    CONSTRAINT entity_type_id FOREIGN KEY (entity_type_id) REFERENCES entity_type(id),
    CONSTRAINT workgroup_id FOREIGN KEY (workgroup_id) REFERENCES workgroup(id)
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
    name text NOT NULL,

    PRIMARY KEY (id)
)
`
