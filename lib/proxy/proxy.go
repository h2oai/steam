package proxy

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

// RProxy
// A map of single host reverse proxies mapped to their corresponding cluster ID
// and their corresponding authorization pieces. This map has RWMutex capacity
// in order to enable concurrency.
type RProxy struct {
	sync.RWMutex
	// Stores a map of format "cloud;user" and a corresponding reverse proxy
	clusters map[string]*httputil.ReverseProxy
}

// NewRProxy
// Instantiates a new reverse proxy with allocated map
func NewRProxy() *RProxy {
	return &RProxy{clusters: make(map[string]*httputil.ReverseProxy)}
}

// NewProxy
// Adds a new proxy to the map based on the specified cluster parameters.
// Properly handles locking.
func (p *RProxy) NewProxy(clusterName, target string) {
	url := &url.URL{
		Scheme: "http",
		Host:   target,
	}

	p.Lock()
	p.clusters[clusterName] = httputil.NewSingleHostReverseProxy(url)
	p.Unlock()
}

// ServeHTTP
// Satisifes the Handler interface for RProxy objects. Searches through the
// proxies map to properly forward based on header information.
func (p *RProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cluster := r.Header.Get("X-Cluster")
	if cluster == "" {
		http.Error(w, "No Cluster specified in header.", http.StatusBadRequest)
		return
	}

	p.RLock()
	if _, ok := p.clusters[cluster]; !ok {
		http.Error(w, fmt.Sprintf("Cluster %s not found", cluster), http.StatusNotFound)
		fmt.Fprintln(w, "Clusters in Steam:")
		for cl, prox := range p.clusters {
			fmt.Fprintf(w, "Cluster %s: %+v", cl, prox)
		}

		log.Println("Cluster", cluster, "not found in proxy")
		p.RUnlock()
		return
	}
	p.clusters[cluster].ServeHTTP(w, r)
	p.RUnlock()
}
