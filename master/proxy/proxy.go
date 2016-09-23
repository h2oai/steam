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

package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"sync"

	"github.com/h2oai/steam/master/az"
	"github.com/h2oai/steam/master/data"
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
		clusterId := r.URL.Query().Get("cluster_id")
		if r.URL.Path == "/flow/" && clusterId != "" {
			r.Header.Set("X-Cluster", clusterId)
			clusterHeader = clusterId
		} else {
			http.Error(w, "Cluster requests via Steam requires a valid X-Cluster HTTP header", http.StatusBadRequest)
			return
		}
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
