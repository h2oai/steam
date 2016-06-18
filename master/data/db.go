package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/h2oai/steamY/master/az"
	"github.com/lib/pq"
	"log"
	"strconv"
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
//
// Project
//   Read               x    x    x
//   Assign Model       x    x
//   Update             x    x
//   Delete             x
//   Share              x
//
// Engine, Model
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
	Version = "1"

	CanView = "view"
	CanEdit = "edit"
	Owns    = "own"

	ForIdentity  = "identity"
	ForWorkgroup = "workgroup"

	RoleEntity      = "role"
	WorkgroupEntity = "workgroup"
	IdentityEntity  = "identity"
	EngineEntity    = "engine"
	ClusterEntity   = "cluster"
	ProjectEntity   = "project"
	ModelEntity     = "model"
	ServiceEntity   = "service"

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

const ( // FIXME change to int64 in postgres
	ManageRole int64 = 1 + iota
	ViewRole
	ManageWorkgroup
	ViewWorkgroup
	ManageIdentity
	ViewIdentity
	ManageEngine
	ViewEngine
	ManageCluster
	ViewCluster
	ManageProject
	ViewProject
	ManageModel
	ViewModel
	ManageService
	ViewService
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
		{0, ManageModel, "Manage models"},
		{0, ViewModel, "View models"},
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
		{0, ModelEntity},
		{0, ServiceEntity},
	}

	ClusterTypes = []ClusterType{
		{0, ClusterExternal},
		{0, ClusterYarn},
	}
}

type metadata map[string]string

type EntityTypeKeys struct {
	Role      int64
	Workgroup int64
	Identity  int64
	Engine    int64
	Cluster   int64
	Project   int64
	Model     int64
	Service   int64
}

type ClusterTypeKeys struct {
	External int64
	Yarn     int64
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
		m[ModelEntity],
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
	db             *sql.DB // Singleton; doesn't actually connect until used, and is pooled internally.
	metadata       metadata
	permissions    []Permission
	permissionMap  map[int64]Permission
	entityTypeMap  map[int64]EntityType
	EntityTypes    *EntityTypeKeys
	clusterTypeMap map[int64]ClusterType
	ClusterTypes   *ClusterTypeKeys
}

func Connect(username, dbname, sslmode string) (*sql.DB, error) {

	log.Println("Connecting to database: user =", username, "db =", dbname, "SSL=", sslmode, "...")

	// Open connection
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=%s", username, dbname, sslmode))
	if err != nil {
		return nil, fmt.Errorf("Database connection failed: %s", err)
	}

	// TODO can use db.SetMaxOpenConns() and db.SetMaxIdleConns() to configure further.

	// Verify connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Database ping failed: %s", err)
	}

	return db, nil
}

// NewDB creates a new instance of a data access object.
//
// Valid values for sslmode are:
//   disable - No SSL
//   require - Always SSL (skip verification)
//   verify-ca - Always SSL (verify that the certificate presented by the server was signed by a trusted CA)
//   verify-full - Always SSL (verify that the certification presented by the server was signed by a
//     trusted CA and the server host name matches the one in the certificate)
func NewDatastore(db *sql.DB) (*Datastore, error) {

	// Read meta information

	metadata, err := readMetadata(db)
	if err != nil {
		return nil, err
	}

	version, ok := metadata["version"]
	if !ok {
		return nil, fmt.Errorf("Failed reading schema version")
	}

	log.Println("Using schema version:", version)

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

	entityTypes, err := readEntityTypes(db)
	if err != nil {
		return nil, err
	}

	entityTypeMap := make(map[int64]EntityType)
	for _, et := range entityTypes {
		entityTypeMap[et.Id] = et
	}

	clusterTypes, err := readClusterTypes(db)
	if err != nil {
		return nil, err
	}

	clusterTypeMap := make(map[int64]ClusterType)
	for _, ct := range clusterTypes {
		clusterTypeMap[ct.Id] = ct
	}

	return &Datastore{
		db,
		metadata,
		permissions,
		permissionMap,
		entityTypeMap,
		toEntityTypeKeys(entityTypes),
		clusterTypeMap,
		toClusterTypeKeys(clusterTypes),
	}, nil
}

func IsPrimed(db *sql.DB) (bool, error) {
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

func Prime(db *sql.DB) error {
	log.Println("Priming database for first time use...")
	if err := createMetadata(db, "version", "1"); err != nil {
		return err
	}
	if err := primePermissions(db, Permissions); err != nil {
		return err
	}
	if err := primeEntityTypes(db, EntityTypes); err != nil {
		return err
	}
	if err := primeClusterTypes(db, ClusterTypes); err != nil {
		return err
	}

	return nil
}

func bulkInsert(db *sql.DB, table string, columns []string, f func(*sql.Stmt) error) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(pq.CopyIn(table, columns...))
	if err != nil {
		return err
	}

	if err := f(stmt); err != nil { // buffer
		return err
	}

	if _, err := stmt.Exec(); err != nil { // flush
		return err
	}

	if err := stmt.Close(); err != nil {
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
				return err
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
				return err
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
				return err
			}
		}
		return nil
	})
}

func upgrade(db *sql.DB, currentVersion string) error {
	if currentVersion == Version {
		return nil
	}

	// TODO add stepwise upgrades

	return nil
}

func truncate(db *sql.DB) error {
	log.Println("Truncating database...")
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
			"project_model",
			"model",
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
			return err
		}

		id, err = createIdentity(tx, name, password, workgroupId)
		if err != nil {
			return err
		}

		if err := linkIdentityAndWorkgroup(tx, id, workgroupId); err != nil {
			return err
		}

		return createPrivilege(tx, Privilege{
			Owns,
			workgroupId,
			ds.EntityTypes.Identity,
			id,
		})
	})
	return id, workgroupId, err
}

func (ds *Datastore) CreateSuperuserRole(pz az.Principal) error {

	roleId, err := ds.CreateRole(pz, "Superuser", "Superuser")
	if err != nil {
		return err
	}

	allPerms := make([]int64, len(ds.permissions))
	for i, permission := range ds.permissions {
		allPerms[i] = permission.Id
	}

	if err := ds.SetRolePermissions(pz, roleId, allPerms); err != nil {
		return err
	}

	if err := ds.LinkIdentityAndRole(pz, pz.Id(), roleId); err != nil {
		return err
	}

	return nil
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

func executeTransaction(db *sql.DB, f func(*sql.Tx) error) (err error) {
	var (
		tx     *sql.Tx
		commit bool
	)

	tx, err = db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if commit {
			err = tx.Commit()
		} else {
			if rberr := tx.Rollback(); rberr != nil {
				err = fmt.Errorf("Rollback failure: %s (after %s)", rberr, err)
			}
		}
	}()
	err = f(tx)
	if err == nil {
		commit = true
	}
	return err
}

func (ds *Datastore) exec(f func(*sql.Tx) error) (err error) {
	return executeTransaction(ds.db, f)
}

func (ds *Datastore) toPermissionDescriptions(ids []int64) ([]string, error) {
	descriptions := make([]string, len(ids))
	for i, id := range ids {
		if p, ok := ds.permissionMap[id]; ok {
			descriptions[i] = p.Description
		} else {
			return descriptions, fmt.Errorf("Invalid permission id: %d", id)
		}
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
			($1,          $2,     $3,             $4,        $5,          now())
		`, pz.Id(), action, entityTypeId, entityId, string(json)); err != nil {
		return err
	}
	return nil
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

func (ds *Datastore) ReadPermissionsForIdentity(pz az.Principal, identityId int64) ([]int64, error) {
	if err := pz.CheckView(ds.EntityTypes.Identity, identityId); err != nil {
		return nil, err
	}

	return ds.readPermissionsForIdentity(identityId)
}

// --- Roles ---

func (ds *Datastore) CreateRole(pz az.Principal, name, description string) (int64, error) {
	var id int64
	err := ds.exec(func(tx *sql.Tx) error {
		row := tx.QueryRow(`
			INSERT INTO
				role
				(name, description, created)
			VALUES
				($1,   $2,          now())
			RETURNING id
			`, name, description)
		if err := row.Scan(&id); err != nil {
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
					workgroup_id IN (SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $3) AND
					entity_type_id = $4
			)
		ORDER BY 
			name
		OFFSET $1
		LIMIT $2
		`, offset, limit, pz.Id(), ds.EntityTypes.Role)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanRoles(rows)
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
					workgroup_id IN (SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $2) AND
					entity_type_id = $3
			)
		ORDER BY
			r.name
		`, identityId, pz.Id(), ds.EntityTypes.Role)

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

func (ds *Datastore) UpdateRole(pz az.Principal, roleId int64, name string) error {
	if err := pz.CheckEdit(ds.EntityTypes.Role, roleId); err != nil {
		return err
	}

	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			UPDATE
				role
			SET
				name = $1
			WHERE
				id = $2
			`, name, roleId); err != nil {
			return err
		}
		return ds.audit(pz, tx, UpdateOp, ds.EntityTypes.Role, roleId, metadata{"name": name})
	})
}

func (ds *Datastore) SetRolePermissions(pz az.Principal, roleId int64, permissionIds []int64) error {
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
			`, roleId,
		); err != nil {
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
		row := tx.QueryRow(`
			INSERT INTO
				workgroup
				(type,          name, description, created)
			VALUES
				('workgroup',   $1,   $2,          now())
			RETURNING id
			`, name, description)
		if err := row.Scan(&id); err != nil {
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
					workgroup_id IN (SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $3) AND
					entity_type_id = $4
			)
		ORDER BY name
		OFFSET $1
		LIMIT $2
		`, offset, limit, pz.Id(), ds.EntityTypes.Workgroup)
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
					workgroup_id IN (SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $2) AND
					entity_type_id = $3
			)
		ORDER BY
			w.name
		`, identityId, pz.Id(), ds.EntityTypes.Workgroup)

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

func (ds *Datastore) UpdateWorkgroup(pz az.Principal, workgroupId int64, name string) error {
	if err := pz.CheckEdit(ds.EntityTypes.Workgroup, workgroupId); err != nil {
		return err
	}
	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			UPDATE
				workgroup
			SET
				name = $1
			WHERE
				id = $2
			`, name, workgroupId); err != nil {
			return err
		}
		return ds.audit(pz, tx, UpdateOp, ds.EntityTypes.Workgroup, workgroupId, metadata{"name": name})
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
	var id int64
	row := tx.QueryRow(`
			INSERT INTO
				workgroup
				(type,       name, description, created)
			VALUES
				('identity', $1,   '',          now())
			RETURNING id
			`, "user:"+name)
	return id, row.Scan(&id)
}

func createIdentity(tx *sql.Tx, name, password string, workgroupId int64) (int64, error) {
	var id int64
	row := tx.QueryRow(`
			INSERT INTO
				identity
				(name, password, workgroup_id, is_active, created)
			VALUES
				($1,   $2,       $3,           $4,        now())
			RETURNING id
			`, name, password, workgroupId, true)

	return id, row.Scan(&id)

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
					workgroup_id IN (SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $3) AND
					entity_type_id = $4
			)
		ORDER BY name
		OFFSET $1
		LIMIT $2
		`, offset, limit, pz.Id(), ds.EntityTypes.Identity)
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
					workgroup_id IN (SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $2) AND
					entity_type_id = $3
			)
		ORDER BY
			i.name
		`, workgroupId, pz.Id(), ds.EntityTypes.Identity)

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
					workgroup_id IN (SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $2) AND
					entity_type_id = $3
			)
		ORDER BY
			i.name
		`, roleId, pz.Id(), ds.EntityTypes.Identity)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanIdentitys(rows)
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
		if _, err := tx.Exec(`
			INSERT INTO
				identity_role
			VALUES
				($1, $2)
			`, identityId, roleId); err != nil {
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
				is_active = false
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

func (ds *Datastore) ReadPrivileges(pz az.Principal, identityId, entityTypeId, entityId int64) ([]string, error) {
	if err := pz.CheckView(ds.EntityTypes.Identity, identityId); err != nil {
		return nil, err
	}

	if err := pz.CheckView(entityTypeId, entityId); err != nil {
		return nil, err
	}

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
	return err
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
		row := tx.QueryRow(`
			INSERT INTO
				engine
				(name, location, created)
			VALUES
				($1,   $2,       now())
			RETURNING id
			`, name, location)
		if err := row.Scan(&id); err != nil {
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
					workgroup_id IN (SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $1) AND
					entity_type_id = $2
			)
		ORDER BY
			name
		`, pz.Id(), ds.EntityTypes.Engine)
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
			`, engineId,
		); err != nil {
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
		row := tx.QueryRow(`
			INSERT INTO
				cluster
				(name, type_id, detail_id, address, state, created)
			VALUES
				($1,   $2,      0,         $3,      $4,    now())
			RETURNING id
			`, name, ds.ClusterTypes.External, address, state)
		if err := row.Scan(&id); err != nil {
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

func (ds *Datastore) CreateYarnCluster(pz az.Principal, name, address, state string, cluster YarnCluster) (int64, error) {
	var clusterId int64
	err := ds.exec(func(tx *sql.Tx) error {
		var yarnClusterId int64
		row := tx.QueryRow(`
			INSERT INTO
				cluster_yarn
				(engine_id, size, application_id, memory, username, output_dir)
			VALUES
				($1,        $2,   $3,             $4,     $5,       $6)
			RETURNING id
			`,
			cluster.EngineId,
			cluster.Size,
			cluster.ApplicationId,
			cluster.Memory,
			cluster.Username,
			cluster.OutputDir,
		)
		if err := row.Scan(&yarnClusterId); err != nil {
			return err
		}

		row = tx.QueryRow(`
			INSERT INTO
				cluster
				(name, type_id, detail_id, address, state, created)
			VALUES
				($1,   $2,      $3,        $4,      $5,    now())
			RETURNING id
			`, name, ds.ClusterTypes.Yarn, yarnClusterId, address, state)
		if err := row.Scan(&clusterId); err != nil {
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

func (ds *Datastore) ReadClusters(pz az.Principal, offset, limit int64) ([]Cluster, error) {
	rows, err := ds.db.Query(`
		SELECT
			id, name, type_id, detail_id, address, state, created
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
					workgroup_id IN (SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $1) AND
					entity_type_id = $2
			)
		ORDER BY
			name
		OFFSET $3
		LIMIT $4
		`, pz.Id(), ds.EntityTypes.Cluster, offset, limit)
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
			id, name, type_id, detail_id, address, state, created
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
			id, name, type_id, detail_id, address, state, created
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
			id, name, type_id, detail_id, address, state, created
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
				state = $2
			WHERE
				id = $1
			`, clusterId, state); err != nil {
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
				`, cluster.DetailId,
			); err != nil {
				return err
			}
		}

		if _, err := tx.Exec(`
			DELETE FROM
				cluster
			WHERE
				id = $1
			`, clusterId,
		); err != nil {
			return err
		}

		if err := deletePrivilegesOn(tx, ds.EntityTypes.Cluster, clusterId); err != nil {
			return err
		}
		return ds.audit(pz, tx, DeleteOp, ds.EntityTypes.Cluster, clusterId, metadata{})
	})
}

// --- Project ---

func (ds *Datastore) CreateProject(pz az.Principal, name, description string) (int64, error) {
	var id int64
	err := ds.exec(func(tx *sql.Tx) error {
		row := tx.QueryRow(`
			INSERT INTO
				project
				(name, description, created)
			VALUES
				($1,   $2,       now())
			RETURNING id
			`, name, description)
		if err := row.Scan(&id); err != nil {
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
			"name":        name,
			"description": description,
		})
	})
	return id, err
}

func (ds *Datastore) LinkProjectAndModel(pz az.Principal, projectId, modelId int64) error {
	if err := pz.CheckView(ds.EntityTypes.Model, modelId); err != nil {
		return err
	}
	if err := pz.CheckEdit(ds.EntityTypes.Project, projectId); err != nil {
		return err
	}

	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			INSERT INTO
				project_model
			VALUES
				($1, $2)
			`, projectId, modelId); err != nil {
			return err
		}
		return ds.audit(pz, tx, LinkOp, ds.EntityTypes.Project, projectId, metadata{
			"id": strconv.FormatInt(modelId, 10),
		})
	})
}

func (ds *Datastore) UnlinkProjectAndModel(pz az.Principal, projectId, modelId int64) error {
	if err := pz.CheckView(ds.EntityTypes.Model, modelId); err != nil {
		return err
	}
	if err := pz.CheckEdit(ds.EntityTypes.Project, projectId); err != nil {
		return err
	}

	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			DELETE FROM
				project_model
			WHERE
				project_id = $1 AND
				model_id = $2
			`, projectId, modelId); err != nil {
			return err
		}
		return ds.audit(pz, tx, UnlinkOp, ds.EntityTypes.Project, projectId, metadata{
			"id": strconv.FormatInt(modelId, 10),
		})
	})
}

func (ds *Datastore) ReadProjects(pz az.Principal, offset, limit int64) ([]Project, error) {
	rows, err := ds.db.Query(`
		SELECT
			id, name, description, created
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
					workgroup_id IN (SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $1) AND
					entity_type_id = $2
			)
		ORDER BY
			name
		OFFSET $3
		LIMIT $4
		`, pz.Id(), ds.EntityTypes.Project, offset, limit)
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
			id, name, description, created
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
			`, projectId,
		); err != nil {
			return err
		}
		if err := deletePrivilegesOn(tx, ds.EntityTypes.Project, projectId); err != nil {
			return err
		}
		return ds.audit(pz, tx, DeleteOp, ds.EntityTypes.Project, projectId, metadata{})
	})
}

// --- Model ---

func (ds *Datastore) CreateModel(pz az.Principal, model Model) (int64, error) {
	var id int64
	err := ds.exec(func(tx *sql.Tx) error {
		row := tx.QueryRow(`
			INSERT INTO
				model
				(name, cluster_name, algorithm, dataset_name, response_column_name, logical_name, location, max_run_time, created)
			VALUES
				($1,   $2,           $3,        $4,           $5,                   $6,           $7,       $8,           now())
			RETURNING id
			`,
			model.Name,
			model.ClusterName,
			model.Algorithm,
			model.DatasetName,
			model.ResponseColumnName,
			model.LogicalName,
			model.Location,
			model.MaxRunTime,
		)
		if err := row.Scan(&id); err != nil {
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
			"algorithm":          model.Algorithm,
			"datasetName":        model.DatasetName,
			"responseColumnName": model.ResponseColumnName,
			"logicalName":        model.LogicalName,
			"location":           model.Location,
			"maxRunTime":         strconv.FormatInt(model.MaxRunTime, 10),
		})
	})
	return id, err
}

func (ds *Datastore) ReadModels(pz az.Principal, offset, limit int64) ([]Model, error) {
	rows, err := ds.db.Query(`
		SELECT
			id, name, cluster_name, algorithm, dataset_name, response_column_name, logical_name, location, max_run_time, created
		FROM
			model
		WHERE
			id IN
			(
				SELECT DISTINCT
					entity_id
				FROM 
					privilege
				WHERE
					workgroup_id IN (SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $1) AND
					entity_type_id = $2
			)
		ORDER BY
			name
		OFFSET $3
		LIMIT $4
		`, pz.Id(), ds.EntityTypes.Model, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanModels(rows)
}

func (ds *Datastore) ReadModelsForProject(pz az.Principal, projectId, offset, limit int64) ([]Model, error) {

	if err := pz.CheckView(ds.EntityTypes.Project, projectId); err != nil {
		return nil, err
	}

	rows, err := ds.db.Query(`
		SELECT
			m.id, m.name, m.cluster_name, m.algorithm, m.dataset_name, m.response_column_name, m.logical_name, m.location, m.max_run_time, m.created
		FROM
			model m,
			project_model pm
		WHERE
			pm.project_id = $1 AND
			pm.model_id = m.id AND
			m.id IN
			(
				SELECT DISTINCT
					entity_id
				FROM 
					privilege
				WHERE
					workgroup_id IN (SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $2) AND
					entity_type_id = $3
			)
		ORDER BY
			m.name
		OFFSET $4
		LIMIT $5
		`, projectId, pz.Id(), ds.EntityTypes.Model, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanModels(rows)
}

func (ds *Datastore) ReadModel(pz az.Principal, modelId int64) (Model, error) {
	if err := pz.CheckView(ds.EntityTypes.Model, modelId); err != nil {
		return Model{}, err
	}

	row := ds.db.QueryRow(`
		SELECT
			id, name, cluster_name, algorithm, dataset_name, response_column_name, logical_name, location, max_run_time, created
		FROM
			model
		WHERE
			id = $1
		`, modelId)
	return ScanModel(row)
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
			`, modelId,
		); err != nil {
			return err
		}
		if err := deletePrivilegesOn(tx, ds.EntityTypes.Model, modelId); err != nil {
			return err
		}
		return ds.audit(pz, tx, DeleteOp, ds.EntityTypes.Model, modelId, metadata{})
	})
}

// --- Service ---

func (ds *Datastore) CreateService(pz az.Principal, service Service) (int64, error) {
	var id int64
	err := ds.exec(func(tx *sql.Tx) error {
		row := tx.QueryRow(`
			INSERT INTO
				service
				(model_id, address, port, process_id, state, created)
			VALUES
				($1,       $2,      $3,   $4,         $5,    now())
			RETURNING id
			`,
			service.ModelId,
			service.Address,
			service.Port,
			service.ProcessId,
			service.State,
		)
		if err := row.Scan(&id); err != nil {
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
			id, model_id, address, port, process_id, state, created
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
					workgroup_id IN (SELECT workgroup_id FROM identity_workgroup WHERE identity_id = $1) AND
					entity_type_id = $2
			)
		ORDER BY
			address, port
		OFFSET $3
		LIMIT $4
		`, pz.Id(), ds.EntityTypes.Service, offset, limit)
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
			id, model_id, address, port, process_id, state, created
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
			id, model_id, address, port, process_id, state, created
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
			`, serviceId,
		); err != nil {
			return err
		}
		if err := deletePrivilegesOn(tx, ds.EntityTypes.Service, serviceId); err != nil {
			return err
		}
		return ds.audit(pz, tx, DeleteOp, ds.EntityTypes.Service, serviceId, metadata{})
	})
}

// --- Datastore-backed Principal Impl ---

func (ds *Datastore) NewPrincipal(name string) (az.Principal, error) {
	identity, err := ds.readIdentityAndPassword(name)
	if err != nil {
		return nil, err
	}

	if identity == nil {
		return nil, nil
	}

	permissionIds, err := ds.readPermissionsForIdentity(identity.Id)
	if err != nil {
		return nil, err
	}

	permissionMap := ds.permissionMap
	permissions := make(map[int64]bool)
	for _, permissionId := range permissionIds {
		permissions[permissionMap[permissionId].Code] = true
	}

	return &Principal{ds, identity, permissions}, nil
}

type Principal struct {
	ds          *Datastore
	identity    *IdentityAndPassword
	permissions map[int64]bool
}

func (pz *Principal) Id() int64 {
	return pz.identity.Id
}

func (pz *Principal) WorkgroupId() int64 {
	return pz.identity.WorkgroupId
}

func (pz *Principal) Name() string {
	return pz.identity.Name
}

func (pz *Principal) Password() string {
	return pz.identity.Password
}

func (pz *Principal) IsActive() bool {
	return pz.identity.IsActive
}

func (pz *Principal) HasPermission(code int64) bool {
	_, ok := pz.permissions[code]
	return ok
}

func (pz *Principal) CheckPermission(code int64) error {
	if pz.HasPermission(code) {
		return nil
	}
	// FIXME return string representation of permission code
	return fmt.Errorf("Identity %s does not have permission %d to perform this operation", pz.Name(), code)
}

// TODO use bitwise ops to simplify this
func (pz *Principal) hasPrivilege(entityTypeId, entityId int64, expectedPrivilege string) (bool, error) {

	owns := false
	canEdit := false
	canView := false

	privileges, err := pz.ds.readPrivileges(pz.identity.Id, entityTypeId, entityId)
	if err != nil {
		return false, err
	}

	if len(privileges) == 0 {
		return false, nil
	}

	for _, p := range privileges {
		switch p {
		case Owns:
			owns = true
			canEdit = true
			canView = true
		case CanEdit:
			canEdit = true
			canView = true
		case CanView:
			canView = true
		}

		switch expectedPrivilege {
		case Owns:
			if owns {
				return true, nil
			}
		case CanEdit:
			if owns || canEdit {
				return true, nil
			}
		case CanView:
			if owns || canEdit || canView {
				return true, nil
			}
		}
	}
	return false, nil
}

func (pz *Principal) checkPrivilege(entityTypeId, entityId int64, expectedPrivilege string) error {
	ok, err := pz.hasPrivilege(entityTypeId, entityId, expectedPrivilege)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("Identity %s does not have privilege '%s' on the entity", pz.Name(), expectedPrivilege)
	}
	return nil
}

func (pz *Principal) Owns(entityTypeId, entityId int64) (bool, error) {
	return pz.hasPrivilege(entityId, entityId, Owns)
}

func (pz *Principal) CanEdit(entityTypeId, entityId int64) (bool, error) {
	return pz.hasPrivilege(entityId, entityId, CanEdit)
}

func (pz *Principal) CanView(entityTypeId, entityId int64) (bool, error) {
	return pz.hasPrivilege(entityId, entityId, CanView)
}

func (pz *Principal) CheckOwns(entityTypeId, entityId int64) error {
	return pz.checkPrivilege(entityId, entityId, Owns)
}

func (pz *Principal) CheckEdit(entityTypeId, entityId int64) error {
	return pz.checkPrivilege(entityId, entityId, CanEdit)
}

func (pz *Principal) CheckView(entityTypeId, entityId int64) error {
	return pz.checkPrivilege(entityId, entityId, CanView)
}

const SystemIdentityName = "system"

func CreateSystemIdentity(db *sql.DB) (int64, int64, error) {
	var id, workgroupId int64
	err := executeTransaction(db, func(tx *sql.Tx) error {
		var err error

		workgroupId, err = createDefaultWorkgroup(tx, SystemIdentityName)
		if err != nil {
			return err
		}

		id, err = createIdentity(tx, SystemIdentityName, "", workgroupId)
		if err != nil {
			return err
		}

		return linkIdentityAndWorkgroup(tx, id, workgroupId)
	})
	return id, workgroupId, err
}

type SystemPrincipal struct {
	identity *IdentityAndPassword
}

func (ds *Datastore) NewSystemPrincipal() (az.Principal, error) {
	identity, err := ds.readIdentityAndPassword(SystemIdentityName)
	if err != nil {
		return nil, err
	}

	if identity == nil {
		return nil, nil
	}

	return &SystemPrincipal{identity}, nil
}

func (pz *SystemPrincipal) Id() int64 {
	return pz.identity.Id
}

func (pz *SystemPrincipal) WorkgroupId() int64 {
	return pz.identity.WorkgroupId
}

func (pz *SystemPrincipal) Name() string {
	return pz.identity.Name
}

func (pz *SystemPrincipal) Password() string {
	return pz.identity.Password
}

func (pz *SystemPrincipal) IsActive() bool {
	return true
}

func (pz *SystemPrincipal) HasPermission(code int64) bool {
	return true
}

func (pz *SystemPrincipal) CheckPermission(code int64) error {
	return nil
}

func (pz *SystemPrincipal) Owns(entityTypeId, entityId int64) (bool, error) {
	return true, nil
}

func (pz *SystemPrincipal) CanEdit(entityTypeId, entityId int64) (bool, error) {
	return true, nil
}

func (pz *SystemPrincipal) CanView(entityTypeId, entityId int64) (bool, error) {
	return true, nil
}

func (pz *SystemPrincipal) CheckOwns(entityTypeId, entityId int64) error {
	return nil
}

func (pz *SystemPrincipal) CheckEdit(entityTypeId, entityId int64) error {
	return nil
}

func (pz *SystemPrincipal) CheckView(entityTypeId, entityId int64) error {
	return nil
}
