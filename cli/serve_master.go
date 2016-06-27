package cli

import (
	"github.com/h2oai/steamY/master"
	"github.com/spf13/cobra"
)

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
		workingDirectory          string
		clusterProxyAddress       string
		compilationServiceAddress string
		scoringServiceHost        string
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
		master.Run(c.version, c.buildDate, &master.Opts{
			webAddress,
			webTLSCertPath,
			webTLSKeyPath,
			workingDirectory,
			clusterProxyAddress,
			compilationServiceAddress,
			scoringServiceHost,
			enableProfiler,
			yarnEnableKerberos,
			yarnUserName,
			yarnKeytab,
			dbName,
			dbUserName,
			dbSSLMode,
			superuserName,
			superuserPassword,
		})
	})

	cmd.Flags().StringVar(&webAddress, "web-address", opts.WebAddress, "Web server address (\"<ip>:<port>\" or \":<port>\").")
	cmd.Flags().StringVar(&webTLSCertPath, "web-tls-cert-path", opts.WebTLSCertPath, "Web server TLS certificate file path (optional).")
	cmd.Flags().StringVar(&webTLSKeyPath, "web-tls-key-path", opts.WebTLSKeyPath, "Web server TLS key file path (optional).")
	cmd.Flags().StringVar(&workingDirectory, "working-directory", opts.WorkingDirectory, "Working directory for application files.")
	cmd.Flags().StringVar(&clusterProxyAddress, "cluster-proxy-address", opts.ClusterProxyAddress, "Cluster proxy address (\"<ip>:<port>\" or \":<port>\")")
	cmd.Flags().StringVar(&compilationServiceAddress, "compilation-service-address", opts.CompilationServiceAddress, "Model compilation service address (\"<ip>:<port>\")")
	cmd.Flags().StringVar(&scoringServiceHost, "scoring-service-address", opts.ScoringServiceHost, "Address to start scoring services on (\"<ip>\")")
	cmd.Flags().BoolVar(&enableProfiler, "profile", opts.EnableProfiler, "Enable Go profiler")
	cmd.Flags().BoolVar(&yarnEnableKerberos, "yarn-enable-kerberos", opts.YarnKerberosEnabled, "Enable Kerberos authentication. Requires username and keytab.") // FIXME: Kerberos authentication is being passed by admin to all
	cmd.Flags().StringVar(&yarnUserName, "yarn-username", opts.YarnUserName, "Username to enable Kerberos")
	cmd.Flags().StringVar(&yarnKeytab, "yarn-keytab", opts.YarnKeytab, "Keytab file to be used with Kerberos authentication")
	cmd.Flags().StringVar(&dbName, "db-name", opts.DBName, "Database name to use for application data storage")
	cmd.Flags().StringVar(&dbUserName, "db-username", opts.DBUserName, "Database username to connect as")
	cmd.Flags().StringVar(&dbSSLMode, "db-ssl-mode", opts.DBSSLMode, "Database connection SSL mode: one of 'disable', 'require', 'verify-ca', 'verify-full'")
	cmd.Flags().StringVar(&superuserName, "superuser-name", opts.SuperuserName, "Set superuser username (required for first-time-use only)")
	cmd.Flags().StringVar(&superuserPassword, "superuser-password", opts.SuperuserPassword, "Set superuser password (required for first-time-use only)")

	return cmd

}
