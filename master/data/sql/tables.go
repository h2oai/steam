//go:generate crudr $GOFILE
//go:generate scaneo $GOFILE

package sql

import (
	"database/sql"
	"time"
)

type Cluster struct {
	Id            int64          `db:"id,pk"`
	Name          string         `db:"name,arg"`
	ClusterTypeId int64          `db:"type_id,arg"`
	DetailId      sql.NullInt64  `db:"detail_id"`
	Address       sql.NullString `db:"address"`
	State         int64          `db:"state,def=q.state.Starting"`
	Created       time.Time      `db:"created,def=time.Now()"`
}

type ClusterType struct {
	Id   int64  `db:"id,pk"`
	Name string `db:"name,arg"`
}

type ClusterYarnDetail struct {
	Id            int64          `db:"id,pk"`
	EngineId      int64          `db:"engine_id,arg"`
	Size          sql.NullInt64  `db:"size"`
	ApplicationId sql.NullString `db:"application_id"`
	Memory        sql.NullString `db:"memory"`
	OutputDir     string         `db:"output_dir,arg"`
}

type Engine struct {
	Id       int64     `db:"id,pk"`
	Name     string    `db:"name,arg"`
	Location string    `db:"location,arg"`
	Created  time.Time `db:"created,def=time.Now()"`
}

type Entity_Type struct {
	Id   int    `db:"id,pk"`
	Name string `db:"name,arg"`
}

type History struct {
	Id           int            `db:"id,pk"`
	Action       string         `db:"action,arg"`
	IdentityId   int            `db:"identity_id,arg"`
	EntityTypeId int            `db:"entity_type_id,arg"`
	EntityId     int            `db:"entity_id,arg"`
	Description  sql.NullString `db:"description"`
	Created      string         `db:"created,def=time.Now()"`
}

type Identity struct {
	Id           int            `db:"id,pk"`
	Name         string         `db:"name,arg"`
	Password     sql.NullString `db:"password"`
	Workgroup_id int            `db:"workgroup_id,arg"`
	Is_active    bool           `db:"is_active,def=true"`
	Last_login   sql.NullInt64  `db:"last_login"`
	Created      time.Time      `db:"created,def=time.Now()"`
}

type IdentityWorkgroup struct {
	IdentityId  int64 `db:"identity_id,arg"`
	WorkgroupId int64 `db:"workgroup_id,arg"`
}

type ModelCategory struct {
	Id   int64  `db:"id,pk"`
	Name string `db:"name,arg"`
}

type Permission struct {
	Id          int64  `db:"id,pk"`
	Code        string `db:"code,arg"`
	Description string `db:"description,arg"`
}

type Privilege struct {
	Id          int64  `db:"id,pk"`
	Typ         string `db:"type,arg"`
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

type State struct {
	Id   int64  `db:"id,pk"`
	Name string `db:"name,arg"`
}

type Service struct {
	Id        int64          `db:"id,pk"`
	ProjectId int64          `db:"project_id,arg"`
	ModelId   int64          `db:"model_id,arg"`
	Name      string         `db:"name,arg"`
	Host      sql.NullString `db:"host"`
	Port      sql.NullString `db:"port"`
	ProcessId sql.NullInt64  `db:"process_id"`
	State     int64          `db:"state,def=q.state.Starting"`
	Created   time.Time      `db:"created,def=time.Now()"`
}
