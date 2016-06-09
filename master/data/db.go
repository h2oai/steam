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

const (
	Version = "1"

	OwnPrivilege  = "own"
	EditPrivilege = "edit"
	ViewPrivilege = "view"

	IdentityPrivilege  = "identity"
	WorkgroupPrivilege = "workgroup"

	RoleLabel      = "role"
	WorkgroupLabel = "workgroup"
	IdentityLabel  = "identity"
	EngineLabel    = "engine"
	ClusterLabel   = "cluster"
	ModelLabel     = "model"
)

var (
	Permissions []Permission
	EntityTypes []EntityType
)

func init() {
	EntityTypes = []EntityType{
		{0, RoleLabel},
		{0, WorkgroupLabel},
		{0, IdentityLabel},
		{0, EngineLabel},
		{0, ClusterLabel},
		{0, ModelLabel},
	}
	Permissions = []Permission{
		{0, "role.manage", "Manage roles"},
		{0, "role.view", "View roles"},
		{0, "workgroup.manage", "Manage workgroups"},
		{0, "workgroup.view", "View workgroups"},
		{0, "identity.manage", "Manage identities"},
		{0, "identity.view", "View identities"},
		{0, "engine.manage", "Manage engines"},
		{0, "engine.view", "View engines"},
		{0, "cluster.manage", "Manage clusters"},
		{0, "cluster.view", "View clusters"},
		{0, "model.manage", "Manage models"},
		{0, "model.view", "View models"},
	}
}

type metadata map[string]string

type Keys struct {
	Role      int64
	Workgroup int64
	Identity  int64
	Engine    int64
	Cluster   int64
	Model     int64
}

func toKeys(entityTypes []EntityType) *Keys {
	m := make(map[string]int64)
	for _, entityType := range entityTypes {
		m[entityType.Name] = entityType.Id
	}

	return &Keys{
		m[RoleLabel],
		m[WorkgroupLabel],
		m[IdentityLabel],
		m[EngineLabel],
		m[ClusterLabel],
		m[ModelLabel],
	}
}

type Datastore struct {
	db            *sql.DB // Singleton; doesn't actually connect until used, and is pooled internally.
	metadata      metadata
	permissions   []Permission
	permissionMap map[int64]Permission
	entityTypeMap map[int64]EntityType
	Keys          *Keys
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

	// Get version; prime if pristine

	version, ok := metadata["version"]
	log.Println("Using schema version:", version)
	if !ok {
		prime(db)
	}

	upgrade(db, version)

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

	return &Datastore{
		db,
		metadata,
		permissions,
		permissionMap,
		entityTypeMap,
		toKeys(entityTypes),
	}, nil
}

func prime(db *sql.DB) error {
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

func primePermissions(db *sql.DB, permissions []Permission) error {
	return bulkInsert(db, "permission", []string{"name", "description"}, func(stmt *sql.Stmt) error {
		for _, permission := range permissions {
			_, err := stmt.Exec(permission.Name, permission.Description)
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
			"workgroup",
			"role",
			"identity",
			"permission",
			"entity_type",
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

func (ds *Datastore) SetupSuperuser(name, password string) (int64, error) {
	roleId, err := ds.CreateRole(nil, "Superuser", "Superuser")
	if err != nil {
		return 0, nil
	}

	allPerms := make([]int64, len(ds.permissions))
	for i, permission := range ds.permissions {
		allPerms[i] = permission.Id
	}

	if err := ds.SetRolePermissions(nil, roleId, allPerms); err != nil {
		return 0, err
	}

	identityId, err := ds.CreateIdentity(nil, name, password)
	if err != nil {
		return identityId, err
	}

	if err := ds.LinkIdentityToRole(nil, identityId, roleId); err != nil {
		return identityId, err
	}

	return identityId, nil
}

// --- Lookup tables (static) ---

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

func (ds *Datastore) toPermissionNames(ids []int64) ([]string, error) {
	names := make([]string, len(ids))
	for i, id := range ids {
		if p, ok := ds.permissionMap[id]; ok {
			names[i] = p.Name
		} else {
			return names, fmt.Errorf("Invalid permission id: %d", id)
		}
	}
	return names, nil
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

func (ds *Datastore) audit(principal *az.Principal, tx *sql.Tx, action string, entityTypeId, entityId int64, metadata metadata) error {
	if principal == nil {
		return nil
	}

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
		`, principal.Id, action, entityTypeId, entityId, string(json)); err != nil {
		return err
	}
	return nil
}

// --- Permissions ---

func readAllPermissions(db *sql.DB) ([]Permission, error) {
	rows, err := db.Query(`
		SELECT
			id, name, description
		FROM
			permission
		ORDER BY
			name
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanPermissions(rows)
}

func (ds *Datastore) ReadAllPermissions(principal *az.Principal) ([]Permission, error) {
	return ds.permissions, nil
}

func (ds *Datastore) ReadPermissionsForRole(principal *az.Principal, roleId int64) ([]Permission, error) {
	rows, err := ds.db.Query(`
		SELECT
			p.id, p.name, p.description
		FROM
			role_permission rp,
			permission p
		WHERE
			rp.role_id = $1 AND
			rp.permission_id = p.id
		ORDER BY
			p.name
		`, roleId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanPermissions(rows)
}

func (ds *Datastore) ReadPermissionsForIdentity(principal *az.Principal, identityId int64) ([]int64, error) {
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

// --- Roles ---

func (ds *Datastore) CreateRole(principal *az.Principal, name, description string) (int64, error) {
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

		return ds.audit(principal, tx, CreateOp, ds.Keys.Role, id, metadata{"name": name, "description": description})
	})
	return id, err
}

func (ds *Datastore) ReadRoles(principal *az.Principal, offset, limit int64) ([]Role, error) {
	rows, err := ds.db.Query(`
		SELECT
			id, name, description, created
		FROM
			role
		ORDER BY name
		OFFSET $1
		LIMIT $2
		`, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanRoles(rows)
}

func (ds *Datastore) ReadRole(principal *az.Principal, roleId int64) (Role, error) {
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

func (ds *Datastore) ReadRolesForIdentity(principal *az.Principal, identityId int64) ([]Role, error) {
	rows, err := ds.db.Query(`
		SELECT
			r.id, r.name, r.description, r.created
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

	return ScanRoles(rows)
}

func (ds *Datastore) UpdateRole(principal *az.Principal, roleId int64, name string) error {
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
		return ds.audit(principal, tx, UpdateOp, ds.Keys.Role, roleId, metadata{"name": name})
	})
}

func (ds *Datastore) SetRolePermissions(principal *az.Principal, roleId int64, permissionIds []int64) error {
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

		permissionNames, err := ds.toPermissionNames(permissionIds)
		if err != nil {
			return err
		}
		permissions, err := json.Marshal(permissionNames)
		if err != nil {
			return err
		}
		return ds.audit(principal, tx, UpdateOp, ds.Keys.Role, roleId, metadata{"permissions": string(permissions)})
	})
}

func (ds *Datastore) DeleteRole(principal *az.Principal, roleId int64) error {
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
		if err := deletePrivilegesOn(tx, ds.Keys.Role, roleId); err != nil {
			return err
		}
		return ds.audit(principal, tx, DeleteOp, ds.Keys.Role, roleId, metadata{})
	})
}

// --- Workgroup ---

func (ds *Datastore) CreateWorkgroup(principal *az.Principal, name, description string) (int64, error) {
	var id int64
	err := ds.exec(func(tx *sql.Tx) error {
		row := tx.QueryRow(`
			INSERT INTO
				workgroup
				(name, description, created)
			VALUES
				($1,   $2,          now())
			RETURNING id
			`, name, description)
		if err := row.Scan(&id); err != nil {
			return err
		}
		return ds.audit(principal, tx, CreateOp, ds.Keys.Workgroup, id, metadata{"name": name, "description": description})
	})
	return id, err
}

func (ds *Datastore) ReadWorkgroups(principal *az.Principal, offset, limit int64) ([]Workgroup, error) {
	rows, err := ds.db.Query(`
		SELECT
			id, name, description, created
		FROM
			workgroup
		ORDER BY name
		OFFSET $1
		LIMIT $2
		`, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanWorkgroups(rows)
}

func (ds *Datastore) ReadWorkgroup(principal *az.Principal, workgroupId int64) (Workgroup, error) {
	row := ds.db.QueryRow(`
		SELECT
			id, name, description, created
		FROM
			workgroup
		WHERE
			id = $1
		ORDER BY
			name
		`, workgroupId)

	return ScanWorkgroup(row)
}

func (ds *Datastore) ReadWorkgroupsForIdentity(principal *az.Principal, identityId int64) ([]Workgroup, error) {
	rows, err := ds.db.Query(`
		SELECT
			w.id, w.name, w.description, w.created
		FROM
			workgroup w,
			identity_workgroup iw
		WHERE
		  iw.identity_id = $1 AND
			iw.workgroup_id = w.id
		ORDER BY
			w.name
		`, identityId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanWorkgroups(rows)
}

func (ds *Datastore) UpdateWorkgroup(principal *az.Principal, workgroupId int64, name string) error {
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
		return ds.audit(principal, tx, UpdateOp, ds.Keys.Workgroup, workgroupId, metadata{"name": name})
	})
}

func (ds *Datastore) DeleteWorkgroup(principal *az.Principal, workgroupId int64) error {
	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			DELETE FROM
				workgroup
			WHERE
				id = $1
			`, workgroupId); err != nil {
			return err
		}
		if err := deletePrivilegesOn(tx, ds.Keys.Workgroup, workgroupId); err != nil {
			return err
		}
		if err := deletePrivilegesFor(tx, WorkgroupPrivilege, workgroupId); err != nil {
			return err
		}
		return ds.audit(principal, tx, DeleteOp, ds.Keys.Workgroup, workgroupId, metadata{})
	})
}

// --- Identity ---

func (ds *Datastore) CreateIdentity(principal *az.Principal, name, password string) (int64, error) {
	var id int64
	err := ds.exec(func(tx *sql.Tx) error {
		row := tx.QueryRow(`
			INSERT INTO
				identity
				(name, password, is_active, created)
			VALUES
				($1,   $2,       $3,        now())
			RETURNING id
			`, name, password, true)
		if err := row.Scan(&id); err != nil {
			return err
		}
		return ds.audit(principal, tx, CreateOp, ds.Keys.Identity, id, metadata{"name": name})
	})
	return id, err
}

func (ds *Datastore) ReadIdentities(principal *az.Principal, offset, limit int64) ([]Identity, error) {
	rows, err := ds.db.Query(`
		SELECT
			id, name, is_active, last_login, created
		FROM
			identity
		ORDER BY name
		OFFSET $1
		LIMIT $2
		`, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanIdentitys(rows)
}

func (ds *Datastore) ReadIdentity(principal *az.Principal, identityId int64) (Identity, error) {
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

func (ds *Datastore) ReadIdentityAndPassword(principal *az.Principal, identityId int64) (IdentityAndPassword, error) {
	row := ds.db.QueryRow(`
		SELECT
			id, name, password, is_active, last_login, created
		FROM
			identity
		WHERE
			id = $1
		`, identityId)
	return ScanIdentityAndPassword(row)
}

func (ds *Datastore) ReadIdentitesForWorkgroup(principal *az.Principal, workgroupId int64) ([]Identity, error) {
	rows, err := ds.db.Query(`
		SELECT
			i.id, i.name, i.is_active, i.last_login, i.created
		FROM
			identity i,
			identity_workgroup iw
		WHERE
			iw.workgroup_id = $1 AND
		  iw.identity_id = i.id
		ORDER BY
			i.name
		`, workgroupId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanIdentitys(rows)
}

func (ds *Datastore) ReadIdentitesForRole(principal *az.Principal, roleId int64) ([]Identity, error) {
	rows, err := ds.db.Query(`
		SELECT
			i.id, i.name, i.is_active, i.last_login, i.created
		FROM
			identity i,
			identity_role ir
		WHERE
			ir.role_id = $1 AND
		  ir.identity_id = i.id
		ORDER BY
			i.name
		`, roleId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanIdentitys(rows)
}

func (ds *Datastore) LinkIdentityToWorkgroup(principal *az.Principal, identityId, workgroupId int64) error {
	workgroup, err := ds.ReadWorkgroup(principal, workgroupId)
	if err != nil {
		return err
	}
	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			INSERT INTO
				identity_workgroup
			VALUES
				($1, $2)
			`, identityId, workgroupId); err != nil {
			return err
		}
		return ds.audit(principal, tx, LinkOp, ds.Keys.Identity, identityId, metadata{
			"type": WorkgroupLabel,
			"id":   strconv.FormatInt(workgroupId, 10),
			"name": workgroup.Name,
		})
	})
}

func (ds *Datastore) UnlinkIdentityFromWorkgroup(principal *az.Principal, identityId, workgroupId int64) error {
	workgroup, err := ds.ReadWorkgroup(principal, workgroupId)
	if err != nil {
		return err
	}
	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			DELETE FROM
				identity_workgroup
			WHERE
				identity_id = $1 AND
				workgroup_id = $2
			`, identityId, workgroupId); err != nil {
			return err
		}
		return ds.audit(principal, tx, UnlinkOp, ds.Keys.Identity, identityId, metadata{
			"type": WorkgroupLabel,
			"id":   strconv.FormatInt(workgroupId, 10),
			"name": workgroup.Name,
		})
	})
}

func (ds *Datastore) LinkIdentityToRole(principal *az.Principal, identityId, roleId int64) error {
	role, err := ds.ReadRole(principal, roleId)
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
		return ds.audit(principal, tx, LinkOp, ds.Keys.Identity, identityId, metadata{
			"type": RoleLabel,
			"id":   strconv.FormatInt(roleId, 10),
			"name": role.Name,
		})
	})
}

func (ds *Datastore) UnlinkIdentityFromRole(principal *az.Principal, identityId, roleId int64) error {
	role, err := ds.ReadRole(principal, roleId)
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
		return ds.audit(principal, tx, UnlinkOp, ds.Keys.Identity, identityId, metadata{
			"type": RoleLabel,
			"id":   strconv.FormatInt(roleId, 10),
			"name": role.Name,
		})
	})
}

func (ds *Datastore) DeactivateIdentity(principal *az.Principal, identityId int64) error {
	identity, err := ds.ReadIdentity(principal, identityId)
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
		return ds.audit(principal, tx, DisableOp, ds.Keys.Identity, identityId, metadata{"name": identity.Name})
	})
	return nil
}

// --- Privileges ---

func readName(tx *sql.Tx, table string, id int64) (string, error) {
	row := tx.QueryRow("SELECT name FROM "+table+" WHERE id = $1", id)
	var name string
	err := row.Scan(&name)
	return name, err
}

func (ds *Datastore) CreatePrivilege(principal *az.Principal, privilege Privilege) error {
	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			INSERT INTO
				privilege
			VALUES
				($1, $2, $3, $4, $5)
			`,
			privilege.Type,
			privilege.IdentityType,
			privilege.IdentityId,
			privilege.EntityType,
			privilege.EntityId,
		); err != nil {
			return err
		}

		identityName, err := readName(tx, privilege.IdentityType, privilege.IdentityId)
		if err != nil {
			return err
		}

		return ds.audit(principal, tx, ShareOp, privilege.EntityType, privilege.EntityId, metadata{
			"type": privilege.IdentityType,
			"id":   strconv.FormatInt(privilege.IdentityId, 10),
			"name": identityName,
		})
	})
}

func (ds *Datastore) ReadPrivileges(principal *az.Principal, identityId, entityTypeId, entityId int64) ([]string, error) {
	rows, err := ds.db.Query(`
		SELECT
			privilege_type
		FROM
			privilege
		WHERE
			entity_id = $1 AND
			entity_type_id = $2 AND
			(
				(
					identity_id = $3 AND
					identity_type = 'identity'
				)
				OR
				(
					identity_id IN 
					(
						SELECT
							workgroup_id
						FROM
							identity_workgroup
						WHERE
							identity_id = $3
					) AND
					identity_type = 'workgroup'
				)
			)
		`, entityId, entityTypeId, identityId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanStrings(rows)
}

func (ds *Datastore) DeletePrivilege(principal *az.Principal, privilege Privilege) error {
	return ds.exec(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			DELETE FROM
				privilege
			WHERE
				privilege_type = $1 AND
				identity_type = $2 AND
				identity_id = $3 AND
				entity_type_id = $4 AND
				entity_id = $5
			`,
			privilege.Type,
			privilege.IdentityType,
			privilege.IdentityId,
			privilege.EntityType,
			privilege.EntityId,
		); err != nil {
			return err
		}

		identityName, err := readName(tx, privilege.IdentityType, privilege.IdentityId)
		if err != nil {
			return err
		}

		return ds.audit(principal, tx, UnshareOp, privilege.EntityType, privilege.EntityId, metadata{
			"type": privilege.IdentityType,
			"id":   strconv.FormatInt(privilege.IdentityId, 10),
			"name": identityName,
		})
	})
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
			identity_type = $1 AND
			identity_id = $2
		`, identityType, identityId)
	return err
}
