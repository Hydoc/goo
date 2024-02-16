package application

import (
	"errors"
	"flag"
	"github.com/Hydoc/goo/internal/command"
	"github.com/Hydoc/goo/internal/controller"
	"github.com/Hydoc/goo/internal/model"
	"github.com/Hydoc/goo/internal/view"
	"os"
	"path/filepath"
	"strings"
)

const (
	defaultFileName = ".goo.json"
	usage           = `How to use goo
  -h, --help
    	Prints this information

  -f, --file
    	Path to a file to use (has to be json, if the file does not exist it gets created)

  -l, --list
        List all todos

  -t, --toggle
        Toggle the state of a todo by its id

  -d, --delete
        Delete a todo by its id

  -e, --edit
        Edit a todo by its id and a new label, use '{}' to insert the old value
             e.g goo --edit 1 {} World!

  -a, --add
        Add a new todo

  --clear
        Clear the whole list`
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

func Main(view view.View, userHomeDir string) {
	file := flag.String("file", "", "Path to a file to use (has to be json, if the file does not exist it gets created)")
	flag.StringVar(file, "f", "", "Path to a file to use (has to be json, if the file does not exist it gets created)")

	list := flag.Bool("list", false, "List all todos")
	flag.BoolVar(list, "l", false, "List all todos")

	toggle := flag.Int("toggle", 0, "Toggle the state of a todo by its id")
	flag.IntVar(toggle, "t", 0, "Toggle the state of a todo by its id")

	add := flag.Bool("add", false, "Add a new todo")
	flag.BoolVar(add, "a", false, "Add a new todo")

	doDelete := flag.Int("delete", 0, "Delete a todo by its id")
	flag.IntVar(doDelete, "d", 0, "Delete a todo by its id")

	edit := flag.Bool("edit", false, "Edit a todo by its id and a new label, use '{}' to insert the old value")
	flag.BoolVar(edit, "e", false, "Edit a todo by its id and a new label, use '{}' to insert the old value")

	doClear := flag.Bool("clear", false, "")

	flag.Usage = func() {
		view.RenderLine(usage)
	}
	flag.Parse()

	if len(*file) > 0 {
		filename = *file
	}

	if len(filename) == 0 {
		filename = filepath.Join(userHomeDir, defaultFileName)
	}

	err := createFileIfNotExists()

	if err != nil {
		view.RenderLine(err.Error())
		return
	}

	todoList, err := model.NewTodoListFromFile(filename)
	if err != nil {
		view.RenderLine(err.Error())
		return
	}

	args := strings.TrimSpace(strings.Join(flag.Args(), " "))
	commandFactory := command.NewFactory()
	ctr := controller.New(view, todoList, commandFactory)
	code, err := ctr.Handle(*list, *toggle, *add, *doDelete, *edit, *doClear, args)
	if err != nil {
		view.RenderLine(err.Error())
	}
	if code == 0 && !*list {
		view.RenderList(todoList)
	}
}
