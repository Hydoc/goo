package controller

import (
	"github.com/Hydoc/goo/internal"
	"github.com/Hydoc/goo/internal/command"
	"github.com/Hydoc/goo/internal/view"
)

type Controller struct {
	view     view.View
	todoList *internal.TodoList
	factory  *command.Factory
}

func (ctr *Controller) Handle(list bool, toggle int, add bool, doDelete int, edit bool, doClear bool, args string) (int, error) {
	switch {
	case list:
		ctr.view.RenderList(ctr.todoList)
	default:
		cmd, err := ctr.factory.Fabricate(ctr.todoList, toggle, add, doDelete, edit, doClear, args)
		if err != nil {
			return 1, err
		}
		if cmd == nil {
			return 0, nil
		}
		cmd.Execute()
		ctr.todoList.SaveToFile()
	}
	return 0, nil
}

func New(view view.View, todoList *internal.TodoList, factory *command.Factory) *Controller {
	return &Controller{
		view,
		todoList,
		factory,
	}
}
