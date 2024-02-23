package command

import (
	"github.com/Hydoc/goo/internal/model"
	"github.com/Hydoc/goo/internal/view"
)

type AddTodo struct {
	todoList  *model.TodoList
	view      view.View
	todoToAdd string
}

func (cmd *AddTodo) Execute() {
	cmd.todoList.Add(model.NewTodo(cmd.todoToAdd, cmd.todoList.NextTodoId()))
	cmd.view.RenderList(cmd.todoList)
	cmd.todoList.SaveToFile()
}

func NewAddTodo(todoList *model.TodoList, view view.View, payload string) (Command, error) {
	if len(payload) == 0 {
		return nil, ErrEmptyTodoNotAllowed
	}

	return &AddTodo{todoList, view, payload}, nil
}
