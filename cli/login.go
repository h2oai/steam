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
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

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
