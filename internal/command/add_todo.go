package command

import (
	"errors"
	"github.com/Hydoc/goo/internal/model"
	"github.com/Hydoc/goo/internal/view"
)

type AddTodo struct {
	todoList  *model.TodoList
	view      view.View
	todoToAdd string
}

func (cmd *AddTodo) Execute() {
	cmd.todoList.Add(model.NewTodo(cmd.todoToAdd, cmd.todoList.NextId()))
	cmd.view.RenderList(cmd.todoList)
	cmd.todoList.SaveToFile()
}

func NewAddTodo(todoList *model.TodoList, view view.View, payload string) (Command, error) {
	if len(payload) == 0 {
		return nil, errors.New("empty todo is not allowed")
	}

	return &AddTodo{todoList, view, payload}, nil
}
