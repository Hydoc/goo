package command

import "github.com/Hydoc/goo/internal"

type Clear struct {
	todoList *internal.TodoList
}

func (cmd *Clear) Execute() {
	cmd.todoList.Clear()
}

func newClear(todoList *internal.TodoList) *Clear {
	return &Clear{todoList}
}
