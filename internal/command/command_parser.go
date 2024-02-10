package command

import (
	"fmt"
	"github.com/Hydoc/goo/internal"
	"slices"
)

type Parser struct {
	validCommands []*StringCommand
	todoList      *internal.TodoList
}

func (par *Parser) Parse(cmd string, payload string) (Command, error) {
	for _, strCmd := range par.validCommands {
		if cmd == strCmd.Abbreviation || slices.Contains(strCmd.Aliases, cmd) {
			return par.fabricate(strCmd.Abbreviation, payload), nil
		}
	}

	return nil, fmt.Errorf("could not find command '%s'", cmd)
}

func (par *Parser) fabricate(input, payload string) Command {
	switch input {
	case QuitAbbr:
		return NewQuit()
	case HelpAbbr:
		return NewHelp(par.validCommands)
	case AddTodoAbbr:
		return NewAddTodo(payload, par.todoList)
	default:
		return nil
	}
}

func NewParser(validCommands []*StringCommand, todoList *internal.TodoList) *Parser {
	return &Parser{
		validCommands: validCommands,
		todoList:      todoList,
	}
}
