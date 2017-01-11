//go:generate crudr $GOFILE
//go:generate scaneo $GOFILE

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
	"time"

	"github.com/lib/pq"
)

type binomialModel struct {
	ModelId  int64   `db:"model_id,arg,pk"`
	Mse      float64 `db:"mse,arg"`
	RSquared float64 `db:"r_squared,arg"`
	Logloss  float64 `db:"logloss,arg"`
	Auc      float64 `db:"auc,arg"`
	Gini     float64 `db:"gini,arg"`
}

type Cluster struct {
	Id            int64          `db:"id,pk"`
	Name          string         `db:"name,arg"`
	ContextPath   sql.NullString `db:"context_path"`
	ClusterTypeId int64          `db:"type_id,arg"`
	DetailId      sql.NullInt64  `db:"detail_id"`
	Address       sql.NullString `db:"address"`
	Token         sql.NullString `db:"token"`
	State         string         `db:"state,def=States.Starting"`
	Created       time.Time      `db:"created,def=time.Now()"`
}

type clusterType struct {
	Id   int64  `db:"id,pk"`
	Name string `db:"name,arg"`
}

type ClusterYarnDetail struct {
	Id            int64  `db:"id,pk"`
	EngineId      int64  `db:"engine_id,arg"`
	Size          int64  `db:"size,arg"`
	ApplicationId string `db:"application_id,arg"`
	Memory        string `db:"memory,arg"`
	OutputDir     string `db:"output_dir,arg"`
}

type Engine struct {
	Id       int64     `db:"id,pk"`
	Name     string    `db:"name,arg"`
	Location string    `db:"location,arg"`
	Created  time.Time `db:"created,def=time.Now()"`
}

type entityType struct {
	Id   int64  `db:"id,pk"`
	Name string `db:"name,arg"`
}

type History struct {
	Id           int64          `db:"id,pk"`
	Action       string         `db:"action,arg"`
	IdentityId   int64          `db:"identity_id,arg"`
	EntityTypeId int64          `db:"entity_type_id,arg"`
	EntityId     int64          `db:"entity_id,arg"`
	Description  sql.NullString `db:"description"`
	Created      time.Time      `db:"created,def=time.Now()"`
}

type Identity struct {
	Id          int64          `db:"id,pk"`
	Name        string         `db:"name,arg"`
	Password    sql.NullString `db:"password"`
	WorkgroupId sql.NullInt64  `db:"workgroup_id"`
	IsActive    bool           `db:"is_active,def=1"`
	LastLogin   pq.NullTime    `db:"last_login"`
	Created     time.Time      `db:"created,def=time.Now()"`
}

type identityRole struct {
	IdentityId int64 `db:"identity_id"`
	RoleId     int64 `db:"role_id"`
}

type identityWorkgroup struct {
	IdentityId  int64 `db:"identity_id,arg"`
	WorkgroupId int64 `db:"workgroup_id,arg"`
}

type Label struct {
	Id          sql.NullInt64  `db:"id,pk"`
	ProjectId   sql.NullInt64  `db:"project_id,arg"`
	ModelId     sql.NullInt64  `db:"model_id"`
	Name        sql.NullString `db:"name,arg"`
	Description sql.NullString `db:"description,arg"`
	Created     pq.NullTime    `db:"created,def=time.Now()"`
}

type meta struct {
	Id    int64  `db:"id,pk"`
	Key   string `db:"key,arg"`
	Value string `db:"value"`
}

type modelCategory struct {
	Id   int64  `db:"id,pk"`
	Name string `db:"name,arg"`
}

type Model struct {
	Id              int64          `db:"id,pk"`
	ProjectId       int64          `db:"project_id"`
	Name            string         `db:"name,arg"`
	ClusterId       int64          `db:"cluster_id"`
	ClusterName     string         `db:"cluster_name"`
	ModelKey        string         `db:"model_key,arg"`
	Algorithm       string         `db:"algorithm,arg"`
	ModelCategory   string         `db:"model_category,arg"`
	DatasetName     sql.NullString `db:"dataset_name"`
	ResponseColumn  string         `db:"response_column_name,arg"`
	LogicalName     sql.NullString `db:"logical_name"`
	Location        sql.NullString `db:"location"`
	ModelObjectType sql.NullString `db:"model_object_type"`
	MaxRunTime      sql.NullInt64  `db:"max_run_time"`
	Schema          sql.NullString `db:"schema"`
	SchemaVersion   sql.NullString `db:"schema_version"`
	Created         time.Time      `db:"created,def=time.Now()"`
}

type multinomialModel struct {
	ModelId  int64   `db:"model_id,arg,pk"`
	Mse      float64 `db:"mse,arg"`
	RSquared float64 `db:"r_squared,arg"`
	Logloss  float64 `db:"logloss,arg"`
}

type Permission struct {
	Id          int64  `db:"id,pk"`
	Code        string `db:"code,arg"`
	Description string `db:"description,arg"`
}

type Privilege struct {
	Type        string `db:"privilege_type,arg"`
	WorkgroupId int64  `db:"workgroup_id,arg"`
	EntityType  int64  `db:"entity_type_id,arg"`
	EntityId    int64  `db:"entity_id,arg"`
}

type Project struct {
	Id            int64     `db:"id,pk"`
	Name          string    `db:"name,arg"`
	Description   string    `db:"description,arg"`
	ModelCategory string    `db:"model_category,arg"`
	Created       time.Time `db:"created,def=time.Now()"`
}

type regressionModel struct {
	ModelId              int64   `db:"model_id,arg,pk"`
	Mse                  float64 `db:"mse,arg"`
	RSquared             float64 `db:"r_squared,arg"`
	MeanResidualDeviance float64 `db:"mean_residual_deviance,arg"`
}

type Role struct {
	Id          int64          `db:"id,pk"`
	Name        string         `db:"name,arg"`
	Description sql.NullString `db:"description"`
	Created     time.Time      `db:"created,def=time.Now()"`
}

type rolePermission struct {
	RoleId       int64 `db:"role_id,arg"`
	PermissionId int64 `db:"permission_id,arg"`
}

type state struct {
	Id   int64  `db:"id,pk"`
	Name string `db:"name,arg"`
}

type Security struct {
	Id    int64  `db:"id,pk"`
	Key   string `db:"key,arg"`
	Value string `db:"value,arg"`
}

type Service struct {
	Id        int64          `db:"id,pk"`
	ProjectId int64          `db:"project_id,arg"`
	ModelId   int64          `db:"model_id,arg"`
	Name      string         `db:"name,arg"`
	Host      sql.NullString `db:"host"`
	Port      sql.NullInt64  `db:"port"`
	ProcessId sql.NullInt64  `db:"process_id"`
	State     string         `db:"state,def=States.Starting"`
	Created   time.Time      `db:"created,def=time.Now()"`
}

type Workgroup struct {
	Id          int64          `db:"id,pk"`
	Type        string         `db:"type,arg"`
	Name        string         `db:"name,arg"`
	Description sql.NullString `db:"description"`
	Created     time.Time      `db:"created,def=time.Now()"`
}
