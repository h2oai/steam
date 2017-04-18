//go:generate scaneo $GOFILE

package data

import (
	"database/sql"
	"log"

	"gopkg.in/doug-martin/goqu.v3"

	"github.com/pkg/errors"
)

type LabelModel struct {
	Model
	Label Label
}

// TODO: FIND AN ALTERNATIVE
func scanLabelModel(r *sql.Row) (LabelModel, error) {
	var s LabelModel
	if err := r.Scan(
		&s.Id,
		&s.ProjectId,
		&s.Name,
		&s.ClusterId,
		&s.ClusterName,
		&s.ModelKey,
		&s.Algorithm,
		&s.ModelCategory,
		&s.DatasetName,
		&s.ResponseColumn,
		&s.LogicalName,
		&s.Location,
		&s.ModelObjectType,
		&s.MaxRunTime,
		&s.Schema,
		&s.SchemaVersion,
		&s.Created,
		&s.Label.Id,
		&s.Label.ProjectId,
		&s.Label.ModelId,
		&s.Label.Name,
		&s.Label.Description,
		&s.Label.Created,
	); err != nil {
		return LabelModel{}, err
	}
	return s, nil
}

func scanLabelModels(rs *sql.Rows) ([]LabelModel, error) {
	structs := make([]LabelModel, 0, 16)
	var err error
	for rs.Next() {
		var s LabelModel
		if err = rs.Scan(
			&s.Id,
			&s.ProjectId,
			&s.Name,
			&s.ClusterId,
			&s.ClusterName,
			&s.ModelKey,
			&s.Algorithm,
			&s.ModelCategory,
			&s.DatasetName,
			&s.ResponseColumn,
			&s.LogicalName,
			&s.Location,
			&s.ModelObjectType,
			&s.MaxRunTime,
			&s.Schema,
			&s.SchemaVersion,
			&s.Created,
			&s.Label.Id,
			&s.Label.ProjectId,
			&s.Label.ModelId,
			&s.Label.Name,
			&s.Label.Description,
			&s.Label.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func (ds *Datastore) ReadLabelModels(options ...QueryOpt) ([]LabelModel, error) {
	tx, err := ds.db.Begin()
	if err != nil {
		return []LabelModel{}, errors.Wrap(err, "beginning transaction")
	}

	var labelModels []LabelModel
	err = tx.Wrap(func() error {
		// Setup query with optional parameters
		q := NewQueryConfig(ds, tx, "", "model", nil)
		q.dataset = q.dataset.LeftOuterJoin(goqu.I("label").As("label"), goqu.On(
			goqu.I("model.id").Eq(goqu.I("label.model_id")),
		))
		for _, option := range options {
			if err := option(q); err != nil {
				return errors.Wrap(err, "setting up query options")
			}
		}
		if debug {
			log.Println(q.dataset.ToSql())
		}
		// Execute query

		rows, err := getRows(tx, q.dataset)
		if err != nil {
			return err
		}
		defer rows.Close()
		labelModels, err = scanLabelModels(rows)
		if err != nil {
			return err
		}
		for _, post := range q.postFunc {
			if err := post(q); err != nil {
				return errors.Wrap(err, "running post functions")
			}
		}
		return nil
	})
	return labelModels, errors.Wrap(err, "committing transaction")
}

func (ds *Datastore) ReadLabelModel(options ...QueryOpt) (LabelModel, bool, error) {
	tx, err := ds.db.Begin()
	if err != nil {
		return LabelModel{}, false, errors.Wrap(err, "beginning transaction")
	}

	var labelModel LabelModel
	var exists bool
	err = tx.Wrap(func() error {
		// Setup query with optional parameters
		q := NewQueryConfig(ds, tx, "", "model", nil)
		q.dataset = q.dataset.LeftOuterJoin(goqu.I("label").As("label"), goqu.On(
			goqu.I("model.id").Eq(goqu.I("label.model_id")),
		))
		for _, option := range options {
			if err := option(q); err != nil {
				return errors.Wrap(err, "setting up query options")
			}
		}
		if debug {
			log.Println(q.dataset.ToSql())
		}
		// Execute query
		row, err := getRow(tx, q.dataset)
		if err != nil {
			return err
		}
		labelModel, err = scanLabelModel(row)
		if err == sql.ErrNoRows {
			return nil
		} else if err == nil {
			exists = true
		}
		if err != nil {
			return err
		}
		for _, post := range q.postFunc {
			if err := post(q); err != nil {
				return errors.Wrap(err, "running post functions")
			}
		}

		return nil
	})

	return labelModel, exists, errors.Wrap(err, "committing transaction")
}

type BinomialModel struct {
	LabelModel
	Binomial binomialModel
}

func scanBinomialModel(r *sql.Row) (BinomialModel, error) {
	var s BinomialModel
	if err := r.Scan(
		&s.Id,
		&s.ProjectId,
		&s.Name,
		&s.ClusterId,
		&s.ClusterName,
		&s.ModelKey,
		&s.Algorithm,
		&s.ModelCategory,
		&s.DatasetName,
		&s.ResponseColumn,
		&s.LogicalName,
		&s.Location,
		&s.ModelObjectType,
		&s.MaxRunTime,
		&s.Schema,
		&s.SchemaVersion,
		&s.Created,
		&s.Label.Id,
		&s.Label.ProjectId,
		&s.Label.ModelId,
		&s.Label.Name,
		&s.Label.Description,
		&s.Label.Created,
		&s.Binomial.ModelId,
		&s.Binomial.Mse,
		&s.Binomial.RSquared,
		&s.Binomial.Logloss,
		&s.Binomial.Auc,
		&s.Binomial.Gini,
	); err != nil {
		return BinomialModel{}, err
	}
	return s, nil
}

func scanBinomialModels(rs *sql.Rows) ([]BinomialModel, error) {
	structs := make([]BinomialModel, 0, 16)
	var err error
	for rs.Next() {
		var s BinomialModel
		if err = rs.Scan(
			&s.Id,
			&s.ProjectId,
			&s.Name,
			&s.ClusterId,
			&s.ClusterName,
			&s.ModelKey,
			&s.Algorithm,
			&s.ModelCategory,
			&s.DatasetName,
			&s.ResponseColumn,
			&s.LogicalName,
			&s.Location,
			&s.ModelObjectType,
			&s.MaxRunTime,
			&s.Schema,
			&s.SchemaVersion,
			&s.Created,
			&s.Label.Id,
			&s.Label.ProjectId,
			&s.Label.ModelId,
			&s.Label.Name,
			&s.Label.Description,
			&s.Label.Created,
			&s.Binomial.ModelId,
			&s.Binomial.Mse,
			&s.Binomial.RSquared,
			&s.Binomial.Logloss,
			&s.Binomial.Auc,
			&s.Binomial.Gini,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func (ds *Datastore) ReadBinomialModels(options ...QueryOpt) ([]BinomialModel, error) {
	tx, err := ds.db.Begin()
	if err != nil {
		return []BinomialModel{}, errors.Wrap(err, "beginning transaction")
	}

	var binomialModels []BinomialModel
	err = tx.Wrap(func() error {
		// Setup query with optional parameters
		q := NewQueryConfig(ds, tx, "", "model", nil)
		q.dataset = q.dataset.LeftOuterJoin(goqu.I("label").As("label"), goqu.On(
			goqu.I("model.id").Eq(goqu.I("label.model_id")),
		)).InnerJoin(goqu.I("binomial_model"), goqu.On(
			goqu.I("model.id").Eq(goqu.I("binomial_model.model_id")),
		))
		for _, option := range options {
			if err := option(q); err != nil {
				return errors.Wrap(err, "setting up query options")
			}
		}
		if debug {
			log.Println(q.dataset.ToSql())
		}
		// Execute query

		rows, err := getRows(tx, q.dataset)
		if err != nil {
			return err
		}
		defer rows.Close()
		binomialModels, err = scanBinomialModels(rows)
		if err != nil {
			return err
		}
		for _, post := range q.postFunc {
			if err := post(q); err != nil {
				return errors.Wrap(err, "running post functions")
			}
		}
		return nil
	})
	return binomialModels, errors.Wrap(err, "committing transaction")
}

func (ds *Datastore) ReadBinomialModel(options ...QueryOpt) (BinomialModel, bool, error) {
	tx, err := ds.db.Begin()
	if err != nil {
		return BinomialModel{}, false, errors.Wrap(err, "beginning transaction")
	}

	var binomialModel BinomialModel
	var exists bool
	err = tx.Wrap(func() error {
		// Setup query with optional parameters
		q := NewQueryConfig(ds, tx, "", "model", nil)
		q.dataset = q.dataset.LeftOuterJoin(goqu.I("label").As("label"), goqu.On(
			goqu.I("model.id").Eq(goqu.I("label.model_id")),
		)).InnerJoin(goqu.I("binomial_model"), goqu.On(
			goqu.I("model.id").Eq(goqu.I("binomial_model.model_id")),
		))
		for _, option := range options {
			if err := option(q); err != nil {
				return errors.Wrap(err, "setting up query options")
			}
		}
		if debug {
			log.Println(q.dataset.ToSql())
		}
		// Execute query
		row, err := getRow(tx, q.dataset)
		if err != nil {
			return err
		}
		binomialModel, err = scanBinomialModel(row)
		if err == sql.ErrNoRows {
			return nil
		} else if err == nil {
			exists = true
		}
		if err != nil {
			return err
		}
		for _, post := range q.postFunc {
			if err := post(q); err != nil {
				return errors.Wrap(err, "running post functions")
			}
		}

		return nil
	})

	return binomialModel, exists, errors.Wrap(err, "committing transaction")
}

type MultinomialModel struct {
	LabelModel
	Multinomial multinomialModel
}

func scanMultinomialModel(r *sql.Row) (MultinomialModel, error) {
	var s MultinomialModel
	if err := r.Scan(
		&s.Id,
		&s.ProjectId,
		&s.Name,
		&s.ClusterId,
		&s.ClusterName,
		&s.ModelKey,
		&s.Algorithm,
		&s.ModelCategory,
		&s.DatasetName,
		&s.ResponseColumn,
		&s.LogicalName,
		&s.Location,
		&s.ModelObjectType,
		&s.MaxRunTime,
		&s.Schema,
		&s.SchemaVersion,
		&s.Created,
		&s.Label.Id,
		&s.Label.ProjectId,
		&s.Label.ModelId,
		&s.Label.Name,
		&s.Label.Description,
		&s.Label.Created,
		&s.Multinomial.ModelId,
		&s.Multinomial.Mse,
		&s.Multinomial.RSquared,
		&s.Multinomial.Logloss,
	); err != nil {
		return MultinomialModel{}, err
	}
	return s, nil
}

func scanMultinomialModels(rs *sql.Rows) ([]MultinomialModel, error) {
	structs := make([]MultinomialModel, 0, 16)
	var err error
	for rs.Next() {
		var s MultinomialModel
		if err = rs.Scan(
			&s.Id,
			&s.ProjectId,
			&s.Name,
			&s.ClusterId,
			&s.ClusterName,
			&s.ModelKey,
			&s.Algorithm,
			&s.ModelCategory,
			&s.DatasetName,
			&s.ResponseColumn,
			&s.LogicalName,
			&s.Location,
			&s.ModelObjectType,
			&s.MaxRunTime,
			&s.Schema,
			&s.SchemaVersion,
			&s.Created,
			&s.Label.Id,
			&s.Label.ProjectId,
			&s.Label.ModelId,
			&s.Label.Name,
			&s.Label.Description,
			&s.Label.Created,
			&s.Multinomial.ModelId,
			&s.Multinomial.Mse,
			&s.Multinomial.RSquared,
			&s.Multinomial.Logloss,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func (ds *Datastore) ReadMultinomialModels(options ...QueryOpt) ([]MultinomialModel, error) {
	tx, err := ds.db.Begin()
	if err != nil {
		return []MultinomialModel{}, errors.Wrap(err, "beginning transaction")
	}

	var multinomialModels []MultinomialModel
	err = tx.Wrap(func() error {
		// Setup query with optional parameters
		q := NewQueryConfig(ds, tx, "", "model", nil)
		q.dataset = q.dataset.LeftOuterJoin(goqu.I("label").As("label"), goqu.On(
			goqu.I("model.id").Eq(goqu.I("label.model_id")),
		)).InnerJoin(goqu.I("multinomial_model"), goqu.On(
			goqu.I("model.id").Eq(goqu.I("multinomial_model.model_id")),
		))
		for _, option := range options {
			if err := option(q); err != nil {
				return errors.Wrap(err, "setting up query options")
			}
		}
		if debug {
			log.Println(q.dataset.ToSql())
		}
		// Execute query

		rows, err := getRows(tx, q.dataset)
		if err != nil {
			return err
		}
		defer rows.Close()
		multinomialModels, err = scanMultinomialModels(rows)
		if err != nil {
			return err
		}
		for _, post := range q.postFunc {
			if err := post(q); err != nil {
				return errors.Wrap(err, "running post functions")
			}
		}
		return nil
	})
	return multinomialModels, errors.Wrap(err, "committing transaction")
}

func (ds *Datastore) ReadMultinomialModel(options ...QueryOpt) (MultinomialModel, bool, error) {
	tx, err := ds.db.Begin()
	if err != nil {
		return MultinomialModel{}, false, errors.Wrap(err, "beginning transaction")
	}

	var multinomialModel MultinomialModel
	var exists bool
	err = tx.Wrap(func() error {
		// Setup query with optional parameters
		q := NewQueryConfig(ds, tx, "", "model", nil)
		q.dataset = q.dataset.LeftOuterJoin(goqu.I("label").As("label"), goqu.On(
			goqu.I("model.id").Eq(goqu.I("label.model_id")),
		)).InnerJoin(goqu.I("multinomial_model"), goqu.On(
			goqu.I("model.id").Eq(goqu.I("multinomial_model.model_id")),
		))
		for _, option := range options {
			if err := option(q); err != nil {
				return errors.Wrap(err, "setting up query options")
			}
		}
		if debug {
			log.Println(q.dataset.ToSql())
		}
		// Execute query
		row, err := getRow(tx, q.dataset)
		if err != nil {
			return err
		}
		multinomialModel, err = scanMultinomialModel(row)
		if err == sql.ErrNoRows {
			return nil
		} else if err == nil {
			exists = true
		}
		if err != nil {
			return err
		}
		for _, post := range q.postFunc {
			if err := post(q); err != nil {
				return errors.Wrap(err, "running post functions")
			}
		}

		return nil
	})

	return multinomialModel, exists, errors.Wrap(err, "committing transaction")
}

type RegressionModel struct {
	LabelModel
	Regression regressionModel
}

func scanRegressionModel(r *sql.Row) (RegressionModel, error) {
	var s RegressionModel
	if err := r.Scan(
		&s.Id,
		&s.ProjectId,
		&s.Name,
		&s.ClusterId,
		&s.ClusterName,
		&s.ModelKey,
		&s.Algorithm,
		&s.ModelCategory,
		&s.DatasetName,
		&s.ResponseColumn,
		&s.LogicalName,
		&s.Location,
		&s.ModelObjectType,
		&s.MaxRunTime,
		&s.Schema,
		&s.SchemaVersion,
		&s.Created,
		&s.Label.Id,
		&s.Label.ProjectId,
		&s.Label.ModelId,
		&s.Label.Name,
		&s.Label.Description,
		&s.Label.Created,
		&s.Regression.ModelId,
		&s.Regression.Mse,
		&s.Regression.RSquared,
		&s.Regression.MeanResidualDeviance,
	); err != nil {
		return RegressionModel{}, err
	}
	return s, nil
}

func scanRegressionModels(rs *sql.Rows) ([]RegressionModel, error) {
	structs := make([]RegressionModel, 0, 16)
	var err error
	for rs.Next() {
		var s RegressionModel
		if err = rs.Scan(
			&s.Id,
			&s.ProjectId,
			&s.Name,
			&s.ClusterId,
			&s.ClusterName,
			&s.ModelKey,
			&s.Algorithm,
			&s.ModelCategory,
			&s.DatasetName,
			&s.ResponseColumn,
			&s.LogicalName,
			&s.Location,
			&s.ModelObjectType,
			&s.MaxRunTime,
			&s.Schema,
			&s.SchemaVersion,
			&s.Created,
			&s.Label.Id,
			&s.Label.ProjectId,
			&s.Label.ModelId,
			&s.Label.Name,
			&s.Label.Description,
			&s.Label.Created,
			&s.Regression.ModelId,
			&s.Regression.Mse,
			&s.Regression.RSquared,
			&s.Regression.MeanResidualDeviance,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func (ds *Datastore) ReadRegressionModels(options ...QueryOpt) ([]RegressionModel, error) {
	tx, err := ds.db.Begin()
	if err != nil {
		return []RegressionModel{}, errors.Wrap(err, "beginning transaction")
	}

	var regressionModels []RegressionModel
	err = tx.Wrap(func() error {
		// Setup query with optional parameters
		q := NewQueryConfig(ds, tx, "", "model", nil)
		q.dataset = q.dataset.LeftOuterJoin(goqu.I("label").As("label"), goqu.On(
			goqu.I("model.id").Eq(goqu.I("label.model_id")),
		)).InnerJoin(goqu.I("regression_model"), goqu.On(
			goqu.I("model.id").Eq(goqu.I("regression_model.model_id")),
		))
		for _, option := range options {
			if err := option(q); err != nil {
				return errors.Wrap(err, "setting up query options")
			}
		}
		if debug {
			log.Println(q.dataset.ToSql())
		}
		// Execute query

		rows, err := getRows(tx, q.dataset)
		if err != nil {
			return err
		}
		defer rows.Close()
		regressionModels, err = scanRegressionModels(rows)
		if err != nil {
			return err
		}
		for _, post := range q.postFunc {
			if err := post(q); err != nil {
				return errors.Wrap(err, "running post functions")
			}
		}
		return nil
	})
	return regressionModels, errors.Wrap(err, "committing transaction")
}

func (ds *Datastore) ReadRegressionModel(options ...QueryOpt) (RegressionModel, bool, error) {
	tx, err := ds.db.Begin()
	if err != nil {
		return RegressionModel{}, false, errors.Wrap(err, "beginning transaction")
	}

	var regressionModel RegressionModel
	var exists bool
	err = tx.Wrap(func() error {
		// Setup query with optional parameters
		q := NewQueryConfig(ds, tx, "", "model", nil)
		q.dataset = q.dataset.LeftOuterJoin(goqu.I("label").As("label"), goqu.On(
			goqu.I("model.id").Eq(goqu.I("label.model_id")),
		)).InnerJoin(goqu.I("regression_model"), goqu.On(
			goqu.I("model.id").Eq(goqu.I("regression_model.model_id")),
		))
		for _, option := range options {
			if err := option(q); err != nil {
				return errors.Wrap(err, "setting up query options")
			}
		}
		if debug {
			log.Println(q.dataset.ToSql())
		}
		// Execute query
		row, err := getRow(tx, q.dataset)
		if err != nil {
			return err
		}
		regressionModel, err = scanRegressionModel(row)
		if err == sql.ErrNoRows {
			return nil
		} else if err == nil {
			exists = true
		}
		if err != nil {
			return err
		}
		for _, post := range q.postFunc {
			if err := post(q); err != nil {
				return errors.Wrap(err, "running post functions")
			}
		}

		return nil
	})

	return regressionModel, exists, errors.Wrap(err, "committing transaction")
}

type EntityPrivilege struct {
	Privilege
	Workgroup Workgroup
}

func scanEntityPrivilege(r *sql.Row) (EntityPrivilege, error) {
	var s EntityPrivilege
	if err := r.Scan(
		&s.Type,
		&s.IdentityId,
		&s.WorkgroupId,
		&s.EntityType,
		&s.EntityId,
		&s.Workgroup.Id,
		&s.Workgroup.Type,
		&s.Workgroup.Name,
		&s.Workgroup.Description,
		&s.Workgroup.Created,
	); err != nil {
		return EntityPrivilege{}, err
	}
	return s, nil
}

func scanEntityPrivileges(rs *sql.Rows) ([]EntityPrivilege, error) {
	structs := make([]EntityPrivilege, 0, 16)
	var err error
	for rs.Next() {
		var s EntityPrivilege
		if err = rs.Scan(
			&s.Type,
			&s.IdentityId,
			&s.WorkgroupId,
			&s.EntityType,
			&s.EntityId,
			&s.Workgroup.Id,
			&s.Workgroup.Type,
			&s.Workgroup.Name,
			&s.Workgroup.Description,
			&s.Workgroup.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func (ds *Datastore) ReadEntityPrivileges(options ...QueryOpt) ([]EntityPrivilege, error) {
	tx, err := ds.db.Begin()
	if err != nil {
		return []EntityPrivilege{}, errors.Wrap(err, "beginning transaction")
	}

	var entityPrivileges []EntityPrivilege
	err = tx.Wrap(func() error {
		// Setup query with optional parameters
		q := NewQueryConfig(ds, tx, "", "privilege", nil)
		q.dataset = q.dataset.LeftOuterJoin(goqu.I("workgroup"), goqu.On(
			goqu.I("privilege.workgroup_id").Eq(goqu.I("workgroup.id")),
		))
		for _, option := range options {
			if err := option(q); err != nil {
				return errors.Wrap(err, "setting up query options")
			}
		}
		if debug {
			log.Println(q.dataset.ToSql())
		}
		// Execute query

		rows, err := getRows(tx, q.dataset)
		if err != nil {
			return err
		}
		defer rows.Close()
		entityPrivileges, err = scanEntityPrivileges(rows)
		if err != nil {
			return err
		}
		for _, post := range q.postFunc {
			if err := post(q); err != nil {
				return errors.Wrap(err, "running post functions")
			}
		}
		return nil
	})
	return entityPrivileges, errors.Wrap(err, "committing transaction")
}

func (ds *Datastore) ReadEntityPrivilege(options ...QueryOpt) (EntityPrivilege, bool, error) {
	tx, err := ds.db.Begin()
	if err != nil {
		return EntityPrivilege{}, false, errors.Wrap(err, "beginning transaction")
	}

	var entityPrivilege EntityPrivilege
	var exists bool
	err = tx.Wrap(func() error {
		// Setup query with optional parameters
		q := NewQueryConfig(ds, tx, "", "privilege", nil)
		q.dataset = q.dataset.LeftOuterJoin(goqu.I("workgroup"), goqu.On(
			goqu.I("privilege.workgroup_id").Eq(goqu.I("workgroup.id")),
		))
		for _, option := range options {
			if err := option(q); err != nil {
				return errors.Wrap(err, "setting up query options")
			}
		}
		if debug {
			log.Println(q.dataset.ToSql())
		}
		// Execute query
		row, err := getRow(tx, q.dataset)
		if err != nil {
			return err
		}
		entityPrivilege, err = scanEntityPrivilege(row)
		if err == sql.ErrNoRows {
			return nil
		} else if err == nil {
			exists = true
		}
		if err != nil {
			return err
		}
		for _, post := range q.postFunc {
			if err := post(q); err != nil {
				return errors.Wrap(err, "running post functions")
			}
		}

		return nil
	})

	return entityPrivilege, exists, errors.Wrap(err, "committing transaction")
}
