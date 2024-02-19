package command

import (
	"github.com/Hydoc/goo/internal/model"
	"github.com/Hydoc/goo/internal/view"
)

type ListTodos struct {
	todoList *model.TodoList
	view     view.View
}

func (cmd *ListTodos) Execute() {
	cmd.view.RenderList(cmd.todoList)
}

func NewListTodos(todoList *model.TodoList, view view.View, _ string) (Command, error) {
	return &ListTodos{todoList, view}, nil
}
