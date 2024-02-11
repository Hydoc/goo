package command

import (
	"errors"
	"fmt"
	"github.com/Hydoc/goo/internal"
	"slices"
	"strconv"
)

type Parser struct {
	validCommands []*StringCommand
}

func (par *Parser) Parse(cmd, payload string, todoList *internal.TodoList, undoStack *UndoStack) (Command, error) {
	for _, strCmd := range par.validCommands {
		if cmd == strCmd.Abbreviation || slices.Contains(strCmd.Aliases, cmd) {
			return par.fabricate(strCmd.Abbreviation, payload, todoList, undoStack)
		}
	}

	return nil, fmt.Errorf("could not find command '%s'", cmd)
}

func (par *Parser) fabricate(input, payload string, todoList *internal.TodoList, undoStack *UndoStack) (Command, error) {
	switch input {
	case QuitAbbr:
		return NewQuit(), nil
	case HelpAbbr:
		return NewHelp(par.validCommands), nil
	case AddTodoAbbr:
		if !CanCreateAddTodo(payload) {
			return nil, errors.New(AddTodoHelp)
		}
		return NewAddTodo(todoList, payload), nil
	case ToggleTodoAbbr:
		if !CanCreateToggleTodo(payload) {
			return nil, errors.New(ToggleTodoHelp)
		}
		id, _ := strconv.Atoi(payload)
		return NewToggleTodo(todoList, id), nil
	case UndoAbbr:
		if !undoStack.HasItems() {
			return nil, errors.New(NothingToUndo)
		}
		return NewUndo(undoStack.Pop()), nil
	default:
		return nil, nil
	}
}

func NewParser(validCommands []*StringCommand) *Parser {
	return &Parser{
		validCommands: validCommands,
	}
}
