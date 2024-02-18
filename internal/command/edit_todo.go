package command

import (
	"errors"
	"fmt"
	"github.com/Hydoc/goo/internal/model"
	"github.com/Hydoc/goo/internal/view"
	"strconv"
	"strings"
)

type EditTodo struct {
	todoList *model.TodoList
	view     view.View
	idToEdit int
	newLabel string
}

func (cmd *EditTodo) Execute() {
	oldLabel := cmd.todoList.Find(cmd.idToEdit).Label
	newLabel := strings.ReplaceAll(cmd.newLabel, "{}", oldLabel)
	cmd.todoList.Edit(cmd.idToEdit, newLabel)
	cmd.view.RenderList(cmd.todoList)
	cmd.todoList.SaveToFile()
}

func NewEditTodo(todoList *model.TodoList, view view.View, payload string) (Command, error) {
	splitBySpace := strings.Split(payload, " ")
	id, err := strconv.Atoi(splitBySpace[0])

	if err != nil {
		return nil, fmt.Errorf("can not edit todo, id is missing")
	}

	if !todoList.Has(id) {
		return nil, fmt.Errorf("there is no todo with id %d", id)
	}

	newLabel := strings.Join(splitBySpace[1:], " ")

	if len(newLabel) == 0 {
		return nil, errors.New("empty todo is not allowed")
	}

	return &EditTodo{todoList, view, id, newLabel}, nil
}
