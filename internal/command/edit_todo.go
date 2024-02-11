package command

import (
	"errors"
	"github.com/Hydoc/goo/internal"
	"strconv"
	"strings"
)

var EditTodoAliases = []string{"e"}

const (
	EditTodoAbbr  = "edit"
	EditTodoDesc  = "Edit a todo (use '{}' to insert the old value)"
	editTodoUsage = "use the command like so: edit <id of the todo> <the new label (use '{}' to insert the old value)>"
	nothingToEdit = "nothing to edit"
	editTemplate  = "{}"
)

type EditTodo struct {
	previousLabel string
	todoList      *internal.TodoList
	idToEdit      int
	newLabel      string
}

func (cmd *EditTodo) Execute() {
	oldLabel := cmd.todoList.Find(cmd.idToEdit).Label
	cmd.previousLabel = oldLabel
	newLabel := strings.ReplaceAll(cmd.newLabel, editTemplate, oldLabel)
	cmd.todoList.Edit(cmd.idToEdit, newLabel)
}

func (cmd *EditTodo) Undo() {
	cmd.todoList.Edit(cmd.idToEdit, cmd.previousLabel)
}

func newEditTodo(todoList *internal.TodoList, payload string) (*EditTodo, error) {
	splitBySpace := strings.Split(payload, " ")
	id, err := strconv.Atoi(splitBySpace[0])

	if err != nil {
		return nil, errors.New(editTodoUsage)
	}

	if !todoList.Has(id) {
		return nil, errors.New(nothingToEdit)
	}

	newLabel := strings.Join(splitBySpace[1:], " ")

	if len(newLabel) == 0 {
		return nil, errors.New(editTodoUsage)
	}

	return &EditTodo{
		todoList: todoList,
		idToEdit: id,
		newLabel: newLabel,
	}, nil
}
