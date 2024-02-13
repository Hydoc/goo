package command

import (
	"fmt"
	"github.com/Hydoc/goo/internal"
)

type DeleteTodo struct {
	todoList   *internal.TodoList
	idToDelete int
}

func (cmd *DeleteTodo) Execute() {
	cmd.todoList.Remove(cmd.idToDelete)
}

func NewDeleteTodo(todoList *internal.TodoList, id int) (*DeleteTodo, error) {
	if !todoList.Has(id) {
		return nil, fmt.Errorf("there is no todo with id %d", id)
	}

	return &DeleteTodo{
		todoList:   todoList,
		idToDelete: id,
	}, nil
}
