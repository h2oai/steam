package cli

import (
	"log"
	"path"

	"github.com/h2oai/steamY/lib/fs"
	"github.com/spf13/cobra"
)

var deployEngineHelp = `
engine [enginePath]
Deploy an H2O engine to Steam.
Examples:

	$ steam deploy engine path/to/engine
`

func deployEngine(c *context) *cobra.Command {
	cmd := newCmd(c, deployEngineHelp, func(c *context, args []string) {
		if len(args) != 1 {
			log.Fatalln("Incorrect number of arguments. See 'steam help deploy engine'.")
		}

		enginePath := args[0]

		if err := c.uploadFile(enginePath, fs.KindEngine); err != nil {
			log.Fatalln(err)
		}

		log.Println("Engine deployed:", path.Base(enginePath))
	})

	return cmd
}
