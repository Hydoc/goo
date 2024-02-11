package command

import (
	"fmt"
	"strings"
)

var HelpAliases = []string{"h"}

const (
	HelpAbbr = "help"
	HelpDesc = "Print help"
)

type Help struct {
	validCommands []*StringCommand
}

func (h *Help) Execute() {
	commandsAsStr := "Here is a list of all possible commands:\r\n"

	for i, strCmd := range h.validCommands {
		commandsAsStr += fmt.Sprintf("%s: %s (aliases: %s)", strCmd.Abbreviation, strCmd.Description, strings.Join(strCmd.Aliases, ", "))
		if i != len(h.validCommands)-1 {
			commandsAsStr += "\r\n"
		}
	}

	fmt.Println(commandsAsStr)
}

func newHelp(validCommands []*StringCommand) *Help {
	return &Help{
		validCommands: validCommands,
	}
}
