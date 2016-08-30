package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"regexp"
	"strconv"
	"sync"
	"syscall"

	"github.com/BurntSushi/toml"
	"github.com/h2oai/steam/lib/rpc"
	"github.com/h2oai/steamY/srv/web"
)

var (
	config tomlConfig
	doInit bool
)

// Initialize config and parse Flags
func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalln(err)
	}
	configfile := filepath.Join(dir, "config.toml")

	if _, err := os.Stat(configfile); err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatalln(err)
	}

	flag.BoolVar(&doInit, "init", false, "Initialize Steam with default roles.")
	flag.Parse()
}

// Run checks, launch services, and handle errors
func main() {
	sigChan := make(chan os.Signal, 1)
	failCh := make(chan error, 3)
	killCh := make(chan bool)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start services in WaitGroup to insure both exit before returning from main
	wg := &sync.WaitGroup{}
	go func() { wg.Add(1); failCh <- launchSteam(killCh); wg.Done() }()
	go func() { wg.Add(1); failCh <- launchScoringService(killCh); wg.Done() }()

	select {
	// Error case: close channel and wait until services exit.
	case err := <-failCh:
		log.Println(err)
		close(killCh)
		wg.Wait()
		return
	case sig := <-sigChan:
		log.Println("Caught signal", sig)
		log.Println("Shut down gracefully.")
		return
	}
}

func launchScoringService(killch chan bool) error {
	log.Println("Launching Scoring Service...")

	// Issue commands and start scoring service
	argv := []string{
		"-jar",
		config.ScoringSerivce.JettyPath,
		"--port",
		strconv.Itoa(config.ScoringSerivce.Port),
		config.ScoringSerivce.WarPath,
	}
	cmd := exec.Command("java", argv...)
	stdErr, err := cmd.StderrPipe() // StdErr for logging
	if err != nil {
		return err
	}
	defer stdErr.Close()
	if err := cmd.Start(); err != nil {
		return err
	}
	// Kill process if main fails
	go func() { <-killch; cmd.Process.Kill() }()

	// Start log and issue any successful startup commands
	okCh := make(chan bool)
	success := regexp.MustCompile(`Started @\d+ms`)
	go writeToLog(stdErr, "service.log", success, okCh)

	go func() {
		ok := <-okCh
		if !ok {
			killch <- true
		}
		log.Println("Scoring service is up.")
	}()

	// Blocking: if this function returns, an error occured.
	return fmt.Errorf("Unexpectedly quit Scoring Service. See service log file.: %v", cmd.Wait())
}

func launchSteam(killch chan bool) error {
	log.Println("Launching Steam...")

	// Issue commands and start steam
	argv := []string{
		"serve",
		"master",
		"--superuser-name=" + config.Superuser.Name,
		"--superuser-password=" + config.Superuser.Pass,
		"--web-address=:" + strconv.Itoa(config.Steam.Port),
		"--scoring-service-address=:" + strconv.Itoa(config.ScoringSerivce.Port),
		"--scoring-service-port-range=" + config.ScoringSerivce.PortRange,
	}
	cmd := exec.Command("./steam", argv...)
	stdErr, err := cmd.StderrPipe() // StdErr for logging
	if err != nil {
		return err
	}
	defer stdErr.Close()
	if err := cmd.Start(); err != nil {
		return err
	}

	// Kill process if main fails
	go func() { <-killch; cmd.Process.Kill() }()

	// Start log and issue any successful startup commands
	okCh := make(chan bool)
	success := regexp.MustCompile("Web server listening at")
	go writeToLog(stdErr, "steam.log", success, okCh)
	go func() { // On success run Initialization if -init specified
		ok := <-okCh
		if !ok {
			killch <- true
		}
		if doInit {
			if err := initSteam(); err != nil {
				log.Println(err)
				killch <- true
			}
		}
		log.Printf("Steam is up on port %d", config.Steam.Port)
	}()

	// Blocking: if this function returns, an error occured.
	return fmt.Errorf("Unexpectedly quit Steam. See steam log file.: %v", cmd.Wait())
}

func initSteam() error {
	log.Println("Initializing Steam roles...")

	// Connect to running Steam instance
	remote := &web.Remote{rpc.NewProc(
		"http",
		"/web",
		"web",
		"localhost:"+strconv.Itoa(config.Steam.Port),
		config.Superuser.Name,
		config.Superuser.Pass),
	}

	// Initialize the permissions map
	perms, err := remote.GetAllPermissions()
	if err != nil {
		return err
	}
	permMap := make(map[string]int64, len(perms))
	for _, perm := range perms {
		permMap[perm.Code] = perm.Id
	}

	// Create roles
	for name, role := range config.DefaultRoles {
		roleId, err := remote.CreateRole(name, role.Description)
		if err != nil {
			return err
		}

		// Map ids to permission codes
		permIds := make([]int64, len(role.Permissions))
		for i, perm := range role.Permissions {
			permIds[i] = permMap[perm]
		}

		if err := remote.LinkRoleWithPermissions(roleId, permIds); err != nil {
			return err
		}
	}

	return nil
}

func writeToLog(stream io.ReadCloser, fileName string, succ *regexp.Regexp, okCh chan bool) error {
	in := bufio.NewScanner(stream)

	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
		okCh <- false
		return err
	}
	defer f.Close()

	for in.Scan() {
		if _, err := f.WriteString(in.Text() + "\n"); err != nil {
			return err
		}
		if err := f.Sync(); err != nil {
			return err
		}

		if succ.Match(in.Bytes()) {
			okCh <- true
		}
	}

	return nil
}
