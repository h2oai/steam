/*
  Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

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

	"github.com/h2oai/steam/lib/fs"
	"github.com/h2oai/steam/master"
	"github.com/h2oai/steam/master/data"
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
		webAddress                   string
		webTLSCertPath               string
		webTLSKeyPath                string
		authProvider                 string
		authConfig                   string
		workingDirectory             string
		clusterProxyAddress          string
		compilationServiceAddress    string
		predictionServiceHost        string
		predictionServicePortsString string
		enableProfiler               bool
		yarnEnableKerberos           bool
		dbName                       string
		dbUserName                   string
		dbPassword                   string
		dbHost                       string
		dbPort                       string
		dbConnectionTimeout          string
		dbSSLMode                    string
		dbSSLCertPath                string
		dbSSLKeyPath                 string
		dbSSLRootCertPath            string
		superuserName                string
		superuserPassword            string
	)

	opts := master.DefaultOpts

	cmd := newCmd(c, serveMasterHelp, func(c *context, args []string) {
		ports := strings.Split(predictionServicePortsString, ":")
		if len(ports) != 2 {
			log.Fatalln("Invalid usage of prediction service ports range. See 'steam help serve master'.")
		}
		var predictionServicePorts [2]int
		for i, port := range ports {
			var err error
			predictionServicePorts[i], err = strconv.Atoi(port)
			if err != nil {
				log.Fatalln("Invalid usage of prediction service ports range. See 'steam help serve master'.")
			}
			if predictionServicePorts[i] < 1025 || predictionServicePorts[i] > 65535 {
				log.Fatalln("Invalid port range.")
			}
		}
		if predictionServicePorts[0] > predictionServicePorts[1] {
			log.Fatalln("Invalid port range.")
		}

		master.Run(c.version, c.buildDate, master.Opts{
			webAddress,
			webTLSCertPath,
			webTLSKeyPath,
			authProvider,
			authConfig,
			workingDirectory,
			clusterProxyAddress,
			compilationServiceAddress,
			predictionServiceHost,
			predictionServicePorts,
			enableProfiler,
			master.YarnOpts{
				yarnEnableKerberos,
			},
			master.DBOpts{
				data.Connection{
					dbName,
					dbUserName,
					dbPassword,
					dbHost,
					dbPort,
					dbConnectionTimeout,
					dbSSLMode,
					dbSSLCertPath,
					dbSSLKeyPath,
					dbSSLRootCertPath,
				},
				superuserName,
				superuserPassword,
			},
		})
	})
	cmd.Flags().StringVar(&webAddress, "web-address", opts.WebAddress, "Web server address (\"<ip>:<port>\" or \":<port>\").")
	cmd.Flags().StringVar(&webTLSCertPath, "web-tls-cert-path", opts.WebTLSCertPath, "Web server TLS certificate file path (optional).")
	cmd.Flags().StringVar(&webTLSKeyPath, "web-tls-key-path", opts.WebTLSKeyPath, "Web server TLS key file path (optional).")
	cmd.Flags().StringVar(&authProvider, "authentication-provider", opts.AuthProvider, "Authentication mechanism for client logins (one of \"basic\", \"digest\"), or \"basic-ldap\"")
	cmd.Flags().StringVar(&authConfig, "authentication-config", opts.AuthConfig, "Configuration file for authentication (used in \"basic-ldap\")")
	cmd.Flags().StringVar(&workingDirectory, "working-directory", opts.WorkingDirectory, "Working directory for application files.")
	cmd.Flags().StringVar(&clusterProxyAddress, "cluster-proxy-address", opts.ClusterProxyAddress, "Cluster proxy address (\"<ip>:<port>\" or \":<port>\")")
	cmd.Flags().StringVar(&compilationServiceAddress, "compilation-service-address", opts.CompilationServiceAddress, "Model compilation service address (\"<ip>:<port>\")")
	cmd.Flags().StringVar(&predictionServiceHost, "scoring-service-address", opts.PredictionServiceHost, "Hostname to start prediction services on (\"<ip>\")")
	cmd.Flags().MarkDeprecated("scoring-service-address", "please use \"prediction-service-host\"")
	cmd.Flags().StringVar(&predictionServiceHost, "prediction-service-host", opts.PredictionServiceHost, "Hostname to start prediction services on (\"<ip>\")")
	cmd.Flags().StringVar(&predictionServicePortsString, "scoring-service-port-range", "1025:65535", "Specified port range to create prediction services on. (\"<from>:<to>\")")
	cmd.Flags().MarkDeprecated("scoring-service-port-range", "please use \"prediction-service-port-range\"")
	cmd.Flags().StringVar(&predictionServicePortsString, "prediction-service-port-range", "1025:65535", "Specified port range to create prediction services on. (\"<from>:<to>\")")
	cmd.Flags().BoolVar(&enableProfiler, "profile", opts.EnableProfiler, "Enable Go profiler")
	cmd.Flags().BoolVar(&yarnEnableKerberos, "yarn-enable-kerberos", opts.Yarn.KerberosEnabled, "Enable Kerberos authentication. Requires username and keytab.") // FIXME: Kerberos authentication is being passed by admin to all
	// cmd.Flags().StringVar(&dbName, "db-name", opts.DB.Connection.DbName, "Database name to use for application data storage (required)")
	// cmd.Flags().StringVar(&dbUserName, "db-username", opts.DB.Connection.User, "Database username (required)")
	// cmd.Flags().StringVar(&dbPassword, "db-password", opts.DB.Connection.Password, "Database password (optional)")
	// cmd.Flags().StringVar(&dbHost, "db-host", opts.DB.Connection.Host, "Database host (optional, defaults to localhost")
	// cmd.Flags().StringVar(&dbPort, "db-port", opts.DB.Connection.Port, "Database port (optional, defaults to 5432)")
	// cmd.Flags().StringVar(&dbConnectionTimeout, "db-connection-timeout", opts.DB.Connection.ConnectionTimeout, "Database connection timeout (optional)")
	// cmd.Flags().StringVar(&dbSSLMode, "db-ssl-mode", opts.DB.Connection.SSLMode, "Database connection SSL mode: one of 'disable', 'require', 'verify-ca', 'verify-full'")
	// cmd.Flags().StringVar(&dbSSLCertPath, "db-ssl-cert-path", opts.DB.Connection.SSLCert, "Database connection SSL certificate path (optional)")
	// cmd.Flags().StringVar(&dbSSLKeyPath, "db-ssl-key-path", opts.DB.Connection.SSLKey, "Database connection SSL key path (optional)")
	// cmd.Flags().StringVar(&dbSSLRootCertPath, "db-ssl-root-cert-path", opts.DB.Connection.SSLRootCert, "Database connection SSL root certificate path (optional)")
	cmd.Flags().StringVar(&superuserName, "superuser-name", opts.DB.SuperuserName, "Set superuser username (required for first-time-use only)")
	cmd.Flags().StringVar(&superuserPassword, "superuser-password", opts.DB.SuperuserPassword, "Set superuser password (required for first-time-use only)")

	return cmd

}

var uploadHelp = `
upload [resource-type]
Upload a resource of the specified type.
Examples:

	$ steam upload file
`

func upload(c *context) *cobra.Command {
	cmd := newCmd(c, uploadHelp, nil)
	cmd.AddCommand(uploadFile(c))
	cmd.AddCommand(uploadEngine(c))
	return cmd
}

var uploadFileHelp = `
file [path]
Upload an asset to Steam. 
Examples:

	$ steam upload file \
		--file-path=? \
		--project-id=?
`

func uploadFile(c *context) *cobra.Command {
	var (
		filePath     string
		projectId    int64
		packageName  string
		relativePath string
	)
	cmd := newCmd(c, uploadFileHelp, func(c *context, args []string) {

		if projectId <= 0 {
			log.Fatalln("Invalid project Id")
		}

		if err := fs.ValidateName(packageName); err != nil {
			log.Fatalln("Invalid package name:", err)
		}

		attrs := map[string]string{
			"type":          fs.KindFile,
			"project-id":    strconv.FormatInt(projectId, 10),
			"package-name":  packageName,
			"relative-path": relativePath,
		}
		if err := c.transmitFile(filePath, attrs); err != nil {
			log.Fatalln(err)
		}

		log.Println("File uploaded:", path.Base(filePath))
	})

	cmd.Flags().StringVar(&filePath, "file-path", "", "File to be uploaded")
	cmd.Flags().Int64Var(&projectId, "project-id", 0, "Target project id")
	cmd.Flags().StringVar(&packageName, "package-name", "", "Target package")
	cmd.Flags().StringVar(&relativePath, "relative-path", "", "Relative path to copy file to")

	return cmd
}

var uploadEngineHelp = `
engine [path]
Upload an engine to Steam. 
Examples:

	$ steam upload engine \
		--file-path=?
`

func uploadEngine(c *context) *cobra.Command {
	var (
		filePath string
	)
	cmd := newCmd(c, uploadEngineHelp, func(c *context, args []string) {
		attrs := map[string]string{
			"type": fs.KindEngine,
		}
		if err := c.transmitFile(filePath, attrs); err != nil {
			log.Fatalln(err)
		}

		log.Println("Engine uploaded:", path.Base(filePath))
	})

	cmd.Flags().StringVar(&filePath, "file-path", "", "File to be uploaded")

	return cmd
}
