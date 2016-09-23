package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path"
	"text/tabwriter"
	"time"

	"github.com/h2oai/steam/lib/rpc"
	"github.com/h2oai/steam/srv/web"
)

var confPath string

type context struct {
	version   string
	buildDate string
	config    *Config
	uploadURL string
	remote    *web.Remote
	trace     *log.Logger
}

func (c *context) getConfigPath() string {
	return path.Join(os.ExpandEnv("$HOME"), ".steam", "config")
}

func (c *context) configure(verbose, requiresAuth bool) {
	if verbose {
		c.trace = log.New(os.Stdout, "", 0)
	}

	confPath = c.getConfigPath()
	c.traceln("Config path:", confPath)

	var conf *Config
	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		c.traceln("Config not found. Creating...")
		confDir := path.Dir(confPath)
		if err := os.MkdirAll(confDir, 0755); err != nil {
			log.Fatalln(err)
		}
		conf = newConfig()
		c.saveConfig(conf)
	} else {
		c.traceln("Config found. Loading...")
		conf = c.loadConfig(confPath)
	}

	c.config = conf
	addr := conf.CurrentHost
	if !requiresAuth {
		return
	}
	host, ok := conf.Hosts[addr]
	if !ok {
		log.Fatalln("You are not authenticated to a Steam server. See 'steam help login' for more details.")
	}
	httpScheme := "http"
	if host.EnableTLS {
		httpScheme = "https"
	}
	c.remote = &web.Remote{rpc.NewProc(httpScheme, "/web", "web", addr, host.Username, host.Password)}
	c.uploadURL = (&url.URL{Scheme: httpScheme, Host: addr, Path: "/upload"}).String()

}

func (c *context) loadConfig(confPath string) *Config {
	c.traceln("Reading config...")
	b, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Failed reading config %s:", confPath), err)
	}

	c.traceln("Decoding config...")
	var conf Config
	if err := json.Unmarshal(b, &conf); err != nil {
		log.Fatalln("Failed decoding config:", err)
	}

	return &conf
}

func (c *context) saveConfig(conf *Config) {
	data, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		log.Fatalln("Failed marshaling config: ", err)
	}

	if err := ioutil.WriteFile(confPath, data, 0755); err != nil {
		log.Fatalln(fmt.Sprintf("Failed writing config file %s:", confPath), err)
	}
}

func (c *context) resetConfig() error {
	confPath = c.getConfigPath()
	if _, err := os.Stat(confPath); err == nil {
		return os.Remove(confPath)
	}
	return nil
}

func (c *context) uploadFile(filepath, kind string) error {
	return uploadFile(c.uploadURL, filepath, kind)
}

func (c *context) traceln(v ...interface{}) {
	c.trace.Println(v)
}

func (c *context) tracef(format string, v ...interface{}) {
	c.trace.Printf(format, v...)
}

func (c *context) printt(header string, lines []string) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 1, '\t', 0)
	fmt.Fprintln(w, header)
	for _, l := range lines {
		fmt.Fprintln(w, l)
	}
	fmt.Fprintln(w)
	w.Flush()
}

func fmtAgo(ts int64) string {
	return fmt.Sprint(time.Unix(ts, 0))
}
