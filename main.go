package main

import (
	"bufio"
	"github.com/Hydoc/goo/internal"
	"github.com/Hydoc/goo/internal/command"
	"github.com/Hydoc/goo/internal/controller"
	"github.com/Hydoc/goo/internal/view"
	"os"
)

func main() {
	quit := command.NewStringCommand(command.QuitAbbr, command.QuitDesc, command.QuitAliases)
	help := command.NewStringCommand(command.HelpAbbr, command.HelpDesc, command.HelpAliases)
	addTodo := command.NewStringCommand(command.AddTodoAbbr, command.AddTodoDesc, command.AddTodoAliases)
	toggleTodo := command.NewStringCommand(command.ToggleTodoAbbr, command.ToggleTodoDesc, command.ToggleTodoAliases)
	undo := command.NewStringCommand(command.UndoAbbr, command.UndoDesc, command.UndoAliases)
	todoList := internal.NewTodoList()
	reader := bufio.NewReader(os.Stdin)
	parser := command.NewParser([]*command.StringCommand{quit, help, addTodo, toggleTodo, undo})
	undoStack := command.NewUndoStack()
	v := view.New(reader)
	cont := controller.New(v, todoList, parser, undoStack)
	cont.Run()
}
