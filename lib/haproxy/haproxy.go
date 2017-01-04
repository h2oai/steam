package haproxy

import (
	"github.com/h2oai/steam/master/data"
	"os/exec"
	"crypto/md5"
	"encoding/hex"
	"encoding/base64"
	"syscall"
	"context"

	"github.com/pkg/errors"

	"os"
	"io/ioutil"
)

func Reload(clusters []data.Cluster, uid, gid uint32) error {
	config :=
		"global\n" +
		"    daemon\n" +
		"    log 127.0.0.1 local0\n\n" +
		"defaults\n" +
		"    log global\n" +
		"    mode http\n" +
		"    timeout connect 5000ms\n" +
		"    timeout client  50000ms\n" +
		"    timeout server  50000ms\n\n" +
		"frontend h2o-clusters\n" +
		"    bind *:9999\n" +
		"    capture request header origin len 128\n"

	for _, c := range clusters {
		if c.ContextPath != "/" {
			config +=
				"    acl cluster_" + c.Name + " path_beg " + c.ContextPath + "\n" +
				"    use_backend " + c.Name + " if " + "cluster_" + c.Name + "\n\n"
		}
	}

	for _, c := range clusters {
		if c.ContextPath != "/" {
			config += "backend " + c.Name + "\n"
			if c.Token != "" {
				config += "    http-response add-header Access-Control-Allow-Origin %[capture.req.hdr(0)] if { capture.req.hdr(0) -m found }\n"
				config += "    rspadd Access-Control-Allow-Headers:\\ Origin,\\ X-Requested-With,\\ Content-Type,\\ Accept  if { capture.req.hdr(0) -m found }\n"
				config += "    http-request set-header Authorization Basic\\ %[req.cook(" + c.Name + ")]\n"
			}
			config += "    server " + c.Name + " " + c.Address + "\n\n"
		}
	}

	if err := ioutil.WriteFile("haproxy.conf",
		[]byte(config), 0644); err != nil {
		return err
	}

	args := []string{
		"-f", "haproxy.conf",
		"-p", "haproxy.pid",
		"-d", "-D",
	}

	pids, _ := ioutil.ReadFile("haproxy.pid")
	if pids != nil {
		args = append(args, "-sf", string(pids))
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmd := exec.CommandContext(ctx, "haproxy", args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uid, Gid: gid}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute command
	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "failed running command %s", cmd.Args)
	}

	return nil
}

func GenRealmFile(username, passwd string) (string, error) {
	entry := username + ": MD5:" + GetMD5Hash(passwd)
	token := base64.StdEncoding.EncodeToString([]byte(username + ":" + passwd))
	if err := ioutil.WriteFile(token + "_realm.properties", []byte(entry), 0644); err != nil {
		return "", err
	}
	return token, nil
}


func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
