package main

import (
	"github.com/pspiagicw/gox/pkg/argparse"
	"github.com/pspiagicw/gox/pkg/handle"
)

var VERSION string

func main() {
	cmd, args := argparse.ParseArguments(VERSION)
	handle.HandleArgs(cmd, args)
	// goreland.LogSuccess("Installing '%s'", argument)

	// binDir, entry, err := compile.CompileProject(argument)
	//
	// if err != nil {
	// 	goreland.LogError("Error while compiling: %v", err)
	// 	os.Exit(1)
	// }
	// installer.InstallBinaries(binDir, entry, argument)
}
