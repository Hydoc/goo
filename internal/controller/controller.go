package controller

import (
	"github.com/Hydoc/goo/internal"
	"github.com/Hydoc/goo/internal/command"
	"github.com/Hydoc/goo/internal/view"
)

type Controller struct {
	view     *view.StdoutView
	todoList *internal.TodoList
	factory  *command.Factory
}

func (ctr *Controller) Handle(list *bool, toggle int, add bool, doDelete int, edit bool, doClear bool, args string) (int, error) {
	defer ctr.todoList.SaveToFile()

	switch {
	case *list:
		ctr.view.RenderList(ctr.todoList.SortedByIdAndState())
	default:
		cmd, err := ctr.factory.Fabricate(ctr.todoList, toggle, add, doDelete, edit, doClear, args)
		if err != nil {
			return 1, err
		}
		cmd.Execute()
	}
	return 0, nil
}

func New(view *view.StdoutView, todoList *internal.TodoList, factory *command.Factory) *Controller {
	return &Controller{
		view,
		todoList,
		factory,
	}
}
