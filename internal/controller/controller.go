package controller

import (
	"github.com/Hydoc/goo/internal"
	"github.com/Hydoc/goo/internal/command"
	"github.com/Hydoc/goo/internal/view"
)

type Controller struct {
	view      *view.StdoutView
	todoList  *internal.TodoList
	parser    *command.Parser
	undoStack *command.UndoStack
}

func (ctr *Controller) Run() {
	var nextError error
	for {
		if ctr.todoList.HasItems() {
			ctr.view.RenderLine(ctr.todoList.String())
		}
		if nextError != nil {
			ctr.view.RenderLine(nextError.Error())
			nextError = nil
		}
		argument := ctr.view.Prompt()
		cmd, err := ctr.parser.Parse(argument.RawCommand, argument.Payload, ctr.todoList, ctr.undoStack)
		if err != nil {
			nextError = err
			continue
		}
		cmd.Execute()
		if undoable, isUndoable := cmd.(command.UndoableCommand); isUndoable {
			ctr.undoStack.Push(&undoable)
		}
	}
}

func New(view *view.StdoutView, list *internal.TodoList, parser *command.Parser, undoStack *command.UndoStack) *Controller {
	return &Controller{
		view:      view,
		todoList:  list,
		parser:    parser,
		undoStack: undoStack,
	}
}