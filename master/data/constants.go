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

const (
	// --- History ---
	CreateOp  = "create"
	UpdateOp  = "update"
	DeleteOp  = "delete"
	EnableOp  = "enable"
	DisableOp = "disable"
	ShareOp   = "share"
	UnshareOp = "unshare"
	LinkOp    = "link"
	UnlinkOp  = "unlink"

	// --- Privilege ---
	Owns = "owns"
	Edit = "edit"
	View = "view"

	// --- Role ---

	AdminRN = "admin"

	// -- Workgroup ---
	AdminWG = "user:admin"

	// --- Security --
	LocalAuth = "local"
	LDAPAuth  = "ldap"
)

var cluster_types_list = []string{
	"external",
	"yarn",
}

type clusterTypeKeys struct {
	External int64
	Yarn     int64
}

func newClusterTypeKeys(clusterTypes []clusterType) clusterTypeKeys {
	m := make(map[string]int64)
	for _, c := range clusterTypes {
		m[c.Name] = c.Id
	}

	return clusterTypeKeys{
		External: m["external"],
		Yarn:     m["yarn"],
	}
}

type entity_types struct {
	Cluster    string
	Engine     string
	Identity   string
	Label      string
	Keytab     string
	Model      string
	Permission string
	Privilege  string
	Project    string
	Role       string
	Service    string
	Workgroup  string
}

func (e *entity_types) init() {
	*e = entity_types{
		"cluster",
		"engine",
		"identity",
		"label",
		"keytab",
		"model",
		"permission",
		"privilege",
		"project",
		"role",
		"service",
		"workgroup",
	}
}

var entity_types_list = []string{
	"cluster",
	"engine",
	"identity",
	"label",
	"keytab",
	"model",
	"permission",
	"privilege",
	"project",
	"role",
	"service",
	"workgroup",
}

type entityTypeKeys struct {
	BinomialModel     int64
	Cluster           int64
	ClusterYarnDetail int64
	Engine            int64
	Label             int64
	Keytab            int64
	Model             int64
	MultinomialModel  int64
	Identity          int64
	Permission        int64
	Privilege         int64
	Project           int64
	Role              int64
	RegressionModel   int64
	Service           int64
	Workgroup         int64
}

func toEntityId(ds *Datastore, name string) int64 {
	switch name {
	case "cluster":
		return ds.EntityType.Cluster
	case "engine":
		return ds.EntityType.Engine
	case "label":
		return ds.EntityType.Label
	case "keytab":
		return ds.EntityType.Keytab
	case "model":
		return ds.EntityType.Model
	case "identity":
		return ds.EntityType.Identity
	case "permission":
		return ds.EntityType.Permission
	case "privilege":
		return ds.EntityType.Privilege
	case "project":
		return ds.EntityType.Project
	case "role":
		return ds.EntityType.Role
	case "service":
		return ds.EntityType.Service
	case "workgroup":
		return ds.EntityType.Workgroup
	}
	return -1
}

func toEntityTypeMap(entityTypes []entityType) map[int64]string {
	m := make(map[int64]string)
	for _, e := range entityTypes {
		m[e.Id] = e.Name
	}
	return m
}

func newEntityTypeKeys(entityTypes []entityType) entityTypeKeys {
	m := make(map[string]int64)
	for _, e := range entityTypes {
		m[e.Name] = e.Id
	}

	return entityTypeKeys{
		Cluster:   m["cluster"],
		Engine:    m["engine"],
		Label:     m["label"],
		Keytab:    m["keytab"],
		Model:     m["model"],
		Identity:  m["identity"],
		Project:   m["project"],
		Role:      m["role"],
		Service:   m["service"],
		Workgroup: m["workgroup"],
	}
}

type permission_map struct {
	Code string
	Desc string
}

type permission_string_map struct {
	Id   int64
	Desc string
}

var permissions_list = []permission_map{
	permission_map{"ManageRole", "Manage role"},
	permission_map{"ViewRole", "View role"},
	permission_map{"ManageWorkgroup", "Manage workgroup"},
	permission_map{"ViewWorkgroup", "View workgroup"},
	permission_map{"ManageIdentity", "Manage identity"},
	permission_map{"ViewIdentity", "View identity"},
	permission_map{"ManageEngine", "Manage engine"},
	permission_map{"ViewEngine", "View engine"},
	permission_map{"ManageCluster", "Manage cluster"},
	permission_map{"ViewCluster", "View cluster"},
	permission_map{"ManageProject", "Manage project"},
	permission_map{"ViewProject", "View project"},
	permission_map{"ManageDatasource", "Manage datasource"},
	permission_map{"ViewDatasource", "View datasource"},
	permission_map{"ManageDataset", "Manage dataset"},
	permission_map{"ViewDataset", "View dataset"},
	permission_map{"ManageModel", "Manage model"},
	permission_map{"ViewModel", "View model"},
	permission_map{"ManageLabel", "Manage label"},
	permission_map{"ViewLabel", "View label"},
	permission_map{"ManageService", "Manage service"},
	permission_map{"ViewService", "View service"},
	permission_map{"ManageKeytab", "Manage keytab"},
	permission_map{"ViewKeytab", "View keytab"},
}

type permissionKeys struct {
	ManageRole       int64
	ViewRole         int64
	ManageWorkgroup  int64
	ViewWorkgroup    int64
	ManageIdentity   int64
	ViewIdentity     int64
	ManageEngine     int64
	ViewEngine       int64
	ManageCluster    int64
	ViewCluster      int64
	ManageProject    int64
	ViewProject      int64
	ManageDatasource int64
	ViewDatasource   int64
	ManageDataset    int64
	ViewDataset      int64
	ManageModel      int64
	ViewModel        int64
	ManageLabel      int64
	ViewLabel        int64
	ManageService    int64
	ViewService      int64
	ManageKeytab     int64
	ViewKeytab       int64
}

func toPermissionMap(permissions []Permission) map[int64]permission_map {
	m := make(map[int64]permission_map)
	for _, p := range permissions {
		pm := permission_map{Code: p.Code, Desc: p.Description}
		m[p.Id] = pm
	}
	return m
}

func toPermissionStringMap(permissions []Permission) map[string]permission_string_map {
	m := make(map[string]permission_string_map)
	for _, p := range permissions {
		pm := permission_string_map{Id: p.Id, Desc: p.Description}
		m[p.Code] = pm
	}
	return m
}

func newPermissionKeys(permissions []Permission) permissionKeys {
	m := make(map[string]int64)
	for _, p := range permissions {
		m[p.Code] = p.Id
	}

	return permissionKeys{
		ManageRole:       m["ManageRole"],
		ViewRole:         m["ViewRole"],
		ManageWorkgroup:  m["ManageWorkgroup"],
		ViewWorkgroup:    m["ViewWorkgroup"],
		ManageIdentity:   m["ManageIdentity"],
		ViewIdentity:     m["ViewIdentity"],
		ManageEngine:     m["ManageEngine"],
		ViewEngine:       m["ViewEngine"],
		ManageCluster:    m["ManageCluster"],
		ViewCluster:      m["ViewCluster"],
		ManageProject:    m["ManageProject"],
		ViewProject:      m["ViewProject"],
		ManageDatasource: m["ManageDatasource"],
		ViewDatasource:   m["ViewDatasource"],
		ManageDataset:    m["ManageDataset"],
		ViewDataset:      m["ViewDataset"],
		ManageModel:      m["ManageModel"],
		ViewModel:        m["ViewModel"],
		ManageLabel:      m["ManageLabel"],
		ViewLabel:        m["ViewLabel"],
		ManageService:    m["ManageService"],
		ViewService:      m["ViewService"],
		ManageKeytab:     m["ManageKeytab"],
		ViewKeytab:       m["ViewKeytab"],
	}
}

var states_list = []string{
	"idle",
	"starting",
	"started",
	"suspending",
	"suspended",
	"stopping",
	"stopped",
	"blocked",
	"disconnected",
	"failed",
	"completed",
}

type states struct {
	Idle         string
	Starting     string
	Started      string
	Suspending   string
	Suspended    string
	Stopping     string
	Stopped      string
	Blocked      string
	Disconnected string
	Failed       string
	Completed    string
}

func (s *states) init() {
	*s = states{
		"idle",
		"starting",
		"started",
		"suspending",
		"suspended",
		"stopping",
		"stopped",
		"blocked",
		"disconnected",
		"failed",
		"completed",
	}
}

type stateKeys struct {
	Idle         int64
	Starting     int64
	Started      int64
	Suspending   int64
	Suspended    int64
	Stopping     int64
	Stopped      int64
	Blocked      int64
	Disconnected int64
	Failed       int64
	Completed    int64
}

func newStateKeys(states []state) stateKeys {
	m := make(map[string]int64)
	for _, s := range states {
		m[s.Name] = s.Id
	}

	return stateKeys{
		Idle:         m["idle"],
		Starting:     m["starting"],
		Started:      m["started"],
		Suspending:   m["suspending"],
		Suspended:    m["suspended"],
		Stopping:     m["stopping"],
		Stopped:      m["stopped"],
		Blocked:      m["blocked"],
		Disconnected: m["disconnected"],
		Failed:       m["failed"],
		Completed:    m["completed"],
	}
}
