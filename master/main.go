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

package master

import (
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"

	"github.com/gorilla/context"
	"github.com/h2oai/steam/lib/fs"
	"github.com/h2oai/steam/lib/ldap"
	"github.com/h2oai/steam/lib/rpc"
	"github.com/h2oai/steam/master/data"
	"github.com/h2oai/steam/master/proxy"
	"github.com/h2oai/steam/master/web"
	srvweb "github.com/h2oai/steam/srv/web"
	"github.com/gorilla/handlers"
)

const (
	defaultWebAddress                   = ":9000"
	defaultClusterProxyAddress          = ":9001"
	defaultCompilationAddress           = ":8080"
	defaultPredictionServiceHost        = ""
	DefaultPredictionServicePortsString = "1025:65535"
)

var defaultPredictionServicePorts = [...]int{1025, 65535}

type DBOpts struct {
	Connection        data.Connection
	SuperuserName     string
	SuperuserPassword string
}

type YarnOpts struct {
	KerberosEnabled bool
}

type Opts struct {
	WebAddress                string
	WebTLSCertPath            string
	WebTLSKeyPath             string
	AuthProvider              string
	AuthConfig                string
	WorkingDirectory          string
	ClusterProxyAddress       string
	CompilationServiceAddress string
	PredictionServiceHost     string
	PredictionServicePorts    [2]int
	EnableProfiler            bool
	Yarn                      YarnOpts
	DB                        DBOpts
}

var DefaultConnection = data.Connection{
	"steam",
	"steam",
	"",
	"",
	"",
	"",
	"disable",
	"",
	"",
	"",
}

var DefaultOpts = &Opts{
	defaultWebAddress,
	"",
	"",
	"basic",
	"ldap.toml",
	path.Join(".", fs.VarDir, "master"),
	defaultClusterProxyAddress,
	defaultCompilationAddress,
	defaultPredictionServiceHost,
	defaultPredictionServicePorts,
	false,
	YarnOpts{false},
	DBOpts{DefaultConnection, "", ""},
}

type AuthProvider interface {
	Secure(handler http.Handler) http.Handler
	Logout() http.Handler
}

func Run(version, buildDate string, opts Opts) {
	log.Printf("steam v%s build %s\n", version, buildDate)

	// --- external ip for base and proxy ---
	webAddress := opts.WebAddress
	proxyAddress := opts.ClusterProxyAddress

	// --- set up wd ---
	wd, err := fs.MkWorkingDirectory(opts.WorkingDirectory)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Working directory:", wd)

	// --- www root ---
	wwwroot := fs.GetWwwRoot(wd)
	if _, err := os.Stat(path.Join(wwwroot, "index.html")); err != nil {
		log.Fatalf("Web root not found at %s: %v\n", wwwroot, err)
	}
	log.Println("WWW root:", wwwroot)

	// --- init storage ---

	ds, err := data.Create(
		path.Join(wd, fs.DbDir, "steam.db"),
		// opts.DB.Connection,
		opts.DB.SuperuserName,
		opts.DB.SuperuserPassword,
	)

	if err != nil {
		log.Fatalln(err)
	}

	// --- create basic auth service ---
	defaultAz := NewDefaultAz(ds)
	var authProvider AuthProvider
	switch opts.AuthProvider {
	case "digest":
		authProvider = newDigestAuthProvider(defaultAz, webAddress)
	case "basic-ldap":
		conn, err := ldap.FromConfig(opts.AuthConfig)
		if err != nil {
			log.Fatalln("Please provide a valid ldap configuration file", err)
		}

		authProvider = NewBasicLdapAuthProvider(webAddress, conn)
	default: // "basic"
		authProvider = newBasicAuthProvider(defaultAz, webAddress)
	}

	// --- set up prediction service launch host

	var predictionServiceHost string
	if opts.PredictionServiceHost != "" {
		predictionServiceHost = opts.PredictionServiceHost
	} else {
		var err error
		predictionServiceHost, err = fs.GetExternalHost()
		if err != nil {
			log.Fatalln(err)
		}
	}

	// --- create web services ---

	webServeMux := http.NewServeMux()
	webService := web.NewService(
		wd,
		ds,
		opts.CompilationServiceAddress,
		predictionServiceHost,
		opts.ClusterProxyAddress,
		opts.PredictionServicePorts,
		opts.Yarn.KerberosEnabled,
	)
	webServiceImpl := &srvweb.Impl{webService, defaultAz}

	webServeMux.Handle("/logout", handlers.CORS()(authProvider.Logout()))
	webServeMux.Handle("/web", handlers.CORS()(authProvider.Secure(rpc.NewServer(rpc.NewService("web", webServiceImpl)))))
	webServeMux.Handle("/upload", handlers.CORS()(authProvider.Secure(newUploadHandler(defaultAz, wd, webServiceImpl.Service, ds))))
	webServeMux.Handle("/download", handlers.CORS()(authProvider.Secure(newDownloadHandler(defaultAz, wd, webServiceImpl.Service, opts.CompilationServiceAddress))))
	webServeMux.Handle("/", handlers.CORS()(authProvider.Secure(http.FileServer(http.Dir(path.Join(wd, "/www"))))))

	if opts.EnableProfiler {
		// --- pprof registrations (no auth) ---
		webServeMux.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
		webServeMux.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
		webServeMux.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
		webServeMux.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	}

	// --- start web server ---

	serverFailChan := make(chan error)
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	certFile := strings.TrimSpace(opts.WebTLSCertPath)
	keyFile := strings.TrimSpace(opts.WebTLSKeyPath)
	enableTLS := !(len(certFile) == 0 && len(keyFile) == 0)

	go func() {
		log.Println("Web server listening at", webAddress)
		prefix := ""
		if len(webAddress) > 1 && webAddress[:1] == ":" {
			prefix = "localhost"
		}
		if enableTLS {
			log.Printf("Point your web browser to https://%s%s/\n", prefix, webAddress)
			if err := http.ListenAndServeTLS(webAddress, certFile, keyFile, context.ClearHandler(webServeMux)); err != nil {
				serverFailChan <- err
			}
		} else {
			log.Printf("Point your web browser to http://%s%s/\n", prefix, webAddress)
			if err := http.ListenAndServe(webAddress, context.ClearHandler(webServeMux)); err != nil {
				serverFailChan <- err
			}
		}
	}()

	// --- start reverse proxy ---

	proxyHandler := authProvider.Secure(proxy.NewProxyHandler(defaultAz, ds))
	proxyFailChan := make(chan error)
	go func() {
		log.Println("Cluster reverse proxy listening at", proxyAddress)
		prefix := ""
		if len(proxyAddress) > 1 && proxyAddress[:1] == ":" {
			prefix = "localhost"
		}
		if enableTLS {
			log.Printf("Point H2O client libraries to https://%s%s/\n", prefix, proxyAddress)
			if err := http.ListenAndServeTLS(proxyAddress, certFile, keyFile, proxyHandler); err != nil {
				proxyFailChan <- err
			}

		} else {
			log.Printf("Point H2O client libraries to http://%s%s/\n", prefix, proxyAddress)
			if err := http.ListenAndServe(proxyAddress, proxyHandler); err != nil {
				proxyFailChan <- err
			}
		}
	}()

	// --- wait for termination ---

	for {
		select {
		case err := <-serverFailChan:
			log.Fatalln("HTTP server startup failed:", err)
			return
		case err := <-proxyFailChan:
			log.Fatalln("Cluster reverse proxy startup failed:", err)
			return
		case sig := <-sigChan:
			log.Println("Caught signal", sig)
			log.Println("Shut down gracefully.")
			return
		}
	}
}
