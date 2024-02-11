package command

import (
	"errors"
	"github.com/Hydoc/goo/internal"
	"strconv"
)

var ToggleTodoAliases = []string{"t"}

const (
	ToggleTodoAbbr  = "toggle"
	ToggleTodoDesc  = "Toggle the done state of a todo"
	toggleTodoUsage = "use the command like so: toggle <id of the todo>"
	nothingToToggle = "nothing to toggle"
)

type ToggleTodo struct {
	todoList       *internal.TodoList
	todoIdToToggle int
}

func (cmd *ToggleTodo) Execute() {
	cmd.todoList.Toggle(cmd.todoIdToToggle)
}

func (cmd *ToggleTodo) Undo() {
	cmd.todoList.Toggle(cmd.todoIdToToggle)
}

func newToggleTodo(todoList *internal.TodoList, payload string) (*ToggleTodo, error) {
	id, err := strconv.Atoi(payload)

	if err != nil {
		return nil, errors.New(toggleTodoUsage)
	}

	if !todoList.Has(id) {
		return nil, errors.New(nothingToToggle)
	}

	return &ToggleTodo{
		todoList:       todoList,
		todoIdToToggle: id,
	}, nil
}
