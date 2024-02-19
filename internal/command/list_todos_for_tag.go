package command

import (
	"fmt"
	"github.com/Hydoc/goo/internal/model"
	"github.com/Hydoc/goo/internal/view"
	"strconv"
)

type ListTodosForTag struct {
	todoList *model.TodoList
	view     view.View
	idOfTag  int
}

func (cmd *ListTodosForTag) Execute() {
	cmd.view.RenderLine("LIST_TODOS_FOR_TAG")
}

func NewListTodosForTag(todoList *model.TodoList, view view.View, payload string) (Command, error) {
	idOfTag, err := strconv.Atoi(payload)
	if err != nil {
		return nil, fmt.Errorf(ErrInvalidId, payload)
	}

	if !todoList.HasTag(model.TagId(idOfTag)) {
		return nil, fmt.Errorf(ErrNoTagWithId, idOfTag)
	}

	return &ListTodosForTag{todoList, view, idOfTag}, nil
}
