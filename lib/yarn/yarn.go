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

package yarn

import (
	"bufio"
	"context"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/h2oai/steam/lib/haproxy"
	"github.com/h2oai/steam/lib/kerberos"
	"github.com/pkg/errors"
)

func randStr(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	r := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		r[i] = chars[rand.Intn(len(chars))]
	}
	return string(r)
}

func cleanDir(dir string, uid, gid uint32) {
	cmd := exec.Command("hadoop", "fs", "-rmdir", dir)
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uid, Gid: gid}

	if out, err := cmd.Output(); err != nil {
		log.Printf("failed to remove outdir %s: %v", dir, string(out))
	}
}

func GetUser(username string) (uint32, uint32, error) {
	u, err := user.Lookup(username)
	if err != nil {
		return 0, 0, errors.Wrap(err, "failed to lookup user")
	}

	uid64, err := strconv.ParseUint(u.Uid, 10, 32)
	if err != nil {
		return 0, 0, errors.Wrap(err, "failed parsing uid to uint32")
	}

	gid64, err := strconv.ParseUint(u.Gid, 10, 32)
	if err != nil {
		return 0, 0, errors.Wrap(err, "failed parsing gid to unit32")
	}
	return uint32(uid64), uint32(gid64), nil
}

func contextParamScan(r io.ReadCloser, ok chan bool) {
	in := bufio.NewScanner(r)
	for in.Scan() {
		if in.Text() != "" {
			if strings.Contains(in.Text(), "-context_path") {
				ok <- true
			}
		}
	}
	ok <- false
}

func yarnScan(r io.Reader, name, username string, appID, address, err *string, cancel context.CancelFunc) {
	// Scan for ip and app_id
	reNode := regexp.MustCompile(`H2O node (\d+\.\d+\.\d+\.\d+:\d+)`)
	reApID := regexp.MustCompile(`application_(\d+_\d+)`)

	in := bufio.NewScanner(r)
	for in.Scan() {
		if in.Text() != "" {
			// Log output
			log.Println("YARN", name, username, in.Text())
			// Find application id
			if appID != nil {
				if s := reNode.FindSubmatch(in.Bytes()); s != nil {
					*address = string(s[1])
				}
			}
			// Find IP address
			if address != nil {
				if s := reApID.FindSubmatch(in.Bytes()); s != nil {
					*appID = string(s[1])
				}
			}
			// Scan for errors
			if err != nil {
				if strings.Contains(in.Text(), "ERROR") {
					*err += in.Text() + "\n"
				}
				if strings.Contains(in.Text(), "Exception") {
					*err += in.Text() + "\n"
					// Exception should kill process
					cancel()
				}
			}
		}
	}
}

func yarnCommand(uid, gid uint32, name, username string, args ...string) (string, string, error) {
	// Create context for killing process if exception encountered
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up hadoop job with user impersonation
	cmd := exec.CommandContext(ctx, "hadoop", args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uid, Gid: gid}

	// Set stdout and stderr
	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		return "", "", errors.Wrap(err, "failed setting standard out")
	}
	defer stdOut.Close()
	stdErr, err := cmd.StderrPipe()
	if err != nil {
		return "", "", errors.Wrap(err, "failed setting standard err")
	}

	defer stdErr.Close()

	// Log output and scan
	var appID, address, cmdErr string
	go yarnScan(stdOut, name, username, &appID, &address, &cmdErr, cancel)
	go yarnScan(stdErr, name, username, nil, nil, &cmdErr, cancel)

	// Execute command
	if err := cmd.Run(); err != nil {
		return "", "", errors.Wrapf(err, "failed running command %s: %v", strings.Join(cmd.Args, " "), cmdErr)
	}

	return appID, address, nil
}

// StartCloud starts a yarn cloud by shelling out to hadoop
//
// This process needs to store the job-ID to kill the process in the future
func StartCloud(size int, kEnable bool, mem, name, enginePath, username, keytab string, secure bool, uid, gid uint32) (string, string, string, string, string, error) {
	if kEnable {
		if err := kerberos.Kinit(keytab, username, uid, gid); err != nil {
			return "", "", "", "", "", errors.Wrap(err, "initializing kerberos")
		}
	}
	// Randomize outfile name
	out := "steam-output/" + name + "_" + randStr(5) + "_out"

	cmdArgs := []string{
		"jar", enginePath,
		"-jobname", "STEAM_" + name,
		"-n", strconv.Itoa(size),
		"-mapperXmx", mem,
		"-output", out,
		"-disown",
	}

	contextPath := ""
	cpe, err := contextPathEnabledEngine(enginePath, uid, gid)
	if err != nil {
		return "", "", "", "", "", err
	}

	token := ""
	if secure {
		if !cpe {
			return "", "", "", "", "", errors.New("this version of h2o is not compatible with Steam secure launch. Please upgrade your h2o version")
		}

		contextPath = "/" + username + "_" + name
		contextPathArgs := []string{"-J", "-context_path", "-J", contextPath}
		cmdArgs = append(cmdArgs, contextPathArgs...)
		passwd := randStr(10)
		token, _ = haproxy.GenRealmFile(username, passwd)
		defer os.Remove(token + "_realm.properties")
		securityArgs := []string{"-hash_login", "-login_conf", token + "_realm.properties"}
		cmdArgs = append(cmdArgs, securityArgs...)
	}

	appID, address, err := yarnCommand(uid, gid, name, username, cmdArgs...)
	if err != nil {
		cleanDir(out, uid, gid)
		return "", "", "", "", "", errors.Wrap(err, "failed executing command")
	}

	return appID, address, out, token, contextPath + "/", nil
}

func contextPathEnabledEngine(enginePath string, uid, gid uint32) (bool, error) {
	// Set up hadoop job with user impersonation
	cmd := exec.Command("java", "-jar", enginePath, "-help")

	// Set stderr
	stdErr, err := cmd.StderrPipe()
	if err != nil {
		return false, errors.Wrap(err, "starting standard error pipe")
	}
	defer stdErr.Close()

	// Log output and scan
	okChan := make(chan bool)
	go contextParamScan(stdErr, okChan)

	// Execute command. NOT checking for err because -help in h2odriver returns always 1...
	cmd.Run()

	return <-okChan, nil
}

// StopCloud kills a hadoop cloud by shelling out a command based on the job-ID
func StopCloud(kEnable bool, name, id, outdir, username, keytab string, uid, gid uint32) error {
	// If kerberos enabled, initialize and defer destroy
	if kEnable {
		if err := kerberos.Kinit(keytab, username, uid, gid); err != nil {
			return errors.Wrap(err, "initializing kerberos")
		}
	}

	if _, _, err := yarnCommand(uid, gid, name, username, "job", "-kill", "job_"+id); err != nil {
		return errors.Wrap(err, "failed executing command")
	}

	cleanDir(outdir, uid, gid)
	return nil
}
