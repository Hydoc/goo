package command

import (
	"errors"
	"github.com/Hydoc/goo/internal"
	"strconv"
)

var DeleteTodoAliases = []string{"d"}

const (
	DeleteTodoAbbr  = "delete"
	DeleteTodoDesc  = "Delete a todo"
	deleteTodoUsage = "use the command like so: delete <id of the todo>"
	nothingToDelete = "nothing to delete"
)

type DeleteTodo struct {
	previousTodoListItems []*internal.Todo
	todoList              *internal.TodoList
	idToDelete            int
}

func (d *DeleteTodo) Execute() {
	d.previousTodoListItems = d.todoList.Items
	d.todoList.Remove(d.idToDelete)
}

func (d *DeleteTodo) Undo() {
	d.todoList.Items = d.previousTodoListItems
}

func newDeleteTodo(todoList *internal.TodoList, payload string) (*DeleteTodo, error) {
	id, err := strconv.Atoi(payload)

	if err != nil {
		return nil, errors.New(deleteTodoUsage)
	}

	if !todoList.Has(id) {
		return nil, errors.New(nothingToDelete)
	}

	return &DeleteTodo{
		todoList:   todoList,
		idToDelete: id,
	}, nil
}
