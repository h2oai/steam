package cli2

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

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

	var verbose bool
	cmd := &cobra.Command{
		Use:               steam,
		Short:             fmt.Sprintf("%s v%s build %s: Command Line Interface to Steam", steam, version, buildDate),
		DisableAutoGenTag: true,

		// CLI configuration / init is in here as a pre-run routine so that
		//   -v / --verbose is captured properly and used during config parsing.
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			var seq string
			getCommandSequence(cmd, &seq)
			c.configure(verbose, requiresAuth(seq))
		},
	}
	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	cmd.AddCommand(
		login(c),
		reset(c),
		serve(c),
		deploy(c),
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
