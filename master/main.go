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
	"github.com/h2oai/steamY/lib/fs"
	"github.com/h2oai/steamY/lib/rpc"
	"github.com/h2oai/steamY/master/auth"
	"github.com/h2oai/steamY/master/data"
	"github.com/h2oai/steamY/master/proxy"
	"github.com/h2oai/steamY/master/web"
)

const (
	defaultWebAddress          = ":9000"
	defaultClusterProxyAddress = ":9001"
	defaultCompilationAddress  = ":8080"
	defaultScoringServiceHost  = ""
)

type Opts struct {
	WebAddress                string
	WebTLSCertPath            string
	WebTLSKeyPath             string
	WorkingDirectory          string
	ClusterProxyAddress       string
	CompilationServiceAddress string
	ScoringServiceHost        string
	EnableProfiler            bool
	YarnKerberosEnabled       bool
	YarnUserName              string
	YarnKeytab                string
	DBName                    string
	DBUserName                string
	DBSSLMode                 string
	SuperuserName             string
	SuperuserPassword         string
}

var DefaultOpts = &Opts{
	defaultWebAddress,
	"",
	"",
	path.Join(".", fs.VarDir, "master"),
	defaultClusterProxyAddress,
	defaultCompilationAddress,
	defaultScoringServiceHost,
	false,
	false,
	"",
	"",
	"steam",
	"steam",
	"disable",
	"",
	"",
}

type AuthProvider interface {
	Secure(handler http.Handler) http.Handler
	Logout() http.Handler
}

func Run(version, buildDate string, opts *Opts) {
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

	db, err := data.Connect(opts.DBUserName, opts.DBName, opts.DBSSLMode)
	if err != nil {
		log.Fatalf("Failed connecting to database %s as user %s (SSL=%s): %s\n", opts.DBName, opts.DBUserName, opts.DBSSLMode, err)
	}

	isPrimed, err := data.IsPrimed(db)
	if err != nil {
		log.Fatalln("Failed database version check:", err)
	}

	if !isPrimed {
		if opts.SuperuserName == "" || opts.SuperuserPassword == "" {
			log.Fatalln("Starting Steam for the first time requires both --superuser-name and --superuser-password arguments to \"steam serve master\".")
		}

		if err := auth.ValidateUsername(opts.SuperuserName); err != nil {
			log.Fatalln("Invalid superuser username:", err)
		}

		if err := auth.ValidatePassword(opts.SuperuserPassword); err != nil {
			log.Fatalln("Invalid superuser password:", err)
		}

		if err := data.Prime(db); err != nil {
			log.Fatalln("Failed priming database:", err)
		}
	}

	ds, err := data.NewDatastore(db)
	if err != nil {
		log.Fatalln("Failed initializing from database:", err)
	}

	if !isPrimed {
		passwordHash, err := auth.HashPassword(opts.SuperuserPassword)
		if err != nil {
			log.Fatalln("Failed hashing superuser password:", err)
		}

		if _, _, err := ds.CreateSuperuser(opts.SuperuserName, passwordHash); err != nil {
			log.Fatalln("Failed superuser identity setup:", err)
		}

		_, err = ds.NewPrincipal(opts.SuperuserName)
		if err != nil {
			log.Fatalln("Failed reading superuser principal:", err)
		}
	}

	// --- create basic auth service ---
	defaultAz := newDefaultAz(ds)
	// TODO add CLI args for other other providers
	// TODO conditionally instantiate auth provider based on CLI arg
	authProvider := newBasicAuthProvider(defaultAz, webAddress)

	// --- create web services ---

	webServeMux := http.NewServeMux()
	webServiceImpl := web.NewService(
		defaultAz,
		wd,
		ds,
		opts.CompilationServiceAddress,
		opts.ScoringServiceHost,
		opts.YarnKerberosEnabled,
		opts.YarnUserName,
		opts.YarnKeytab,
	)

	webServeMux.Handle("/logout", authProvider.Logout())
	webServeMux.Handle("/web", authProvider.Secure(rpc.NewServer(rpc.NewService("web", webServiceImpl))))
	webServeMux.Handle("/upload", authProvider.Secure(newUploadHandler(defaultAz, wd, webServiceImpl.Service)))
	webServeMux.Handle("/", authProvider.Secure(http.FileServer(http.Dir(path.Join(wd, "/www")))))

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
