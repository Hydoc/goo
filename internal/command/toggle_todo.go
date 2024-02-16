package command

import (
	"fmt"
	"github.com/Hydoc/goo/internal/model"
)

type ToggleTodo struct {
	todoList   *model.TodoList
	idToToggle int
}

func (cmd *ToggleTodo) Execute() {
	cmd.todoList.Toggle(cmd.idToToggle)
}

func newToggleTodo(todoList *model.TodoList, id int) (*ToggleTodo, error) {
	if !todoList.Has(id) {
		return nil, fmt.Errorf("there is no todo with id %d", id)
	}

	return &ToggleTodo{todoList, id}, nil
}
