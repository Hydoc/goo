package command

import (
	"fmt"
	"github.com/Hydoc/goo/internal/model"
	"github.com/Hydoc/goo/internal/view"
	"strconv"
	"strings"
)

type RemoveTagFromTodo struct {
	todoList      *model.TodoList
	view          view.View
	idOfTodo      int
	tagIdToRemove model.TagId
}

func (cmd *RemoveTagFromTodo) Execute() {
	cmd.todoList.RemoveTagFromTodo(cmd.tagIdToRemove, cmd.idOfTodo)
	cmd.view.RenderList(cmd.todoList)
	cmd.todoList.SaveToFile()
}

func NewRemoveTagFromTodo(todoList *model.TodoList, view view.View, payload string) (Command, error) {
	splitBySpace := strings.Split(payload, " ")

	if len(splitBySpace) < 2 || len(splitBySpace) > 2 {
		return nil, fmt.Errorf("can not untag todo, need id of tag as first argument, the second has to be the id of the todo")
	}

	tagId, err := strconv.Atoi(splitBySpace[0])
	if err != nil {
		return nil, errInvalidId(splitBySpace[0])
	}

	todoId, err := strconv.Atoi(splitBySpace[1])
	if err != nil {
		return nil, errInvalidId(splitBySpace[1])
	}

	todo := todoList.Find(todoId)
	if todo == nil {
		return nil, fmt.Errorf(ErrNoTodoWithId, todoId)
	}

	tag := todoList.FindTag(tagId)
	if tag == nil {
		return nil, fmt.Errorf(ErrNoTagWithId, tagId)
	}

	if !todo.HasTag(model.TagId(tagId)) {
		return nil, fmt.Errorf(ErrTodoDoesNotHaveTag, todo.Id, tagId)
	}

	return &RemoveTagFromTodo{todoList, view, todoId, tag.Id}, nil
}
