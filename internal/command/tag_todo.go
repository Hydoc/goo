package command

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Hydoc/goo/internal/model"
	"github.com/Hydoc/goo/internal/view"
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
		return nil, fmt.Errorf("can not tag todo, need id of tag as first argument, the second has to be the id of the todo")
	}

	tagId, err := strconv.Atoi(splitBySpace[0])
	if err != nil {
		return nil, errInvalidId(splitBySpace[0])
	}

	todoIdToTag, err := strconv.Atoi(splitBySpace[1])
	if err != nil {
		return nil, errInvalidId(splitBySpace[1])
	}

	todo := todoList.Find(todoIdToTag)
	if todo == nil {
		return nil, errNoTodoWithId(todoIdToTag)
	}

	tag := todoList.FindTag(tagId)
	if tag == nil {
		return nil, errNoTagWithId(tagId)
	}

	if todo.HasTag(tag.Id) {
		return nil, errTodoAlreadyHasTag(todoIdToTag, tagId)
	}

	return &TagTodo{todoList, view, todoIdToTag, tag.Id}, nil
}
