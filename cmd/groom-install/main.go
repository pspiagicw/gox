package main

import (
	"flag"
	"os"

	"github.com/golang-groom/groom-install/pkg/compiler"
	"github.com/golang-groom/groom-install/pkg/installer"
	"github.com/pspiagicw/goreland"
)

var VERSION string

func main() {
	flags := compileFlags()
	argument := parseFlags(flags)
    goreland.LogSuccess("Installing '%s'", argument)
	binDir, entries, err := compiler.CompileProject(argument)
	if err != nil {
		goreland.LogError("Error while compiling: %v", err)
		os.Exit(1)
	}
	installer.InstallBinaries(binDir, entries)
}

func compileFlags() []string {
	version := flag.Bool("version", false, "Print version info")

	flag.Parse()

	if *version {
		goreland.LogInfo("groom-install version %s", VERSION)
		os.Exit(0)
	}
	return flag.Args()

}
func printUsage() {
    goreland.LogError("No argument provided!")
    goreland.LogSuccess("Provide a Go repo/location to install. Example...")
    goreland.LogExec("groom install github.com/golang-groom/groom/cmd/groom@latest")
    os.Exit(1)
}
func parseFlags(flags []string) string {
	if len(flags) == 0 {
		printUsage()
	} else if len(flags) > 1 {
		goreland.LogError("Only 1 argument expected!")
	}

	return flags[0]
}
