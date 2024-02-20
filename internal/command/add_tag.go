package command

import (
	"github.com/Hydoc/goo/internal/model"
	"github.com/Hydoc/goo/internal/view"
)

type AddTag struct {
	todoList        *model.TodoList
	view            view.View
	tagNameToCreate string
}

func (cmd *AddTag) Execute() {
	cmd.view.RenderLine("ADD_TAG")
}

func NewAddTag(todoList *model.TodoList, view view.View, payload string) (Command, error) {
	return &AddTag{todoList, view, payload}, nil
}
