package command

import (
	"fmt"
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
		return nil, fmt.Errorf("%s is an invalid id", payload)
	}

	if !todoList.Has(id) {
		return nil, fmt.Errorf("there is no todo with id %d", id)
	}

	return &ToggleTodo{todoList, view, id}, nil
}
