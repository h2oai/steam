package master

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"path"
	"syscall"

	"github.com/gorilla/context"
	"github.com/h2oai/steamY/lib/fs"
	"github.com/h2oai/steamY/lib/proxy"
	"github.com/h2oai/steamY/lib/rpc"
	"github.com/h2oai/steamY/master/db"
	"github.com/h2oai/steamY/master/web"
	srvweb "github.com/h2oai/steamY/srv/web"
)

const (
	defaultWebAddress          = "0.0.0.0:9000"
	defaultClusterProxyAddress = "0.0.0.0:9001"
	defaultCompilationAddress  = "0.0.0.0:8080"
	defaultScoringService      = "0.0.0.0"
)

type Opts struct {
	WebAddress                string
	WorkingDirectory          string
	ClusterProxyAddress       string
	CompilationServiceAddress string
	ScoringServiceAddress     string
	EnableProfiler            bool
	KerberosEnabled           bool
	Username                  string
	Keytab                    string
}

var DefaultOpts = &Opts{
	defaultWebAddress,
	path.Join(".", fs.VarDir, "master"),
	defaultClusterProxyAddress,
	defaultCompilationAddress,
	defaultScoringService,
	false,
	false,
	"",
	"",
}

type UploadHandler struct {
	workingDirectory string
	webService       srvweb.Service
}

func newUploadHandler(wd string, webService srvweb.Service) *UploadHandler {
	return &UploadHandler{wd, webService}
}

func (s *UploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("File upload request received.")

	r.ParseMultipartForm(0)

	kind := r.FormValue("kind")

	src, handler, err := r.FormFile("file")
	if err != nil {
		log.Println("Upload form parse failed:", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Malformed request: %v", err)
		return
	}
	defer src.Close()

	log.Println("Remote file: ", handler.Filename)

	fileBaseName := path.Base(handler.Filename)

	dstPath := path.Join(s.workingDirectory, fs.LibDir, kind, fileBaseName)
	if err := os.MkdirAll(path.Dir(dstPath), fs.DirPerm); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "%v", err)
		return
	}

	dst, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE, fs.FilePerm)
	if err != nil {
		log.Println("Upload file open operation failed:", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error writing uploaded file to disk: %v", err)
		return
	}
	defer dst.Close()
	io.Copy(dst, src)

	if err := s.webService.AddEngine(fileBaseName, dstPath); err != nil {
		log.Println("Failed saving engine to datastore", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error saving engine to datastore: %v", err)
		return
	}

	log.Println("Engine uploaded:", dstPath)

}

func Run(version, buildDate string, opts *Opts) {
	log.Printf("steam v%s build %s\n", version, buildDate)

	// --- external ip for base and proxy ---
	webAddress := opts.WebAddress
	proxAddress := opts.ClusterProxyAddress
	if webAddress == defaultWebAddress {
		webAddress = fs.GetExternalIP(webAddress)
		proxAddress = fs.GetExternalIP(proxAddress)
	}

	// --- set up wd ---
	wd, err := fs.MkWorkingDirectory(opts.WorkingDirectory)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Working directory: %s", wd)

	// --- www root ---
	wwwroot := fs.GetWwwRoot(wd)
	if _, err := os.Stat(path.Join(wwwroot, "index.html")); err != nil {
		log.Fatalf("Web root not found at %s: %v\n", wwwroot, err)
	}
	log.Printf("WWW root: %s", wwwroot)

	// --- init storage ---
	dbPath := fs.GetDbPath(wd)
	ds, err := db.Open(dbPath)
	if err != nil {
		log.Fatalln(err)
	}
	primed, err := isPrimed(ds)
	if err != nil {
		log.Fatalln(err)
	}
	if !primed {
		if err := prime(ds); err != nil {
			log.Fatalln(err)
		}
	}
	ds.Init()
	log.Println("Datastore location:", dbPath)

	// --- create domain services ---
	// jt := job.NewTracker(webAddress)

	// --- create proxy services ---
	rp := proxy.NewRProxy()
	rpFailch := make(chan error)

	go func() {
		log.Println()
		if err := http.ListenAndServe(proxAddress, rp); err != nil {
			rpFailch <- err
		}
	}()

	// --- create front end api services ---

	webServeMux := http.NewServeMux()
	webServiceImpl := web.NewService(
		wd,
		ds,
		opts.CompilationServiceAddress,
		opts.ScoringServiceAddress,
		rp,
		opts.KerberosEnabled,
		opts.Username,
		opts.Keytab,
	)
	webServeMux.Handle("/web", rpc.NewServer(rpc.NewService("web", webServiceImpl)))
	webServeMux.Handle("/upload", newUploadHandler(wd, webServiceImpl.Service))
	webServeMux.Handle("/", http.FileServer(http.Dir(path.Join(wd, "/www")))) // no auth

	if opts.EnableProfiler {
		// --- pprof registrations (no auth) ---
		webServeMux.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
		webServeMux.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
		webServeMux.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
		webServeMux.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	}

	// --- start http servers ---

	failch := make(chan error)
	sigch := make(chan os.Signal)
	signal.Notify(sigch, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Web server listening at", webAddress)
		log.Printf("Point your web browser to http://%s/\n", webAddress)
		if err := http.ListenAndServe(webAddress, context.ClearHandler(webServeMux)); err != nil {
			failch <- err
		}
	}()

	webServiceImpl.Service.InitClusterProxy() // Initialize reverse proxies for clusters
	webServiceImpl.Service.ActivityPoll(true) // Poll clouds for activity

	for {
		select {
		case err := <-rpFailch:
			log.Fatalln("HTTP proxy server startup failed:", err)
			return
		case err := <-failch:
			ds.Close()
			log.Fatalln("HTTP server startup failed:", err)
			return
		case sig := <-sigch:
			ds.Close()
			log.Println(sig)
			log.Println("Shut down gracefully.")
			return
		}
	}
}

const (
	SchemaVersion uint32 = 1
)

func isPrimed(ds *db.DS) (bool, error) {
	primed, err := ds.HasBucket("Sys")
	if err != nil {
		return false, err
	}

	if primed {
		return true, nil
	}

	return false, nil
}

func prime(ds *db.DS) error {
	log.Println("Priming datastore for first time use...")

	if err := ds.CreateBuckets(db.Buckets); err != nil {
		return err
	}

	if err := ds.CreateSys(db.NewSys("System", SchemaVersion)); err != nil {
		return err
	}

	return nil
}
