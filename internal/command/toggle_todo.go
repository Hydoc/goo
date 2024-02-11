package command

import (
	"github.com/Hydoc/goo/internal"
	"strconv"
)

var ToggleTodoAliases = []string{"t"}

const (
	ToggleTodoAbbr = "toggle"
	ToggleTodoDesc = "Toggle the done state of a todo"
	ToggleTodoHelp = "use the command like so: toggle <id of the todo>"
)

type ToggleTodo struct {
	todoList       *internal.TodoList
	todoIdToToggle int
}

func (toggle *ToggleTodo) Execute() {
	toggle.todoList.Toggle(toggle.todoIdToToggle)
}

func (toggle *ToggleTodo) Undo() {
	toggle.todoList.Toggle(toggle.todoIdToToggle)
}

func NewToggleTodo(list *internal.TodoList, idToToggle int) *ToggleTodo {
	return &ToggleTodo{
		todoList:       list,
		todoIdToToggle: idToToggle,
	}
}

func CanCreateToggleTodo(payload string) bool {
	_, err := strconv.Atoi(payload)
	return len(payload) > 0 && err == nil
}
