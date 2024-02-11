package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"github.com/Hydoc/goo/internal"
	"github.com/Hydoc/goo/internal/command"
	"github.com/Hydoc/goo/internal/controller"
	"github.com/Hydoc/goo/internal/view"
	"os"
)

var filename string

func main() {
	flag.StringVar(&filename, "file", "", "-file path to file to read (has to be json, if the file does not exists it gets created)")
	flag.Parse()

	if len(filename) == 0 {
		fmt.Println("file is missing, run goo -file <path to file>")
		return
	}

	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		err = os.WriteFile(filename, []byte("[]"), 0644)
		if err != nil {
			panic(err)
		}
	}

	quit := command.NewStringCommand(command.QuitAbbr, command.QuitDesc, command.QuitAliases)
	help := command.NewStringCommand(command.HelpAbbr, command.HelpDesc, command.HelpAliases)
	addTodo := command.NewStringCommand(command.AddTodoAbbr, command.AddTodoDesc, command.AddTodoAliases)
	toggleTodo := command.NewStringCommand(command.ToggleTodoAbbr, command.ToggleTodoDesc, command.ToggleTodoAliases)
	undo := command.NewStringCommand(command.UndoAbbr, command.UndoDesc, command.UndoAliases)
	deleteTodo := command.NewStringCommand(command.DeleteTodoAbbr, command.DeleteTodoDesc, command.DeleteTodoAliases)
	validCommands := []*command.StringCommand{quit, help, addTodo, toggleTodo, undo, deleteTodo}
	todoList, err := internal.NewTodoListFromFile(filename)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(os.Stdin)
	parser := command.NewParser(validCommands)
	factory := command.NewFactory(validCommands)
	undoStack := command.NewUndoStack()
	v := view.New(reader)
	ctr := controller.New(v, todoList, parser, undoStack, factory)
	ctr.Run()
}
