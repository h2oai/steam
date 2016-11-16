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
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/h2oai/steam/master/auth"
	"github.com/h2oai/steam/master/az"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

//
// -----  Privilege / Sharing rules -----
//
// --------------------------------------
// Entity               Own  Edit View
// --------------------------------------
// Role
//   Read               x    x    x
//   Update             x    x
//   Assign Permission  x    x
//   Delete             x
//   Share              x
//
// Workgroup
//   Read               x    x    x
//   Update             x    x
//   Delete             x
//   Share              x
//
// Identity
//   Read               x    x    x
//   Assign Role        x    x
//   Assign Workgroup   x    x
//   Update             x    x
//   Delete             x
//   Share              x
//
// Cluster
//   Read               x    x    x
//   Start/Stop         x
//
// Project
//   Read               x    x    x
//   Assign Model       x    x
//   Update             x    x
//   Delete             x
//   Share              x
//
// Engine, Datasource, Dataset, Model
//   Read               x    x    x
//   Update             x    x
//   Delete             x
//   Share              x
//
// --------------------------------------

const (
	View = 1 << iota
	Edit // 2
	Own  // 4
)

const (
	Version = "1.1.0"

	SuperuserRoleName = "Superuser"

	CanView = "view"
	CanEdit = "edit"
	Owns    = "own"

	ForIdentity  = "identity"
	ForWorkgroup = "workgroup"

	RoleEntity       = "role"
	WorkgroupEntity  = "workgroup"
	IdentityEntity   = "identity"
	EngineEntity     = "engine"
	ClusterEntity    = "cluster"
	ProjectEntity    = "project"
	DatasourceEntity = "datasource"
	DatasetEntity    = "dataset"
	ModelEntity      = "model"
	LabelEntity      = "label"
	ServiceEntity    = "service"

	ClusterExternal = "external"
	ClusterYarn     = "yarn"
)

const (
	IdleState         = "idle"
	StartingState     = "starting"
	StartedState      = "started"
	SuspendingState   = "suspending"
	SuspendedState    = "suspended"
	StoppingState     = "stopping"
	StoppedState      = "stopped"
	BlockedState      = "blocked"
	DisconnectedState = "disconnected"
	FailedState       = "failed"
	CompletedState    = "completed"
)

const (
	ManageRole       = "ManageRole"
	ViewRole         = "ViewRole"
	ManageWorkgroup  = "ManageWorkgroup"
	ViewWorkgroup    = "ViewWorkgroup"
	ManageIdentity   = "ManageIdentity"
	ViewIdentity     = "ViewIdentity"
	ManageEngine     = "ManageEngine"
	ViewEngine       = "ViewEngine"
	ManageCluster    = "ManageCluster"
	ViewCluster      = "ViewCluster"
	ManageProject    = "ManageProject"
	ViewProject      = "ViewProject"
	ManageDatasource = "ManageDatasource"
	ViewDatasource   = "ViewDatasource"
	ManageDataset    = "ManageDataset"
	ViewDataset      = "ViewDataset"
	ManageModel      = "ManageModel"
	ViewModel        = "ViewModel"
	ManageLabel      = "ManageLabel"
	ViewLabel        = "ViewLabel"
	ManageService    = "ManageService"
	ViewService      = "ViewService"
)

var (
	Privileges   map[string]int
	Permissions  []Permission
	EntityTypes  []EntityType
	ClusterTypes []ClusterType
)

func init() {
	Privileges = map[string]int{
		CanView: View,
		CanEdit: Edit,
		Owns:    Own,
	}

	Permissions = []Permission{
		{0, ManageRole, "Manage roles"},
		{0, ViewRole, "View roles"},
		{0, ManageWorkgroup, "Manage workgroups"},
		{0, ViewWorkgroup, "View workgroups"},
		{0, ManageIdentity, "Manage identities"},
		{0, ViewIdentity, "View identities"},
		{0, ManageEngine, "Manage engines"},
		{0, ViewEngine, "View engines"},
		{0, ManageCluster, "Manage clusters"},
		{0, ViewCluster, "View clusters"},
		{0, ManageProject, "Manage projects"},
		{0, ViewProject, "View projects"},
		{0, ManageDatasource, "Manage datasources"},
		{0, ViewDatasource, "View datasources"},
		{0, ManageDataset, "Manage datasets"},
		{0, ViewDataset, "View datasets"},
		{0, ManageModel, "Manage models"},
		{0, ViewModel, "View models"},
		{0, ManageLabel, "Manage labels"},
		{0, ViewLabel, "View labels"},
		{0, ManageService, "Manage services"},
		{0, ViewService, "View services"},
	}

	EntityTypes = []EntityType{
		{0, RoleEntity},
		{0, WorkgroupEntity},
		{0, IdentityEntity},
		{0, EngineEntity},
		{0, ClusterEntity},
		{0, ProjectEntity},
		{0, DatasourceEntity},
		{0, DatasetEntity},
		{0, ModelEntity},
		{0, LabelEntity},
		{0, ServiceEntity},
	}

	ClusterTypes = []ClusterType{
		{0, ClusterExternal},
		{0, ClusterYarn},
	}
}

type metadata map[string]string

type PermissionKeys struct {
	ManageRole       int64
	ViewRole         int64
	ManageWorkgroup  int64
	ViewWorkgroup    int64
	ManageIdentity   int64
	ViewIdentity     int64
	ManageEngine     int64
	ViewEngine       int64
	ManageCluster    int64
	ViewCluster      int64
	ManageProject    int64
	ViewProject      int64
	ManageDatasource int64
	ViewDatasource   int64
	ManageDataset    int64
	ViewDataset      int64
	ManageModel      int64
	ViewModel        int64
	ManageLabel      int64
	ViewLabel        int64
	ManageService    int64
	ViewService      int64
}

type EntityTypeKeys struct {
	Role       int64
	Workgroup  int64
	Identity   int64
	Engine     int64
	Cluster    int64
	Project    int64
	Datasource int64
	Dataset    int64
	Model      int64
	Label      int64
	Service    int64
}

type ClusterTypeKeys struct {
	External int64
	Yarn     int64
}

func toPermissionKeys(permissions []Permission) *PermissionKeys {
	m := make(map[string]int64)
	for _, p := range permissions {
		m[p.Code] = p.Id
	}

	return &PermissionKeys{
		m[ManageRole],
		m[ViewRole],
		m[ManageWorkgroup],
		m[ViewWorkgroup],
		m[ManageIdentity],
		m[ViewIdentity],
		m[ManageEngine],
		m[ViewEngine],
		m[ManageCluster],
		m[ViewCluster],
		m[ManageProject],
		m[ViewProject],
		m[ManageDatasource],
		m[ViewDatasource],
		m[ManageDataset],
		m[ViewDataset],
		m[ManageModel],
		m[ViewModel],
		m[ManageLabel],
		m[ViewLabel],
		m[ManageService],
		m[ViewService],
	}
}

func toEntityTypeKeys(entityTypes []EntityType) *EntityTypeKeys {
	m := make(map[string]int64)
	for _, et := range entityTypes {
		m[et.Name] = et.Id
	}

	return &EntityTypeKeys{
		m[RoleEntity],
		m[WorkgroupEntity],
		m[IdentityEntity],
		m[EngineEntity],
		m[ClusterEntity],
		m[ProjectEntity],
		m[DatasourceEntity],
		m[DatasetEntity],
		m[ModelEntity],
		m[LabelEntity],
		m[ServiceEntity],
	}
}

func toClusterTypeKeys(clusterTypes []ClusterType) *ClusterTypeKeys {
	m := make(map[string]int64)
	for _, ct := range clusterTypes {
		m[ct.Name] = ct.Id
	}
	return &ClusterTypeKeys{
		m[ClusterExternal],
		m[ClusterYarn],
	}
}

type Datastore struct {
	db                *sql.DB // Singleton; doesn't actually connect until used, and is pooled internally.
	metadata          metadata
	permissions       []Permission
	permissionMap     map[int64]Permission
	Permissions       *PermissionKeys
	entityTypes       []EntityType
	entityTypeMap     map[int64]EntityType
	EntityTypes       *EntityTypeKeys
	clusterTypes      []ClusterType
	clusterTypeMap    map[int64]ClusterType
	ClusterTypes      *ClusterTypeKeys
	ViewPermissions   map[int64]int64
	ManagePermissions map[int64]int64
}

func Create(dbPath, suname, supass string) (*Datastore, error) {
	// connectionString := createConnectionString(connection)
	db, err := connect(dbPath)
	if err != nil {
		return nil, fmt.Errorf("Failed connecting to database using %s: %s", dbPath, err)
	}

	primed, err := isPrimed(db)
	if err != nil {
		return nil, fmt.Errorf("Failed database version check:", err)
	}

	if !primed {
		if suname == "" || supass == "" {
			return nil, fmt.Errorf("Starting Steam for the first time requires both --superuser-name and --superuser-password arguments to \"steam serve master\".")
		}

		if err := auth.ValidateUsername(suname); err != nil {
			return nil, fmt.Errorf("Invalid superuser username: %s", err)
		}

		if err := auth.ValidatePassword(supass); err != nil {
			return nil, fmt.Errorf("Invalid superuser password: %s", err)
		}

		if err := prime(db); err != nil {
			return nil, fmt.Errorf("Failed priming database: %s", err)
		}
	}

	ds, err := newDatastore(db)
	if err != nil {
		return nil, fmt.Errorf("Failed initializing from database: %s", err)
	}

	if !primed {
		passwordHash, err := auth.HashPassword(supass)
		if err != nil {
			return nil, fmt.Errorf("Failed hashing superuser password: %s", err)
		}

		if _, _, err := ds.CreateSuperuser(suname, passwordHash); err != nil {
			return nil, fmt.Errorf("Failed superuser identity setup: %s", err)
		}

		_, err = ds.Lookup(suname)
		if err != nil {
			return nil, fmt.Errorf("Failed reading superuser principal: %s", err)
		}
	}

	return ds, nil
}

func Destroy(dbPath string) error {
	// connectionString := createConnectionString(connection)
	db, err := connect(dbPath)
	if err != nil {
		return fmt.Errorf("Failed connecting to database using %s: %s", dbPath, err)
	}
	return truncate(db)
}

type Connection struct {
	DbName            string
	User              string
	Password          string
	Host              string
	Port              string
	ConnectionTimeout string
	SSLMode           string
	SSLCert           string
	SSLKey            string
	SSLRootCert       string
}

func createConnectionString(c Connection) string {
	s := fmt.Sprintf("dbname=%s", c.DbName)

	m := map[string]string{
		"user":            c.User,
		"password":        c.Password,
		"host":            c.Host,
		"port":            c.Port,
		"connect_timeout": c.ConnectionTimeout,
		"sslmode":         c.SSLMode,
		"sslcert":         c.SSLCert,
		"sslkey":          c.SSLKey,
		"sslrootcert":     c.SSLRootCert,
	}

	for k, v := range m {
		if len(v) > 0 {
			s = s + " " + k + "=" + v
		}
	}

	return s
}

func connect(dbPath string) (*sql.DB, error) {
	// Open connection
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed opening database")
	}
	// Set configurations (eg. use fk constraints)
	if _, err := db.Exec(`
		PRAGMA foreign_keys = ON
		`); err != nil {
		return nil, errors.Wrap(err, "failed configuring database")
	}
	// Verify connection
	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "failed pinging database")
	}
	// TODO can use db.SetMaxOpenConns() and db.SetMaxIdleConns() to configure further.
	return db, nil
}

// TODO deprecated?
// 	// FIXME logging need to be handled for testing
// 	// log.Println("Connecting to database: user =", username, "db =", dbname, "SSL=", sslmode, "...")

// 	// Open connection
// 	db, err := sql.Open("postgres", connection)
// 	if err != nil {
// 		return nil, fmt.Errorf("Database connection failed: %s", err)
// 	}

// 	// Verify connection
// 	if err := db.Ping(); err != nil {
// 		return nil, fmt.Errorf("Database ping failed: %s", err)
// 	}

// 	return db, nil
// }

// newDatastore creates a new instance of a data access object.
//
// Valid values for sslmode are:
//   disable - No SSL
//   require - Always SSL (skip verification)
//   verify-ca - Always SSL (verify that the certificate presented by the server was signed by a trusted CA)
//   verify-full - Always SSL (verify that the certification presented by the server was signed by a
//     trusted CA and the server host name matches the one in the certificate)
func newDatastore(db *sql.DB) (*Datastore, error) {

	// Read meta information

	metadata, err := readMetadata(db)
	if err != nil {
		return nil, err
	}

	version, ok := metadata["version"]
	if !ok {
		return nil, fmt.Errorf("Failed reading schema version")
	}

	// FIXME logging needs to be handled for testing
	// log.Println("Using schema version:", version)

	if err := upgrade(db, version); err != nil {
		return nil, err
	}

	permissions, err := readAllPermissions(db)
	if err != nil {
		return nil, err
	}

	permissionMap := make(map[int64]Permission)
	for _, permission := range permissions {
		permissionMap[permission.Id] = permission
	}

	permissionKeys := toPermissionKeys(permissions)

	entityTypes, err := readEntityTypes(db)
	if err != nil {
		return nil, err
	}

	entityTypeMap := make(map[int64]EntityType)
	for _, et := range entityTypes {
		entityTypeMap[et.Id] = et
	}

	entityTypeKeys := toEntityTypeKeys(entityTypes)

	clusterTypes, err := readClusterTypes(db)
	if err != nil {
		return nil, err
	}

	clusterTypeMap := make(map[int64]ClusterType)
	for _, ct := range clusterTypes {
		clusterTypeMap[ct.Id] = ct
	}

	viewPermissions := map[int64]int64{
		entityTypeKeys.Engine:     permissionKeys.ViewEngine,
		entityTypeKeys.Cluster:    permissionKeys.ViewCluster,
		entityTypeKeys.Project:    permissionKeys.ViewProject,
		entityTypeKeys.Datasource: permissionKeys.ViewDatasource,
		entityTypeKeys.Dataset:    permissionKeys.ViewDataset,
		entityTypeKeys.Model:      permissionKeys.ViewModel,
		entityTypeKeys.Label:      permissionKeys.ViewLabel,
		entityTypeKeys.Service:    permissionKeys.ViewService,
		entityTypeKeys.Identity:   permissionKeys.ViewIdentity,
		entityTypeKeys.Role:       permissionKeys.ViewRole,
		entityTypeKeys.Workgroup:  permissionKeys.ViewWorkgroup,
	}

	managePermissions := map[int64]int64{
		entityTypeKeys.Engine:     permissionKeys.ManageEngine,
		entityTypeKeys.Cluster:    permissionKeys.ManageCluster,
		entityTypeKeys.Project:    permissionKeys.ManageProject,
		entityTypeKeys.Datasource: permissionKeys.ManageDatasource,
		entityTypeKeys.Dataset:    permissionKeys.ManageDataset,
		entityTypeKeys.Model:      permissionKeys.ManageModel,
		entityTypeKeys.Label:      permissionKeys.ManageLabel,
		entityTypeKeys.Service:    permissionKeys.ManageService,
		entityTypeKeys.Identity:   permissionKeys.ManageIdentity,
		entityTypeKeys.Role:       permissionKeys.ManageRole,
		entityTypeKeys.Workgroup:  permissionKeys.ManageWorkgroup,
	}

	return &Datastore{
		db,
		metadata,
		permissions,
		permissionMap,
		permissionKeys,
		entityTypes,
		entityTypeMap,
		entityTypeKeys,
		clusterTypes,
		clusterTypeMap,
		toClusterTypeKeys(clusterTypes),
		viewPermissions,
		managePermissions,
	}, nil
}

func isPrimed(db *sql.DB) (bool, error) {
	row := db.QueryRow(`
		SELECT 
			count(1)
		FROM
			meta
		`)
	count, err := scanInt(row)
	if err != nil {
		return false, err
	}
	return count > 0, err
}

func prime(db *sql.DB) error {
	// FIXME logging needs to be handled for testing
	// log.Println("Priming database for first time use...")
	if err := createMetadata(db, "version", Version); err != nil {
		return err
	}
	if err := primePermissions(db, Permissions); err != nil {
		return errors.Wrap(err, "failed priming permissions")
	}
	if err := primeEntityTypes(db, EntityTypes); err != nil {
		return errors.Wrap(err, "failed priming entity_types")
	}
	if err := primeClusterTypes(db, ClusterTypes); err != nil {
		return errors.Wrap(err, "failed priming cluster_types")
	}

	return nil
}

func insertIn(table string, columns ...string) string {
	stmt := "INSERT INTO '" + table + "' ("
	for i, col := range columns {
		if i != 0 {
			stmt += ", "
		}
		stmt += "'" + col + "'"
	}
	stmt += ") VALUES ("
	for i := range columns {
		if i != 0 {
			stmt += ", "
		}
		stmt += "?"
	}
	stmt += ")"
	return stmt
}

func bulkInsert(db *sql.DB, table string, columns []string, f func(*sql.Stmt) error) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(insertIn(table, columns...))
	if err != nil {
		return err
	}
	defer stmt.Close()

	if err := f(stmt); err != nil { // buffer
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func createMetadata(db *sql.DB, key, value string) error {
	_, err := db.Exec(`
		INSERT INTO
			meta
			(key, value)
		VALUES
			($1,  $2)
		`, key, value)
	return err
}

func primeEntityTypes(db *sql.DB, entityTypes []EntityType) error {
	return bulkInsert(db, "entity_type", []string{"name"}, func(stmt *sql.Stmt) error {
		for _, entityType := range entityTypes {
			_, err := stmt.Exec(entityType.Name)
			if err != nil {
				return errors.Wrap(err, "failed entity_types bulk insert")
			}
		}
		return nil
	})
}

func primeClusterTypes(db *sql.DB, clusterTypes []ClusterType) error {
	return bulkInsert(db, "cluster_type", []string{"name"}, func(stmt *sql.Stmt) error {
		for _, clusterType := range clusterTypes {
			_, err := stmt.Exec(clusterType.Name)
			if err != nil {
				return errors.Wrap(err, "failed cluster_types bulk insert")
			}
		}
		return nil
	})
}

func primePermissions(db *sql.DB, permissions []Permission) error {
	return bulkInsert(db, "permission", []string{"code", "description"}, func(stmt *sql.Stmt) error {
		for _, permission := range permissions {
			_, err := stmt.Exec(permission.Code, permission.Description)
			if err != nil {
				return errors.Wrap(err, "failed permissions bulk insert")
			}
		}
		return nil
	})
}

func upgrade(db *sql.DB, currentVersion string) error {
	for currentVersion != Version {
		var err error
		switch {
		case currentVersion == "1":
			log.Println("Upgrading database to 1.1.0")
			currentVersion, err = upgradeTo_1_1_0(db)
		}

		if err != nil {
			return errors.Wrap(err, "upgrading database")
		}
	}

	return nil
}

func truncate(db *sql.DB) error {
	// FIXME logging needs to be handled for testing
	// log.Println("Truncating database...")
	return executeTransaction(db, func(tx *sql.Tx) error {
		tables := []string{
			"history",
			"privilege",
			"role_permission",
			"identity_role",
			"identity_workgroup",
			"identity",
			"workgroup",
			"role",
			"permission",
			"entity_type",
			"service",
			"label",
			"binomial_model",
			"multinomial_model",
			"regression_model",
			"model",
			"dataset",
			"datasource",
			"project",
			"cluster",
			"cluster_yarn",
			"cluster_type",
			"engine",
			"meta",
		}
		for _, table := range tables {
			if _, err := tx.Exec("DELETE FROM " + table); err != nil {
				return err
			}
		}
		return nil
	})
}

// --- Superuser ---

func (ds *Datastore) CreateSuperuser(name, password string) (int64, int64, error) {
	var id, workgroupId int64
	err := ds.exec(func(tx *sql.Tx) error {
		var err error

		workgroupId, err = createDefaultWorkgroup(tx, name)
		if err != nil {
			return errors.Wrap(err, "creating workgroup")
		}

		id, err = createIdentity(tx, name, password, workgroupId)
		if err != nil {
			return errors.Wrap(err, "creating identity")
		}

		if err := linkIdentityAndWorkgroup(tx, id, workgroupId); err != nil {
			return errors.Wrap(err, "linking identity and workgroup")
		}

		if err := createPrivilege(tx, Privilege{
			Owns,
			workgroupId,
			ds.EntityTypes.Identity,
			id,
		}); err != nil {
			return errors.Wrap(err, "creating identity privilege")
		}

		if err := createPrivilege(tx, Privilege{
			Owns,
			workgroupId,
			ds.EntityTypes.Workgroup,
			workgroupId,
		}); err != nil {
			return errors.Wrap(err, "creating workgroup privilege")
		}

		roleId, err := createRole(tx, SuperuserRoleName, SuperuserRoleName)
		if err != nil {
			return errors.Wrap(err, "creating role")
		}

		if err := linkIdentityAndRole(tx, id, roleId); err != nil {
			return nil
		}

		return nil
	})
	return id, workgroupId, err
}

// --- Lookup tables (static) ---

func readMetadataValue(db *sql.DB, key string) (string, error) {
	row := db.QueryRow(`
		SELECT
			value
		FROM
			meta
		WHERE
			name = $1
		`, key)
	return scanString(row)
}

func readMetadata(db *sql.DB) (map[string]string, error) {
	rows, err := db.Query(`
		SELECT
			id, key, value
		FROM
			meta
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	entries, err := ScanMetas(rows)
	if err != nil {
		return nil, err
	}

	lookup := make(map[string]string)
	for _, entry := range entries {
		lookup[entry.Key] = entry.Value
	}

	return lookup, nil
}

func readEntityTypes(db *sql.DB) ([]EntityType, error) {
	rows, err := db.Query(`
		SELECT
			id, name
		FROM
			entity_type
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanEntityTypes(rows)
}

func readClusterTypes(db *sql.DB) ([]ClusterType, error) {
	rows, err := db.Query(`
		SELECT
			id, name
		FROM
			cluster_type
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanClusterTypes(rows)
}

func executeTransaction(db *sql.DB, f func(*sql.Tx) error) error {
	// Open transaction; Rollback in event
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Execute transaction and return corresponding value (usually id)
	if err := f(tx); err != nil {
		return err
	}

	return tx.Commit()
}

func (ds *Datastore) exec(f func(*sql.Tx) error) error {
	return executeTransaction(ds.db, f)
}

func (ds *Datastore) toPermissionDescription(id int64) (string, error) {
	if p, ok := ds.permissionMap[id]; ok {
		return p.Description, nil
	} else {
		return "", fmt.Errorf("Invalid permission id: %d", id)
	}
}

func (ds *Datastore) toPermissionDescriptions(ids []int64) ([]string, error) {
	descriptions := make([]string, len(ids))
	for i, id := range ids {
		description, err := ds.toPermissionDescription(id)
		if err != nil {
			return nil, err
		}
		descriptions[i] = description
	}
	return descriptions, nil
}

func scanInt(r *sql.Row) (int64, error) {
	var value int64
	if err := r.Scan(&value); err != nil {
		return value, err
	}
	return value, nil
}

func scanInts(rs *sql.Rows) ([]int64, error) {
	values := make([]int64, 0, 16)
	var err error
	for rs.Next() {
		var value int64
		if err = rs.Scan(&value); err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return values, nil
}

func scanString(r *sql.Row) (string, error) {
	var value string
	if err := r.Scan(&value); err != nil {
		return value, err
	}
	return value, nil
}

func scanStrings(rs *sql.Rows) ([]string, error) {
	values := make([]string, 0, 16)
	var err error
	for rs.Next() {
		var value string
		if err = rs.Scan(&value); err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return values, nil
}

// --- History ---

const (
	CreateOp  string = "create"
	UpdateOp  string = "update"
	DeleteOp  string = "delete"
	EnableOp  string = "enable"
	DisableOp string = "disable"
	ShareOp   string = "share"
	UnshareOp string = "unshare"
	LinkOp    string = "link"
	UnlinkOp  string = "unlink"
)

func (ds *Datastore) audit(pz az.Principal, tx *sql.Tx, action string, entityTypeId, entityId int64, metadata metadata) error {
	json, err := json.Marshal(metadata)
	if err != nil {
		return err
	}

	if _, err := tx.Exec(`
		INSERT INTO
			history
			(identity_id, action, entity_type_id, entity_id, description, created)
		VALUES
			($1,          $2,     $3,             $4,        $5,          datetime('now'))
		`, pz.Id(), action, entityTypeId, entityId, string(json)); err != nil {
		return err
	}
	return nil
}

func (ds *Datastore) ReadEntityTypes(pz az.Principal) []EntityType {
	return ds.entityTypes
}

func (ds *Datastore) ReadHistoryForEntity(pz az.Principal, entityTypeId, entityId, offset, limit int64) ([]EntityHistory, error) {
	if err := pz.CheckView(entityTypeId, entityId); err != nil {
		return nil, err
	}

	rows, err := ds.db.Query(`
		SELECT
			identity_id, action, description, created
		FROM
			history
		WHERE
			entity_id = $1 AND
			entity_type_id = $2
		ORDER BY
			created DESC
		OFFSET $3
		LIMIT $4
	`, entityId, entityTypeId, offset, limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanEntityHistorys(rows)
}

// --- Permissions ---

func readAllPermissions(db *sql.DB) ([]Permission, error) {
	rows, err := db.Query(`
		SELECT
			id, code, description
		FROM
			permission
		ORDER BY
			code
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanPermissions(rows)
}

func (ds *Datastore) ReadAllPermissions(pz az.Principal) ([]Permission, error) {
	return ds.permissions, nil
}

func (ds *Datastore) ReadPermissionsForRole(pz az.Principal, roleId int64) ([]Permission, error) {

	if err := pz.CheckView(ds.EntityTypes.Role, roleId); err != nil {
		return nil, err
	}

	rows, err := ds.db.Query(`
		SELECT
			p.id, p.code, p.description
		FROM
			role_permission rp,
			permission p
		WHERE
			rp.role_id = $1 AND
			rp.permission_id = p.id
		ORDER BY
			p.code
		`, roleId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanPermissions(rows)
}

func (ds *Datastore) readPermissionsForIdentity(identityId int64) ([]int64, error) {
	rows, err := ds.db.Query(`
		SELECT DISTINCT
			p.id
		FROM
		  identity_role ir,
			role_permission rp,
			permission p
		WHERE
			ir.identity_id = $1 AND
			ir.role_id = rp.role_id AND
			rp.permission_id = p.id
	`, identityId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanInts(rows)
}

func (ds *Datastore) ReadPermissionsForIdentity(pz az.Principal, identityId int64) ([]Permission, error) {
	if err := pz.CheckView(ds.EntityTypes.Identity, identityId); err != nil {
		return nil, err
	}

	rows, err := ds.db.Query(`
		SELECT DISTINCT
			p.id, p.code, p.description
		FROM
		  identity_role ir,
			role_permission rp,
			permission p
		WHERE
			ir.identity_id = $1 AND
			ir.role_id = rp.role_id AND
			rp.permission_id = p.id
	`, identityId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanPermissions(rows)
}

// --- Roles ---

func createRole(tx *sql.Tx, name, description string) (int64, error) {
	res, err := tx.Exec(`
			INSERT INTO
				role
				(name, description, created)
			VALUES
				($1,   $2,          datetime('now'))
			`, name, description)
	if err != nil {
		return 0, errors.Wrap(err, "failed creating role")
	}

	return res.LastInsertId()
}

func (ds *Datastore) CreateRole(pz az.Principal, name, description string) (int64, error) {
	var id int64
	err := ds.exec(func(tx *sql.Tx) error {
		var err error

		id, err = createRole(tx, name, description)
		if err != nil {
			return err
		}

		if err := createPrivilege(tx, Privilege{
			Owns,
			pz.WorkgroupId(),
			ds.EntityTypes.Role,
			id,
		}); err != nil {
			return err
		}

		return ds.audit(pz, tx, CreateOp, ds.EntityTypes.Role, id, metadata{"name": name, "description": description})
	})
	return id, err
}

func (ds *Datastore) ReadRoles(pz az.Principal, offset, limit int64) ([]Role, error) {
	rows, err := ds.db.Query(`
		SELECT
			id, name, description, created
		FROM
			role
		WHERE
			id IN
			(
				SELECT DISTINCT
					entity_id
				FROM 
					privilege
				WHERE
				 	$1 OR
					(
						workgroup_id IN 
						(
							SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $2
						) AND 
						entity_type_id = $3
					)
			)
		ORDER BY 
			name
		LIMIT $4
		OFFSET $5
		`, pz.IsSuperuser(), pz.Id(), ds.EntityTypes.Role, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanRoles(rows)
}

func (ds *Datastore) readRoleNamesForIdentity(identityId int64) ([]string, error) {
	rows, err := ds.db.Query(`
		SELECT DISTINCT
			r.name
		FROM
			role r,
			identity_role ir
		WHERE
		  ir.identity_id = $1 AND
			ir.role_id = r.id
		ORDER BY
			r.name
		`, identityId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanStrings(rows)
}

func (ds *Datastore) ReadRolesForIdentity(pz az.Principal, identityId int64) ([]Role, error) {
	rows, err := ds.db.Query(`
		SELECT
			r.id, r.name, r.description, r.created
		FROM
			role r,
			identity_role ir
		WHERE
		  ir.identity_id = $1 AND
			ir.role_id = r.id AND
			r.id IN
			(
				SELECT DISTINCT
					entity_id
				FROM 
					privilege
				WHERE
				  $2 OR
					(
						workgroup_id IN 
						(
							SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $3
						) AND 
						entity_type_id = $4
					)
			)
		ORDER BY
			r.name
		`, identityId, pz.IsSuperuser(), pz.Id(), ds.EntityTypes.Role)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanRoles(rows)
}

func (ds *Datastore) ReadRole(pz az.Principal, roleId int64) (Role, error) {
	if err := pz.CheckView(ds.EntityTypes.Role, roleId); err != nil {
		return Role{}, err
	}

	row := ds.db.QueryRow(`
		SELECT
			id, name, description, created
		FROM
			role
		WHERE
			id = $1
		`, roleId)
	return ScanRole(row)
}

func (ds *Datastore) ReadRoleByName(pz az.Principal, name string) (Role, error) {
	row := ds.db.QueryRow(`
		SELECT
			id, name, description, created
		FROM
			role
		WHERE
			name = $1
		`, name)

	role, err := ScanRole(row)
	if err != nil {
		return Role{}, err
	}

	if err := pz.CheckView(ds.EntityTypes.Role, role.Id); err != nil {
		return Role{}, err
	}

	return role, nil
}

func (ds *Datastore) UpdateRole(pz az.Principal, roleId int64, name, description string) error {
	if err := pz.CheckEdit(ds.EntityTypes.Role, roleId); err != nil {
		return err
	}

	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			UPDATE
				role
			SET
				name = $1,
				description = $2
			WHERE
				id = $3
			`, name, description, roleId); err != nil {
			return err
		}
		return ds.audit(pz, tx, UpdateOp, ds.EntityTypes.Role, roleId, metadata{
			"name":        name,
			"description": description,
		})
	})
}

func (ds *Datastore) LinkRoleAndPermissions(pz az.Principal, roleId int64, permissionIds []int64) error {
	if err := pz.CheckEdit(ds.EntityTypes.Role, roleId); err != nil {
		return err
	}

	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			DELETE FROM
				role_permission
			WHERE
				role_id = $1
			`, roleId); err != nil {
			return err
		}

		for _, permissionId := range permissionIds {
			if _, err := tx.Exec(`
				INSERT INTO
					role_permission
				VALUES
					($1, $2)
				`, roleId, permissionId); err != nil {
				return err
			}
		}

		permissionDescriptions, err := ds.toPermissionDescriptions(permissionIds)
		if err != nil {
			return err
		}
		permissions, err := json.Marshal(permissionDescriptions)
		if err != nil {
			return err
		}
		return ds.audit(pz, tx, UpdateOp, ds.EntityTypes.Role, roleId, metadata{"permissions": string(permissions)})
	})
}

func (ds *Datastore) LinkRoleWithPermission(pz az.Principal, roleId int64, permissionId int64) error {
	if err := pz.CheckEdit(ds.EntityTypes.Role, roleId); err != nil {
		return err
	}

	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			INSERT INTO
				role_permission
			VALUES
				($1, $2)
			`, roleId, permissionId); err != nil {
			return err
		}

		permissionDescription, err := ds.toPermissionDescription(permissionId)
		if err != nil {
			return err
		}
		permissions, err := json.Marshal([]string{permissionDescription})
		if err != nil {
			return err
		}
		return ds.audit(pz, tx, LinkOp, ds.EntityTypes.Role, roleId, metadata{"permissions": string(permissions)})
	})
}

func (ds *Datastore) UnlinkRoleFromPermission(pz az.Principal, roleId int64, permissionId int64) error {
	if err := pz.CheckEdit(ds.EntityTypes.Role, roleId); err != nil {
		return err
	}

	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			DELETE FROM
				role_permission
			WHERE
				role_id = $1 AND
				permission_id = $2
			`, roleId, permissionId); err != nil {
			return err
		}

		permissionDescription, err := ds.toPermissionDescription(permissionId)
		if err != nil {
			return err
		}
		permissions, err := json.Marshal([]string{permissionDescription})
		if err != nil {
			return err
		}
		return ds.audit(pz, tx, UnlinkOp, ds.EntityTypes.Role, roleId, metadata{"permissions": string(permissions)})
	})
}

func (ds *Datastore) DeleteRole(pz az.Principal, roleId int64) error {
	if err := pz.CheckOwns(ds.EntityTypes.Role, roleId); err != nil {
		return err
	}

	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			DELETE FROM
				role
			WHERE
				id = $1
            `, roleId); err != nil {
			return err
		}
		if err := deletePrivilegesOn(tx, ds.EntityTypes.Role, roleId); err != nil {
			return err
		}
		return ds.audit(pz, tx, DeleteOp, ds.EntityTypes.Role, roleId, metadata{})
	})
}

// --- Workgroup ---

func (ds *Datastore) CreateWorkgroup(pz az.Principal, name, description string) (int64, error) {
	var id int64
	err := ds.exec(func(tx *sql.Tx) error {
		res, err := tx.Exec(`
			INSERT INTO
				workgroup
				(type,          name, description, created)
			VALUES
				('workgroup',   $1,   $2,          datetime('now'))
			`, name, description)
		if err != nil {
			return errors.Wrapf(err, "failed creating workgroup %s", name)
		}
		id, err = res.LastInsertId()
		if err != nil {
			return err
		}
		if err := createPrivilege(tx, Privilege{
			Owns,
			pz.WorkgroupId(),
			ds.EntityTypes.Workgroup,
			id,
		}); err != nil {
			return err
		}
		return ds.audit(pz, tx, CreateOp, ds.EntityTypes.Workgroup, id, metadata{
			"name":        name,
			"description": description,
		})
	})
	return id, err
}

func (ds *Datastore) ReadWorkgroups(pz az.Principal, offset, limit int64) ([]Workgroup, error) {
	rows, err := ds.db.Query(`
		SELECT
			id, type, name, description, created
		FROM
			workgroup
		WHERE
			type = 'workgroup' AND
			id IN
			(
				SELECT DISTINCT
					entity_id
				FROM 
					privilege
				WHERE
				  	$1 OR
					(
						workgroup_id IN 
						(
							SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $2
						) AND
						entity_type_id = $3
					)
			)
		ORDER BY name
		LIMIT $4
		OFFSET $5
		`, pz.IsSuperuser(), pz.Id(), ds.EntityTypes.Workgroup, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanWorkgroups(rows)
}

func (ds *Datastore) ReadWorkgroupsForIdentity(pz az.Principal, identityId int64) ([]Workgroup, error) {
	if err := pz.CheckView(ds.EntityTypes.Identity, identityId); err != nil {
		return nil, err
	}

	rows, err := ds.db.Query(`
		SELECT
			w.id, w.type, w.name, w.description, w.created
		FROM
			workgroup w,
			identity_workgroup iw
		WHERE
		  iw.identity_id = $1 AND
			iw.workgroup_id = w.id AND
			w.type = 'workgroup' AND
			w.id IN
			(
				SELECT DISTINCT
					entity_id
				FROM 
					privilege
				WHERE
				  	$2 OR
					(
						workgroup_id IN 
						(
							SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $3
						) AND 
						entity_type_id = $4
					)
			)
		ORDER BY
			w.name
		`, identityId, pz.IsSuperuser(), pz.Id(), ds.EntityTypes.Workgroup)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanWorkgroups(rows)
}

func (ds *Datastore) ReadWorkgroup(pz az.Principal, workgroupId int64) (Workgroup, error) {
	if err := pz.CheckView(ds.EntityTypes.Workgroup, workgroupId); err != nil {
		return Workgroup{}, err
	}

	row := ds.db.QueryRow(`
		SELECT
			id, type, name, description, created
		FROM
			workgroup
		WHERE
			id = $1 AND
			type = 'workgroup'
		ORDER BY
			name
		`, workgroupId)

	return ScanWorkgroup(row)
}

func (ds *Datastore) ReadWorkgroupByName(pz az.Principal, name string) (Workgroup, error) {
	row := ds.db.QueryRow(`
		SELECT
			id, type, name, description, created
		FROM
			workgroup
		WHERE
			name = $1
		`, name)

	workgroup, err := ScanWorkgroup(row)
	if err != nil {
		return Workgroup{}, err
	}

	if err := pz.CheckView(ds.EntityTypes.Workgroup, workgroup.Id); err != nil {
		return Workgroup{}, err
	}

	return workgroup, nil
}

func (ds *Datastore) UpdateWorkgroup(pz az.Principal, workgroupId int64, name, description string) error {
	if err := pz.CheckEdit(ds.EntityTypes.Workgroup, workgroupId); err != nil {
		return err
	}
	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			UPDATE
				workgroup
			SET
				name = $1,
				description = $2
			WHERE
				id = $3
			`, name, description, workgroupId); err != nil {
			return err
		}
		return ds.audit(pz, tx, UpdateOp, ds.EntityTypes.Workgroup, workgroupId, metadata{
			"name":        name,
			"description": description,
		})
	})
}

func (ds *Datastore) DeleteWorkgroup(pz az.Principal, workgroupId int64) error {
	if err := pz.CheckOwns(ds.EntityTypes.Workgroup, workgroupId); err != nil {
		return err
	}
	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			DELETE FROM
				workgroup
			WHERE
				id = $1
			`, workgroupId); err != nil {
			return err
		}
		if err := deletePrivilegesOn(tx, ds.EntityTypes.Workgroup, workgroupId); err != nil {
			return err
		}
		if err := deletePrivilegesFor(tx, ForWorkgroup, workgroupId); err != nil {
			return err
		}
		return ds.audit(pz, tx, DeleteOp, ds.EntityTypes.Workgroup, workgroupId, metadata{})
	})
}

// --- Identity ---

func createDefaultWorkgroup(tx *sql.Tx, name string) (int64, error) {
	res, err := tx.Exec(`
			INSERT INTO
				workgroup
				(type,       name, description, created)
			VALUES
				('identity', $1,   '',          datetime('now'))
			`, "user:"+name)
	if err != nil {
		return 0, errors.Wrap(err, "failed creating default workgroup")
	}
	return res.LastInsertId()
}

func createIdentity(tx *sql.Tx, name, password string, workgroupId int64) (int64, error) {
	res, err := tx.Exec(`
			INSERT INTO
				identity
				(name, password, workgroup_id, is_active, created)
			VALUES
				($1,   $2,       $3,           $4,        datetime('now'))
			`, name, password, workgroupId, true)
	if err != nil {
		return 0, errors.Wrap(err, "failed creating identity")
	}
	return res.LastInsertId()
}

func linkIdentityAndWorkgroup(tx *sql.Tx, identityId, workgroupId int64) error {
	_, err := tx.Exec(`
			INSERT INTO
				identity_workgroup
			VALUES
				($1, $2)
	`, identityId, workgroupId)
	return err
}

func unlinkIdentityAndWorkgroup(tx *sql.Tx, identityId, workgroupId int64) error {
	_, err := tx.Exec(`
			DELETE FROM
				identity_workgroup
			WHERE
				identity_id = $1 AND
				workgroup_id = $2
	`, identityId, workgroupId)
	return err
}

func (ds *Datastore) CreateIdentity(pz az.Principal, name, password string) (int64, int64, error) {
	var id, workgroupId int64
	err := ds.exec(func(tx *sql.Tx) error {
		var err error

		workgroupId, err = createDefaultWorkgroup(tx, name)
		if err != nil {
			return err
		}

		id, err = createIdentity(tx, name, password, workgroupId)
		if err != nil {
			return err
		}

		if err := linkIdentityAndWorkgroup(tx, id, workgroupId); err != nil {
			return err
		}

		if err := createPrivilege(tx, Privilege{
			Owns,
			workgroupId,
			ds.EntityTypes.Identity,
			id,
		}); err != nil {
			return err
		}

		if err := createPrivilege(tx, Privilege{
			Owns,
			workgroupId,
			ds.EntityTypes.Workgroup,
			workgroupId,
		}); err != nil {
			return err
		}

		return ds.audit(pz, tx, CreateOp, ds.EntityTypes.Identity, id, metadata{"name": name})
	})
	return id, workgroupId, err
}

func (ds *Datastore) ReadIdentities(pz az.Principal, offset, limit int64) ([]Identity, error) {
	rows, err := ds.db.Query(`
		SELECT
			id, name, is_active, last_login, created
		FROM
			identity
		WHERE
			id IN
			(
				SELECT DISTINCT
					entity_id
				FROM 
					privilege
				WHERE
				  	$1 OR
					(
						workgroup_id IN 
						(
							SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $2
						) AND 
						entity_type_id = $3
					)
			)
		ORDER BY name
		LIMIT $4
		OFFSET $5
		`, pz.IsSuperuser(), pz.Id(), ds.EntityTypes.Identity, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanIdentitys(rows)
}

func (ds *Datastore) ReadIdentity(pz az.Principal, identityId int64) (Identity, error) {
	if err := pz.CheckView(ds.EntityTypes.Identity, identityId); err != nil {
		return Identity{}, err
	}

	row := ds.db.QueryRow(`
		SELECT
			id, name, is_active, last_login, created
		FROM
			identity
		WHERE
			id = $1
		`, identityId)

	return ScanIdentity(row)
}

func (ds *Datastore) ReadIdentityByName(pz az.Principal, name string) (Identity, error) {
	row := ds.db.QueryRow(`
		SELECT
			id, name, is_active, last_login, created
		FROM
			identity
		WHERE
			name = $1
		`, name)

	identity, err := ScanIdentity(row)
	if err != nil {
		return Identity{}, err
	}

	if err := pz.CheckView(ds.EntityTypes.Identity, identity.Id); err != nil {
		return Identity{}, err
	}

	return identity, nil
}

func (ds *Datastore) readIdentityAndPassword(name string) (*IdentityAndPassword, error) {
	rows, err := ds.db.Query(`
		SELECT
			id, name, password, workgroup_id, is_active, last_login, created
		FROM
			identity
		WHERE
			name = $1
		`, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	identities, err := ScanIdentityAndPasswords(rows)
	if err != nil {
		return nil, err
	}

	if len(identities) == 0 {
		return nil, nil
	}

	return &identities[0], nil
}

func (ds *Datastore) ReadIdentitiesForWorkgroup(pz az.Principal, workgroupId int64) ([]Identity, error) {
	if err := pz.CheckView(ds.EntityTypes.Workgroup, workgroupId); err != nil {
		return nil, err
	}

	rows, err := ds.db.Query(`
		SELECT
			i.id, i.name, i.is_active, i.last_login, i.created
		FROM
			identity i,
			identity_workgroup iw
		WHERE
			iw.workgroup_id = $1 AND
		  iw.identity_id = i.id AND
			i.id IN
			(
				SELECT DISTINCT
					entity_id
				FROM 
					privilege
				WHERE
					$2 OR
					(
						workgroup_id IN 
						(
							SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $3
						) AND
						entity_type_id = $4
					)
			)
		ORDER BY
			i.name
		`, workgroupId, pz.IsSuperuser(), pz.Id(), ds.EntityTypes.Identity)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanIdentitys(rows)
}

func (ds *Datastore) ReadIdentitiesForRole(pz az.Principal, roleId int64) ([]Identity, error) {
	if err := pz.CheckView(ds.EntityTypes.Role, roleId); err != nil {
		return nil, err
	}
	rows, err := ds.db.Query(`
		SELECT
			i.id, i.name, i.is_active, i.last_login, i.created
		FROM
			identity i,
			identity_role ir
		WHERE
			ir.role_id = $1 AND
		  ir.identity_id = i.id AND
			i.id IN
			(
				SELECT DISTINCT
					entity_id
				FROM 
					privilege
				WHERE
				  	$2 OR
					(
						workgroup_id IN 
						(
							SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $3
						) AND
						entity_type_id = $4
					)
			)
		ORDER BY
			i.name
		`, roleId, pz.IsSuperuser(), pz.Id(), ds.EntityTypes.Identity)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanIdentitys(rows)
}

func (ds *Datastore) ReadUsersForEntity(pz az.Principal, entityTypeId, entityId int64) ([]IdentityAndRole, error) {
	rows, err := ds.db.Query(`
		SELECT 
			privilege.privilege_type,
			identity.id, 
			identity.name, 
			role.id,
			role.name
		FROM
			identity, role,
			identity_role,
			privilege,
			identity_workgroup
		WHERE 
			identity_role.identity_id = identity.id AND
			identity_role.role_id = role.id AND
			privilege.workgroup_id = identity_workgroup.workgroup_id AND
  			identity_workgroup.identity_id = identity.id AND
			privilege.entity_id = $1 AND
			privilege.entity_type_id = $2;
		`, entityId, entityTypeId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanIdentityAndRoles(rows)
}

func (ds *Datastore) LinkIdentityAndWorkgroup(pz az.Principal, identityId, workgroupId int64) error {
	if err := pz.CheckView(ds.EntityTypes.Workgroup, workgroupId); err != nil {
		return err
	}
	if err := pz.CheckEdit(ds.EntityTypes.Identity, identityId); err != nil {
		return err
	}

	workgroup, err := ds.ReadWorkgroup(pz, workgroupId)
	if err != nil {
		return err
	}
	return ds.exec(func(tx *sql.Tx) error {
		err := linkIdentityAndWorkgroup(tx, identityId, workgroupId)
		if err != nil {
			return err
		}
		return ds.audit(pz, tx, LinkOp, ds.EntityTypes.Identity, identityId, metadata{
			"type": WorkgroupEntity,
			"id":   strconv.FormatInt(workgroupId, 10),
			"name": workgroup.Name,
		})
	})
}

func (ds *Datastore) UnlinkIdentityAndWorkgroup(pz az.Principal, identityId, workgroupId int64) error {
	if err := pz.CheckView(ds.EntityTypes.Workgroup, workgroupId); err != nil {
		return err
	}
	if err := pz.CheckEdit(ds.EntityTypes.Identity, identityId); err != nil {
		return err
	}

	workgroup, err := ds.ReadWorkgroup(pz, workgroupId)
	if err != nil {
		return err
	}
	return ds.exec(func(tx *sql.Tx) error {
		err := unlinkIdentityAndWorkgroup(tx, identityId, workgroupId)
		if err != nil {
			return err
		}
		return ds.audit(pz, tx, UnlinkOp, ds.EntityTypes.Identity, identityId, metadata{
			"type": WorkgroupEntity,
			"id":   strconv.FormatInt(workgroupId, 10),
			"name": workgroup.Name,
		})
	})
}

func linkIdentityAndRole(tx *sql.Tx, identityId, roleId int64) error {
	_, err := tx.Exec(`
		INSERT INTO
			identity_role
		VALUES
			($1, $2)
		`, identityId, roleId)
	return err
}

func (ds *Datastore) LinkIdentityAndRole(pz az.Principal, identityId, roleId int64) error {
	if err := pz.CheckView(ds.EntityTypes.Role, roleId); err != nil {
		return err
	}
	if err := pz.CheckEdit(ds.EntityTypes.Identity, identityId); err != nil {
		return err
	}

	role, err := ds.ReadRole(pz, roleId)
	if err != nil {
		return err
	}
	return ds.exec(func(tx *sql.Tx) error {
		if err := linkIdentityAndRole(tx, identityId, roleId); err != nil {
			return err
		}
		return ds.audit(pz, tx, LinkOp, ds.EntityTypes.Identity, identityId, metadata{
			"type": RoleEntity,
			"id":   strconv.FormatInt(roleId, 10),
			"name": role.Name,
		})
	})
}

func (ds *Datastore) UnlinkIdentityAndRole(pz az.Principal, identityId, roleId int64) error {
	if err := pz.CheckView(ds.EntityTypes.Role, roleId); err != nil {
		return err
	}
	if err := pz.CheckEdit(ds.EntityTypes.Identity, identityId); err != nil {
		return err
	}

	role, err := ds.ReadRole(pz, roleId)
	if err != nil {
		return err
	}
	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			DELETE FROM
				identity_role
			WHERE
				identity_id = $1 AND
				role_id = $2
			`, identityId, roleId); err != nil {
			return err
		}
		return ds.audit(pz, tx, UnlinkOp, ds.EntityTypes.Identity, identityId, metadata{
			"type": RoleEntity,
			"id":   strconv.FormatInt(roleId, 10),
			"name": role.Name,
		})
	})
}

func (ds *Datastore) UpdateIdentity(pz az.Principal, identityId int64, password string) error {
	if err := pz.CheckEdit(ds.EntityTypes.Identity, identityId); err != nil {
		return err
	}

	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			UPDATE
				identity
			SET
				password = $1
			WHERE
				id = $2
			`, password, identityId); err != nil {
			return err
		}

		return ds.audit(pz, tx, UpdateOp, ds.EntityTypes.Identity, identityId, metadata{"password": "(changed)"})
	})
}

func (ds *Datastore) ActivateIdentity(pz az.Principal, identityId int64) error {
	if err := pz.CheckOwns(ds.EntityTypes.Identity, identityId); err != nil {
		return err
	}

	identity, err := ds.ReadIdentity(pz, identityId)
	if err != nil {
		return err
	}
	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			UPDATE
				identity
			SET
				is_active = 1
			WHERE
				id = $1
			`, identityId); err != nil {
			return err
		}

		return ds.audit(pz, tx, EnableOp, ds.EntityTypes.Identity, identityId, metadata{"name": identity.Name})
	})
}

func (ds *Datastore) DeactivateIdentity(pz az.Principal, identityId int64) error {
	if err := pz.CheckOwns(ds.EntityTypes.Identity, identityId); err != nil {
		return err
	}

	identity, err := ds.ReadIdentity(pz, identityId)
	if err != nil {
		return err
	}
	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			UPDATE
				identity
			SET
				is_active = 0
			WHERE
				id = $1
			`, identityId); err != nil {
			return err
		}
		return ds.audit(pz, tx, DisableOp, ds.EntityTypes.Identity, identityId, metadata{"name": identity.Name})
	})
	return nil
}

// --- Privileges ---

func readWorkgroupName(tx *sql.Tx, workgroupId int64) (string, error) {
	row := tx.QueryRow("SELECT name FROM workgroup WHERE id = $1", workgroupId)
	var name string
	err := row.Scan(&name)
	return name, err
}

func (ds *Datastore) CreatePrivilege(pz az.Principal, privilege Privilege) error {
	if err := pz.CheckView(ds.EntityTypes.Workgroup, privilege.WorkgroupId); err != nil {
		return err
	}

	if err := pz.CheckOwns(privilege.EntityType, privilege.EntityId); err != nil {
		return err
	}

	return ds.exec(func(tx *sql.Tx) error {
		if err := createPrivilege(tx, privilege); err != nil {
			return err
		}

		identityName, err := readWorkgroupName(tx, privilege.WorkgroupId)
		if err != nil {
			return err
		}

		return ds.audit(pz, tx, ShareOp, privilege.EntityType, privilege.EntityId, metadata{
			"id":   strconv.FormatInt(privilege.WorkgroupId, 10),
			"name": identityName,
		})
	})
}

func (ds *Datastore) ReadEntityPrivileges(pz az.Principal, entityTypeId, entityId int64) ([]EntityPrivilege, error) {
	if err := pz.CheckView(entityTypeId, entityId); err != nil {
		return nil, err
	}

	rows, err := ds.db.Query(`
		SELECT
		  p.privilege_type, w.id, w.name, w.description
		FROM
			privilege p,
			workgroup w
		WHERE
			p.entity_id = $1 AND
			p.entity_type_id = $2 AND
			w.id = p.workgroup_id
		`, entityId, entityTypeId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanEntityPrivileges(rows)
}

func (ds *Datastore) readPrivileges(identityId, entityTypeId, entityId int64) ([]string, error) {
	rows, err := ds.db.Query(`
		SELECT DISTINCT
			privilege_type
		FROM
			privilege
		WHERE
			entity_id = $1 AND
			entity_type_id = $2 AND
			workgroup_id IN (SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $3)
		`, entityId, entityTypeId, identityId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanStrings(rows)
}

func (ds *Datastore) DeletePrivilege(pz az.Principal, privilege Privilege) error {
	if err := pz.CheckView(ds.EntityTypes.Workgroup, privilege.WorkgroupId); err != nil {
		return err
	}

	if err := pz.CheckOwns(privilege.EntityType, privilege.EntityId); err != nil {
		return err
	}

	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			DELETE FROM
				privilege
			WHERE
				privilege_type = $1 AND
				workgroup_id = $2 AND
				entity_type_id = $3 AND
				entity_id = $4
			`,
			privilege.Type,
			privilege.WorkgroupId,
			privilege.EntityType,
			privilege.EntityId,
		); err != nil {
			return err
		}

		identityName, err := readWorkgroupName(tx, privilege.WorkgroupId)
		if err != nil {
			return err
		}

		return ds.audit(pz, tx, UnshareOp, privilege.EntityType, privilege.EntityId, metadata{
			"id":   strconv.FormatInt(privilege.WorkgroupId, 10),
			"name": identityName,
		})
	})
}

func createPrivilege(tx *sql.Tx, privilege Privilege) error {

	_, err := tx.Exec(`
			INSERT INTO
				privilege
			VALUES
				($1, $2, $3, $4)
			`,
		privilege.Type,
		privilege.WorkgroupId,
		privilege.EntityType,
		privilege.EntityId,
	)
	return errors.Wrap(err, "executing query")
}

func deletePrivilegesOn(tx *sql.Tx, entityTypeId, entityId int64) error {
	_, err := tx.Exec(`
		DELETE FROM
			privilege
		WHERE
			entity_type_id = $1 AND
			entity_id = $2
		`, entityTypeId, entityId)
	return err
}

func deletePrivilegesFor(tx *sql.Tx, identityType string, identityId int64) error {
	_, err := tx.Exec(`
		DELETE FROM
			privilege
		WHERE
			workgroup_id = (SELECT workgroup_id FROM identity WHERE id = $1)
		`, identityId)
	return err
}

// --- Engine ---

func (ds *Datastore) CreateEngine(pz az.Principal, name, location string) (int64, error) {
	var id int64
	err := ds.exec(func(tx *sql.Tx) error {
		res, err := tx.Exec(`
			INSERT INTO
				engine
				(name, location, created)
			VALUES
				($1,   $2,       datetime('now'))
			`, name, location)
		if err != nil {
			return err
		}

		id, err = res.LastInsertId()
		if err != nil {
			return err
		}

		if err := createPrivilege(tx, Privilege{
			Owns,
			pz.WorkgroupId(),
			ds.EntityTypes.Engine,
			id,
		}); err != nil {
			return err
		}

		return ds.audit(pz, tx, CreateOp, ds.EntityTypes.Engine, id, metadata{
			"name":     name,
			"location": location,
		})

	})
	return id, err
}

func (ds *Datastore) ReadEngines(pz az.Principal) ([]Engine, error) {
	rows, err := ds.db.Query(`
		SELECT
			id, name, location, created
		FROM
			engine
		WHERE
			id IN
			(
				SELECT DISTINCT
					entity_id
				FROM 
					privilege
				WHERE
				  	$1 OR
					(
						workgroup_id IN 
						(
							SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $2
						) AND
						entity_type_id = $3
					)
			)
		ORDER BY
			name
		`, pz.IsSuperuser(), pz.Id(), ds.EntityTypes.Engine)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanEngines(rows)
}

func (ds *Datastore) ReadEngine(pz az.Principal, engineId int64) (Engine, error) {
	if err := pz.CheckView(ds.EntityTypes.Engine, engineId); err != nil {
		return Engine{}, err
	}

	row := ds.db.QueryRow(`
		SELECT
			id, name, location, created
		FROM
			engine
		WHERE
			id = $1
		`, engineId)
	return ScanEngine(row)
}

func (ds *Datastore) DeleteEngine(pz az.Principal, engineId int64) error {
	if err := pz.CheckOwns(ds.EntityTypes.Engine, engineId); err != nil {
		return err
	}

	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			DELETE FROM
				engine
			WHERE
				id = $1
			`, engineId); err != nil {
			return err
		}
		if err := deletePrivilegesOn(tx, ds.EntityTypes.Engine, engineId); err != nil {
			return err
		}
		return ds.audit(pz, tx, DeleteOp, ds.EntityTypes.Engine, engineId, metadata{})
	})
}

// --- Cluster ---

func (ds *Datastore) CreateExternalCluster(pz az.Principal, name, address, state string) (int64, error) {
	var id int64
	err := ds.exec(func(tx *sql.Tx) error {
		res, err := tx.Exec(`
			INSERT INTO
				cluster
				(name, type_id, detail_id, address, token, state, created)
			VALUES
				($1,   $2,      0,         $3,      $4,      $5,    datetime('now'))
			`, name, ds.ClusterTypes.External, address, "", state)
		// TODO add token?!
		if err != nil {
			return err
		}

		id, err = res.LastInsertId()
		if err != nil {
			return err
		}

		if err := createPrivilege(tx, Privilege{
			Owns,
			pz.WorkgroupId(),
			ds.EntityTypes.Cluster,
			id,
		}); err != nil {
			return err
		}

		return ds.audit(pz, tx, CreateOp, ds.EntityTypes.Cluster, id, metadata{
			"name":    name,
			"type":    ClusterExternal,
			"address": address,
			"state":   state,
		})
	})
	return id, err
}

func (ds *Datastore) CreateYarnCluster(pz az.Principal, name, address, token, state string, cluster YarnCluster) (int64, error) {
	var clusterId int64
	err := ds.exec(func(tx *sql.Tx) error {
		var yarnClusterId int64
		if res, err := tx.Exec(`
			INSERT INTO
				cluster_yarn
				(engine_id, size, application_id, memory, username, output_dir)
			VALUES
				($1,        $2,   $3,             $4,     $5,       $6)
			`,
			cluster.EngineId,
			cluster.Size,
			cluster.ApplicationId,
			cluster.Memory,
			cluster.Username,
			cluster.OutputDir,
		); err != nil {
			return err
		} else {
			yarnClusterId, err = res.LastInsertId()
			if err != nil {
				return err
			}
		}

		res, err := tx.Exec(`
			INSERT INTO
				cluster
				(name, type_id, detail_id, address, token, state, created)
			VALUES
				($1,   $2,      $3,        $4,      $5,      $6,    datetime('now'))
			`, name, ds.ClusterTypes.Yarn, yarnClusterId, address, token, state)
		if err != nil {
			return err
		}

		clusterId, err = res.LastInsertId()
		if err != nil {
			return err
		}

		if err := createPrivilege(tx, Privilege{
			Owns,
			pz.WorkgroupId(),
			ds.EntityTypes.Cluster,
			clusterId,
		}); err != nil {
			return err
		}

		return ds.audit(pz, tx, CreateOp, ds.EntityTypes.Cluster, clusterId, metadata{
			"name":            name,
			"type":            ClusterYarn,
			"address":         address,
			"token": 	   token,
			"state":           state,
			"engineId":        strconv.FormatInt(cluster.EngineId, 10),
			"size":            strconv.FormatInt(cluster.Size, 10),
			"applicationId":   cluster.ApplicationId,
			"memory":          cluster.Memory,
			"username":        cluster.Username,
			"outputDirectory": cluster.OutputDir,
		})
	})
	return clusterId, err
}

func (ds *Datastore) ReadClusterTypes(pz az.Principal) []ClusterType {
	return ds.clusterTypes
}

func (ds *Datastore) ReadClusters(pz az.Principal, offset, limit int64) ([]Cluster, error) {
	rows, err := ds.db.Query(`
		SELECT
			id, name, type_id, detail_id, address, token, state, created
		FROM
			cluster
		WHERE
			id IN
			(
				SELECT DISTINCT
					entity_id
				FROM 
					privilege
				WHERE
					$1 OR
					(
						workgroup_id IN 
						(
							SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $2
						) AND
						entity_type_id = $3
					)
			)
		ORDER BY
			name
		LIMIT $4
		OFFSET $5
		`, pz.IsSuperuser(), pz.Id(), ds.EntityTypes.Cluster, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanClusters(rows)
}

func (ds *Datastore) ReadCluster(pz az.Principal, clusterId int64) (Cluster, error) {
	if err := pz.CheckView(ds.EntityTypes.Cluster, clusterId); err != nil {
		return Cluster{}, err
	}
	row := ds.db.QueryRow(`
		SELECT
			id, name, type_id, detail_id, address, token, state, created
		FROM
			cluster
		WHERE
			id = $1
		`, clusterId)

	return ScanCluster(row)
}

func (ds *Datastore) ReadClusterByAddress(pz az.Principal, address string) (Cluster, bool, error) {
	var cluster Cluster
	rows, err := ds.db.Query(`
		SELECT
			id, name, type_id, detail_id, address, token, state, created
		FROM
			cluster
		WHERE
			address = $1
		`, address)

	if err != nil {
		return cluster, false, err
	}
	defer rows.Close()

	return scanCluster(rows)
}

func (ds *Datastore) ReadClusterByName(pz az.Principal, name string) (Cluster, bool, error) {
	var cluster Cluster
	rows, err := ds.db.Query(`
		SELECT
			id, name, type_id, detail_id, address, token, state, created
		FROM
			cluster
		WHERE
			name = $1
		`, name)
	if err != nil {
		return cluster, false, err
	}
	defer rows.Close()

	return scanCluster(rows)
}

func (ds *Datastore) ReadAllClusters(pz az.Principal) ([]Cluster, error) {
	rows, err := ds.db.Query(`
		SELECT
			id, name, type_id, detail_id, address, token, state, created
		FROM
			cluster
		`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanClusters(rows)
}

func scanCluster(rows *sql.Rows) (Cluster, bool, error) {
	var cluster Cluster

	clusters, err := ScanClusters(rows)
	if err != nil {
		return cluster, false, err
	}

	if len(clusters) == 0 {
		return cluster, false, nil
	}

	return clusters[0], true, nil
}

func (ds *Datastore) ReadYarnCluster(pz az.Principal, clusterId int64) (YarnCluster, error) {

	if err := pz.CheckView(ds.EntityTypes.Cluster, clusterId); err != nil {
		return YarnCluster{}, err
	}

	row := ds.db.QueryRow(`
		SELECT
			y.id, y.engine_id, y.size, y.application_id, y.memory, y.username, y.output_dir
		FROM
			cluster c,
			cluster_yarn y
		WHERE
			c.id = $1 AND
			c.detail_id = y.id
		`, clusterId)

	return ScanYarnCluster(row)
}

func (ds *Datastore) UpdateClusterState(pz az.Principal, clusterId int64, state string) error {
	if err := pz.CheckEdit(ds.EntityTypes.Cluster, clusterId); err != nil {
		return err
	}

	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			UPDATE
				cluster
			SET
				state = $1
			WHERE
				id = $2
			`, state, clusterId); err != nil {
			return err
		}
		return ds.audit(pz, tx, UpdateOp, ds.EntityTypes.Cluster, clusterId, metadata{"state": state})
	})
}

func (ds *Datastore) DeleteCluster(pz az.Principal, clusterId int64) error {
	if err := pz.CheckOwns(ds.EntityTypes.Cluster, clusterId); err != nil {
		return err
	}

	cluster, err := ds.ReadCluster(pz, clusterId)
	if err != nil {
		return err
	}

	return ds.exec(func(tx *sql.Tx) error {
		if cluster.TypeId == ds.ClusterTypes.Yarn {
			if _, err := tx.Exec(`
				DELETE FROM
					cluster_yarn
				WHERE
					id = $1
				`, cluster.DetailId); err != nil {
				return err
			}
		}

		if _, err := tx.Exec(`
			DELETE FROM
				cluster
			WHERE
				id = $1
			`, clusterId); err != nil {
			return err
		}

		if err := deletePrivilegesOn(tx, ds.EntityTypes.Cluster, clusterId); err != nil {
			return err
		}
		return ds.audit(pz, tx, DeleteOp, ds.EntityTypes.Cluster, clusterId, metadata{})
	})
}

// --- Project ---

func (ds *Datastore) CreateProject(pz az.Principal, name, description, modelCategory string) (int64, error) {
	var id int64
	err := ds.exec(func(tx *sql.Tx) error {
		res, err := tx.Exec(`
			INSERT INTO
				project
				(name, description, model_category, created)
			VALUES
				($1,   $2,          $3,             datetime('now'))
			`, name, description, modelCategory)
		if err != nil {
			return err
		}

		id, err = res.LastInsertId()
		if err != nil {
			return err
		}

		if err := createPrivilege(tx, Privilege{
			Owns,
			pz.WorkgroupId(),
			ds.EntityTypes.Project,
			id,
		}); err != nil {
			return err
		}

		return ds.audit(pz, tx, CreateOp, ds.EntityTypes.Project, id, metadata{
			"name":           name,
			"description":    description,
			"model_category": modelCategory,
		})
	})
	return id, err
}

func (ds *Datastore) ReadProjects(pz az.Principal, offset, limit int64) ([]Project, error) {
	rows, err := ds.db.Query(`
		SELECT
			id, name, description, model_category, created
		FROM
			project
		WHERE
			id IN
			(
				SELECT DISTINCT
					entity_id
				FROM 
					privilege
				WHERE
				  	$1 OR
					(
						workgroup_id IN 
						(
							SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $2
						) AND 
						entity_type_id = $3
					)
			)
		ORDER BY
			name
		LIMIT $4
		OFFSET $5
		`, pz.IsSuperuser(), pz.Id(), ds.EntityTypes.Project, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanProjects(rows)
}

func (ds *Datastore) ReadProject(pz az.Principal, projectId int64) (Project, error) {
	if err := pz.CheckView(ds.EntityTypes.Project, projectId); err != nil {
		return Project{}, err
	}

	row := ds.db.QueryRow(`
		SELECT
			id, name, description, model_category, created
		FROM
			project
		WHERE
			id = $1
		`, projectId)
	return ScanProject(row)
}

func (ds *Datastore) DeleteProject(pz az.Principal, projectId int64) error {
	if err := pz.CheckOwns(ds.EntityTypes.Project, projectId); err != nil {
		return err
	}

	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			DELETE FROM
				project
			WHERE
				id = $1
			`, projectId); err != nil {
			return err
		}
		if err := deletePrivilegesOn(tx, ds.EntityTypes.Project, projectId); err != nil {
			return err
		}
		return ds.audit(pz, tx, DeleteOp, ds.EntityTypes.Project, projectId, metadata{})
	})
}

// --- Datasource ---
func (ds *Datastore) CreateDatasource(pz az.Principal, datasource Datasource) (int64, error) {
	var id int64
	err := ds.exec(func(tx *sql.Tx) error {
		res, err := tx.Exec(`
			INSERT INTO
				datasource
				(project_id, name, description, kind, configuration, created)
			VALUES
				($1,         $2,   $3,          $4,   $5,            datetime('now'))
			`,
			datasource.ProjectId,
			datasource.Name,
			datasource.Description,
			datasource.Kind,
			datasource.Configuration,
		)
		if err != nil {
			return err
		}

		id, err = res.LastInsertId()
		if err != nil {
			return err
		}

		if err := createPrivilege(tx, Privilege{
			Owns,
			pz.WorkgroupId(),
			ds.EntityTypes.Datasource,
			id,
		}); err != nil {
			return err
		}

		return ds.audit(pz, tx, CreateOp, ds.EntityTypes.Datasource, id, metadata{
			"name":          datasource.Name,
			"description":   datasource.Description,
			"kind":          datasource.Kind,
			"configuration": datasource.Configuration,
		})
	})
	return id, err
}

func (ds *Datastore) ReadDatasources(pz az.Principal, projectId, offset, limit int64) ([]Datasource, error) {
	rows, err := ds.db.Query(`
			SELECT
				id, project_id, name, description, kind, configuration, created
			FROM
				datasource
			WHERE
				id IN
				(
					SELECT DISTINCT
						entity_id
					FROM
						privilege
					WHERE
						$1
						OR
						(
							workgroup_id IN
							(
								SELECT 
									workgroup_id 
								FROM
									identity_workgroup
								WHERE
									identity_id = $2
							)
							AND
							entity_type_id = $3
						)
				)
				AND
				project_id = $4
			ORDER BY
				name
			LIMIT $5
			OFFSET $6
			`, pz.IsSuperuser(), pz.Id(), ds.EntityTypes.Datasource, projectId, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanDatasources(rows)
}

func (ds *Datastore) ReadDatasource(pz az.Principal, datasourceId int64) (Datasource, error) {
	if err := pz.CheckView(ds.EntityTypes.Datasource, datasourceId); err != nil {
		return Datasource{}, err
	}

	row := ds.db.QueryRow(`
		SELECT
			id, project_id, name, description, kind, configuration, created
		FROM
			datasource
		WHERE
			id = $1
		`, datasourceId)
	return ScanDatasource(row)
}

func (ds *Datastore) ReadDatasourceByProject(pz az.Principal, projectId int64) (Datasource, bool, error) {
	var datasource Datasource

	rows, err := ds.db.Query(`
		SELECT
			id, project_id, name, description, kind, configuration, created
		FROM
			datasource
		WHERE
			project_id = $1
		`, projectId)
	if err != nil {
		return datasource, false, err
	}
	defer rows.Close()

	return scanDatasources(rows)
}

func scanDatasources(rows *sql.Rows) (Datasource, bool, error) {
	datasources, err := ScanDatasources(rows)
	if err != nil {
		return Datasource{}, false, err
	}

	if len(datasources) == 0 {
		return Datasource{}, false, nil
	}

	return datasources[0], true, nil

}

func (ds *Datastore) UpdateDatasource(pz az.Principal, datasourceId int64, datasource Datasource) error {
	if err := pz.CheckEdit(ds.EntityTypes.Datasource, datasourceId); err != nil {
		return err
	}

	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			UPDATE
				datasource
			SET
				name = $1,
				description = $2,
				kind = $3,
				configuration = $4
			WHERE
				id = $5
			`,
			datasource.Name,
			datasource.Description,
			datasource.Kind,
			datasource.Configuration,
			datasourceId,
		); err != nil {
			return err
		}
		return ds.audit(pz, tx, UpdateOp, ds.EntityTypes.Datasource, datasourceId, metadata{
			"name":          datasource.Name,
			"description":   datasource.Description,
			"kind":          datasource.Kind,
			"configuration": datasource.Configuration,
		})
	})
}

func (ds *Datastore) DeleteDatasource(pz az.Principal, datasourceId int64) error {
	if err := pz.CheckOwns(ds.EntityTypes.Datasource, datasourceId); err != nil {
		return err
	}

	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			DELETE FROM
				datasource
			WHERE
				id = $1
			`, datasourceId); err != nil {
			return err
		}
		if err := deletePrivilegesOn(tx, ds.EntityTypes.Datasource, datasourceId); err != nil {
			return err
		}
		return ds.audit(pz, tx, DeleteOp, ds.EntityTypes.Datasource, datasourceId, metadata{})
	})
}

// --- Dataset ---

func (ds *Datastore) CreateDataset(pz az.Principal, dataset Dataset) (int64, error) {
	var id int64
	err := ds.exec(func(tx *sql.Tx) error {
		res, err := tx.Exec(`
			INSERT INTO
				dataset
				(datasource_id, name, description, frame_name, response_column_name, properties, properties_version, created)
			VALUES
				($1,            $2,   $3,          $4,         $5,                   $6,         $7,                 datetime('now'))
			`,
			dataset.DatasourceId,
			dataset.Name,
			dataset.Description,
			dataset.FrameName,
			dataset.ResponseColumnName,
			dataset.Properties,
			dataset.PropertiesVersion,
		)
		if err != nil {
			return err
		}

		id, err = res.LastInsertId()
		if err != nil {
			return err
		}

		if err := createPrivilege(tx, Privilege{
			Owns,
			pz.WorkgroupId(),
			ds.EntityTypes.Dataset,
			id,
		}); err != nil {
			return err
		}

		return ds.audit(pz, tx, CreateOp, ds.EntityTypes.Dataset, id, metadata{
			"name":               dataset.Name,
			"description":        dataset.Description,
			"responseColumnName": dataset.ResponseColumnName,
		})
	})
	return id, err
}

func (ds *Datastore) ReadDatasets(pz az.Principal, datasourceId, offset, limit int64) ([]Dataset, error) {
	rows, err := ds.db.Query(`
			SELECT
				id, datasource_id, name, description, frame_name, response_column_name, properties, properties_version, created
			FROM
				dataset
			WHERE
				id IN
				(
					SELECT DISTINCT
						entity_id
					FROM
						privilege
					WHERE
						$1
						OR
						(
							workgroup_id IN
							(
								SELECT 
									workgroup_id 
								FROM
									identity_workgroup
								WHERE
									identity_id = $2
							)
							AND
							entity_type_id = $3
						)
				)
				AND
				datasource_id = $4
			ORDER BY
				name
			LIMIT $5
			OFFSET $6
			`, pz.IsSuperuser(), pz.Id(), ds.EntityTypes.Dataset, datasourceId, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanDatasets(rows)
}

func (ds *Datastore) ReadDataset(pz az.Principal, datasetId int64) (Dataset, error) {
	if err := pz.CheckView(ds.EntityTypes.Dataset, datasetId); err != nil {
		return Dataset{}, err
	}

	row := ds.db.QueryRow(`
		SELECT
			id, datasource_id, name, description, frame_name, response_column_name, properties, properties_version, created
		FROM
			dataset
		WHERE
			id = $1
		`, datasetId)
	return ScanDataset(row)
}

func (ds *Datastore) ReadDatasetByDatasource(pz az.Principal, datasourceId int64) (Dataset, bool, error) {
	var dataset Dataset
	rows, err := ds.db.Query(`
		SELECT
			id, datasource_id, name, description, frame_name, response_column_name, properties, properties_version, created
		FROM
			dataset
		WHERE
			datasource_id = $1
		`, datasourceId)
	if err != nil {
		return dataset, false, err
	}
	defer rows.Close()

	return scanDatasets(rows)
}

func scanDatasets(rows *sql.Rows) (Dataset, bool, error) {
	var dataset Dataset

	datasets, err := ScanDatasets(rows)
	if err != nil {
		return dataset, false, err
	}

	if len(datasets) == 0 {
		return dataset, false, nil
	}

	return datasets[0], true, nil
}

func (ds *Datastore) UpdateDataset(pz az.Principal, datasetId int64, dataset Dataset) error {
	if err := pz.CheckEdit(ds.EntityTypes.Dataset, datasetId); err != nil {
		return err
	}

	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			UPDATE
				dataset
			SET
				name = $1,
				description = $2,
				responseColumnName = $3
			WHERE
				id = $4
			`,
			dataset.Name,
			dataset.Description,
			dataset.ResponseColumnName,
			datasetId); err != nil {
			return err
		}
		return ds.audit(pz, tx, UpdateOp, ds.EntityTypes.Dataset, datasetId, metadata{
			"name":               dataset.Name,
			"description":        dataset.Description,
			"responseColumnName": dataset.ResponseColumnName,
		})
	})
}

func (ds *Datastore) DeleteDataset(pz az.Principal, datasetId int64) error {
	if err := pz.CheckOwns(ds.EntityTypes.Dataset, datasetId); err != nil {
		return err
	}

	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			DELETE FROM
				dataset
			WHERE
				id = $1
			`, datasetId); err != nil {
			return err
		}
		if err := deletePrivilegesOn(tx, ds.EntityTypes.Dataset, datasetId); err != nil {
			return err
		}
		return ds.audit(pz, tx, DeleteOp, ds.EntityTypes.Dataset, datasetId, metadata{})
	})
}

// --- Model ---

func (ds *Datastore) CreateModel(pz az.Principal, model Model) (int64, error) {
	var id int64
	err := ds.exec(func(tx *sql.Tx) error {
		res, err := tx.Exec(`
			INSERT INTO
				model
				(
					project_id,
					training_dataset_id,
					validation_dataset_id,
					name,
					cluster_name,
					cluster_id,
					model_key,
					algorithm,
					model_category,
					dataset_name,
					response_column_name,
					logical_name,
					location,
					max_run_time,
					metrics,
					metrics_version,
					created
				)
			VALUES
				($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, datetime('now'))
			`,
			model.ProjectId,           //$1
			model.TrainingDatasetId,   //$2
			model.ValidationDatasetId, //$3
			model.Name,                //$4
			model.ClusterName,         //$5
			model.ClusterId,           //$6
			model.ModelKey,            //$7
			model.Algorithm,           //$8
			model.ModelCategory,       //$9
			model.DatasetName,         //$10
			model.ResponseColumnName,  //$11
			model.LogicalName,         //$12
			model.Location,            //$13
			model.MaxRunTime,          //$14
			model.Metrics,             //$15
			model.MetricsVersion,      //$16
		)
		if err != nil {
			return err
		}

		id, err = res.LastInsertId()
		if err != nil {
			return err
		}

		if err := createPrivilege(tx, Privilege{
			Owns,
			pz.WorkgroupId(),
			ds.EntityTypes.Model,
			id,
		}); err != nil {
			return err
		}

		return ds.audit(pz, tx, CreateOp, ds.EntityTypes.Model, id, metadata{
			"name":               model.Name,
			"clusterName":        model.ClusterName,
			"modelKey":           model.ModelKey,
			"algorithm":          model.Algorithm,
			"datasetName":        model.DatasetName,
			"responseColumnName": model.ResponseColumnName,
			"logicalName":        model.LogicalName.String,
			"location":           model.Location,
			"maxRunTime":         strconv.FormatInt(model.MaxRunTime, 10),
		})
	})
	return id, err
}

func (ds *Datastore) CreateBinomialModel(pz az.Principal, modelId int64, mse, rSquared, logloss, auc, gini float64) error {
	var id int64
	err := ds.exec(func(tx *sql.Tx) error {
		res, err := tx.Exec(`
			INSERT INTO
				binomial_model
				(model_id, mse, r_squared, logloss, auc, gini)
			VALUES
				($1,      $2,  $3,        $4,      $5,  $6)
			`,
			modelId,
			mse,
			rSquared,
			logloss,
			auc,
			gini,
		)
		if err != nil {
			return err
		}

		id, err = res.LastInsertId()
		if err != nil {
			return err
		}

		return nil
	})
	return err
}

func (ds *Datastore) CreateMultinomialModel(pz az.Principal, modelId int64, mse, rSquared, logloss float64) error {
	var id int64
	err := ds.exec(func(tx *sql.Tx) error {
		res, err := tx.Exec(`
			INSERT INTO
				multinomial_model
				(model_id, mse, r_squared, logloss)
			VALUES
				($1,      $2,  $3,        $4)
			`,
			modelId,
			mse,
			rSquared,
			logloss,
		)
		if err != nil {
			return err
		}

		id, err = res.LastInsertId()
		if err != nil {
			return err
		}

		return nil
	})
	return err
}

func (ds *Datastore) CreateRegressionModel(pz az.Principal, modelId int64, mse, rSquared, deviance float64) error {
	var id int64
	err := ds.exec(func(tx *sql.Tx) error {
		res, err := tx.Exec(`
			INSERT INTO
				regression_model
				(model_id, mse, r_squared, mean_residual_deviance)
			VALUES
				($1,      $2,  $3,        $4)
			`,
			modelId,
			mse,
			rSquared,
			deviance,
		)
		if err != nil {
			return err
		}

		id, err = res.LastInsertId()
		if err != nil {
			return err
		}

		return nil
	})
	return err
}

// TODO: Deprecate
func (ds *Datastore) ReadModels(pz az.Principal, offset, limit int64) ([]Model, error) {
	rows, err := ds.db.Query(`
		SELECT
			model.*,
			label.id,
			label.name
		FROM
			model
		LEFT OUTER JOIN
			label ON label.model_id = model.id
		WHERE
			model.id IN
			(
				SELECT DISTINCT
					entity_id
				FROM 
					privilege
				WHERE
					$1 OR
					(
						workgroup_id IN 
						(
							SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $2
						) AND 
						entity_type_id = $3
					)
			)
		ORDER BY
			model.name
		LIMIT $4
		OFFSET $5
		`, pz.IsSuperuser(), pz.Id(), ds.EntityTypes.Model, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanModels(rows)
}

func (ds *Datastore) ReadModelsForProject(pz az.Principal, projectId, offset, limit int64) ([]Model, bool, error) {
	if err := pz.CheckView(ds.EntityTypes.Project, projectId); err != nil {
		return nil, false, err
	}

	rows, err := ds.db.Query(`
		SELECT
			model.*,
			label.id,
			label.name
		FROM
			model
		LEFT OUTER JOIN
			label on label.model_id = model.id
		WHERE
			model.project_id = $1 AND
			model.id IN
			(
				SELECT DISTINCT
					entity_id
				FROM 
					privilege
				WHERE
				  	$2 OR
					(
						workgroup_id IN 
						(
							SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $3
						) AND
						entity_type_id = $4
					)
			)
		ORDER BY
			model.name
		LIMIT $5
		OFFSET $6
		`, projectId, pz.IsSuperuser(), pz.Id(), ds.EntityTypes.Model, limit, offset)
	if err != nil {
		return nil, false, err
	}
	defer rows.Close()
	return scanModels(rows)
}

func (ds *Datastore) CountModelsForProject(pz az.Principal, projectId int64) (int64, error) {
	if err := pz.CheckView(ds.EntityTypes.Project, projectId); err != nil {
		return 0, err
	}

	var ct int64
	err := ds.exec(func(tx *sql.Tx) error {
		row := tx.QueryRow(`
			SELECT
				count(id)
			FROM
				model
			WHERE
				project_id = $1	
			`, projectId)

		return row.Scan(&ct)
	})

	return ct, err
}

func scanModels(rows *sql.Rows) ([]Model, bool, error) {
	models, err := ScanModels(rows)
	if err != nil {
		return nil, false, err
	}

	if len(models) == 0 {
		return nil, false, nil
	}

	return models, true, nil
}

func scanModel(rows *sql.Rows) (Model, bool, error) {
	models, err := ScanModels(rows)
	if err != nil {
		return Model{}, false, err
	}

	if len(models) == 0 {
		return Model{}, false, nil
	}

	return models[0], true, nil
}

func (ds *Datastore) ReadModelByDataset(pz az.Principal, datasetId int64) (Model, bool, error) {
	rows, err := ds.db.Query(`
		SELECT
			model.*,
			label.id,
			label.name
		FROM
			model
		LEFT OUTER JOIN
			label ON label.model_id = model.id
		WHERE
			model.training_dataset_id = $1
			OR
			model.validation_dataset_id = $1
		`, datasetId)
	if err != nil {
		return Model{}, false, err
	}
	defer rows.Close()

	return scanModel(rows)
}

func (ds *Datastore) ReadBinomialModels(pz az.Principal, projectId int64, namePart, sortBy string, ascending bool, offset, limit int64) ([]BinomialModel, error) {
	if err := pz.CheckView(ds.EntityTypes.Project, projectId); err != nil {
		return nil, err
	}

	dir := "ASC"
	if !ascending {
		dir = "DESC"
	}

	var filter string
	switch sortBy {
	case "mse", "r_squared", "logloss", "auc", "gini":
		filter = "bm." + sortBy + " " + dir
	default:
		filter = "model.name " + dir
	}

	rows, err := ds.db.Query(`
		SELECT
			model.*,
			label.id,
			label.name,
			bm.mse, bm.r_squared, bm.logloss, bm.auc, bm.gini
		FROM
			model
		INNER JOIN 
			binomial_model bm ON bm.model_id = model.id
		LEFT OUTER JOIN
			label ON label.model_id = model.id
		WHERE
			model.project_id = $1 AND
			model.id IN
			( 
				SELECT DISTINCT
					entity_id
				FROM
					privilege
				WHERE
					$2 OR
					(
						workgroup_id IN 
						(
							SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $3
							) AND 
							entity_type_id = $4
						)
			) AND
			model.name LIKE '%' || $5 || '%'
		ORDER BY
			`+filter+`
		LIMIT $6
		OFFSET $7
		`,
		projectId,
		pz.IsSuperuser(),
		pz.Id(),
		ds.EntityTypes.Model,
		namePart,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanBinomialModels(rows)
}

func (ds *Datastore) ReadBinomialModel(pz az.Principal, modelId int64) (BinomialModel, error) {
	if err := pz.CheckView(ds.EntityTypes.Model, modelId); err != nil {
		return BinomialModel{}, err
	}

	row := ds.db.QueryRow(`
		SELECT
			model.*, 
			label.id,
			label.name,
			bm.mse, bm.r_squared, bm.logloss, bm.auc, bm.gini
		FROM
			model
		INNER JOIN
			binomial_model bm ON bm.model_id = model.id
		LEFT OUTER JOIN
			label ON label.model_id = model.id
		WHERE
			model.id = $1
		`, modelId)
	return ScanBinomialModel(row)
}

func (ds *Datastore) ReadMultinomialModels(pz az.Principal, projectId int64, namePart, sortBy string, ascending bool, offset, limit int64) ([]MultinomialModel, error) {
	if err := pz.CheckView(ds.EntityTypes.Project, projectId); err != nil {
		return nil, err
	}

	dir := "ASC"
	if !ascending {
		dir = "DESC"
	}

	var filter string
	switch sortBy {
	case "mse", "r_squared", "logloss":
		filter = "mm." + sortBy + " " + dir
	default:
		filter = "model.name " + dir
	}

	rows, err := ds.db.Query(`
		SELECT
			model.*,
			label.id,
			label.name,
			mm.mse, mm.r_squared, mm.logloss
		FROM
			model
		INNER JOIN
			multinomial_model mm on mm.model_id = model.id
		LEFT OUTER JOIN
			label ON label.model_id = model.id
		WHERE
			model.project_id = $1 AND
			model.id IN
			( 
				SELECT DISTINCT
					entity_id
				FROM
					privilege
				WHERE
					$2 OR
					(
						workgroup_id IN 
						(
							SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $3
							) AND 
							entity_type_id = $4
						)
			) AND
			model.name LIKE '%' || $5 || '%'
		ORDER BY
			`+filter+`
		LIMIT $6
		OFFSET $7
		`,
		projectId,
		pz.IsSuperuser(),
		pz.Id(),
		ds.EntityTypes.Model,
		namePart,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanMultinomialModels(rows)
}

func (ds *Datastore) ReadMultinomialModel(pz az.Principal, modelId int64) (MultinomialModel, error) {
	if err := pz.CheckView(ds.EntityTypes.Model, modelId); err != nil {
		return MultinomialModel{}, err
	}

	row := ds.db.QueryRow(`
		SELECT
			model.*, 
			label.id,
			label.name,
			mm.mse, mm.r_squared, mm.logloss
		FROM
			model
		INNER JOIN 
			multinomial_model mm on mm.model_id = model.id
		LEFT OUTER JOIN
			label ON label.model_id = model.id
		WHERE
			model.id = $1
		`, modelId)
	return ScanMultinomialModel(row)
}

func (ds *Datastore) ReadRegressionModels(pz az.Principal, projectId int64, namePart, sortBy string, ascending bool, offset, limit int64) ([]RegressionModel, error) {
	if err := pz.CheckView(ds.EntityTypes.Project, projectId); err != nil {
		return nil, err
	}

	dir := "ASC"
	if !ascending {
		dir = "DESC"
	}

	var filter string
	switch sortBy {
	case "mse", "r_squared", "mean_residual_deviance":
		filter = "rm." + sortBy + " " + dir
	default:
		filter = "model.name " + dir
	}

	rows, err := ds.db.Query(`
		SELECT
			model.*,
			label.id,
			label.name,
			rm.mse, rm.r_squared, rm.mean_residual_deviance
		FROM
			model
		INNER JOIN 
			regression_model rm ON rm.model_id = model.id
		LEFT OUTER JOIN
			label ON label.model_id = model.id
		WHERE
			model.project_id = $1 AND
			model.id IN
			( 
				SELECT DISTINCT
					entity_id
				FROM
					privilege
				WHERE
					$2 OR
					(
						workgroup_id IN 
						(
							SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $3
							) AND 
							entity_type_id = $4
						)
			) AND
			model.name LIKE '%' || $5 || '%'
		ORDER BY
			`+filter+`
		LIMIT $6
		OFFSET $7
		`,
		projectId,
		pz.IsSuperuser(),
		pz.Id(),
		ds.EntityTypes.Model,
		namePart,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanRegressionModels(rows)
}

func (ds *Datastore) ReadRegressionModel(pz az.Principal, modelId int64) (RegressionModel, error) {
	if err := pz.CheckView(ds.EntityTypes.Model, modelId); err != nil {
		return RegressionModel{}, err
	}

	row := ds.db.QueryRow(`
		SELECT
			model.*, 
			label.id,
			label.name,
			rm.mse, rm.r_squared, rm.mean_residual_deviance
		FROM
			model
		INNER JOIN
			regression_model rm ON rm.model_id = model.id
		LEFT OUTER JOIN
			label ON label.model_id = model.id
		WHERE
			model.id = $1
		`, modelId)
	return ScanRegressionModel(row)
}

func (ds *Datastore) ReadModel(pz az.Principal, modelId int64) (Model, error) {
	if err := pz.CheckView(ds.EntityTypes.Model, modelId); err != nil {
		return Model{}, err
	}

	row := ds.db.QueryRow(`
		SELECT
			model.*,
			label.id,
			label.name
		FROM
			model
		LEFT OUTER JOIN
			label ON label.model_id = model.id
		WHERE
			model.id = $1
		`, modelId)
	return ScanModel(row)
}

func (ds *Datastore) UpdateModelLocation(pz az.Principal, modelId int64, location, logicalName string) error {
	if err := pz.CheckEdit(ds.EntityTypes.Model, modelId); err != nil {
		return err
	}

	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			UPDATE
				model
			SET
				location = $1,
				logical_name = $2
			WHERE
				id = $3
			`, location, logicalName, modelId); err != nil {
			return err
		}
		return ds.audit(pz, tx, UpdateOp, ds.EntityTypes.Model, modelId, metadata{
			"location": location,
		})
	})
}

func (ds *Datastore) UpdateModelObjectType(pz az.Principal, modelId int64, typ string) error {
	if err := pz.CheckEdit(ds.EntityTypes.Model, modelId); err != nil {
		return errors.Wrap(err, "failed checking edit privilege")
	}

	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			UPDATE
				model
			SET
				model_object_type = $1
			WHERE
				id = $2	
			`, typ, modelId); err != nil {
			return errors.Wrap(err, "failed executing transaction")
		}
		return ds.audit(pz, tx, UpdateOp, ds.EntityTypes.Model, modelId, metadata{
			"has_pojo": typ,
		})
	})
}

func (ds *Datastore) UpdateModelName(pz az.Principal, modelId int64, name string) error {
	if err := pz.CheckEdit(ds.EntityTypes.Model, modelId); err != nil {
		return err
	}
	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
 			UPDATE
 				model
 			SET
 				name = $1
 			WHERE
 				id = $2
 			`, name, modelId); err != nil {
			return err
		}
		return ds.audit(pz, tx, UpdateOp, ds.EntityTypes.Model, modelId, metadata{"name": name})
	})
}

func (ds *Datastore) DeleteModel(pz az.Principal, modelId int64) error {
	if err := pz.CheckOwns(ds.EntityTypes.Model, modelId); err != nil {
		return err
	}

	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			DELETE FROM
				model
			WHERE
				id = $1
			`, modelId); err != nil {
			return err
		}
		if err := deletePrivilegesOn(tx, ds.EntityTypes.Model, modelId); err != nil {
			return err
		}
		return ds.audit(pz, tx, DeleteOp, ds.EntityTypes.Model, modelId, metadata{})
	})
}

// --- Label ---

func (ds *Datastore) CreateLabel(pz az.Principal, projectId int64, name, description string) (int64, error) {
	if err := pz.CheckEdit(ds.EntityTypes.Project, projectId); err != nil {
		return 0, err
	}

	var id int64
	err := ds.exec(func(tx *sql.Tx) error {
		res, err := tx.Exec(`
			INSERT INTO
				label
				(project_id, name, description, created)
			VALUES
				($1,         $2,   $3,          datetime('now'))
			`,
			projectId,
			name,
			description,
		)
		if err != nil {
			return err
		}

		id, err = res.LastInsertId()
		if err != nil {
			return err
		}

		if err := createPrivilege(tx, Privilege{
			Owns,
			pz.WorkgroupId(),
			ds.EntityTypes.Label,
			id,
		}); err != nil {
			return err
		}

		return ds.audit(pz, tx, CreateOp, ds.EntityTypes.Label, id, metadata{
			"projectId":   strconv.FormatInt(projectId, 10),
			"name":        name,
			"description": description,
		})
	})
	return id, err
}

func (ds *Datastore) UpdateLabel(pz az.Principal, labelId int64, name, description string) error {
	if err := pz.CheckEdit(ds.EntityTypes.Label, labelId); err != nil {
		return err
	}

	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			UPDATE
				label
			SET
				name = $1,
				description = $2
			WHERE
				id = $3
			`, name, description, labelId); err != nil {
			return err
		}
		return ds.audit(pz, tx, UpdateOp, ds.EntityTypes.Label, labelId, metadata{"name": name, "description": description})
	})
}

func (ds *Datastore) DeleteLabel(pz az.Principal, labelId int64) error {
	if err := pz.CheckOwns(ds.EntityTypes.Label, labelId); err != nil {
		return err
	}

	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			DELETE FROM
				label
			WHERE
				id = $1
			`, labelId); err != nil {
			return err
		}

		if err := deletePrivilegesOn(tx, ds.EntityTypes.Label, labelId); err != nil {
			return err
		}
		return ds.audit(pz, tx, DeleteOp, ds.EntityTypes.Label, labelId, metadata{})
	})
}

func readModelName(tx *sql.Tx, modelId int64) (string, error) {
	row := tx.QueryRow("SELECT name FROM model WHERE id = $1", modelId)
	var name string
	err := row.Scan(&name)
	return name, err
}

func (ds *Datastore) LinkLabelWithModel(pz az.Principal, labelId, modelId int64) error {
	if err := pz.CheckEdit(ds.EntityTypes.Label, labelId); err != nil {
		return err
	}

	if err := pz.CheckView(ds.EntityTypes.Model, modelId); err != nil {
		return err
	}

	return ds.exec(func(tx *sql.Tx) error {
		modelName, err := readModelName(tx, modelId)
		if err != nil {
			return err
		}

		if _, err := tx.Exec(`
			UPDATE
				label
			SET
				model_id = $1
			WHERE
				id = $2
			`, modelId, labelId); err != nil {
			return err
		}
		return ds.audit(pz, tx, UpdateOp, ds.EntityTypes.Label, labelId, metadata{"modelName": modelName})
	})
}

func (ds *Datastore) UnlinkLabelFromModel(pz az.Principal, labelId, modelId int64) error {
	if err := pz.CheckEdit(ds.EntityTypes.Label, labelId); err != nil {
		return err
	}

	if err := pz.CheckView(ds.EntityTypes.Model, modelId); err != nil {
		return err
	}

	return ds.exec(func(tx *sql.Tx) error {
		modelName, err := readModelName(tx, modelId)
		if err != nil {
			return err
		}
		if _, err := tx.Exec(`
			UPDATE
				label
			SET
				model_id = null
			WHERE
				id = $1
			`, labelId); err != nil {
			return err
		}
		return ds.audit(pz, tx, UpdateOp, ds.EntityTypes.Label, labelId, metadata{"modelName": modelName})
	})
}

func (ds *Datastore) ReadLabelsForProject(pz az.Principal, projectId int64) ([]Label, error) {
	rows, err := ds.db.Query(`
		SELECT
			id, project_id, model_id, name, description, created
		FROM
			label
		WHERE
			project_id = $1 AND
			id IN
			(
				SELECT DISTINCT
					entity_id
				FROM 
					privilege
				WHERE
					$2 OR
					(
						workgroup_id IN 
						(
							SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $3
						) AND 
						entity_type_id = $4
					)
			)
		ORDER BY
			name
		`, projectId, pz.IsSuperuser(), pz.Id(), ds.EntityTypes.Label)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanLabels(rows)
}

func scanLabel(rows *sql.Rows) (Label, bool, error) {
	var label Label

	labels, err := ScanLabels(rows)
	if err != nil {
		return label, false, err
	}

	if len(labels) == 0 {
		return label, false, nil
	}

	return labels[0], true, nil
}

func (ds *Datastore) ReadLabelByModel(pz az.Principal, modelId int64) (Label, bool, error) {
	rows, err := ds.db.Query(`
		SELECT
			id, project_id, model_id, name, description, created
		FROM
			label
		WHERE
			model_id = $1
		`, modelId)
	if err != nil {
		return Label{}, false, err
	}
	defer rows.Close()

	return scanLabel(rows)
}

func (ds *Datastore) ReadLabel(pz az.Principal, labelId int64) (Label, error) {
	rows, err := ds.db.Query(`
		SELECT
			id, project_id, model_id, name, description, created
		FROM
			label
		WHERE
			id = $1 AND
			id IN
			(
				SELECT DISTINCT
					entity_id
				FROM 
					privilege
				WHERE
					$2 OR
					(
						workgroup_id IN 
						(
							SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $3
						) AND
						entity_type_id = $4
					)
			)
		ORDER BY
			name
		`, labelId, pz.IsSuperuser(), pz.Id(), ds.EntityTypes.Label)
	if err != nil {
		return Label{}, err
	}
	defer rows.Close()

	labels, err := ScanLabels(rows)
	if err != nil {
		return Label{}, err
	}
	if len(labels) > 0 {
		return labels[0], nil
	}
	return Label{}, fmt.Errorf("Label %d not found", labelId)
}

// --- Service ---

func (ds *Datastore) CreateService(pz az.Principal, service Service) (int64, error) {
	var id int64
	err := ds.exec(func(tx *sql.Tx) error {
		res, err := tx.Exec(`
			INSERT INTO
				service
				(project_id, model_id, name, address, port, process_id, state, created)
			VALUES
				($1,       $2,        $3,   $4,      $5,   $6,         $7,   datetime('now'))
			`,
			service.ProjectId,
			service.ModelId,
			service.Name,
			service.Address,
			service.Port,
			service.ProcessId,
			service.State,
		)
		if err != nil {
			return err
		}

		id, err = res.LastInsertId()
		if err != nil {
			return err
		}

		if err := createPrivilege(tx, Privilege{
			Owns,
			pz.WorkgroupId(),
			ds.EntityTypes.Service,
			id,
		}); err != nil {
			return err
		}

		return ds.audit(pz, tx, CreateOp, ds.EntityTypes.Service, id, metadata{
			"modelId":   strconv.FormatInt(service.ModelId, 10),
			"address":   service.Address,
			"port":      strconv.FormatInt(service.Port, 10),
			"processId": strconv.FormatInt(service.ProcessId, 10),
			"state":     service.State,
		})
	})
	return id, err
}

func (ds *Datastore) ReadServices(pz az.Principal, offset, limit int64) ([]Service, error) {
	rows, err := ds.db.Query(`
		SELECT
			id, project_id, model_id, name, address, port, process_id, state, created
		FROM
			service
		WHERE
			id IN
			(
				SELECT DISTINCT
					entity_id
				FROM 
					privilege
				WHERE
					$1 OR
					(
						workgroup_id IN 
						(
							SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $2
							) AND 
							entity_type_id = $3
						)
			)
		ORDER BY
			address, port
		LIMIT $4
		OFFSET $5
		`, pz.IsSuperuser(), pz.Id(), ds.EntityTypes.Service, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanServices(rows)
}

func (ds *Datastore) ReadServicesForProjectId(pz az.Principal, projectId, offset, limit int64) ([]Service, error) {
	rows, err := ds.db.Query(`
		SELECT
			id, project_id, model_id, name, address, port, process_id, state, created
		FROM
			service
		WHERE
			project_id = $1 AND
			id IN
			(
				SELECT DISTINCT
					entity_id
				FROM 
					privilege
				WHERE
					$2 OR
					(
						workgroup_id IN 
						(
							SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $3
						) AND
						entity_type_id = $4
					)
			)
		ORDER BY
			address, port
		LIMIT $5
		OFFSET $6
		`, projectId, pz.IsSuperuser(), pz.Id(), ds.EntityTypes.Service, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanServices(rows)
}

func (ds *Datastore) ReadServicesForModelId(pz az.Principal, modelId int64) ([]Service, error) {
	if err := pz.CheckView(ds.EntityTypes.Model, modelId); err != nil {
		return nil, err
	}

	rows, err := ds.db.Query(`
		SELECT
			id, project_id, model_id, name, address, port, process_id, state, created
		FROM
			service
		WHERE
			model_id = $1
		ORDER BY
			address, port
		`, modelId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanServices(rows)
}

func (ds *Datastore) ReadService(pz az.Principal, serviceId int64) (Service, error) {
	if err := pz.CheckView(ds.EntityTypes.Service, serviceId); err != nil {
		return Service{}, err
	}

	row := ds.db.QueryRow(`
		SELECT
			id, project_id, model_id, name, address, port, process_id, state, created
		FROM
			service
		WHERE
			id = $1
		`, serviceId)
	return ScanService(row)
}

func (ds *Datastore) UpdateServiceState(pz az.Principal, serviceId int64, state string) error {
	if err := pz.CheckEdit(ds.EntityTypes.Service, serviceId); err != nil {
		return err
	}

	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			UPDATE
				service
			SET
				state = $1
			WHERE
				id = $2
			`, state, serviceId); err != nil {
			return err
		}
		return ds.audit(pz, tx, UpdateOp, ds.EntityTypes.Service, serviceId, metadata{"state": state})
	})
}

func (ds *Datastore) UpdateServiceName(pz az.Principal, serviceId int64, name string) error {
	if err := pz.CheckEdit(ds.EntityTypes.Service, serviceId); err != nil {
		return err
	}

	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
 			UPDATE
 				service
 			SET
 				name = $1
 			WHERE
 				id = $2
 			`, name, serviceId); err != nil {
			return err
		}
		return ds.audit(pz, tx, UpdateOp, ds.EntityTypes.Service, serviceId, metadata{"name": name})
	})
}

func (ds *Datastore) DeleteService(pz az.Principal, serviceId int64) error {
	if err := pz.CheckOwns(ds.EntityTypes.Service, serviceId); err != nil {
		return err
	}

	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			DELETE FROM
				service
			WHERE
				id = $1
			`, serviceId); err != nil {
			return err
		}

		if err := deletePrivilegesOn(tx, ds.EntityTypes.Service, serviceId); err != nil {
			return err
		}
		return ds.audit(pz, tx, DeleteOp, ds.EntityTypes.Service, serviceId, metadata{})
	})
}
