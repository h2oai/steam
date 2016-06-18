//go:generate scaneo $GOFILE

package data

import (
	"github.com/lib/pq"
	"time"
)

type Meta struct {
	Id    int64
	Key   string
	Value string
}

type Privilege struct {
	Type        string
	WorkgroupId int64
	EntityType  int64
	EntityId    int64
}

type Permission struct {
	Id          int64
	Code        int64
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

type Model struct {
	Id                 int64
	Name               string
	ClusterName        string
	Algorithm          string
	DatasetName        string
	ResponseColumnName string
	LogicalName        string
	Location           string
	MaxRunTime         int64
	Created            time.Time
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
