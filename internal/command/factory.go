package command

import (
	"github.com/Hydoc/goo/internal"
)

type Factory struct {
	validCommands []*StringCommand
}

func (f *Factory) Fabricate(parsedCmd *ParsedCommand, todoList *internal.TodoList, undoStack *UndoStack) (Command, error) {
	switch parsedCmd.abbreviation {
	case QuitAbbr:
		return newQuit(), nil
	case HelpAbbr:
		return newHelp(f.validCommands), nil
	case AddTodoAbbr:
		return newAddTodo(todoList, parsedCmd.payload)
	case ToggleTodoAbbr:
		return newToggleTodo(todoList, parsedCmd.payload)
	case UndoAbbr:
		return newUndo(undoStack)
	case DeleteTodoAbbr:
		return newDeleteTodo(todoList, parsedCmd.payload)
	default:
		return nil, nil
	}
}

func NewFactory(validCommands []*StringCommand) *Factory {
	return &Factory{
		validCommands: validCommands,
	}
}
