package argparse

import (
	"flag"

	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/gox/pkg/help"
)

func ParseArguments(VERSION string) (string, []string) {

	flag.Usage = func() {
		help.PrintUsage(VERSION)
	}

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		help.PrintUsage(VERSION)
		goreland.LogFatal("No subcommands provided!")
	}

	cmd, args := args[0], args[1:]

	return cmd, args
}
