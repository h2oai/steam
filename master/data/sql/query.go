package sql

import (
	"fmt"

	"github.com/h2oai/steam/master/az"
	"github.com/pkg/errors"
	"gopkg.in/doug-martin/goqu.v3"
)

type QueryConfig struct {
	tx       *goqu.TxDatabase
	table    string
	dataset  *goqu.Dataset
	fields   map[string]interface{}
	postFunc []QueryOpt
	result   int64
	audit    bool
	pz       az.Principal

	clusterType clusterTypeKeys
	state       stateKeys
	permission  permissionKeys
	entityType  entityTypeKeys
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

	return &QueryConfig{
		tx:       tx,
		table:    table,
		dataset:  data,
		fields:   make(map[string]interface{}),
		postFunc: make([]QueryOpt, 0),

		clusterType: clusterTypes,
		state:       states,
		permission:  permissions,
		entityType:  entityTypes,
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

// WithAudit adds a history entry for the provided principal
func WithAudit(pz az.Principal) QueryOpt {
	return func(q *QueryConfig) error {
		if q.pz == nil {
			return errors.New("WithAudit: no principal provided")
		}
		q.pz = pz
		q.audit = true
		return nil
	}
}

// WithDescription adds a description value to the query
func WithDescription(description string) QueryOpt {
	return func(q *QueryConfig) (err error) { q.fields["description"] = description; return }
}

// WithDefaultIdentityWorkgroup creates and links a default workgroup for an identity
func WithDefaultIdentityWorkgroup(q *QueryConfig) error {
	// Fetch identity name
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
		_, err := createIdentityWorkgroup(c.tx, c.result, workgroupId)
		return errors.Wrap(err, "WithDefaultIdentityWorkgroup: linking identity to workgroup")
	})
	q.AddPostFunc(func(c *QueryConfig) error {
		_, err := createPrivilege(c.tx, Owns, workgroupId, c.entityType.Identity, c.result)
		return errors.Wrap(err, "WithDefaultIdentityWorkgroup: creating identity privilege")
	})
	q.AddPostFunc(func(c *QueryConfig) error {
		_, err := createPrivilege(c.tx, Owns, workgroupId, c.entityType.Workgroup, workgroupId)
		return errors.Wrap(err, "WithDefaultIdentityWorkgroup: creating workgroup privilege")
	})
	return nil
}

// WithExternal adds a type_id of external
func WithExternal(q *QueryConfig) (err error) {
	q.fields["type_id"] = q.clusterType.External
	return
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

		q.fields["type_id"] = q.clusterType.Yarn
		q.fields["detail_id"] = yarnId
		return nil
	}
}

// WithPrivilege adds a privilege to an entity with the corresponding principal
func WithPrivilege(pz az.Principal, typ string, entityTypeId int64) QueryOpt {
	return func(q *QueryConfig) error {
		q.AddPostFunc(func(c *QueryConfig) error {
			_, err := createPrivilege(c.tx, typ, pz.WorkgroupId(), entityTypeId, c.result)
			return err
		})
		return nil
	}
}

// CheckPrivilege checks the privilege of the principal with an entity
func CheckPrivilege(pz az.Principal, entityTypeId int64) QueryOpt {
	return func(q *QueryConfig) error {
		if pz.IsSuperuser() {
			return nil
		}
		x := q.tx.From("identity_workgroup").Select("workgroup_id").Where(
			goqu.I("identity_id").Eq(pz.Id()),
		)
		aux := q.tx.From("privilege").SelectDistinct("entity_id").Where(
			goqu.I("workgroup_id").In(x),
			goqu.I("entity_type").Eq(entityTypeId),
		)

		q.dataset = q.dataset.Where(goqu.I("id").In(aux))
		return nil
	}
}
