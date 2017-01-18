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
	"io/ioutil"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/gorilla/context"
	"github.com/h2oai/steam/lib/fs"
	"github.com/h2oai/steam/lib/ldap"
	"github.com/h2oai/steam/lib/rpc"
	"github.com/h2oai/steam/lib/yarn"
	"github.com/h2oai/steam/master/data"
	"github.com/h2oai/steam/master/web"
	srvweb "github.com/h2oai/steam/srv/web"
	"github.com/rs/cors"
	"github.com/pkg/errors"
)

const (
	defaultWebAddress                   = ":9000"
	defaultClusterProxyAddress          = ":9001"
	defaultCompilationAddress           = ":8080"
	defaultPredictionServiceHost        = ""
	defaultPredictionServicePortsString = "1025:65535"
)

var defaultPredictionServicePorts = [...]int{1025, 65535}

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
	DBOpts                    data.DBOpts
	Dev		          bool
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
	data.DBOpts{
		Driver: "sqlite3",

		Path: filepath.Join(".", fs.VarDir, "master", fs.DbDir, "steam.db"),

		Name:    "steam",
		Pass:    "steam",
		Host:    "localhost",
		Port:    "5432",
		SSLMode: "disable",
	},
	false,
}

type AuthProvider interface {
	Secure(handler http.Handler) http.Handler
	Logout() http.Handler
}

func handlerStrategy(handler http.Handler, opts Opts) http.Handler {
	if (opts.Dev) {
		c := cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"POST", "OPTIONS", "GET"},
			AllowCredentials: true,
		})
		return c.Handler(handler)
	}
	return handler
}

func Run(version, buildDate string, opts Opts) {
	log.Printf("steam v%s build %s\n", version, buildDate)

	// --- external ip for base and proxy ---
	webAddress := opts.WebAddress

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

	opts.DBOpts.Path = path.Join(wd, fs.DbDir, "steam.db")
	ds, err := data.NewDatastore(opts.DBOpts, false)
	if err != nil {
		log.Fatalln(err)
	}

	// --- create basic auth service ---
	defaultAz := NewDefaultAz(ds)
	var authProvider AuthProvider

	var set string
	if auth, exists, err := ds.ReadAuthentication(data.ByEnabled); err != nil {
		log.Fatalln("Reading authentication setting from database:", err)
	} else if exists {
		set = auth.Key
	}
	switch {
	case opts.AuthProvider == "digest":
		authProvider = newDigestAuthProvider(defaultAz, webAddress)
	case opts.AuthProvider == "basic-ldap", set == data.LDAPAuth:
		conn, err := ldap.FromDatabase(ds)
		if err != nil {
			log.Fatalln("Invalid configuration:", err)
		}

		authProvider = NewBasicLdapAuthProvider(ds, webAddress, conn)
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
		version,
		wd,
		ds,
		opts.CompilationServiceAddress,
		predictionServiceHost,
		opts.ClusterProxyAddress,
		opts.PredictionServicePorts,
		opts.Yarn.KerberosEnabled,
	)
	webServiceImpl := &srvweb.Impl{webService, defaultAz}

	webServeMux.Handle("/logout", handlerStrategy(authProvider.Logout(), opts))
	webServeMux.Handle("/web", authProvider.Secure(handlerStrategy(rpc.NewServer(rpc.NewService("web", webServiceImpl)), opts)))
	webServeMux.Handle("/upload", handlerStrategy(authProvider.Secure(newUploadHandler(defaultAz, wd, webServiceImpl.Service, ds)), opts))
	webServeMux.Handle("/download", handlerStrategy(authProvider.Secure(newDownloadHandler(defaultAz, wd, webServiceImpl.Service, opts.CompilationServiceAddress)), opts))
	webServeMux.Handle("/", handlerStrategy(authProvider.Secure(http.FileServer(http.Dir(path.Join(wd, "/www")))), opts))

	if opts.EnableProfiler {
		// --- pprof registrations (no auth) ---
		webServeMux.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
		webServeMux.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
		webServeMux.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
		webServeMux.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	}

	// --- launch polling job
	pollFailChan := make(chan error)
	go func() { pollFailChan <- yarn.StartPoll(ds, data.States.Started, data.States.Stopped) }()

	// --- start web server ---

	serverFailChan := make(chan error)
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	certFile := strings.TrimSpace(opts.WebTLSCertPath)
	keyFile := strings.TrimSpace(opts.WebTLSKeyPath)

	enableTLS := !(len(certFile) == 0 && len(keyFile) == 0)

	go func() {
		log.Println("Web server listening at", webAddress)
		if enableTLS {
			kb, err := ioutil.ReadFile(keyFile)
			if err != nil {
				log.Fatalln(err) //FIXME format error
			}
			cb, err := ioutil.ReadFile(certFile)
			if err != nil {
				log.Fatalln(err) //FIXME format error
			}
			if err := ioutil.WriteFile(fs.GetAssetsPath(opts.WorkingDirectory, "cert.pem"), append(cb, kb...), 0622); err != nil {
				log.Fatalln(err)
			}

			if err := http.ListenAndServeTLS(webAddress, certFile, keyFile, context.ClearHandler(handlerStrategy(webServeMux, opts))); err != nil {
				serverFailChan <- err
			}
		} else {
			if err := http.ListenAndServe(webAddress, context.ClearHandler(handlerStrategy(webServeMux, opts))); err != nil {
				serverFailChan <- err
			}
		}
	}()

	// --- wait for termination ---

	for {
		select {
		case err := <-pollFailChan:
			if err != nil {
				log.Fatalln("Poll startup failure", err)
				return
			}
		case err := <-serverFailChan:
			log.Fatalln("HTTP server startup failed:", err)
			return
		case sig := <-sigChan:
			log.Println("Caught signal", sig)
			log.Println("Shut down gracefully.")
			return
		}
	}
}

func SetAdmin(workingDirectory string, dbOpts data.DBOpts) error {

	// --- set up wd ---
	wd, err := fs.MkWorkingDirectory(workingDirectory)
	if err != nil {
		return err
	}

	// --- init storage ---
	dbOpts.Path = path.Join(wd, fs.DbDir, "steam.db")
	_, err = data.NewDatastore(dbOpts, true)
	return err
}

func CheckAdmin(workingDirectory string, dbOpts data.DBOpts) error {

	// --- set up wd ---
	wd, err := fs.MkWorkingDirectory(workingDirectory)
	if err != nil {
		return err
	}

	// --- init storage ---

	switch dbOpts.Driver {
	case "sqlite3":
		dbOpts.Path = path.Join(wd, fs.DbDir, "steam.db")
		if _, err := os.Stat(dbOpts.Path); os.IsNotExist(err) {
			return err
		}

	}

	_, err = data.NewDatastore(dbOpts, false)
	return errors.Wrap(err, "setting up datastore")
}
