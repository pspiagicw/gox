package argparse

import (
	"flag"

	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/gox/pkg/help"
)

func ParseArguments(VERSION string) (string, []string) {
	flag.Usage = help.PrintUsage
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		help.PrintUsage()
		goreland.LogFatal("No subcommands provided!")
	}

	cmd, args := args[0], args[1:]

	return cmd, args
}
