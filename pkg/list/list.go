package list

import (
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/gox/pkg/database"
)

func ListPackage(args []string) {
	gdb, err := database.ParseDatabase()

	if err != nil {
		goreland.LogFatal("Error parsing database: %v", err)
	}

	headers := []string{"Package Name", "Package Path"}
	rows := [][]string{}

	for name, pack := range gdb.Packages {
		rows = append(rows, []string{name, pack.Name})
	}

	if len(rows) == 0 {
		goreland.LogSuccess("No packages installed yet!")

	} else {
		goreland.LogTable(headers, rows)

	}

}
