package command

import (
	"errors"
	"github.com/Hydoc/goo/internal/model"
)

type AddTodo struct {
	todoList  *model.TodoList
	todoToAdd string
}

func (cmd *AddTodo) Execute() {
	cmd.todoList.Add(model.NewTodo(cmd.todoToAdd, cmd.todoList.NextId()))
}

func newAddTodo(todoList *model.TodoList, payload string) (*AddTodo, error) {
	if len(payload) == 0 {
		return nil, errors.New("empty todo is not allowed")
	}

	return &AddTodo{todoList, payload}, nil
}
