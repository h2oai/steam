package yarn

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os/exec"
	"os/user"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/pkg/errors"
)

func kInit(username, keytab string, uid, gid uint32) error {
	cmd := exec.Command("kinit", username, "-k", "-t", keytab)
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uid, Gid: gid}

	if err := cmd.Run(); err != nil {
		return errors.Wrap(err, "failed executing kinit")
	}

	return nil
}

func kDest(uid, gid uint32) {
	cmd := exec.Command("kdestroy")
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uid, Gid: gid}

	if err := cmd.Run(); err != nil {
		panic(fmt.Sprintf("failed executing kdestroy: %v", err))
	}
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

func cleanDir(dir string) {
	// log.Printf("Removing empty directory %s", dir)
	cmdClean := exec.Command("hadoop", "fs", "-rmdir", dir)
	if out, err := cmdClean.Output(); err != nil {
		panic(fmt.Sprintf("failed to remove outdir: %v", string(out)))
	}
}

func getUser(username string) (uint32, uint32, error) {
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

func yarnScan(r io.Reader, name string, appID, address, err *string) {
	reNode := regexp.MustCompile(`H2O node (\d+\.\d+\.\d+\.\d+:\d+)`)
	reApID := regexp.MustCompile(`application_(\d+_\d+)`)

	in := bufio.NewScanner(r)
	for in.Scan() {
		log.Println(fmt.Sprintf("YARN CLUSTER %s %v", name, in.Text()))

		if appID != nil {
			if s := reNode.FindSubmatch(in.Bytes()); s != nil {
				*address = string(s[1])
			}
		}
		if address != nil {
			if s := reApID.FindSubmatch(in.Bytes()); s != nil {
				*appID = string(s[1])
			}
		}
		if err != nil {
			if strings.Contains(in.Text(), "ERROR") {
				*err += in.Text() + "\n"
			}
		}
	}
}

// StartCloud starts a yarn cloud by shelling out to hadoop
//
// This process needs to store the job-ID to kill the process in the future
func StartCloud(size int, kerberos bool, mem, name, enginePath, username, keytab string) (string, string, string, error) {
	// Get user information for Kerberos and Yarn reasons
	uid, gid, err := getUser(username)
	if err != nil {
		return "", "", "", errors.Wrap(err, "failed getting user")
	}

	// If kerberos enabled, initialize and defer destroy
	if kerberos {
		if err := kInit(username, keytab, uid, gid); err != nil {
			return "", "", "", errors.Wrap(err, "failed initializing kerberos")
		}
		defer kDest(uid, gid)
	}

	// Randomize outfile name
	out := "steam/" + name + "_" + randStr(5) + "_out"

	cmdArgs := []string{
		"jar", enginePath,
		"-jobname", "STEAM_" + name,
		"-n", strconv.Itoa(size),
		"-mapperXmx", mem,
		"-output", out,
		"-disown",
	}

	// Execute command as user
	cmd := exec.Command("hadoop", cmdArgs...)
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uid, Gid: gid}

	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		return "", "", "", err
	}
	defer stdOut.Close()
	stdErr, err := cmd.StderrPipe()
	if err != nil {
		return "", "", "", err
	}
	defer stdErr.Close()

	if err := cmd.Start(); err != nil {
		return "", "", "", err
	}

	var appID, address, cmdErr string
	go yarnScan(stdOut, name, &appID, &address, &cmdErr)
	go yarnScan(stdOut, name, nil, nil, nil)

	if err := cmd.Wait(); err != nil {
		// cleanDir(out)
		// recover()
		return "", "", "", errors.Wrapf(err, "failed waiting on exec: %v", cmdErr)
	}

	return appID, address, out, nil
}

// StopCloud kills a hadoop cloud by shelling out a command based on the job-ID
//
// In the future this
func StopCloud(kerberos bool, name, id, outdir, username, keytab string) error {
	uid, gid, err := getUser(username)
	if err != nil {
		return errors.Wrap(err, "failed getting user")
	}

	if kerberos {
		if err := kInit(username, keytab, uid, gid); err != nil {
			return errors.Wrap(err, "failed initializing kerberos")
		}
		defer kDest(uid, gid)
	}

	cmd := exec.Command("hadoop", "job", "-kill", "job_"+id)
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uid, Gid: gid}
	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	defer stdOut.Close()
	stdErr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	defer stdErr.Close()
	go yarnScan(stdOut, name, nil, nil, nil)
	go yarnScan(stdErr, name, nil, nil, nil)

	defer cleanDir(outdir)
	return cmd.Run()
}
