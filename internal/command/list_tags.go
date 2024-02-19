package command

import (
	"github.com/Hydoc/goo/internal/model"
	"github.com/Hydoc/goo/internal/view"
)

type ListTags struct {
	todoList *model.TodoList
	view     view.View
}

func (cmd *ListTags) Execute() {
	cmd.view.RenderTags(cmd.todoList)
}

func NewListTags(todoList *model.TodoList, view view.View, _ string) (Command, error) {
	return &ListTags{todoList, view}, nil
}
