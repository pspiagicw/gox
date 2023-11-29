package handle

import (
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/gox/pkg/help"
	"github.com/pspiagicw/gox/pkg/install"
	"github.com/pspiagicw/gox/pkg/list"
	"github.com/pspiagicw/gox/pkg/remove"
	"github.com/pspiagicw/gox/pkg/update"
)

func HandleArgs(cmd string, args []string, version string) {

	handlers := map[string]func([]string){
		"help":       help.HelpArgs,
		"install":    install.InstallPackage,
		"remove":     remove.RemovePackage,
		"list":       list.ListPackage,
		"version":    notImplemented,
		"update":     update.UpdatePackage,
		"update-all": update.UpdateAll,
	}

	handlerFunc, exists := handlers[cmd]

	if exists {
		handlerFunc(args)
	} else {
        help.PrintUsage(version)
	}
}

func notImplemented(args []string) {
	goreland.LogFatal("Not implemented this command yet!.")
}
