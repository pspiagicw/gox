package list

import (
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/gox/pkg/database"
)

func ListPackage(args []string) {
	gdb := database.ParseDatabase()

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
