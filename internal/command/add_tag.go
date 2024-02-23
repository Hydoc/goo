package command

import (
	"github.com/Hydoc/goo/internal/model"
	"github.com/Hydoc/goo/internal/view"
	"strings"
)

type AddTag struct {
	todoList        *model.TodoList
	view            view.View
	tagNameToCreate string
}

func (cmd *AddTag) Execute() {
	cmd.todoList.AddTag(model.NewTag(cmd.todoList.NextTagId(), cmd.tagNameToCreate))
	cmd.view.RenderTags(cmd.todoList.TagList)
	cmd.todoList.SaveToFile()
}

func NewAddTag(todoList *model.TodoList, view view.View, payload string) (Command, error) {
	normalizedTagName := strings.TrimSpace(strings.ToLower(payload))

	if todoList.HasTagWith(normalizedTagName) {
		return nil, errTagAlreadyExists(normalizedTagName)
	}

	return &AddTag{todoList, view, normalizedTagName}, nil
}
