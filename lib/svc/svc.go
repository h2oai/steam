package svc

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// Start starts a scoring service.
func Start(warfile string, port int) (int, error) {

	argv := []string{"10000"} // FIXME use jetty-runner args

	cmd := exec.Command("/bin/sleep", argv...) // FIXME use jetty-runner

	if err := cmd.Start(); err != nil {
		return 0, fmt.Errorf("Failed starting scoring service for %s at port %d: %v", warfile, port, err)
	}

	return cmd.Process.Pid, nil
}

// Stop stops a scoring service.
func Stop(pid int) error {
	p, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("Failed locating scoring service with pid %d: %v", pid, err)
	}

	const sigintInterval = time.Second
	const sigintTimeout = time.Second * 10

	// retry SIGINTs;  SIGKILL on timeout
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

func isProcessFinished(err error) bool {
	return err.Error() == "os: process already finished"
}
