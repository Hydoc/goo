package command

import (
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
		return nil, errInvalidId(splitBySpace[0])
	}

	if !todoList.Has(id) {
		return nil, fmt.Errorf(ErrNoTodoWithId, id)
	}

	newLabel := strings.Join(splitBySpace[1:], " ")

	if len(newLabel) == 0 {
		return nil, errEmptyTodoNotAllowed()
	}

	return &EditTodo{todoList, view, id, newLabel}, nil
}
