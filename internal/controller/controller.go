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
	factory   *command.Factory
}

func (ctr *Controller) Run() {
	var nextError error
	doClearScreen := true
	doQuit := false
	for !doQuit {
		if doClearScreen {
			ctr.view.ClearScreen()
		}
		if ctr.todoList.HasItems() {
			ctr.view.RenderLine("Your list:")
			ctr.view.RenderLine(ctr.todoList.String())
		} else {
			ctr.view.RenderLine("You currently have no todos")
			ctr.view.RenderLine("Type 'help' for an overview of all commands")
		}
		if nextError != nil {
			ctr.view.RenderLine(nextError.Error())
			nextError = nil
		}
		argument := ctr.view.Prompt()
		parsedCmd, err := ctr.parser.Parse(argument)
		if err != nil {
			nextError = err
			doClearScreen = true
			continue
		}
		cmd, err := ctr.factory.Fabricate(parsedCmd, ctr.todoList, ctr.undoStack, ctr.view)
		if err != nil {
			nextError = err
			doClearScreen = true
			continue
		}
		cmd.Execute()
		switch cmd.(interface{}).(type) {
		case *command.Help:
			// do not clear screen when command is help otherwise it vanishes
			doClearScreen = false
		case *command.Quit:
			doQuit = true
		}
		if undoable, isUndoable := cmd.(command.UndoableCommand); isUndoable {
			ctr.undoStack.Push(undoable)
		}
	}
}

func New(view *view.StdoutView, list *internal.TodoList, parser *command.Parser, undoStack *command.UndoStack, factory *command.Factory) *Controller {
	return &Controller{
		view:      view,
		todoList:  list,
		parser:    parser,
		undoStack: undoStack,
		factory:   factory,
	}
}
