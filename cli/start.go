package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

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
		kerberos, details     bool
	)

	cmd := newCmd(c, startCloudHelp, func(c *context, args []string) {
		if len(args) != 2 {
			log.Fatalln("Incorrect number of arguments. See 'steam help start cloud'.")
		}

		name := args[0]
		engine := args[1]

		// --- add additional args here ---

		log.Println("Attempting to start cluster...")
		cloud, err := c.remote.StartCloud(name, engine, size, mem, username)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Cluster started at:", cloud.Address)

		if details {
			info, err := c.remote.GetCloudStatus(cloud.Name)
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Printf(
				`
	      Engine: %v
	     Version: %v
	 Total Nodes: %v
	Total Memory: %v
`, info.EngineName, info.EngineVersion, info.Size, info.Memory)
		}
		// TODO: name corresponds to id for purpose of stopCloud

	})

	cmd.Flags().IntVarP(&size, "size", "n", 1, "The number of nodes to provision.")
	cmd.Flags().StringVarP(&mem, "memory", "m", "10g", "The max amount of memory to use per node.")
	cmd.Flags().BoolVar(&kerberos, "kerberos", true, "Set false on systems with no kerberos authentication.")
	cmd.Flags().StringVar(&username, "username", "", "The valid kerberos username.")
	cmd.Flags().StringVar(&keytab, "keytab", "", "The name of the keytab file to use.")
	cmd.Flags().BoolVarP(&details, "details", "d", false, "Detailed cluster infomration")

	return cmd
}
