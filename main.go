package main

import (
	"github.com/pspiagicw/gox/pkg/argparse"
	"github.com/pspiagicw/gox/pkg/handle"
)

var VERSION string

func main() {
	cmd, args := argparse.ParseArguments(VERSION)
	handle.HandleArgs(cmd, args, VERSION)
}
