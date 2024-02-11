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

func (cmd *DeleteTodo) Execute() {
	cmd.previousTodoListItems = make([]*internal.Todo, len(cmd.todoList.Items))
	copy(cmd.previousTodoListItems, cmd.todoList.Items)
	cmd.todoList.Remove(cmd.idToDelete)
}

func (cmd *DeleteTodo) Undo() {
	cmd.todoList.Items = cmd.previousTodoListItems
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
