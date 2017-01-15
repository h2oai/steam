package haproxy

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os/exec"
	"text/template"

	"github.com/h2oai/steam/master/data"

	"github.com/pkg/errors"

	"io/ioutil"
	"os"
)

var config = `global
	daemon
defaults
    mode http
    timeout connect 5000ms
    timeout client  50000ms
    timeout server  50000ms
    option forwardfor
    option http-server-close

frontend h2o-clusters
    bind *{{.Port}}{{if .Cert}} ssl crt {{.Cert}}{{end}}
    reqadd X-Forwarded-Proto:\ https
    {{- range .Clus}}
    acl cluster_{{.Name}} path_beg {{.Ctxt}}
    use_backend {{.Name}} if cluster_{{.Name}}
    {{- end}}
{{range .Clus}}
backend {{.Name}}{{if .Toke}}
	http-request set-header Authorization Basic\ %[req.cook({{.Name}})]
	redirect scheme https if !{ ssl_fc }{{end}}
	server {{.Name}} {{.Addr}}
{{end}}
`

func (h *haProxyConfig) writeConfig() (*bytes.Buffer, error) {
	tmpl := template.New("Config")

	buf := new(bytes.Buffer)
	configTemplate, err := tmpl.Parse(config)
	if err != nil {
		return nil, err //FIXME format error
	}
	if err := configTemplate.Execute(buf, h); err != nil {
		return nil, err //FIXME: format error
	}

	return buf, nil

}

type haProxyConfig struct {
	Port string
	Cert string
	Clus []haCluster
}

type haCluster struct{ Name, Addr, Ctxt, Toke string }

func Reload(clusters []data.Cluster, port, certFilePath string) error {
	conf := haProxyConfig{
		Port: port,
		Cert: certFilePath,
		Clus: make([]haCluster, len(clusters)),
	}
	for i, cluster := range clusters {
		conf.Clus[i] = haCluster{Name: fmt.Sprintf("%s_%s", cluster.Name, cluster.Username.String),
			Addr: cluster.Address.String, Ctxt: cluster.ContextPath.String,
			Toke: cluster.Token.String}
	}

	buf, err := conf.writeConfig()
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile("haproxy.conf", buf.Bytes(), 0644); err != nil {
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
