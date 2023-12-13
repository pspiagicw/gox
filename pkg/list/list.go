package list

import (
	"flag"
	"fmt"

	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/gox/pkg/database"
	"github.com/pspiagicw/gox/pkg/help"
)

type options struct {
	simple bool
}

func parseListArgs(args []string) ([]string, *options) {
	flag := flag.NewFlagSet("gox list", flag.ExitOnError)
	options := new(options)
	flag.Usage = help.HelpList
	flag.BoolVar(&options.simple, "simple", false, "Print simple listing")
	flag.Parse(args)
	return flag.Args(), options
}

func ListPackage(args []string) {
	args, options := parseListArgs(args)
	gdb := database.ParseDatabase()

	if !options.simple {
		listDefault(gdb)
	} else {
		listSimple(gdb)
	}
}
func listDefault(gdb *database.GoxDatabase) {
	headers := []string{"Package Name", "Package Path", "Package URL"}
	rows := [][]string{}

	for name, pack := range gdb.Packages {
		rows = append(rows, []string{name, pack.Path, pack.URL})
	}

	if len(rows) == 0 {
		goreland.LogSuccess("No packages installed yet!")

	} else {
		goreland.LogTable(headers, rows)
	}

}
func listSimple(gdb *database.GoxDatabase) {
	fmt.Println("INSTALLED PACKAGES:\n")
	names := []string{}
	for name, _ := range gdb.Packages {
		names = append(names, name)
	}
	for _, name := range names {
		fmt.Printf("- [%s]\n", name)
	}
}
