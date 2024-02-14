package command

import (
	"errors"
	"github.com/Hydoc/goo/internal"
)

type AddTodo struct {
	todoList  *internal.TodoList
	todoToAdd string
}

func (cmd *AddTodo) Execute() {
	cmd.todoList.Add(internal.NewTodo(cmd.todoToAdd, cmd.todoList.NextId()))
}

func newAddTodo(todoList *internal.TodoList, payload string) (*AddTodo, error) {
	if len(payload) == 0 {
		return nil, errors.New("empty todo is not allowed")
	}

	return &AddTodo{todoList, payload}, nil
}
