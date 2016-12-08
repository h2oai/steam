package sql

import (
	"encoding/json"
	"fmt"

	"github.com/h2oai/steam/master/az"
	"github.com/pkg/errors"
	"gopkg.in/doug-martin/goqu.v3"
)

type QueryConfig struct {
	// Sql query setting
	tx       *goqu.TxDatabase
	table    string
	dataset  *goqu.Dataset
	fields   map[string]interface{}
	postFunc []QueryOpt
	// Entity options
	entityTypeId int64
	entityId     int64
	audit        string
	// Enum references
	clusterTypes clusterTypeKeys
	state        stateKeys
	permissions  permissionKeys
	entityTypes  entityTypeKeys
}

func NewQueryConfig(ds *Datastore, tx *goqu.TxDatabase, table string, data *goqu.Dataset) *QueryConfig {
	var (
		clusterTypes clusterTypeKeys
		states       stateKeys
		permissions  permissionKeys
		entityTypes  entityTypeKeys
	)
	if ds != nil {
		clusterTypes = ds.ClusterType
		states = ds.State
		permissions = ds.Permission
		entityTypes = ds.EntityType
	}

	var entityTypeId int64
	if ds != nil {
		entityTypeId = toEntityId(ds, table)
	}

	return &QueryConfig{
		tx:       tx,
		table:    table,
		dataset:  data,
		fields:   make(map[string]interface{}),
		postFunc: make([]QueryOpt, 0),

		entityTypeId: entityTypeId,

		clusterTypes: clusterTypes,
		state:        states,
		permissions:  permissions,
		entityTypes:  entityTypes,
	}
}

func (q *QueryConfig) AddFields(fs goqu.Record) {
	for key, value := range fs {
		q.fields[key] = value
	}
}

func (q *QueryConfig) AddPostFunc(opt QueryOpt) {
	q.postFunc = append(q.postFunc, opt)
}

type QueryOpt func(*QueryConfig) error

// ------------- ------------- -------------
// ------------- Query Options -------------
// ------------- ------------- -------------

// ByAddress queries the database for matching address columns
func ByAddress(address string) QueryOpt {
	return func(q *QueryConfig) (err error) { q.dataset = q.dataset.Where(goqu.I("address").Eq(address)); return }
}

// WithAddress adds an address value to the query
func WithAddress(address string) QueryOpt {
	return func(q *QueryConfig) (err error) { q.fields["address"] = address; return }
}

// WithDefaultIdentityWorkgroup creates and links a default workgroup for an identity
func WithDefaultIdentityWorkgroup(q *QueryConfig) error {
	// Fetch identity name
	if q.entityTypeId != q.entityTypes.Identity {
		return errors.New("WithDefaultIdentityWorkgroup: entity must be of type 'Identity'")
	}
	name := q.fields["name"]
	val, ok := name.(string)
	if !ok || val == "" {
		return errors.New("WithDefaultIdentityWorkgroup: identity must have a name for workgroup")
	}
	// Create workgroup
	workgroupId, err := createWorkgroup(q.tx, "identity", "user:"+val)
	if err != nil {
		return errors.Wrap(err, "WithDefaultIdentityWorkgroup: creating default workgroup")
	}
	// Add workgroup to query and add post functions
	q.fields["workgroup_id"] = workgroupId
	q.AddPostFunc(func(c *QueryConfig) error {
		_, err := createIdentityWorkgroup(c.tx, c.entityId, workgroupId)
		return errors.Wrap(err, "WithDefaultIdentityWorkgroup: linking identity to workgroup")
	})
	q.AddPostFunc(func(c *QueryConfig) error {
		_, err := createPrivilege(c.tx, Owns, workgroupId, c.entityTypes.Identity, c.entityId)
		return errors.Wrap(err, "WithDefaultIdentityWorkgroup: creating identity privilege")
	})
	q.AddPostFunc(func(c *QueryConfig) error {
		_, err := createPrivilege(c.tx, Owns, workgroupId, c.entityTypes.Workgroup, workgroupId)
		return errors.Wrap(err, "WithDefaultIdentityWorkgroup: creating workgroup privilege")
	})
	return nil
}

// WithDescription adds a description value to the query
func WithDescription(description string) QueryOpt {
	return func(q *QueryConfig) (err error) { q.fields["description"] = description; return }
}

// ByEntityId queries the database for matching entity_id columns
func ByEntityId(entityId int64) QueryOpt {
	return func(q *QueryConfig) (err error) {
		q.dataset = q.dataset.Where(goqu.I("entity_id").Eq(entityId))
		return
	}
}

// ByEntityTypeId queries the database for matching entity_type_id columns
func ByEntityTypeId(entityTypeId int64) QueryOpt {
	return func(q *QueryConfig) (err error) {
		q.dataset = q.dataset.Where(goqu.I("entity_type_id").Eq(entityTypeId))
		return
	}
}

// WithExternal adds a type_id of external
func WithExternal(q *QueryConfig) (err error) {
	q.fields["type_id"] = q.clusterTypes.External
	return
}

// ById queries the database for matching id columns
func ById(id int64) QueryOpt {
	return func(q *QueryConfig) (err error) { q.dataset = q.dataset.Where(goqu.I("id").Eq(id)); return }
}

// ForIdentity queries an identity relation table
func ForIdentity(identityId int64) QueryOpt {
	return func(q *QueryConfig) error {
		// There must exist a identity_TABLE
		var isPerm bool
		switch q.table {
		case "role", "workgroup":
		// Permission must map identity_role -> role_permission -> permission
		case "permission":
			q.table = "role"
			isPerm = true
		default:
			return fmt.Errorf("ForIdentity: no identity for %s table", q.table)
		}
		// Create a nested select table
		crossTbl := "identity_" + q.table
		crossCol := q.table + "_id"
		ds := q.tx.From(crossTbl).SelectDistinct(crossCol).Where(
			goqu.I("identity_id").Eq(identityId),
		)
		if isPerm {
			ds = q.tx.From("role_permission").SelectDistinct("permission_id").Where(
				goqu.I("role_id").Eq(ds),
			)
		}
		// Add to dataset
		q.dataset = q.dataset.Where(goqu.I("id").Eq(ds))
		return nil
	}
}

// WithLimit adds a limit value to the query
func WithLimit(limit uint) QueryOpt {
	return func(q *QueryConfig) (err error) { q.dataset = q.dataset.Limit(limit); return }
}

// ByName queries the database for matching name columns
func ByName(name string) QueryOpt {
	return func(q *QueryConfig) (err error) { q.dataset = q.dataset.Where(goqu.I("id").Eq(name)); return }
}

// WithOffset adds a offset value to the query
func WithOffset(offset uint) QueryOpt {
	return func(q *QueryConfig) (err error) { q.dataset = q.dataset.Offset(offset); return }
}

// WithSize adds a size value to the query
func WithSize(size string) QueryOpt {
	return func(q *QueryConfig) (err error) { q.fields["size"] = size; return }
}

// ByState queries the database for a matching state column
func ByState(state int64) QueryOpt {
	return func(q *QueryConfig) (err error) { q.dataset = q.dataset.Where(goqu.I("state").Eq(state)); return }
}

// WithState adds a state value to the query
func WithState(state int64) QueryOpt {
	return func(q *QueryConfig) (err error) { q.fields["state"] = state; return }
}

// WithYarnDetail adds a type_id of yarn and provides the corresponding detail
func WithYarnDetail(engineId, size int64, applicationId, memory, outputDir string) QueryOpt {
	return func(q *QueryConfig) error {
		yarnId, err := createClusterYarnDetail(q.tx, engineId, size, applicationId, memory, outputDir)
		if err != nil {
			return errors.Wrap(err, "WithYarnDetail: creating cluster yarn details in database")
		}

		q.fields["type_id"] = q.clusterTypes.Yarn
		q.fields["detail_id"] = yarnId
		return nil
	}
}

// --------- --------- ---------
// --------- Principal ---------
// --------- --------- ---------

// WithAudit adds a history entry for the provided principal
func WithAudit(pz az.Principal) QueryOpt {
	return func(q *QueryConfig) error {
		q.AddPostFunc(func(c *QueryConfig) error {
			if pz == nil {
				return errors.New("WithAudit: no principal provided")
			}
			json, err := json.Marshal(c.fields)
			if err != nil {
				return errors.Wrap(err, "WithAudit: serializing metadata")
			}
			_, err = createHistory(c.tx, c.audit, pz.Id(), c.entityTypeId, c.entityId,
				WithDescription(string(json)),
			)
			return errors.Wrap(err, "WithAudit: creating audit entry")
		})
		return nil
	}
}

// WithPrivilege adds a privilege to an entity with the corresponding principal
func WithPrivilege(pz az.Principal, typ string) QueryOpt {
	return func(q *QueryConfig) error {
		q.AddPostFunc(func(c *QueryConfig) error {
			if pz == nil {
				return errors.New("WithPrivilege: no principal provided")
			}
			_, err := createPrivilege(c.tx, typ, pz.WorkgroupId(), c.entityTypeId, c.entityId)
			return errors.Wrap(err, "WithPrivilege: creating privilege")
		})
		return nil
	}
}

// ByPrivilege returns only entries that match the privileges of the provided lprincipa
func ByPrivilege(pz az.Principal) QueryOpt {
	return func(q *QueryConfig) error {
		if pz == nil {
			return errors.New("CheckPrivilege: no principal provided")
		}
		// Noop if isSuperuser
		if pz.IsSuperuser() {
			return nil
		}
		x := q.tx.From("identity_workgroup").Select("workgroup_id").Where(
			goqu.I("identity_id").Eq(pz.Id()),
		)
		aux := q.tx.From("privilege").SelectDistinct("entity_id").Where(
			goqu.I("workgroup_id").In(x),
			goqu.I("entity_type").Eq(q.entityTypeId),
		)

		q.dataset = q.dataset.Where(goqu.I("id").In(aux))
		return nil
	}
}
