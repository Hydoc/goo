package command

import (
	"fmt"
	"github.com/Hydoc/goo/internal"
)

type ToggleTodo struct {
	todoList   *internal.TodoList
	idToToggle int
}

func (cmd *ToggleTodo) Execute() {
	cmd.todoList.Toggle(cmd.idToToggle)
}

func newToggleTodo(todoList *internal.TodoList, id int) (*ToggleTodo, error) {
	if !todoList.Has(id) {
		return nil, fmt.Errorf("there is no todo with id %d", id)
	}

	return &ToggleTodo{
		todoList:   todoList,
		idToToggle: id,
	}, nil
}
