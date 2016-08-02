package cli2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"syscall"

	"github.com/h2oai/steamY/lib/fs"
	"github.com/h2oai/steamY/master"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var loginHelp = `
login [address:port] --username=[username] --password=[password]
Sign in to a Steam server.
Examples:

	$ steam login 192.168.42.42:9000 \
			--username=arthur
			--password=beeblebrox
`

func login(c *context) *cobra.Command {
	var (
		username             string
		password             string
		authenticationMethod string
		enableTLS            bool
	)
	cmd := newCmd(c, loginHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("*** Missing address. See 'steam help login'.")
		}
		address := args[0]

		if len(strings.TrimSpace(username)) == 0 {
			var err error
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Username: ")
			username, err = reader.ReadString('\n')
			if err != nil {
				log.Fatalln(err)
			}
			username = strings.TrimSpace(username)
		}

		if len(strings.TrimSpace(password)) == 0 {
			fmt.Print("Password: ")
			passwordBytes, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil {
				log.Fatalln(err)
			}
			password = strings.TrimSpace(string(passwordBytes))
		}

		c.config.CurrentHost = address
		c.config.Hosts[address] = &Host{
			username,
			password,
			authenticationMethod,
			enableTLS,
		}
		c.saveConfig(c.config)
		fmt.Println("Login credentials saved for server", address)
	})

	cmd.Flags().StringVar(&username, "username", "", "Login username")
	cmd.Flags().StringVar(&password, "password", "", "Login password")
	cmd.Flags().StringVar(&authenticationMethod, "authentication", "basic", "Authentication method")
	cmd.Flags().BoolVar(&enableTLS, "secure", false, "Enable TLS")

	return cmd
}

var resetHelp = `
reset
Reset Steam client configuration.
Examples:

    $ steam reset
`

func reset(c *context) *cobra.Command {
	cmd := newCmd(c, resetHelp, func(c *context, args []string) {
		if err := c.resetConfig(); err != nil {
			log.Fatalln("Failed configuration reset:", err)
			return
		}
		fmt.Println("Configuration reset successfully. Use 'steam login <server-address>' to re-authenticate to steam")
	})
	return cmd
}

var serveHelp = `
serve [agent-type]
Launch a new service.
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
		webTLSCertPath            string
		webTLSKeyPath             string
		authProvider              string
		workingDirectory          string
		clusterProxyAddress       string
		compilationServiceAddress string
		scoringServiceHost        string
		scoringServicePortsString string
		enableProfiler            bool
		yarnEnableKerberos        bool
		yarnUserName              string
		yarnKeytab                string
		dbName                    string
		dbUserName                string
		dbSSLMode                 string
		superuserName             string
		superuserPassword         string
	)

	opts := master.DefaultOpts

	cmd := newCmd(c, serveMasterHelp, func(c *context, args []string) {
		ports := strings.Split(scoringServicePortsString, ":")
		if len(ports) != 2 {
			log.Fatalln("Invalid usage of scoring service ports range. See 'steam help serve master'.")
		}
		var scoringServicePorts [2]int
		for i, port := range ports {
			var err error
			scoringServicePorts[i], err = strconv.Atoi(port)
			if err != nil {
				log.Fatalln("Invalid usage of scoring service ports range. See 'steam help serve master'.")
			}
			if scoringServicePorts[i] < 1025 || scoringServicePorts[i] > 65535 {
				log.Fatalln("Invalid port range.")
			}
		}
		if scoringServicePorts[0] > scoringServicePorts[1] {
			log.Fatalln("Invalid port range.")
		}

		master.Run(c.version, c.buildDate, master.Opts{
			webAddress,
			webTLSCertPath,
			webTLSKeyPath,
			authProvider,
			workingDirectory,
			clusterProxyAddress,
			compilationServiceAddress,
			scoringServiceHost,
			scoringServicePorts,
			enableProfiler,
			master.YarnOpts{
				yarnEnableKerberos,
				yarnUserName,
				yarnKeytab,
			},
			master.DBOpts{
				dbName,
				dbUserName,
				dbSSLMode,
				superuserName,
				superuserPassword,
			},
		})
	})

	cmd.Flags().StringVar(&webAddress, "web-address", opts.WebAddress, "Web server address (\"<ip>:<port>\" or \":<port>\").")
	cmd.Flags().StringVar(&webTLSCertPath, "web-tls-cert-path", opts.WebTLSCertPath, "Web server TLS certificate file path (optional).")
	cmd.Flags().StringVar(&webTLSKeyPath, "web-tls-key-path", opts.WebTLSKeyPath, "Web server TLS key file path (optional).")
	cmd.Flags().StringVar(&authProvider, "authentication-provider", opts.AuthProvider, "Authentication mechanismfor client logins (one of \"basic\" or \"digest\")")
	cmd.Flags().StringVar(&workingDirectory, "working-directory", opts.WorkingDirectory, "Working directory for application files.")
	cmd.Flags().StringVar(&clusterProxyAddress, "cluster-proxy-address", opts.ClusterProxyAddress, "Cluster proxy address (\"<ip>:<port>\" or \":<port>\")")
	cmd.Flags().StringVar(&compilationServiceAddress, "compilation-service-address", opts.CompilationServiceAddress, "Model compilation service address (\"<ip>:<port>\")")
	cmd.Flags().StringVar(&scoringServiceHost, "scoring-service-address", opts.ScoringServiceHost, "Address to start scoring services on (\"<ip>\")")
	// TODO: this uses a hardcoded port range, not the default const
	cmd.Flags().StringVar(&scoringServicePortsString, "scoring-service-port-range", "1025:65535", "Specified port range to create scoring services on. (\"<from>:<to>\")")
	cmd.Flags().BoolVar(&enableProfiler, "profile", opts.EnableProfiler, "Enable Go profiler")
	cmd.Flags().BoolVar(&yarnEnableKerberos, "yarn-enable-kerberos", opts.Yarn.KerberosEnabled, "Enable Kerberos authentication. Requires username and keytab.") // FIXME: Kerberos authentication is being passed by admin to all
	cmd.Flags().StringVar(&yarnUserName, "yarn-username", opts.Yarn.Username, "Username to enable Kerberos")
	cmd.Flags().StringVar(&yarnKeytab, "yarn-keytab", opts.Yarn.Keytab, "Keytab file to be used with Kerberos authentication")
	cmd.Flags().StringVar(&dbName, "db-name", opts.DB.Name, "Database name to use for application data storage")
	cmd.Flags().StringVar(&dbUserName, "db-username", opts.DB.Username, "Database username to connect as")
	cmd.Flags().StringVar(&dbSSLMode, "db-ssl-mode", opts.DB.SSLMode, "Database connection SSL mode: one of 'disable', 'require', 'verify-ca', 'verify-full'")
	cmd.Flags().StringVar(&superuserName, "superuser-name", opts.DB.SuperuserName, "Set superuser username (required for first-time-use only)")
	cmd.Flags().StringVar(&superuserPassword, "superuser-password", opts.DB.SuperuserPassword, "Set superuser password (required for first-time-use only)")

	return cmd

}

var deployHelp = `
deploy [resource-type]
Deploy a resource of the specified type.
Examples:

	$ steam deploy engine
`

func deploy(c *context) *cobra.Command {
	cmd := newCmd(c, deployHelp, nil)
	cmd.AddCommand(deployEngine(c))
	return cmd
}

var deployEngineHelp = `
engine [enginePath]
Deploy an H2O engine to Steam.
Examples:

	$ steam deploy engine --file-path=path/to/engine
`

func deployEngine(c *context) *cobra.Command {
	var (
		filePath string
	)
	cmd := newCmd(c, deployEngineHelp, func(c *context, args []string) {
		attrs := map[string]string{
			"type": fs.KindEngine,
		}

		if err := c.transmitFile(filePath, attrs); err != nil {
			log.Fatalln(err)
		}

		log.Println("Engine deployed:", path.Base(filePath))
	})

	cmd.Flags().StringVar(&filePath, "file-path", "", "Path to engine")

	return cmd
}
			log.Fatalln(err)
		}

	})

	return cmd
}
