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

package svc

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/h2oai/steam/srv/web"
	"github.com/pkg/errors"
)

func svcScan(r io.Reader, name, username string, err *string, pass chan struct{}) {
	in := bufio.NewScanner(r)
	for in.Scan() {
		log.Println("PRED", name, username, in.Text())

		if strings.Contains(in.Text(), "Started @") {
			close(pass)
		}
		if strings.Contains(in.Text(), "FAILED") {
			s := strings.SplitAfter(in.Text(), "FAILED")
			*err += s[1] + "\n"
		}
	}
}

// Start starts a scoring service.
func Start(warfile, jetty, host string, port int, name, username string) (int, error) {

	argv := []string{"-jar", jetty, "--port", strconv.Itoa(port)}

	if len(host) > 0 {
		argv = append(argv, "--host", host)
	}

	argv = append(argv, warfile)

	cmd := exec.Command("java", argv...)
	stdErr, err := cmd.StderrPipe()
	if err != nil {
		return 0, errors.Wrap(err, "failed setting standard out")
	}
	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		return 0, errors.Wrap(err, "failed setting standard err")
	}
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	pass := make(chan struct{})
	fail := make(chan struct{})
	var cmdErr string
	go svcScan(stdErr, name, username, &cmdErr, pass)
	go svcScan(stdOut, name, username, &cmdErr, pass)

	go func() { err = cmd.Run(); close(fail) }()

	select {
	case <-pass:
		pid := cmd.Process.Pid
		return pid, nil
	case <-fail:
		e := fmt.Errorf("%v", cmdErr)
		return 0, errors.Wrap(e, "starting service")
	}
}

// Stop stops a scoring service.
func Stop(pid int) error {
	p, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("Failed locating scoring service with pid %d: %v", pid, err)
	}

	// Verify if this pid belongs to a scoring service
	pscmd, err := getProcessCommand(pid)
	if err != nil {
		return fmt.Errorf("Failed inspecting pid %d: %v", pid, err)
	}

	isJetty := regexp.MustCompile(`java -jar .*/jetty.*\.jar .*\.war`)
	if isJetty.Find([]byte(pscmd)) == nil {
		return fmt.Errorf("Process %d is not a scoring service", pid)
	}

	const sigintInterval = time.Second
	const sigintTimeout = time.Second * 10

	// Retry SIGINTs;  SIGKILL on timeout
	ticker := time.NewTicker(sigintInterval)
	timeout := time.NewTimer(sigintTimeout)
	for {
		select {
		case <-timeout.C:
			ticker.Stop()
			if err := p.Signal(os.Kill); err != nil {
				return fmt.Errorf("Failed terminating scoring service at port %d: %v", pid, err)

			}
		case <-ticker.C:
			err := p.Signal(os.Interrupt)
			if err == nil || (err != nil && isProcessFinished(err)) {
				ticker.Stop()
				return nil
			}
		}
	}
}

// Polls a scoring service to determine the last time it was used
//
// Returns a timestamp (int64) of the Unix time since the scoring services was
//		created or last used.
func Poll(s *web.ScoringService) (int64, error) {
	u := (&url.URL{
		Scheme: "http",
		Host:   s.Address + ":" + strconv.Itoa(s.Port),
		Path:   "stats",
	}).String()

	r, err := http.Get(u)
	if err != nil {
		return 0, fmt.Errorf("Service request failed: %s: %v", u, err)
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return 0, fmt.Errorf("Service response read failed: %s: %v", u, err)
	}

	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return 0, fmt.Errorf("Error unmarshaling response: %s: %v", string(b), err)
	}

	last := int64(m["lastTime"].(float64))

	return last / 1000, nil
}

func getProcessCommand(pid int) (string, error) {
	lines, err := exec.Command("/bin/ps", "-o", "command", "-p", strconv.Itoa(pid)).Output()
	if err != nil {
		return "", fmt.Errorf("Failed inspecting process with pid %d: %v", pid, err)
	}
	return string(lines), nil
}

func isProcessFinished(err error) bool {
	return err.Error() == "os: process already finished"
}
