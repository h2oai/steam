package sql

import (
	"github.com/h2oai/steam/master/az"
	"gopkg.in/doug-martin/goqu.v3"
)

type QueryConfig struct {
	tx       *goqu.TxDatabase
	dataset  *goqu.Dataset
	fields   map[string]interface{}
	postFunc []QueryOpt
	result   int64

	clusterType clusterTypeKeys
	state       stateKeys
	permission  permissionKeys
}

func NewQueryConfig(ds *Datastore, tx *goqu.TxDatabase, data *goqu.Dataset) *QueryConfig {
	var (
		clusterTypes clusterTypeKeys
		states       stateKeys
		permissions  permissionKeys
	)
	if ds != nil {
		clusterTypes = ds.ClusterType
		states = ds.State
		permissions = ds.Permission
	}

	return &QueryConfig{
		tx:       tx,
		dataset:  data,
		fields:   make(map[string]interface{}),
		postFunc: make([]QueryOpt, 0),

		clusterType: clusterTypes,
		state:       states,
		permission:  permissions,
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

func ByAddress(address string) QueryOpt {
	return func(q *QueryConfig) (err error) { q.dataset = q.dataset.Where(goqu.I("address").Eq(address)); return }
}

func WithAddress(address string) QueryOpt {
	return func(q *QueryConfig) (err error) { q.fields["address"] = address; return }
}

func WithAudit(pz az.Principal, action string, entityTypeId int64, metadata metadata) QueryOpt {
	json, err := json.Marshal(metadata)
	if err != nil {
		return errors.Wrap(err, "serializing metadata")
	}
}

func WithExternal(q *QueryConfig) (err error) {
	q.fields["type_id"] = q.clusterType.External
	return
}

func ById(id int64) QueryOpt {
	return func(q *QueryConfig) (err error) { q.dataset = q.dataset.Where(goqu.I("id").Eq(id)); return }
}

func ByName(name string) QueryOpt {
	return func(q *QueryConfig) (err error) { q.dataset = q.dataset.Where(goqu.I("id").Eq(name)); return }
}

func ByState(state int64) QueryOpt {
	return func(q *QueryConfig) (err error) { q.dataset = q.dataset.Where(goqu.I("state").Eq(state)); return }
}

func WithState(state int64) QueryOpt {
	return func(q *QueryConfig) (err error) { q.fields["state"] = state; return }
}

func WithYarnDetail() QueryOpt {
	return func(q *QueryConfig) error {
		yarnId := 0
		q.fields["type_id"] = q.clusterType.Yarn
		q.fields["detail_id"] = yarnId
		return nil
	}
}

func WithPrivilege(pz az.Principal, typ string, entityTypeId int64) QueryOpt {
	return func(q *QueryConfig) error {
		q.AddPostFunc(func(c *QueryConfig) error {
			_, err := createPrivilege(c.tx, typ, pz.WorkgroupId(), entityTypeId, c.result)
			return err
		})
		return nil
	}
}

func CheckPrivilege(pz az.Principal, entityId int64) QueryOpt {
	return func(q *QueryConfig) error {
		if pz.IsSuperuser() {
			return nil
		}
		x := q.tx.From("identity_workgroup").Select("workgroup_id").Where(
			goqu.I("identity_id").Eq(pz.Id()),
		)
		aux := q.tx.From("privilege").SelectDistinct("entity_id").Where(
			goqu.I("workgroup_id").In(x),
			goqu.I("entity_type").Eq(entityId),
		)

		q.dataset = q.dataset.Where(goqu.I("id").In(aux))
		return nil
	}
}
