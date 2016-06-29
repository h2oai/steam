package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"sync"

	"github.com/h2oai/steamY/master/az"
	"github.com/h2oai/steamY/master/data"
)

type reverseProxy struct {
	clusterId int64
	host      string
	proxy     *httputil.ReverseProxy
}

func newReverseProxy(clusterId int64, host string) *reverseProxy {
	return &reverseProxy{
		clusterId,
		host,
		httputil.NewSingleHostReverseProxy(&url.URL{
			Scheme: "http",
			Host:   host,
		}),
	}
}

type ProxyHandler struct {
	mu      *sync.RWMutex
	proxies map[int64]*reverseProxy
	az      az.Az
	ds      *data.Datastore
}

func NewProxyHandler(az az.Az, ds *data.Datastore) *ProxyHandler {
	return &ProxyHandler{
		&sync.RWMutex{},
		make(map[int64]*reverseProxy),
		az,
		ds,
	}
}

func (pm *ProxyHandler) getOrCreateReverseProxy(clusterId int64, host string) *reverseProxy {
	pm.mu.RLock()
	rp, ok := pm.proxies[clusterId]
	pm.mu.RUnlock()

	if ok {
		return rp
	}

	rp = newReverseProxy(clusterId, host)
	pm.mu.Lock()
	pm.proxies[clusterId] = rp
	pm.mu.Unlock()
	return rp
}

func (pm *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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
		return
	}

	// Check if principal is allowed to use clusters

	if err := pz.CheckPermission(pm.ds.Permissions.ViewCluster); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	// Read cluster from database.
	// This also checks if the principal has privileges to view this specific cluster.

	cluster, err := pm.ds.ReadCluster(pz, clusterId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	// Get existing proxy, or create one if missing.

	rp := pm.getOrCreateReverseProxy(clusterId, cluster.Address)

	// Forward

	rp.proxy.ServeHTTP(w, r)
}
