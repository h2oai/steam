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

package main

import (
	"bufio"
	"context"
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
	"syscall"

	"github.com/BurntSushi/toml"
	"github.com/h2oai/steam/lib/rpc"
	"github.com/h2oai/steam/srv/web"
	"github.com/pkg/errors"
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
	var err error
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() { err = errors.Wrap(launchSteam(ctx), "failed launching steam"); cancel() }()
	go func() {
		err = errors.Wrap(launchPredictionBuilder(ctx), "failed launching scoring service")
		cancel()
	}()

	select {
	case <-ctx.Done():
		log.Fatalln(err)
	case sig := <-sigChan:
		log.Println("Caught signal", sig)
		log.Println("Shut down gracefully.")
		return
	}
}

func launchPredictionBuilder(ctx context.Context) error {
	log.Println("Launching Scoring Service...")

	// Issue commands and start scoring service
	argv := []string{
		"-jar", config.PredictionBuilder.JettyPath,
		"--port", strconv.Itoa(config.PredictionBuilder.Port),
		config.PredictionBuilder.WarPath,
	}
	cmd := exec.CommandContext(ctx, "java", argv...)
	stdErr, err := cmd.StderrPipe() // StdErr for logging
	if err != nil {
		return errors.Wrap(err, "failed to set stderr pipe")
	}
	defer stdErr.Close()
	stdOut, err := cmd.StdoutPipe() // StdOut for logging
	if err != nil {
		return errors.Wrap(err, "failed to set stdout pipe")
	}
	defer stdOut.Close()
	if err := cmd.Start(); err != nil {
		return errors.Wrap(err, "failed to start")
	}

	pass := make(chan struct{})
	re := regexp.MustCompile(`Started @\d+`)
	go logScan(stdErr, "BUILDER", re, pass)
	go logScan(stdOut, "BUILDER", re, pass)

	return errors.Wrap(cmd.Wait(), "unexpetedly quit")
}

func launchSteam(ctx context.Context) error {
	log.Println("Launching Steam...")

	// Issue commands and start steam
	argv := []string{
		"serve", "master",
		"--admin-name=" + config.Admin.Name,
		"--admin-password=" + config.Admin.Pass,
		"--web-address=:" + strconv.Itoa(config.Steam.Port),
		"--compilation-service-address=" + config.PredictionBuilder.Host + ":" + strconv.Itoa(config.PredictionBuilder.Port),
		"--scoring-service-port-range=" + config.PredictionService.PortRange,
	}
	cmd := exec.CommandContext(ctx, "./steam", argv...)
	stdErr, err := cmd.StderrPipe() // StdErr for logging
	if err != nil {
		return errors.Wrap(err, "failed setting stderr pipe")
	}
	defer stdErr.Close()
	stdOut, err := cmd.StdoutPipe() // StdOut for logging
	if err != nil {
		return errors.Wrap(err, "failed setting stdout pipe")
	}
	defer stdOut.Close()

	pass := make(chan struct{})
	re := regexp.MustCompile(`Web server listening at :\d+`)
	go logScan(stdErr, "STEAM", re, pass)
	go logScan(stdOut, "STEAM", re, pass)
	defer close(pass)
	go func() {
		_ = <-pass
		if doInit {
			//FIXME: Not checking for errors here should be fixed
			initSteam()
		}
	}()

	return errors.Wrap(cmd.Run(), "unexpectedly quit")
}

func initSteam() error {
	log.Println("Initializing Steam roles...")

	// Connect to running Steam instance
	remote := &web.Remote{rpc.NewProc(
		"http",
		"/web",
		"web",
		"localhost:"+strconv.Itoa(config.Steam.Port),
		config.Admin.Name,
		config.Admin.Pass),
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

func logScan(stream io.ReadCloser, src string, re *regexp.Regexp, pass chan struct{}) error {

	in := bufio.NewScanner(stream)
	for in.Scan() {
		fmt.Println(src, in.Text())

		if re.Match(in.Bytes()) {
			close(pass)
		}
	}

	return nil
}
