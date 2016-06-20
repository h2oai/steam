package cli

import (
	"fmt"
	"log"

	// "github.com/h2oai/steamY/srv/web"
	"github.com/spf13/cobra"
)

var getHelp = `
get [resource-type]
List or view resources of the specified type.
Examples:

    $ steam get clouds
`

func get(c *context) *cobra.Command {
	cmd := newCmd(c, getHelp, nil)
	cmd.AddCommand(getClouds(c))
	cmd.AddCommand(getEngines(c))
	cmd.AddCommand(getModels(c))
	cmd.AddCommand(getServices(c))
	return cmd
}

var getCloudsHelp = `
clouds
List all clouds.
Examples:

	$ steam get clouds
`

func getClouds(c *context) *cobra.Command {
	var details, silent bool
	cmd := newCmd(c, getCloudsHelp, func(c *context, args []string) {

		// FIXME

		// cs, err := c.remote.GetClouds()
		// if err != nil {
		// 	log.Fatalln(err)
		// }

		// lines := make([]string, len(cs))
		// if details {
		// 	for i, cl := range cs {
		// 		var info *web.Cloud
		// 		if cl.State == web.CloudStopped {
		// 			info = cl
		// 		} else {
		// 			info, err = c.remote.GetCloudStatus(cl.Name)
		// 			if err != nil && silent {
		// 				log.Println(err)
		// 				info = cl
		// 				info.State = web.CloudUnknown
		// 			} else if err != nil {
		// 				log.Fatalln(err)
		// 			}
		// 		}
		// 		lines[i] = fmt.Sprintf("%s\t%s\t%s\t%d\t%s", info.Name, info.EngineName, info.Memory, info.Size, info.State)
		// 	}
		// 	c.printt("NAME\tENGINE\tMEMORY\tSIZE\tSTATE", lines)
		// } else {
		// 	lines := make([]string, len(cs))
		// 	for i, cl := range cs {
		// 		lines[i] = fmt.Sprintf("%s\t%s\t%s", cl.Name, cl.EngineName, cl.State)
		// 	}
		// 	c.printt("NAME\tENGINE\tSTATE", lines)
		// }
	})

	cmd.Flags().BoolVarP(&details, "details", "d", false, "Detailed cluster information")
	cmd.Flags().BoolVarP(&silent, "silent", "s", false, "Silence errors")

	return cmd
}

var getEnginesHelp = `
engines
List all engines.
Examples:

	$ steam get engines
`

func getEngines(c *context) *cobra.Command {
	cmd := newCmd(c, getEnginesHelp, func(c *context, args []string) {
		es, err := c.remote.GetEngines()
		if err != nil {
			log.Fatalln(err)
		}

		lines := make([]string, len(es))
		for i, e := range es {
			lines[i] = fmt.Sprintf("%s\t%s", e.Name, fmtAgo(e.CreatedAt))
		}
		c.printt("NAME\tAGE", lines)
	})

	return cmd
}

// FIXME: getCloud requires storage of all nodes in cluster

var getModelsHelp = `
models
List all models.
Examples:

	$ steam get models
`

func getModels(c *context) *cobra.Command {
	cmd := newCmd(c, getModelsHelp, func(c *context, args []string) {

		// FIXME

		ms, err := c.remote.GetModels(0, 10000)
		if err != nil {
			log.Fatalln(err)
		}

		lines := make([]string, len(ms))
		for i, m := range ms {
			lines[i] = fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t", m.Name, m.Algorithm, m.DatasetName, m.ResponseColumnName, fmtAgo(m.CreatedAt))
		}
		c.printt("NAME\tALGO\tDATASET\tTARGET\tAGE\t", lines)
	})

	return cmd
}

var getServicesHelp = `
services
List all services.
Examples:

	$ steam get services
`

func getServices(c *context) *cobra.Command {
	cmd := newCmd(c, getServicesHelp, func(c *context, args []string) {

		// FIXME

		// ss, err := c.remote.GetScoringServices()
		// if err != nil {
		// 	log.Fatalln(err)
		// }

		// lines := make([]string, len(ss))
		// for i, s := range ss {
		// 	lines[i] = fmt.Sprintf("%s\t%s\t%d\t%s\t%s", s.ModelName, s.Address, s.ModelName, s.Port, fmtAgo(s.CreatedAt))
		// }

		// c.printt("M.NAME\tADDRESS\tPORT\tSTATE\tAGE", lines)
	})

	return cmd
}
