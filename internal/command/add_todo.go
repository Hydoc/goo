package command

import (
	"errors"
	"github.com/Hydoc/goo/internal"
)

var AddTodoAliases = []string{"a"}

const (
	AddTodoAbbr  = "add"
	AddTodoDesc  = "Add a new todo"
	addTodoUsage = "use the command like so: add <label of the todo>"
)

type AddTodo struct {
	previousTodoListItems []*internal.Todo
	todoList              *internal.TodoList
	todoToAdd             string
}

func (cmd *AddTodo) Execute() {
	cmd.previousTodoListItems = make([]*internal.Todo, len(cmd.todoList.Items))
	copy(cmd.previousTodoListItems, cmd.todoList.Items)
	cmd.todoList.Add(internal.NewTodo(cmd.todoToAdd, cmd.todoList.NextId()))
}

func (cmd *AddTodo) Undo() {
	cmd.todoList.Items = cmd.previousTodoListItems
}

func newAddTodo(todoList *internal.TodoList, payload string) (*AddTodo, error) {
	if len(payload) == 0 {
		return nil, errors.New(addTodoUsage)
	}

	return &AddTodo{
		todoList:  todoList,
		todoToAdd: payload,
	}, nil
}
