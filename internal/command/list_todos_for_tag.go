package command

import (
	"github.com/Hydoc/goo/internal/model"
	"github.com/Hydoc/goo/internal/view"
	"strconv"
)

type ListTodosForTag struct {
	todoList *model.TodoList
	view     view.View
	idOfTag  model.TagId
}

func (cmd *ListTodosForTag) Execute() {
	cmd.view.RenderList(cmd.todoList.TodosForTag(cmd.idOfTag))
}

func NewListTodosForTag(todoList *model.TodoList, view view.View, payload string) (Command, error) {
	idOfTag, err := strconv.Atoi(payload)
	if err != nil {
		return nil, errInvalidId(payload)
	}

	if !todoList.HasTag(model.TagId(idOfTag)) {
		return nil, errNoTagWithId(idOfTag)
	}

	return &ListTodosForTag{todoList, view, model.TagId(idOfTag)}, nil
}
