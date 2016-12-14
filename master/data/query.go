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
	"encoding/json"
	"fmt"
	"strconv"

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
	permissions  permissionKeys
	entityTypes  entityTypeKeys
}

func NewQueryConfig(ds *Datastore, tx *goqu.TxDatabase, table string, data *goqu.Dataset) *QueryConfig {
	var (
		clusterTypes clusterTypeKeys
		permissions  permissionKeys
		entityTypes  entityTypeKeys
	)
	if ds != nil {
		clusterTypes = ds.ClusterType
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

func (q *QueryConfig) I(column string) goqu.IdentifierExpression {
	return goqu.I(q.table + "." + column)
}

type QueryOpt func(*QueryConfig) error

// ------------- ------------- -------------
// ------------- Query Options -------------
// ------------- ------------- -------------

func WithActivity(activate bool) QueryOpt {
	return func(q *QueryConfig) (err error) { q.fields["is_active"] = activate; return }
}

// ByAddress queries the database for matching address columns
func ByAddress(address string) QueryOpt {
	return func(q *QueryConfig) (err error) { q.dataset = q.dataset.Where(q.I("address").Eq(address)); return }
}

// WithAddress adds an address value to the query
func WithAddress(address string) QueryOpt {
	return func(q *QueryConfig) (err error) { q.fields["address"] = address; return }
}

// WithBinomial model creates an entry in the binomial_metrics table and links it to the model
func WithBinomialModel(mse, rSquared, logloss, auc, gini float64) QueryOpt {
	return func(q *QueryConfig) error {
		if q.entityTypeId != q.entityTypes.Model {
			return errors.New("WithBinomialModel: entity must be of type 'Model'")
		}
		q.AddPostFunc(func(c *QueryConfig) error {
			_, err := createBinomialModel(c.tx, c.entityId, mse, rSquared, logloss, auc, gini)
			return errors.Wrap(err, "WithBinomialModel: creating binomial model")
		})
		return nil
	}
}

// WithCluster adds cluster_id and cluster_name values to the query
func WithCluster(clusterId int64, clusterName string) QueryOpt {
	return func(q *QueryConfig) (err error) {
		q.fields["cluster_id"] = clusterId
		q.fields["cluster_name"] = clusterName
		return
	}
}

// WithClusterId adds a cluster_id value to the query
func WithClusterId(clusterId int64) QueryOpt {
	return func(q *QueryConfig) (err error) { q.fields["cluster_id"] = clusterId; return }
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

func ForEntity(entityTypeId, entityId int64) QueryOpt {
	return func(q *QueryConfig) error {
		return nil
	}
}

// ByEntityId queries the database for matching entity_id columns
func ByEntityId(entityId int64) QueryOpt {
	return func(q *QueryConfig) (err error) {
		q.dataset = q.dataset.Where(q.I("entity_id").Eq(entityId))
		return
	}
}

// ByEntityTypeId queries the database for matching entity_type_id columns
func ByEntityTypeId(entityTypeId int64) QueryOpt {
	return func(q *QueryConfig) (err error) {
		q.dataset = q.dataset.Where(q.I("entity_type_id").Eq(entityTypeId))
		return
	}
}

func WithFilterByName(filter string) QueryOpt {
	return func(q *QueryConfig) (err error) {
		q.dataset = q.dataset.Where(q.I("name").Like("%" + filter + "%"))
		return
	}
}

// WithIdentityId adds an identity_id value to the query
func WithIdentityId(identityId int64) QueryOpt {
	return func(q *QueryConfig) (err error) { q.fields["identity_id"] = identityId; return }
}

// ById queries the database for matching id columns
func ById(id int64) QueryOpt {
	return func(q *QueryConfig) (err error) { q.dataset = q.dataset.Where(q.I("id").Eq(id)); return }
}

func ByIdentityId(identityId int64) QueryOpt {
	return func(q *QueryConfig) (err error) {
		q.dataset = q.dataset.Where(q.I("identity_id").Eq(identityId))
		return
	}
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

func WithLocation(modelId int64, logicalName string) QueryOpt {
	location := strconv.FormatInt(modelId, 10)
	return func(q *QueryConfig) (err error) {
		q.fields["location"] = location
		q.fields["logical_name"] = logicalName
		return
	}
}

func ByModelId(modelId int64) QueryOpt {
	return func(q *QueryConfig) (err error) { q.dataset = q.dataset.Where(q.I("model_id").Eq(modelId)); return }
}

func WithModelId(modelId int64) QueryOpt {
	return func(q *QueryConfig) (err error) { q.fields["model_id"] = modelId; return }

}

func WithModelObjectType(typ string) QueryOpt {
	return func(q *QueryConfig) (err error) { q.fields["model_object_type"] = typ; return }
}

func WithMultinomialModel(mse, rSquared, logloss float64) QueryOpt {
	return func(q *QueryConfig) error {
		if q.entityTypeId != q.entityTypes.Model {
			return errors.New("WithMultinomialModel: entity must of type model")
		}
		q.AddPostFunc(func(c *QueryConfig) error {
			_, err := createMultinomialModel(c.tx, c.entityId, mse, rSquared, logloss)
			return errors.Wrap(err, "WithMultinomialModel: creating multinomial model")
		})
		return nil
	}
}

// ByName queries the database for matching name columns
func ByName(name string) QueryOpt {
	return func(q *QueryConfig) (err error) { q.dataset = q.dataset.Where(q.I("name").Eq(name)); return }
}

func WithName(name string) QueryOpt {
	return func(q *QueryConfig) (err error) { q.fields["name"] = name; return }
}

func WithNil(column string) QueryOpt {
	return func(q *QueryConfig) (err error) { q.fields[column] = nil; return }
}

// WithOffset adds a offset value to the query
func WithOffset(offset uint) QueryOpt {
	return func(q *QueryConfig) (err error) { q.dataset = q.dataset.Offset(offset); return }
}

func WithOrderBy(column string, asc bool) QueryOpt {
	return func(q *QueryConfig) (err error) {
		if asc {
			q.dataset = q.dataset.Order(goqu.I(column).Asc())
			return
		}
		q.dataset = q.dataset.Order(goqu.I(column).Desc())
		return
	}
}

// WithPassword adds a password value to the query
func WithPassword(password string) QueryOpt {
	return func(q *QueryConfig) (err error) { q.fields["password"] = password; return }
}

func ByPermissionId(permissionId int64) QueryOpt {
	return func(q *QueryConfig) (err error) {
		q.dataset = q.dataset.Where(q.I("permission_id").Eq(permissionId))
		return
	}
}

func LinkPermissions(reset bool, permissionIds ...int64) QueryOpt {
	return func(q *QueryConfig) error {
		if q.entityTypeId != q.entityTypes.Role {
			return errors.New("LinkPermission: permission can only be linked to a role")
		}
		q.AddPostFunc(func(c *QueryConfig) error {
			if reset {
				if err := deleteRolePermission(c.tx, ByRoleId(c.entityId)); err != nil {
					return errors.Wrap(err, "LinkPermission: deleting role permission")
				}
			}
			for _, permissionId := range permissionIds {
				if _, err := createRolePermission(c.tx, c.entityId, permissionId); err != nil {
					return errors.Wrap(err, "LinkPermission: creating role permission")
				}
			}
			return nil
		})
		return nil
	}
}

func UnlinkPermissions(permissionIds ...int64) QueryOpt {
	return func(q *QueryConfig) error {
		if q.entityTypeId != q.entityTypes.Role {
			return errors.New("UnlinkPermission: permission can only be unlinked from a role")
		}
		q.AddPostFunc(func(c *QueryConfig) error {
			for _, permissionId := range permissionIds {
				if err := deleteRolePermission(c.tx, ByRoleId(c.entityId), ByPermissionId(permissionId)); err != nil {
					return errors.Wrap(err, "UnlinkPermission: deleting role permission")
				}
			}
			return nil
		})
		return nil
	}
}

// ByProject queries the database for a matching state column
func ByProjectId(projectId int64) QueryOpt {
	return func(q *QueryConfig) (err error) {
		q.dataset = q.dataset.Where(q.I("project_id").Eq(projectId))
		return
	}
}

func WithPort(port int) QueryOpt {
	return func(q *QueryConfig) (err error) { q.fields["port"] = port; return }
}

// WithProjectId adds a project_id value to the query
func WithProjectId(projectId int64) QueryOpt {
	return func(q *QueryConfig) (err error) { q.fields["project_id"] = projectId; return }
}

func WithProcessId(processId int) QueryOpt {
	return func(q *QueryConfig) (err error) { q.fields["process_id"] = processId; return }
}

// WithRawSchema adds schema and schema_version values to the query
func WithRawSchema(schema, schemaVersion string) QueryOpt {
	return func(q *QueryConfig) (err error) {
		q.fields["schema"] = schema
		q.fields["schema_version"] = schemaVersion
		return
	}
}

func WithRegressionModel(mse, rSquared, meanResidualDeviance float64) QueryOpt {
	return func(q *QueryConfig) error {
		if q.entityTypeId != q.entityTypes.Model {
			return errors.New("WithRegressionModel: entity must of type model")
		}

		q.AddPostFunc(func(c *QueryConfig) error {
			_, err := createRegressionModel(c.tx, c.entityId, mse, rSquared, meanResidualDeviance)
			return errors.Wrap(err, "WithRegressionModel: creating regression model")
		})
		return nil
	}
}

func ByRoleId(roleId int64) QueryOpt {
	return func(q *QueryConfig) (err error) { q.dataset = q.dataset.Where(q.I("role_id").Eq(roleId)); return }
}

func ForRole(roleId int64) QueryOpt {
	return func(q *QueryConfig) error {
		crossCol := q.table + "_id"
		var crossTbl string
		switch q.table {
		case "identity":
			crossTbl = "identity_role"
		case "permission":
			crossTbl = "role_permission"
		default:
			return fmt.Errorf("ForRole: unable to search %s for roles", q.table)
		}
		// Create relational dataset
		ds := q.tx.From(crossTbl).SelectDistinct(crossCol).Where(
			goqu.I("role_id").Eq(roleId),
		)
		q.dataset = q.dataset.Where(goqu.I("id").Eq(ds))
		return nil
	}
}

func LinkRole(roleId int64) QueryOpt {
	return func(q *QueryConfig) error {
		if q.entityTypeId != q.entityTypes.Identity {
			return errors.New("LinkRole: roles may only be linked with identities")
		}
		q.AddPostFunc(func(c *QueryConfig) error {
			_, err := createIdentityRole(c.tx, WithIdentityId(c.entityId), WithRoleId(roleId))
			return errors.Wrap(err, "creating identity role in database")
		})
		return nil
	}
}

func UnlinkRole(roleId int64) QueryOpt {
	return func(q *QueryConfig) error {
		if q.entityTypeId != q.entityTypes.Identity {
			return errors.New("UnlinkRole: roles may only be unlinked with identities")
		}
		q.AddPostFunc(func(c *QueryConfig) error {
			err := deleteIdentityRole(c.tx, ByIdentityId(c.entityId), ByRoleId(roleId))
			return errors.Wrap(err, "creating identity role in database")
		})
		return nil
	}
}

// WithRoleId adds an role_id value to the query
func WithRoleId(roleId int64) QueryOpt {
	return func(q *QueryConfig) (err error) { q.fields["role_id"] = roleId; return }
}

// WithSize adds a size value to the query
func WithSize(size string) QueryOpt {
	return func(q *QueryConfig) (err error) { q.fields["size"] = size; return }
}

// ByState queries the database for a matching state column
func ByState(state int64) QueryOpt {
	return func(q *QueryConfig) (err error) { q.dataset = q.dataset.Where(q.I("state").Eq(state)); return }
}

// WithState adds a state value to the query
func WithState(state string) QueryOpt {
	return func(q *QueryConfig) (err error) { q.fields["state"] = state; return }
}

func ByWorkgroupId(workgroupId int64) QueryOpt {
	return func(q *QueryConfig) (err error) {
		q.dataset = q.dataset.Where(q.I("workgroup_id").Eq(workgroupId))
		return
	}
}

func ForWorkgroup(workgroupId int64) QueryOpt {
	return func(q *QueryConfig) error {
		if q.entityTypeId != q.entityTypes.Identity {
			errors.New("ForWorkgroup: workgroups can only be searched from identities")
		}
		// Create relational dataset
		ds := q.tx.From("identity_workgroup").SelectDistinct("identity_id").Where(
			goqu.I("workgroup_id").Eq(workgroupId),
		)
		q.dataset = q.dataset.Where(goqu.I("identity_id").Eq(ds))
		return nil
	}
}

func LinkWorkgroup(workgroupId int64) QueryOpt {
	return func(q *QueryConfig) error {
		if q.entityTypeId != q.entityTypes.Identity {
			return errors.New("LinkWorkgroup: workgroups can only be linked to identities")
		}
		// Insert into relation
		q.AddPostFunc(func(c *QueryConfig) error {
			_, err := createIdentityWorkgroup(c.tx, c.entityId, workgroupId)
			return errors.Wrap(err, "creating identity workgroup in database")
		})
		return nil
	}
}

func UnlinkWorkgroup(workgroupId int64) QueryOpt {
	return func(q *QueryConfig) error {
		if q.entityTypeId != q.entityTypes.Identity {
			return errors.New("LinkWorkgroup: workgroups can only be linked to identities")
		}
		// Insert into relation
		q.AddPostFunc(func(c *QueryConfig) error {
			err := deleteIdentityWorkgroup(c.tx, ByIdentityId(c.entityId), ByWorkgroupId(workgroupId))
			return errors.Wrap(err, "creating identity workgroup in database")
		})
		return nil
	}
}

// WithYarnDetail adds a type_id of yarn and provides the corresponding detail
func WithYarnDetail(engineId, size int64, applicationId, memory, outputDir string) QueryOpt {
	return func(q *QueryConfig) error {
		if q.fields["type_id"] != q.clusterTypes.Yarn {
			return errors.New("cannot add yarn details for a non-yarn cluster")
		}
		yarnId, err := createClusterYarnDetail(q.tx, engineId, size, applicationId, memory, outputDir)
		if err != nil {
			return errors.Wrap(err, "WithYarnDetail: creating cluster yarn details in database")
		}

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

func WithLinkAudit(pz az.Principal) QueryOpt {
	return func(q *QueryConfig) error {
		q.AddPostFunc(func(c *QueryConfig) error {
			if pz == nil {
				return errors.New("WithAudit: no principal provided")
			}
			json, err := json.Marshal(c.fields)
			if err != nil {
				return errors.Wrap(err, "WithAudit: serializing metadata")
			}
			_, err = createHistory(c.tx, LinkOp, pz.Id(), c.entityTypeId, c.entityId,
				WithDescription(string(json)),
			)
			return errors.Wrap(err, "WithAudit: creating audit entry")
		})
		return nil
	}
}

func WithUnlinkAudit(pz az.Principal) QueryOpt {
	return func(q *QueryConfig) error {
		q.AddPostFunc(func(c *QueryConfig) error {
			if pz == nil {
				return errors.New("WithAudit: no principal provided")
			}
			json, err := json.Marshal(c.fields)
			if err != nil {
				return errors.Wrap(err, "WithAudit: serializing metadata")
			}
			_, err = createHistory(c.tx, UnlinkOp, pz.Id(), c.entityTypeId, c.entityId,
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
