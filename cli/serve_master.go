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

package cli

import (
	"github.com/h2oai/steam/master"
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
		authProvider              string
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
		adminName             string
		adminPassword         string
	)

	opts := master.DefaultOpts

	cmd := newCmd(c, serveMasterHelp, func(c *context, args []string) {
		master.Run(c.version, c.buildDate, master.Opts{
			webAddress,
			webTLSCertPath,
			webTLSKeyPath,
			authProvider,
			workingDirectory,
			clusterProxyAddress,
			compilationServiceAddress,
			scoringServiceHost,
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
				adminName,
				adminPassword,
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
	cmd.Flags().BoolVar(&enableProfiler, "profile", opts.EnableProfiler, "Enable Go profiler")
	cmd.Flags().BoolVar(&yarnEnableKerberos, "yarn-enable-kerberos", opts.Yarn.KerberosEnabled, "Enable Kerberos authentication. Requires username and keytab.") // FIXME: Kerberos authentication is being passed by admin to all
	cmd.Flags().StringVar(&yarnUserName, "yarn-username", opts.Yarn.Username, "Username to enable Kerberos")
	cmd.Flags().StringVar(&yarnKeytab, "yarn-keytab", opts.Yarn.Keytab, "Keytab file to be used with Kerberos authentication")
	cmd.Flags().StringVar(&dbName, "db-name", opts.DB.Name, "Database name to use for application data storage")
	cmd.Flags().StringVar(&dbUserName, "db-username", opts.DB.Username, "Database username to connect as")
	cmd.Flags().StringVar(&dbSSLMode, "db-ssl-mode", opts.DB.SSLMode, "Database connection SSL mode: one of 'disable', 'require', 'verify-ca', 'verify-full'")
	cmd.Flags().StringVar(&adminName, "admin-name", opts.DB.AdminName, "Set admin username (required for first-time-use only)")
	cmd.Flags().StringVar(&adminPassword, "admin-password", opts.DB.AdminPassword, "Set admin password (required for first-time-use only)")

	return cmd

}
