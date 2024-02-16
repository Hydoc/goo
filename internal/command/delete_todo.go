package command

import (
	"fmt"
	"github.com/Hydoc/goo/internal/model"
)

type DeleteTodo struct {
	todoList   *model.TodoList
	idToDelete int
}

func (cmd *DeleteTodo) Execute() {
	cmd.todoList.Remove(cmd.idToDelete)
}

func newDeleteTodo(todoList *model.TodoList, id int) (*DeleteTodo, error) {
	if !todoList.Has(id) {
		return nil, fmt.Errorf("there is no todo with id %d", id)
	}

	return &DeleteTodo{todoList, id}, nil
}
