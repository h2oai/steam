package cli

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/h2oai/steamY/lib/svc"
	"github.com/h2oai/steamY/lib/yarn"
	"github.com/h2oai/steamY/master"
	"github.com/spf13/cobra"
)

const (
	steam = "steam"
)

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
	}

	cmd := &cobra.Command{
		Use:               steam,
		Short:             fmt.Sprintf("%s v%s build %s: Command Line Interface to Steam", steam, version, buildDate),
		DisableAutoGenTag: true,
	}

	cmd.AddCommand(
		start(c), // temporary; will not be accessible from the CLI in the future
		stop(c),  // temporary; will not be accessible from the CLI in the future
		serve(c),
	)
	return cmd
}

type context struct {
	version   string
	buildDate string
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
	cmd.Flags().StringVar(&compilationServiceAddress, "compilation-service-address", opts.CompilationServiceAddress, "Compilation service address")
	cmd.Flags().StringVar(&scoringServiceAddress, "scoring-service-address", opts.ScoringServiceAddress, "Address to start scoring service on")
	cmd.Flags().BoolVar(&enableProfiler, "profile", opts.EnableProfiler, "Enable Go profiler")
	cmd.Flags().BoolVar(&enableKerberos, "kerberos", opts.KerberosEnabled, "Enable Kerberos authentication. Requires username and keytab.") // FIXME: Kerberos authentication is being passed by admin to all
	cmd.Flags().StringVar(&username, "username", opts.Username, "Username to enable Kerberos")
	cmd.Flags().StringVar(&keytab, "keytab", opts.Keytab, "Keytab file to be used with Kerberos authentication")
	return cmd

}

var startHelp = `
start [resource-type]
Start a new resource.
Examples:

    $ steam start cloud
`

func start(c *context) *cobra.Command {
	cmd := newCmd(c, startHelp, nil)
	cmd.AddCommand(startCloud(c))
	cmd.AddCommand(startService(c))
	return cmd
}

var startCloudHelp = `
cloud [cloud-name]
Start a new cloud using the specified H2O package.
Examples:

Start a 4 node H2O 3.2.0.9 cloud

    $ steam start cloud42 h2odriver.jar --size=4
`

func startCloud(c *context) *cobra.Command {
	var (
		size                  int
		mem, keytab, username string
		kerberos              bool
	)

	cmd := newCmd(c, startCloudHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Incorrect number of arguments. See 'steam help start cloud'.")
		}

		name := args[0]
		engine := args[1]

		// --- add additional args here ---

		if _, _, err := yarn.StartCloud(size, kerberos, mem, name, engine, username, keytab); err != nil {
			log.Fatalln(err)
		}

		// TODO: name corresponds to id for purpose of stopCloud

	})
	cmd.Flags().IntVar(&size, "size", 1, "The number of nodes to provision.")
	cmd.Flags().StringVar(&mem, "mem", "10g", "The max amount of memory to use per node.")
	cmd.Flags().BoolVar(&kerberos, "kerberos", true, "Set false on systems with no kerberos authentication.")
	cmd.Flags().StringVar(&username, "username", "", "The valid kerberos username.")
	cmd.Flags().StringVar(&keytab, "keytab", "", "The name of the keytab file to use")

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

var stopHelp = `
stop [resource-type]
Stop the specified resource.
Examples:

    $ steam stop cloud
`

func stop(c *context) *cobra.Command {
	cmd := newCmd(c, stopHelp, nil)
	cmd.AddCommand(stopCloud(c))
	cmd.AddCommand(stopService(c))
	return cmd
}

var stopCloudHelp = `
cloud [cloud-name] [cloud-id]
Stop a cloud.
Examples:

    $ steam stop cloud cloud42 1457562501251_0543
`

func stopCloud(c *context) *cobra.Command {
	var (
		kerberos, force  bool
		username, keytab string
	)

	cmd := newCmd(c, stopCloudHelp, func(c *context, args []string) {
		if len(args) != 2 {
			log.Fatalln("Missing cloud-name. See 'steam help stop cloud'.")
		}

		name := args[0]
		id := args[1] // FIXME: This should be a function of the name
		// --- add additional args here ---

		if err := yarn.StopCloud(kerberos, name, id, username, keytab); err != nil {
			log.Fatalln(err)
		}

	})

	cmd.Flags().BoolVar(&kerberos, "kerberos", true, "Set false on systems with no kerberos authentication.")
	cmd.Flags().StringVar(&username, "username", "", "The valid kerberos username.")
	cmd.Flags().StringVar(&keytab, "keytab", "", "The name of the keytab file to use")
	cmd.Flags().BoolVar(&force, "force", false, "Force-kill all H2O instances in the cloud")

	return cmd
}

var stopServiceHelp = `
service
Stop a scoring service.
Examples:

    $ steam stop service --pid=67997
`

func stopService(c *context) *cobra.Command {
	var (
		pid int
	)

	cmd := newCmd(c, stopServiceHelp, func(c *context, args []string) {
		if pid == 0 {
			log.Fatalln("Invalid pid. See 'steam help stop service'")
		}
		if err := svc.Stop(pid); err != nil {
			log.Fatalln(err)
		}
		log.Println("Service stopped:", pid)
	})

	cmd.Flags().IntVar(&pid, "pid", 0, "The pid of the service to kill.")

	return cmd
}
