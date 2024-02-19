package command

import (
	"github.com/Hydoc/goo/internal/model"
	"github.com/Hydoc/goo/internal/view"
)

type Clear struct {
	todoList *model.TodoList
	view     view.View
}

func (cmd *Clear) Execute() {
	cmd.todoList.Clear()
	cmd.view.RenderList(cmd.todoList)
	cmd.todoList.SaveToFile()
}

func NewClear(todoList *model.TodoList, view view.View, _ string) (Command, error) {
	return &Clear{todoList, view}, nil
}
