package haproxy

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"os/exec"

	"github.com/h2oai/steam/master/data"

	"github.com/pkg/errors"

	"io/ioutil"
	"os"
)

func Reload(clusters []data.Cluster, uid, gid uint32) error {
	config :=
		"global\n" +
			"    daemon\n\n" +
			"defaults\n" +
			"    mode http\n" +
			"    timeout connect 5000ms\n" +
			"    timeout client  50000ms\n" +
			"    timeout server  50000ms\n" +
			"    option forwardfor\n" +
			"    option http-server-close\n\n" +
			"frontend h2o-clusters\n" +
			"    bind *:443 ssl crt ./steam_haproxy.pem\n" +
			"    reqadd X-Forwarded-Proto:\\ https\n"

	for _, c := range clusters {
		if c.ContextPath.String != "/" {
			config +=
				"    acl cluster_" + c.Name + " path_beg " + c.ContextPath.String + "\n" +
					"    use_backend " + c.Name + " if " + "cluster_" + c.Name + "\n\n"
		}
	}

	for _, c := range clusters {
		if c.ContextPath.String != "/" {
			config += "backend " + c.Name + "\n"
			if c.Token.String != "" {
				config += "    http-request set-header Authorization Basic\\ %[req.cook(" + c.Name + ")]\n"
				config += "    redirect scheme https if !{ ssl_fc }\n"
			}
			config += "    server " + c.Name + " " + c.Address.String + "\n\n"
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
	if err := ioutil.WriteFile(token+"_realm.properties", []byte(entry), 0644); err != nil {
		return "", err
	}
	return token, nil
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
