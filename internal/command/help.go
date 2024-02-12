package command

import (
	"fmt"
	"github.com/Hydoc/goo/internal/view"
	"strings"
)

var HelpAliases = []string{"h"}

const (
	HelpAbbr = "help"
	HelpDesc = "Print help"
)

type Help struct {
	view          *view.StdoutView
	validCommands []*StringCommand
}

func (cmd *Help) Execute() {
	commandsAsStr := "Here is a list of all possible commands:\r\n"

	for i, strCmd := range cmd.validCommands {
		commandsAsStr += fmt.Sprintf("%s: %s (aliases: %s)", strCmd.Abbreviation, strCmd.Description, strings.Join(strCmd.Aliases, ", "))
		if i != len(cmd.validCommands)-1 {
			commandsAsStr += "\r\n"
		}
	}
	cmd.view.RenderLine(commandsAsStr)
}

func newHelp(view *view.StdoutView, validCommands []*StringCommand) *Help {
	return &Help{
		view:          view,
		validCommands: validCommands,
	}
}
