package command

import (
	"github.com/Hydoc/goo/internal"
	"os"
)

var QuitAliases = []string{"q"}

const (
	QuitAbbr = "quit"
	QuitDesc = "Quit the app"
)

type Quit struct {
	todoList *internal.TodoList
}

func (cmd *Quit) Execute() {
	cmd.todoList.SaveToFile()
	os.Exit(0)
}

func newQuit(todoList *internal.TodoList) *Quit {
	return &Quit{todoList: todoList}
}
