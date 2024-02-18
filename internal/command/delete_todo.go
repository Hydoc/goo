package command

import (
	"fmt"
	"github.com/Hydoc/goo/internal/model"
	"github.com/Hydoc/goo/internal/view"
	"strconv"
)

type DeleteTodo struct {
	todoList   *model.TodoList
	view       view.View
	idToDelete int
}

func (cmd *DeleteTodo) Execute() {
	cmd.todoList.Remove(cmd.idToDelete)
	cmd.view.RenderList(cmd.todoList)
	cmd.todoList.SaveToFile()
}

func NewDeleteTodo(todoList *model.TodoList, view view.View, payload string) (Command, error) {
	id, err := strconv.Atoi(payload)
	if err != nil {
		return nil, fmt.Errorf("could not parse %s to int", payload)
	}

	if !todoList.Has(id) {
		return nil, fmt.Errorf("there is no todo with id %d", id)
	}

	return &DeleteTodo{todoList, view, id}, nil
}
