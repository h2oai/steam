package yarn

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// If there is a open ticket, does nothing, else kinits a new session
func kCheck(username, keytab string) error {
	if err := exec.Command("klist", "-s").Run(); err != nil { // returns no err if there's an open ticket
		if err = kInit(username, keytab); err != nil {
			return err
		}
	}

	return nil
}

func kInit(username, keytab string) error {

	if len(username) < 1 {
		return fmt.Errorf("A username is required to authenticate with Kerberos.")
	}
	if len(keytab) < 0 {
		return fmt.Errorf("A keytab is required to authenticate with Kerberos.")
	}

	cmd := exec.Command("kinit", username, "-k", "-t", keytab)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}

	// Scanning for errors while cmd is running
	go func() {
		in := bufio.NewScanner(stderr)
		for in.Scan() {
			log.Printf(in.Text())
		}
	}()

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("Failed to authenticate: %v", err)
	}

	return nil
}

func randStr(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	r := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		r[i] = chars[rand.Intn(len(chars))]
	}
	return string(r)
}

func cleanDir(dir string) error {
	log.Printf("Removing empty directory %s", dir)
	cmdClean := exec.Command("hadoop", "fs", "-rmdir", dir)
	if out, err := cmdClean.Output(); err != nil {
		log.Printf("Failed to remove outdir.\n%s", string(out))
		return err
	}
	return nil
}

// StartCloud starts a yarn cloud by shelling out to hadoop
//
// This process needs to store the job-ID to kill the process in the future
func StartCloud(size int, kerberos bool, mem, name, enginePath, username, keytab string) (string, string, string, error) {
	if kerberos {
		if err := kCheck(username, keytab); err != nil {
			return "", "", "", err
		}
	}

	out := "steam/" + name + "_" + randStr(5) + "_out"

	cmdArgs := []string{
		"jar",              //
		enginePath,         //
		"-jobname",         //
		"STEAM_" + name,    //
		"-n",               //
		strconv.Itoa(size), //
		"-mapperXmx",       //
		mem,                //
		"-output",          //
		out,                //
		"-disown",          //
	}

	log.Println("Attempting to start cloud...")

	cmd := exec.Command("hadoop", cmdArgs...)

	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		return "", "", "", err
	}
	stdErr, err := cmd.StderrPipe()
	if err != nil {
		return "", "", "", err
	}

	if err := cmd.Start(); err != nil {
		return "", "", "", err
	}

	reNode := regexp.MustCompile(`H2O node (\d+\.\d+\.\d+\.\d+:\d+)`)
	reApID := regexp.MustCompile(`application_(\d+_\d+)`)
	var cmdErr, apID, addr string
	go func() {
		in := bufio.NewScanner(stdOut)
		for in.Scan() {
			log.Printf("--YARN-- %s\n", in.Text())
			if s := reNode.FindSubmatch([]byte(in.Text())); s != nil {
				addr = string(s[1])
			}
			if s := reApID.FindSubmatch([]byte(in.Text())); s != nil {
				apID = string(s[1])
			}
			if strings.Contains(in.Text(), "ERROR") {
				cmdErr += in.Text() + "\n"
			}
		}
	}()
	go func() {
		in := bufio.NewScanner(stdErr)
		for in.Scan() {
			cmdErr += in.Text() + "\n"
		}
	}()

	// TODO should be a ticket system, not halting
	if err := cmd.Wait(); err != nil {
		cleanDir(name + "_out")
		return "", "", "", fmt.Errorf("Failed to launch hadoop.\n%s", cmdErr)
	}
	return apID, addr, out, nil
}

// StopCloud kills a hadoop cloud by shelling out a command based on the job-ID
//
// In the future this
func StopCloud(kerberos bool, name, id, outdir, username, keytab string) error {
	if kerberos {
		if err := kCheck(username, keytab); err != nil {
			return err
		}
	}

	log.Println("Attempting to stop cloud...")
	cmd := exec.Command("hadoop", "job", "-kill", "job_"+id)
	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stdErr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	e := make(chan error)
	var cmdOut, cmdErr string
	go func() {
		in := bufio.NewScanner(stdOut)
		for in.Scan() {
			log.Printf("--YARN-- %s\n", in.Text())
			cmdOut += in.Text() + "\n"
		}
	}()
	go func() {
		in := bufio.NewScanner(stdErr)
		for in.Scan() {
			cmdErr = cmdErr + in.Text() + "\n"
			if strings.Contains(in.Text(), "Exception") {
				e <- fmt.Errorf("Failed to shutdown hadoop.\n%s", in.Text())
			}
		}
	}()
	go func() {
		err := cmd.Wait()
		if err != nil {
			log.Printf("Failed to shutdown hadoop.\n%s", cmdErr)
			e <- fmt.Errorf("Failed to shutdown hadoop.\n%s", cmdErr)
		}
		e <- nil
	}()
	if err := <-e; err != nil {
		return err
	}

	if err := cleanDir(outdir); err != nil {
		return err
	}

	return nil
}
