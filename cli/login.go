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
login [address]
Access a Steam server.
Examples:

	$ steam login http://192.168.42.42:9000
	$ steam login http://192.168.42.42:9000 \
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
