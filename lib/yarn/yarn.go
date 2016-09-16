package yarn

import (
	"bufio"
	"context"
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

	if out, err := cmd.CombinedOutput(); err != nil {
		return errors.Wrapf(err, "failed executing kinit: %v", string(out))
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

func cleanDir(dir string, uid, gid uint32) {
	cmd := exec.Command("hadoop", "fs", "-rmdir", dir)
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uid, Gid: gid}

	if out, err := cmd.Output(); err != nil {
		log.Printf("failed to remove outdir %s: %v", dir, string(out))
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
		return "", "", errors.Wrapf(err, "failed running command: %v", cmdErr)
	}

	return appID, address, nil
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
	appID, address, err := yarnCommand(uid, gid, name, username, cmdArgs...)
	if err != nil {
		cleanDir(out, uid, gid)
		return "", "", "", errors.Wrap(err, "failed executing command")
	}

	return appID, address, out, nil
}

// StopCloud kills a hadoop cloud by shelling out a command based on the job-ID
func StopCloud(kerberos bool, name, id, outdir, username, keytab string) error {
	uid, gid, err := getUser(username)
	if err != nil {
		return errors.Wrap(err, "failed getting user")
	}

	// If kerberos enabled, initialize and defer destroy
	if kerberos {
		if err := kInit(username, keytab, uid, gid); err != nil {
			return errors.Wrap(err, "failed initializing kerberos")
		}
		defer kDest(uid, gid)
	}

	if _, _, err := yarnCommand(uid, gid, name, username, "job", "-kill", "job_"+id); err != nil {
		return errors.Wrap(err, "failed executing command")
	}

	cleanDir(outdir, uid, gid)
	return nil
}
