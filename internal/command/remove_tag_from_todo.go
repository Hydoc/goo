package command

import (
	"github.com/Hydoc/goo/internal/model"
	"github.com/Hydoc/goo/internal/view"
)

type RemoveTagFromTodo struct {
	todoList      *model.TodoList
	view          view.View
	idOfTodo      int
	tagIdToRemove model.TagId
}

func (cmd *RemoveTagFromTodo) Execute() {
	cmd.view.RenderLine("REMOVE_TAG_FROM_TODO")
}

func NewRemoveTagFromTodo(todoList *model.TodoList, view view.View, payload string) (Command, error) {
	return &RemoveTagFromTodo{todoList, view, 0, 0}, nil
}
