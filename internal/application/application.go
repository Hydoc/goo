package application

import (
	"errors"
	"flag"
	"fmt"
	"github.com/Hydoc/goo/internal/command"
	"github.com/Hydoc/goo/internal/model"
	"github.com/Hydoc/goo/internal/view"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	defaultFileName = ".goo.json"
	usage           = `How to use goo
  -f, --file
    Path to a file to use (has to be json, if the file does not exist it gets created, has to be the first argument before the subcommands)
    goo -f path/to/file.json list

  list: List all todos
    goo list

  toggle: Toggle the state of a todo by its id
    goo toggle <ID of todo>

  rm: Delete a todo by its id
    goo rm <ID of todo>

  edit: Edit a todo by its id and a new label, use '{}' to insert the old value
    goo edit <ID of todo> <new label>

  add: Add a new todo
    goo add <label>

  clear: Clear the whole list (no confirmation required)
    goo clear

  swap: Swap the labels of two todos by their id
    goo swap <ID of the first todo> <ID of the second todo>

  tags: List all tags
    goo tags
       -tid <ID of todo> show all tags on this todo
       -id <ID of tag> show all todos with this tag

  tag: tag handling
    goo tag <ID of todo> <label of the tag>
       creates the tag if it does not exist and adds it to the todo

    goo tag -rm
       remove a tag or a tag from a todo
       -rm <ID of tag> removes the tag from all todos and the tag itself
       -rm <ID of tag> <ID of todo> removes the specific tag from the todo`
)

var filename string

func createFileIfNotExists() error {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		err = os.WriteFile(filename, []byte("{}"), 0644)
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
	tags := flag.NewFlagSet("tags", flag.ExitOnError)
	showTagsOnTodo := tags.Int("tid", 0, "<ID of todo> show all tags on this todo")
	showTodosForTag := tags.Int("id", 0, "<ID of tag> show all todos with this tag")
	tag := flag.NewFlagSet("tag", flag.ExitOnError)
	tagRm := tag.Bool("rm", false, "remove a tag")

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

	var fabricateCommand command.FabricateCommand
	args := strings.TrimSpace(strings.Join(flag.Args()[1:], " "))
	fmt.Println(flag.Args())
	switch flag.Args()[0] {
	case list.Name():
		fabricateCommand = command.NewListTodos
	case add.Name():
		fabricateCommand = command.NewAddTodo
	case doDelete.Name():
		fabricateCommand = command.NewDeleteTodo
	case toggle.Name():
		fabricateCommand = command.NewToggleTodo
	case edit.Name():
		fabricateCommand = command.NewEditTodo
	case doClear.Name():
		fabricateCommand = command.NewClear
	case swap.Name():
		fabricateCommand = command.NewSwap
	case tags.Name():
		tags.Parse(flag.Args()[1:])
		switch {
		case *showTagsOnTodo > 0:
			args = strconv.Itoa(*showTagsOnTodo)
			fabricateCommand = command.NewListTagsOnTodo
		case *showTodosForTag > 0:
			args = strconv.Itoa(*showTodosForTag)
			fabricateCommand = command.NewListTodosForTag
		default:
			fabricateCommand = command.NewListTags
		}
	case tag.Name():
		tag.Parse(flag.Args()[1:])
		switch {
		case *tagRm && len(tag.Args()) > 1:
			fabricateCommand = command.NewRemoveTagFromTodo
		case *tagRm && len(tag.Args()) == 1:
			fabricateCommand = command.NewRemoveTag
		default:
			fabricateCommand = command.NewTagTodo
		}
	default:
		flag.Usage()
		return 2
	}

	cmd, err := fabricateCommand(todoList, view, args)
	if err != nil {
		view.RenderLine(err.Error())
		return 1
	}
	cmd.Execute()
	return 0
}
