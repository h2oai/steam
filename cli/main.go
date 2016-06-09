package cli

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/h2oai/steamY/lib/svc"
	"github.com/h2oai/steamY/master"
	"github.com/spf13/cobra"
)

const (
	steam = "steam"
)

type Config struct {
	Version     string
	Kind        string
	CurrentHost string
}

func newConfig() *Config {
	return &Config{
		"1.0.0",
		"Config",
		"",
	}
}

func Run(version, buildDate string) {
	cmd := Steam(version, buildDate, os.Stdout, os.Stdin, ioutil.Discard)
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func Steam(version, buildDate string, stdout, stderr, trace io.Writer) *cobra.Command {
	c := &context{
		version:   version,
		buildDate: buildDate,
		trace:     log.New(trace, "", 0),
		// remote:    &web.Remote{rpc.NewProc("http", "/web", "web", "172.16.2.103:9000", "", "")},
	}

	var verbose bool
	cmd := &cobra.Command{
		Use:               steam,
		Short:             fmt.Sprintf("%s v%s build %s: Command Line Interface to Steam", steam, version, buildDate),
		DisableAutoGenTag: true,

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			c.configure(verbose)
		},
	}
	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	cmd.AddCommand(
		delete(c),
		deploy(c),
		get(c),
		login(c),
		retrieve(c),
		serve(c),
		start(c),
		stop(c),
	)
	return cmd
}

func newCmd(c *context, help string, run func(c *context, args []string)) *cobra.Command {
	doc, err := parseHelp(help)
	if err != nil {
		log.Fatalln("Could not parse help:", err)
	}
	cmd := &cobra.Command{
		Use:   doc.Usage,
		Short: doc.Short,
		Long:  doc.Long,
	}
	if run != nil {
		cmd.Run = func(cmd *cobra.Command, args []string) {
			run(c, args)
		}
	}
	return cmd
}

type Doc struct {
	Usage string
	Short string
	Long  string
}

func parseHelp(text string) (*Doc, error) {
	d := strings.SplitN(strings.TrimSpace(text), "\n", 3)
	if len(d) != 3 {
		return nil, fmt.Errorf("Expected usage, short, long; found %d tokens: %s", len(d), text)
	}
	return &Doc{
		d[0],
		d[1],
		d[1] + "\n\n" + d[2],
	}, nil
}

//
// Commands
//

var serveHelp = `
serve [agent-type]
Lauch a new service.
Examples:

    $ steam serve master
`

func serve(c *context) *cobra.Command {
	cmd := newCmd(c, serveHelp, nil)
	cmd.AddCommand(serveMaster(c))
	return cmd
}

var serveMasterHelp = `
master
Launch the Steam master.
Examples:

    $ steam serve master
`

func serveMaster(c *context) *cobra.Command {
	var (
		webAddress                string
		workingDirectory          string
		clusterProxyAddress       string
		compilationServiceAddress string
		scoringServiceAddress     string
		enableProfiler            bool
		enableKerberos            bool
		username, keytab          string
	)

	opts := master.DefaultOpts

	cmd := newCmd(c, serveMasterHelp, func(c *context, args []string) {
		master.Run(c.version, c.buildDate, &master.Opts{
			webAddress,
			workingDirectory,
			clusterProxyAddress,
			compilationServiceAddress,
			scoringServiceAddress,
			enableProfiler,
			enableKerberos,
			username,
			keytab,
		})
	})

	cmd.Flags().StringVar(&webAddress, "web-address", opts.WebAddress, "Web server address.")
	cmd.Flags().StringVar(&workingDirectory, "working-directory", opts.WorkingDirectory, "Working directory for application files.")
	cmd.Flags().StringVar(&clusterProxyAddress, "cluster-proxy-address", opts.ClusterProxyAddress, "Address for cluster proxy")
	cmd.Flags().StringVar(&compilationServiceAddress, "compilation-service-address", opts.CompilationServiceAddress, "Compilation service address")
	cmd.Flags().StringVar(&scoringServiceAddress, "scoring-service-address", opts.ScoringServiceAddress, "Address to start scoring service on")
	cmd.Flags().BoolVar(&enableProfiler, "profile", opts.EnableProfiler, "Enable Go profiler")
	cmd.Flags().BoolVar(&enableKerberos, "kerberos", opts.KerberosEnabled, "Enable Kerberos authentication. Requires username and keytab.") // FIXME: Kerberos authentication is being passed by admin to all
	cmd.Flags().StringVar(&username, "username", opts.Username, "Username to enable Kerberos")
	cmd.Flags().StringVar(&keytab, "keytab", opts.Keytab, "Keytab file to be used with Kerberos authentication")
	return cmd

}

var startServiceHelp = `
service
Start a new scoring service
Examples:

Start a new scoring service instance using foo.war listening on port 8888
    $ steam start service --warfile=foo.war --port=8888
`

func startService(c *context) *cobra.Command {
	var (
		warfile string
		jetty   string
		address string
		port    int
	)
	cmd := newCmd(c, startServiceHelp, func(c *context, args []string) {
		pid, err := svc.Start(warfile, jetty, address, port)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Started process:", pid)
	})
	cmd.Flags().StringVar(&warfile, "warfile", "", "The WAR file to launch.")
	cmd.Flags().StringVar(&jetty, "jetty-runner", "", "The jetty runner jar.")
	cmd.Flags().StringVar(&address, "address", "0.0.0.0", "The ip of the host to launch the scoring service.")
	cmd.Flags().IntVar(&port, "port", 8000, "The port to listen on.")
	return cmd
}
