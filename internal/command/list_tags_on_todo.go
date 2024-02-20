package command

import (
	"fmt"
	"github.com/Hydoc/goo/internal/model"
	"github.com/Hydoc/goo/internal/view"
	"strconv"
)

type ListTagsOnTodo struct {
	todoList *model.TodoList
	view     view.View
	idOfTodo int
}

func (cmd *ListTagsOnTodo) Execute() {
	cmd.view.RenderLine("LIST_TAGS_ON_TODO")
}

func NewListTagsOnTodo(todoList *model.TodoList, view view.View, payload string) (Command, error) {
	idOfTodo, err := strconv.Atoi(payload)
	if err != nil {
		return nil, errInvalidId(payload)
	}

	if !todoList.Has(idOfTodo) {
		return nil, fmt.Errorf(ErrNoTodoWithId, idOfTodo)
	}

	return &ListTagsOnTodo{todoList, view, idOfTodo}, nil
}
