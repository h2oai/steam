package sql

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
	Superuser = "superuser"
)

var CLUSTER_TYPES = []string{
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

var ENTITY_TYPES = []string{
	"cluster",
	"engine",
	"identity",
	"permission",
	"privilege",
	"project",
	"role",
	"service",
	"workgroup",
}

type entityTypeKeys struct {
	Cluster           int64
	ClusterYarnDetail int64
	Engine            int64
	Identity          int64
	Permission        int64
	Privilege         int64
	Project           int64
	Role              int64
	Service           int64
	Workgroup         int64
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
		Cluster:           m["cluster"],
		ClusterYarnDetail: m["clusterYarnDetail"],
		Engine:            m["engine"],
		Identity:          m["identity"],
		Permission:        m["permission"],
		Privilege:         m["privilege"],
		Project:           m["project"],
		Role:              m["role"],
		Service:           m["service"],
		Workgroup:         m["workgroup"],
	}
}

var PERMISSIONS = []struct{ code, desc string }{
	struct{ code, desc string }{"ManageRole", "Manage role"},
	struct{ code, desc string }{"ViewRole", "View role"},
	struct{ code, desc string }{"ManageWorkgroup", "Manage workgroup"},
	struct{ code, desc string }{"ViewWorkgroup", "View workgroup"},
	struct{ code, desc string }{"ManageIdentity", "Manage identity"},
	struct{ code, desc string }{"ViewIdentity", "View identity"},
	struct{ code, desc string }{"ManageEngine", "Manage engine"},
	struct{ code, desc string }{"ViewEngine", "View engine"},
	struct{ code, desc string }{"ManageCluster", "Manage cluster"},
	struct{ code, desc string }{"ViewCluster", "View cluster"},
	struct{ code, desc string }{"ManageProject", "Manage project"},
	struct{ code, desc string }{"ViewProject", "View project"},
	struct{ code, desc string }{"ManageDatasource", "Manage datasource"},
	struct{ code, desc string }{"ViewDatasource", "View datasource"},
	struct{ code, desc string }{"ManageDataset", "Manage dataset"},
	struct{ code, desc string }{"ViewDataset", "View dataset"},
	struct{ code, desc string }{"ManageModel", "Manage model"},
	struct{ code, desc string }{"ViewModel", "View model"},
	struct{ code, desc string }{"ManageLabel", "Manage label"},
	struct{ code, desc string }{"ViewLabel", "View label"},
	struct{ code, desc string }{"ManageService", "Manage service"},
	struct{ code, desc string }{"ViewService", "View service"},
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
}

func toPermissionMap(permissions []Permission) map[int64]string {
	m := make(map[int64]string)
	for _, p := range permissions {
		m[p.Id] = p.Description
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
	}
}

var STATES = []string{
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
