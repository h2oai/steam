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
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/h2oai/steam/master"
	"github.com/h2oai/steam/master/data"
	"github.com/spf13/cobra"
)

var debug bool

const (
	steam = "steam"
)

type Config struct {
	Version     string
	Kind        string
	CurrentHost string
	Hosts       map[string]*Host
}

type Host struct {
	Username             string
	Password             string
	AuthenticationMethod string
	EnableTLS            bool
}

func newConfig() *Config {
	return &Config{
		"1.0.0",
		"Config",
		"",
		make(map[string]*Host),
	}
}

func Run(version, buildDate string) {
	cmd := Steam(version, buildDate, os.Stdout, os.Stdin, ioutil.Discard)
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

var unauthenticatedCommands = []string{
	"steam reset",
	"steam login",
	"steam serve",
}

func requiresAuth(seq string) bool {
	for _, c := range unauthenticatedCommands {
		if strings.Contains(seq, c) {
			return false
		}
	}
	return true
}

func getCommandSequence(cmd *cobra.Command, seq *string) {
	*seq = cmd.Use + " " + *seq
	if cmd.HasParent() {
		cmd.VisitParents(func(p *cobra.Command) {
			getCommandSequence(p, seq)
		})
	}
}

func Steam(version, buildDate string, stdout, stderr, trace io.Writer) *cobra.Command {
	c := &context{
		version:   version,
		buildDate: buildDate,
		trace:     log.New(trace, "", 0),
	}

	var (
		verbose, setAdmin   bool
		workingDirectory    string
		dbDriver            string
		dbPath              string
		dbName              string
		dbUsername          string
		dbPassword          string
		dbHost              string
		dbPort              string
		dbConnectionTimeout string
		dbSSLMode           string
		dbSSLCertPath       string
		dbSSLKeyPath        string
		dbSSLRootCertPath   string
	)
	cmd := &cobra.Command{
		Use:               steam,
		Short:             fmt.Sprintf("%s v%s build %s: Command Line Interface to Steam", steam, version, buildDate),
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			dbOpts := data.DBOpts{
				Driver:      dbDriver,
				Path:        dbPath,
				Name:        dbName,
				User:        dbUsername,
				Pass:        dbPassword,
				Port:        dbPort,
				Host:        dbHost,
				SSLMode:     dbSSLMode,
				SSLCert:     dbSSLCertPath,
				SSLKey:      dbSSLKeyPath,
				SSLRootCert: dbSSLRootCertPath,
			}
			if err := master.SetAdmin(workingDirectory, dbOpts); err != nil {
				log.Fatalln(err)
			}
		},

		// CLI configuration / init is in here as a pre-run routine so that
		//   -v / --verbose is captured properly and used during config parsing.
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			var seq string
			getCommandSequence(cmd, &seq)
			c.configure(verbose, requiresAuth(seq))
		},
	}
	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	cmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Set this to debug")

	opts := master.DefaultOpts
	cmd.Flags().BoolVar(&setAdmin, "set-admin", false, "Set this flag to set the Steam local admin")
	cmd.Flags().StringVar(&dbDriver, "db-driver", opts.DBOpts.Driver, "Driver for sql implementation. (Supported types are \"sqlite3\" or \"postgres\")")
	cmd.Flags().StringVar(&dbPath, "db-path", opts.DBOpts.Path, "Set the path to a local database")
	cmd.Flags().StringVar(&dbName, "db-name", opts.DBOpts.Name, "Database name to use for application data storage (required)")
	cmd.Flags().StringVar(&dbUsername, "db-username", opts.DBOpts.User, "Database username (required)")
	cmd.Flags().StringVar(&dbPassword, "db-password", opts.DBOpts.Pass, "Database password (optional)")
	cmd.Flags().StringVar(&dbHost, "db-host", opts.DBOpts.Host, "Database host (optional, defaults to localhost")
	cmd.Flags().StringVar(&dbPort, "db-port", opts.DBOpts.Port, "Database port (optional, defaults to 5432)")
	cmd.Flags().StringVar(&dbConnectionTimeout, "db-connection-timeout", opts.DBOpts.ConnectionTimeout, "Database connection timeout (optional)")
	cmd.Flags().StringVar(&dbSSLMode, "db-ssl-mode", opts.DBOpts.SSLMode, "Database connection SSL mode: one of 'disable', 'require', 'verify-ca', 'verify-full'")
	cmd.Flags().StringVar(&dbSSLCertPath, "db-ssl-cert-path", opts.DBOpts.SSLCert, "Database connection SSL certificate path (optional)")
	cmd.Flags().StringVar(&dbSSLKeyPath, "db-ssl-key-path", opts.DBOpts.SSLKey, "Database connection SSL key path (optional)")
	cmd.Flags().StringVar(&dbSSLRootCertPath, "db-ssl-root-cert-path", opts.DBOpts.SSLRootCert, "Database connection SSL root certificate path (optional)")

	cmd.AddCommand(
		login(c),
		reset(c),
		serve(c),
		upload(c),
	)
	registerGeneratedCommands(c, cmd)
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
