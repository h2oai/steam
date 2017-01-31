package kerberos

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"

	"github.com/h2oai/steam/lib/fs"
	"github.com/h2oai/steam/master/data"

	"github.com/pkg/errors"
)

// Kinit starts or renews a kerberos session depending on the result of klist
func Kinit(keytabFile, principal string, uid, gid uint32) error {
	args := make([]string, 0)
	if klist(uid, gid) {
		args = append(args, "-R")
	} else {
		args = append(args, "-k", "-t", keytabFile, "-l", "1m")
	}

	cmd := exec.Command("kinit", append(args, principal)...)
	// Impersonate to verify that ticket is valid for user
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uid, Gid: gid}
	msg, err := cmd.CombinedOutput()
	return errors.Wrapf(err, "invalid keytab. Please delete keytab and upload a valid keytab: %s", msg)
}

// klist is primarily used to verify whether a user has a valid ticket or not
func klist(uid, gid uint32) bool {
	cmd := exec.Command("klist", "-s")
	// Impersonate to verify that ticket is valid for user
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uid, Gid: gid}
	return cmd.Run() == nil
}

func WriteKeytab(keytab data.Keytab, workingDirectory string, uid, gid int) (string, error) {
	fPath := path.Join(workingDirectory, fs.LibDir, fs.KindKeytab, strconv.FormatInt(keytab.Id, 10)+".keytab")
	if err := ioutil.WriteFile(fPath, keytab.File, 0600); err != nil {
		return "", errors.Wrap(err, "writing file")
	}
	return fPath, errors.Wrap(os.Chown(fPath, uid, gid), "changing ownership of file")
}

func DeleteKeytab(keytabFile, workingDirectory string) error {
	fPath := path.Join(workingDirectory, fs.LibDir, fs.KindKeytab, keytabFile)
	if _, err := os.Stat(fPath); os.IsNotExist(err) {
		return nil
	}
	return os.Remove(fPath)
}
