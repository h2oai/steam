package proxy

import (
	"github.com/h2oai/steamY/master/az"
	"github.com/h2oai/steamY/master/data"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"sync"
)

type ReverseProxy struct {
	clusterId int64
	host      string
	proxy     *httputil.ReverseProxy
}

func NewReverseProxy(clusterId int64, host string) *ReverseProxy {
	return &ReverseProxy{
		clusterId,
		host,
		httputil.NewSingleHostReverseProxy(&url.URL{
			Scheme: "http",
			Host:   host,
		}),
	}
}

func NewProxy(az az.Az, ds *data.Datastore) *ProxyManager {
	return &ProxyManager{
		&sync.RWMutex{},
		make(map[int64]*ReverseProxy),
		az,
		ds,
	}
}

type ProxyManager struct {
	mu      *sync.RWMutex
	proxies map[int64]*ReverseProxy
	az      az.Az
	ds      *data.Datastore
}

func (pm *ProxyManager) getOrCreateReverseProxy(clusterId int64, host string) *ReverseProxy {
	pm.mu.RLock()
	rp, ok := pm.proxies[clusterId]
	pm.mu.RUnlock()

	if ok {
		return rp
	}

	rp = NewReverseProxy(clusterId, host)
	pm.mu.Lock()
	pm.proxies[clusterId] = rp
	pm.mu.Unlock()
	return rp
}

func (pm *ProxyManager) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// All proxy requests require a header with key=X-Cluster; value=cluster-id (in)

	clusterHeader := r.Header.Get("X-Cluster")
	if clusterHeader == "" {
		http.Error(w, "Cluster requests via Steam requires a valid X-Cluster HTTP header", http.StatusBadRequest)
		return
	}

	clusterId, err := strconv.ParseInt(clusterHeader, 10, 64)
	if err != nil {
		http.Error(w, "Invalid X-Cluster HTTP header, expected integer", http.StatusBadRequest)
		return
	}

	// Identify the principal

	pz, azerr := pm.az.Identify(r)
	if azerr != nil {
		http.Error(w, azerr.Error(), http.StatusForbidden)
	}

	// Check if principal is allowed to use clusters

	if err := pz.CheckPermission(pm.ds.Permissions.ViewCluster); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}

	// Read cluster from database.
	// This also checks if the principal has privileges to view this specific cluster.

	cluster, err := pm.ds.ReadCluster(pz, clusterId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}

	// Get existing proxy, or create one if missing.

	rp := pm.getOrCreateReverseProxy(clusterId, cluster.Address)

	// Forward

	rp.proxy.ServeHTTP(w, r)
}
