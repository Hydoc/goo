package application

import (
	"errors"
	"flag"
	"github.com/Hydoc/goo/internal/command"
	"github.com/Hydoc/goo/internal/model"
	"github.com/Hydoc/goo/internal/view"
	"os"
	"path/filepath"
	"strings"
)

const (
	defaultFileName = ".goo.json"
	usage           = `How to use goo
  -f, --file
    	Path to a file to use (has to be json, if the file does not exist it gets created)

  list: List all todos
  toggle: Toggle the state of a todo by its id
  rm: Delete a todo by its id
  edit: Edit a todo by its id and a new label, use '{}' to insert the old value
  add: Add a new todo
  clear: Clear the whole list
  swap: Swap the labels of two todos by their id`
)

var filename string

func createFileIfNotExists() error {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		err = os.WriteFile(filename, []byte("[]"), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func Main(view view.View, userHomeDir func() (string, error)) int {
	file := flag.String("file", "", "Path to a file to use (has to be json, if the file does not exist it gets created)")
	flag.StringVar(file, "f", "", "Path to a file to use (has to be json, if the file does not exist it gets created)")

	list := flag.NewFlagSet("list", flag.ExitOnError)
	add := flag.NewFlagSet("add", flag.ExitOnError)
	doDelete := flag.NewFlagSet("rm", flag.ExitOnError)
	toggle := flag.NewFlagSet("toggle", flag.ExitOnError)
	edit := flag.NewFlagSet("edit", flag.ExitOnError)
	doClear := flag.NewFlagSet("clear", flag.ExitOnError)
	swap := flag.NewFlagSet("swap", flag.ExitOnError)

	cmdMap := map[string]command.FabricateCommand{
		list.Name():     command.NewListTodos,
		add.Name():      command.NewAddTodo,
		doDelete.Name(): command.NewDeleteTodo,
		toggle.Name():   command.NewToggleTodo,
		edit.Name():     command.NewEditTodo,
		doClear.Name():  command.NewClear,
		swap.Name():     command.NewSwap,
	}

	flag.Usage = func() {
		view.RenderLine(usage)
	}
	flag.Parse()

	if len(*file) > 0 {
		filename = *file
	}

	if len(filename) == 0 {
		homeDir, err := userHomeDir()
		if err != nil {
			view.RenderLine(err.Error())
			return 1
		}
		filename = filepath.Join(homeDir, defaultFileName)
	}

	err := createFileIfNotExists()

	if err != nil {
		view.RenderLine(err.Error())
		return 1
	}

	todoList, err := model.NewTodoListFromFile(filename)
	if err != nil {
		view.RenderLine(err.Error())
		return 1
	}

	if len(os.Args) < 2 || (len(*file) > 0 && len(os.Args) < 4) {
		flag.Usage()
		return 2
	}

	fabricateCommand, ok := cmdMap[flag.Args()[0]]
	if !ok {
		flag.Usage()
		return 2
	}

	args := strings.TrimSpace(strings.Join(flag.Args()[1:], " "))
	cmd, err := fabricateCommand(todoList, view, args)
	if err != nil {
		view.RenderLine(err.Error())
		return 1
	}
	cmd.Execute()
	return 0
}
