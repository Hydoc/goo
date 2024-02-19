package command

import (
	"github.com/Hydoc/goo/internal/model"
	"github.com/Hydoc/goo/internal/view"
)

type RemoveTag struct {
	todoList        *model.TodoList
	view            view.View
	tagNameToRemove string
}

func (cmd *RemoveTag) Execute() {
	cmd.view.RenderLine("REMOVE_TAG")
}

func NewRemoveTag(todoList *model.TodoList, view view.View, payload string) (Command, error) {
	return &RemoveTag{todoList, view, payload}, nil
}
