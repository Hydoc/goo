package command

import "github.com/Hydoc/goo/internal/model"

type Clear struct {
	todoList *model.TodoList
}

func (cmd *Clear) Execute() {
	cmd.todoList.Clear()
}

func newClear(todoList *model.TodoList) *Clear {
	return &Clear{todoList}
}
