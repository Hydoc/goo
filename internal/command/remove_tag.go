package command

import (
	"fmt"
	"github.com/Hydoc/goo/internal/model"
	"github.com/Hydoc/goo/internal/view"
	"strconv"
	"strings"
)

type RemoveTag struct {
	todoList      *model.TodoList
	view          view.View
	tagIdToRemove model.TagId
}

func (cmd *RemoveTag) Execute() {
	cmd.todoList.RemoveTag(cmd.tagIdToRemove)
	cmd.view.RenderTags(cmd.todoList.TagList)
	cmd.todoList.SaveToFile()
}

func NewRemoveTag(todoList *model.TodoList, view view.View, payload string) (Command, error) {
	splitBySpace := strings.Split(payload, " ")
	rawTagId, err := strconv.Atoi(splitBySpace[0])
	if err != nil {
		return nil, errInvalidId(payload)
	}

	tagId := model.TagId(rawTagId)
	if !todoList.HasTag(tagId) {
		return nil, fmt.Errorf(ErrNoTagWithId, tagId)
	}

	return &RemoveTag{todoList, view, tagId}, nil
}
