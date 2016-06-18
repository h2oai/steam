package svc

import (
	"bufio"
	"encoding/json"
	"fmt"
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

	"github.com/h2oai/steamY/srv/web"
)

// Start starts a scoring service.
func Start(warfile, jetty, address string, port int) (int, error) {

	argv := []string{
		"-jar",
		jetty,
		"--host",
		address,
		"--port",
		strconv.Itoa(port),
		warfile,
	}

	cmd := exec.Command("java", argv...)
	stdErr, err := cmd.StderrPipe()
	if err != nil {
		return 0, err
	}
	// stdOut, err := cmd.StdoutPipe()
	// if err != nil {
	// 	return 0, err
	// }

	defer stdErr.Close() // Wait may not necessarily close, so pipe should close

	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	var errText string
	if err := cmd.Start(); err != nil {
		return 0, err
	}

	e := make(chan error)
	success := false
	// the cool bits
	// goroutines will only pipe to e if their condition is met
	go func() { // This is the success condition
		in := bufio.NewScanner(stdErr)
		for in.Scan() {
			errText = errText + in.Text() + "\n"
			if strings.Contains(in.Text(), "Started @") {
				e <- nil
				success = true
			}
		}
	}()
	go func() { // This is the fail condition
		if err := cmd.Wait(); err != nil {
			if !success {
				log.Printf("Failed starting scoring service for %s at  %s:%d:\n%v", warfile, address, port, errText)
				e <- fmt.Errorf("Failed starting scoring service for %s at  %s:%d:\n%v", warfile, address, port, errText)
			}
		}
	}()
	// Blocking here, until either goroutine returns with their condition
	if err := <-e; err != nil {
		return 0, err
	}
	pid := cmd.Process.Pid
	cmd.Process.Release()
	return pid, nil
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
