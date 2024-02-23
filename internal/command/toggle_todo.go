package command

import (
	"github.com/Hydoc/goo/internal/model"
	"github.com/Hydoc/goo/internal/view"
	"strconv"
)

type ToggleTodo struct {
	todoList   *model.TodoList
	view       view.View
	idToToggle int
}

func (cmd *ToggleTodo) Execute() {
	cmd.todoList.Toggle(cmd.idToToggle)
	cmd.view.RenderList(cmd.todoList)
	cmd.todoList.SaveToFile()
}

func NewToggleTodo(todoList *model.TodoList, view view.View, payload string) (Command, error) {
	id, err := strconv.Atoi(payload)
	if err != nil {
		return nil, errInvalidId(payload)
	}

	if !todoList.Has(id) {
		return nil, errNoTodoWithId(id)
	}

	return &ToggleTodo{todoList, view, id}, nil
}
