package command

import (
	"fmt"
	"github.com/Hydoc/goo/internal/view"
	"slices"
)

type Parser struct {
	validCommands []*StringCommand
}

type ParsedCommand struct {
	abbreviation string
	payload      string
}

func (par *Parser) Parse(arg *view.Argument) (*ParsedCommand, error) {
	for _, strCmd := range par.validCommands {
		if arg.RawCommand == strCmd.Abbreviation || slices.Contains(strCmd.Aliases, arg.RawCommand) {
			return &ParsedCommand{abbreviation: strCmd.Abbreviation, payload: arg.Payload}, nil
		}
	}

	return nil, fmt.Errorf("could not find command '%s'", arg.RawCommand)
}

func NewParser(validCommands []*StringCommand) *Parser {
	return &Parser{
		validCommands: validCommands,
	}
}
