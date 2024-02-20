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
	tagId    model.TagId
}

func (cmd *TagTodo) Execute() {
	cmd.todoList.TagTodo(cmd.idToTag, cmd.tagId)
	cmd.view.RenderList(cmd.todoList)
	cmd.todoList.SaveToFile()
}

func NewTagTodo(todoList *model.TodoList, view view.View, payload string) (Command, error) {
	splitBySpace := strings.Split(payload, " ")

	if len(splitBySpace) < 2 || len(splitBySpace) > 2 {
		return nil, fmt.Errorf("can not tag todo, need id of todo as first argument, the second has to be the id of the tag")
	}

	todoIdToTag, err := strconv.Atoi(splitBySpace[0])
	if err != nil {
		return nil, fmt.Errorf(ErrInvalidId, splitBySpace[0])
	}

	tagId, err := strconv.Atoi(splitBySpace[1])
	if err != nil {
		return nil, fmt.Errorf(ErrInvalidId, splitBySpace[1])
	}

	if !todoList.Has(todoIdToTag) {
		return nil, fmt.Errorf(ErrNoTodoWithId, todoIdToTag)
	}

	tag := todoList.FindTag(tagId)
	if tag == nil {
		return nil, fmt.Errorf(ErrNoTagWithId, tagId)
	}

	return &TagTodo{todoList, view, todoIdToTag, tag.Id}, nil
}
