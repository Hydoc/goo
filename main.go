package main

import (
	"errors"
	"flag"
	"github.com/Hydoc/goo/internal"
	"github.com/Hydoc/goo/internal/command"
	"github.com/Hydoc/goo/internal/controller"
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

func main() {
	v := view.New(os.Stdout)
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
		v.RenderLine(usage)
	}
	flag.Parse()

	if len(*file) > 0 {
		filename = *file
	}

	if len(filename) == 0 {
		home, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		filename = filepath.Join(home, defaultFileName)
	}

	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		err = os.WriteFile(filename, []byte("[]"), 0644)
		if err != nil {
			v.RenderLine(err.Error())
			os.Exit(1)
		}
	}

	todoList, err := internal.NewTodoListFromFile(filename)
	if err != nil {
		v.RenderLine(err.Error())
		os.Exit(1)
	}

	args := strings.TrimSpace(strings.Join(flag.Args(), " "))
	commandFactory := command.NewFactory()
	ctr := controller.New(v, todoList, commandFactory)
	code, err := ctr.Handle(list, toggle, add, doDelete, edit, doClear, args)
	if err != nil {
		v.RenderLine(err.Error())
	}
	if code == 0 && !*list {
		v.RenderList(todoList)
	}
	os.Exit(code)
}
