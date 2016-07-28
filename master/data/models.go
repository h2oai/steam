//go:generate scaneo $GOFILE

package data

import (
	"time"

	"github.com/lib/pq"
)

type Meta struct {
	Id    int64
	Key   string
	Value string
}

type EntityHistory struct {
	IdentityId  int64
	Action      string
	Description string
	Created     time.Time
}

type Privilege struct {
	Type        string
	WorkgroupId int64
	EntityType  int64
	EntityId    int64
}

type EntityPrivilege struct {
	Type                 string
	WorkgroupId          int64
	WorkgroupName        string
	WorkgroupDescription string
}

type Permission struct {
	Id          int64
	Code        string
	Description string
}

type EntityType struct {
	Id   int64
	Name string
}

type Role struct {
	Id          int64
	Name        string
	Description string
	Created     time.Time
}

type Workgroup struct {
	Id          int64
	Type        string
	Name        string
	Description string
	Created     time.Time
}

type Identity struct {
	Id        int64
	Name      string
	IsActive  bool
	LastLogin pq.NullTime
	Created   time.Time
}

type IdentityAndPassword struct {
	Id          int64
	Name        string
	Password    string
	WorkgroupId int64
	IsActive    bool
	LastLogin   pq.NullTime
	Created     time.Time
}

type Engine struct {
	Id       int64
	Name     string
	Location string
	Created  time.Time
}

type ClusterType struct {
	Id   int64
	Name string
}

type Cluster struct {
	Id       int64
	Name     string
	TypeId   int64
	DetailId int64
	Address  string
	State    string
	Created  time.Time
}

type YarnCluster struct {
	Id            int64
	EngineId      int64
	Size          int64
	ApplicationId string
	Memory        string
	Username      string
	OutputDir     string
}

type Project struct {
	Id          int64
	Name        string
	Description string
	Created     time.Time
}

type Datasource struct {
	Id            int64
	ProjectId     int64
	Name          string
	Description   string
	Kind          string
	Configuration string
	Created       time.Time
}

type Dataset struct {
	Id                 int64
	DatasourceId       int64
	Name               string
	Description        string
	FrameName          string
	ResponseColumnName string
	Properties         string
	PropertiesVersion  string
	Created            time.Time
}

type Model struct {
	Id                  int64
	TrainingDatasetId   int64
	ValidationDatasetId int64
	Name                string
	ClusterName         string
	ModelKey            string
	Algorithm           string
	DatasetName         string
	ResponseColumnName  string
	LogicalName         string
	Location            string
	MaxRunTime          int64
	Metrics             string
	MetricsVersion      string
	Created             time.Time
}

type Label struct {
	Id          int64
	ProjectId   int64
	ModelId     int64
	Name        string
	Description string
	Created     time.Time
}

type Service struct {
	Id        int64
	ModelId   int64
	Address   string
	Port      int64
	ProcessId int64
	State     string
	Created   time.Time
}
