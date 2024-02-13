package controller

import (
	"github.com/Hydoc/goo/internal"
	"github.com/Hydoc/goo/internal/command"
	"github.com/Hydoc/goo/internal/view"
)

type Controller struct {
	view     *view.StdoutView
	todoList *internal.TodoList
}

func (ctr *Controller) Handle(list *bool, toggle *int, add *bool, doDelete *int, edit *bool, args string) (int, error) {
	defer ctr.todoList.SaveToFile()

	switch {
	case *list:
		ctr.view.RenderList(ctr.todoList)
	case *add:
		cmd, err := command.NewAddTodo(ctr.todoList, args)
		if err != nil {
			return 1, err
		}
		cmd.Execute()
	case *edit:
		cmd, err := command.NewEditTodo(ctr.todoList, args)
		if err != nil {
			return 1, err
		}
		cmd.Execute()
	case *doDelete > 0:
		cmd, err := command.NewDeleteTodo(ctr.todoList, *doDelete)
		if err != nil {
			return 1, err
		}
		cmd.Execute()
	case *toggle > 0:
		cmd, err := command.NewToggleTodo(ctr.todoList, *toggle)
		if err != nil {
			return 1, err
		}
		cmd.Execute()
	}
	return 0, nil
}

func New(view *view.StdoutView, todoList *internal.TodoList) *Controller {
	return &Controller{
		view:     view,
		todoList: todoList,
	}
}
