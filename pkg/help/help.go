package help

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/pspiagicw/goreland"
)

func PrintUsage(version string) {
    goreland.LogInfo("gox version: %s", version)
    fmt.Println("Manage binary Go packages better")
    fmt.Println()
    fmt.Println("USAGE")
    fmt.Println("  gox [command] [args]")
    fmt.Println()
    fmt.Println("COMMANDS")

    commands := `
install:
remove:
list:
update:
update-all:
version:
help:`
    messages := `
Install a package
Remove a package
List all installed packages
Update a package
Update all packages
Show version info
Show this message`

    commandCol := lipgloss.NewStyle().Align(lipgloss.Left).SetString(commands).MarginLeft(2).String()
    messageCol := lipgloss.NewStyle().Align(lipgloss.Left).SetString(messages).MarginLeft(5).String()

    fmt.Println(lipgloss.JoinHorizontal(lipgloss.Bottom, commandCol, messageCol))

    fmt.Println()
    fmt.Println("MORE HELP")
    fmt.Println("  Use 'gox help [command]' for more info about a command.")

    fmt.Println()
    fmt.Println("EXAMPLES")
    fmt.Println("  $ gox install github.com/pspiagicw/groom@latest")
    fmt.Println("  $ gox install .")
    fmt.Println("  $ gox update groom")
    fmt.Println()


}

func HelpInstall() {
    fmt.Println("Install golang packages")
    fmt.Println()
    fmt.Println("USAGE")
    fmt.Println("  gox install [flags] <package>")
    fmt.Println()
    fmt.Println("ARGUMENTS")
    fmt.Println("  The package to install can be specified in any of the formats")
    fmt.Println(`  - A URL with a version specifier, e.g "golang.org/x/tools/gopls@latest"`)
    fmt.Println(`  - A PATH, either relative or absolute (avoid ~), e.g "cmd/mytool"`)
    fmt.Println()
    fmt.Println("EXAMPLES")
    fmt.Println("  $ gox install github.com/pspiagicw/gox@latest")
    fmt.Println("  $ gox install .")
    fmt.Println("  $ gox install /path/to/your/project")
    fmt.Println()

}
func HelpRemove() {
    fmt.Println("Remove golang packages")
    fmt.Println()
    fmt.Println("USAGE")
    fmt.Println("  gox remove [flags] <package-name>")
    fmt.Println()
    fmt.Println("ARGUMENTS")
    fmt.Println("  The package to remove is to be specified by the package name shown on `gox list`")
    fmt.Println()
    fmt.Println("EXAMPLES")
    fmt.Println("  $ gox remove gopls")
    fmt.Println()
}
func HelpUpdate() {
	// TODO: Print Help message for Update of package.
	goreland.LogInfo("Printing help for update command not implemented yet!")
}
func HelpArgs(args []string) {
    if len(args) == 0 {
        return
    }
    cmd := args[0]

    handlers := map[string]func() {
        "install": HelpInstall,
        "remove": HelpRemove,
        "update": HelpUpdate,
    }

    handlerFunc, exists := handlers[cmd]
    if exists {
        handlerFunc()
    } else {
        goreland.LogFatal("No help for command %s found", cmd)
    }
}
