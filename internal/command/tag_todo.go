package command

import (
	"fmt"
	"github.com/Hydoc/goo/internal/model"
	"github.com/Hydoc/goo/internal/view"
	"strconv"
	"strings"
)

type TagTodo struct {
	todoList *model.TodoList
	view     view.View
	idToTag  int
	tag      model.Tag
}

func (cmd *TagTodo) Execute() {
	cmd.todoList.Tag(cmd.idToTag, cmd.tag.Id)
	cmd.view.RenderList(cmd.todoList)
	cmd.todoList.SaveToFile()
}

func NewTagTodo(todoList *model.TodoList, view view.View, payload string) (Command, error) {
	splitBySpace := strings.Split(payload, " ")

	if len(splitBySpace) < 2 {
		return nil, fmt.Errorf("can not tag todo, need id of todo as first argument, the second has to be the name of the label")
	}

	todoIdToTag, err := strconv.Atoi(splitBySpace[0])
	if err != nil {
		return nil, fmt.Errorf(ErrInvalidId, splitBySpace[0])
	}

	if !todoList.Has(todoIdToTag) {
		return nil, fmt.Errorf(ErrNoTodoWithId, todoIdToTag)
	}

	tagName := strings.TrimSpace(strings.ToLower(strings.Join(splitBySpace[1:], " ")))
	return &TagTodo{todoList, view, todoIdToTag, todoList.CreateTagIfNotExist(tagName)}, nil
}
