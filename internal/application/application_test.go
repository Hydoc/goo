package application

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Hydoc/goo/internal/command"
	"github.com/Hydoc/goo/internal/model"
	"github.com/Hydoc/goo/internal/view"
)

func setUpFile(t *testing.T, filename string, content interface{}) func() {
	jsonContent, err := json.Marshal(content)

	if err != nil {
		t.Errorf("there was an error marshaling the content %s", err)
	}

	err = os.WriteFile(filename, jsonContent, 0644)
	if err != nil {
		t.Errorf("there was an error creating the file %s", err)
	}

	return func() {
		err := os.Remove(filename)
		if err != nil {
			t.Errorf("there was an error removing the file %s", err)
		}
	}
}

func userHomeDirWith(homeDir string, err error) func() (string, error) {
	return func() (string, error) {
		return homeDir, err
	}
}

func Test_Main(t *testing.T) {
	t.Run("when file in home dir does not exist, it should get created", func(t *testing.T) {
		buffer := bytes.NewBuffer(make([]byte, 0))
		oldArgs := os.Args
		defer func() {
			os.Args = oldArgs
		}()
		v := view.New(buffer)
		userHomeDir := "./"
		expectedFile := filepath.Join(userHomeDir, defaultFileName)
		os.Args = []string{"when file in home dir does not exist, it should get created"}
		Main(v, userHomeDirWith(userHomeDir, nil))

		if _, err := os.Stat(expectedFile); errors.Is(err, os.ErrNotExist) {
			t.Errorf("expected file %s to be created", expectedFile)
		}
		os.Remove(filename)
	})

	t.Run("without arguments should print usage", func(t *testing.T) {
		tearDown := setUpFile(t, defaultFileName, &model.TodoList{})
		oldArgs := os.Args
		defer func() {
			os.Args = oldArgs
		}()
		defer tearDown()
		buffer := bytes.NewBuffer(make([]byte, 0))
		v := view.New(buffer)

		flag.CommandLine = flag.NewFlagSet("with add flag", flag.ExitOnError)
		os.Args = []string{"without arguments should print usage"}
		Main(v, userHomeDirWith("./", nil))

		want := "How to use goo\n  -f, --file\n    Path to a file to use (has to be json, if the file does not exist it gets created, has to be the first argument before the subcommands)\n    goo -f path/to/file.json list\n\n  list: List all todos\n    goo list\n\n  toggle: Toggle the state of a todo by its id\n    goo toggle <ID of todo:int>\n\n  rm: Delete a todo by its id\n    goo rm <ID of todo:int>\n\n  edit: Edit a todo by its id and a new label, use '{}' to insert the old value\n    goo edit <ID of todo:int> <new label:string>\n\n  add: Add a new todo\n    goo add <label:string>\n\n  clear: Clear the whole list (no confirmation required)\n    goo clear\n\n  swap: Swap the labels of two todos by their id\n    goo swap <ID of the first todo:int> <ID of the second todo:int>\n\n  tags: List all tags\n    goo tags\n       -tid <ID of todo:int> show all tags on this todo\n       -id <ID of tag:int> show all todos with this tag\n\n  tag: tag handling\n    goo tag <ID of tag:int> <ID of the todo:int>\n       adds the given tag to the todo\n\n    goo tag -c <Label of the tag:string>\n       creates a new tag\n\n    goo tag -rm\n       remove a tag or a tag from a todo\n       -rm <ID of tag:int> removes the tag from all todos and the tag itself\n       -rm <ID of tag:int> <ID of todo:int> removes the specific tag from the todo\n  \n  version: prints the version\n    goo version\n"

		if buffer.String() != want {
			t.Errorf("want %#v, got %#v", want, buffer.String())
		}
	})

	t.Run("with add subcommand", func(t *testing.T) {
		tearDown := setUpFile(t, defaultFileName, &model.TodoList{})
		oldArgs := os.Args
		defer func() {
			os.Args = oldArgs
		}()
		defer tearDown()
		buffer := bytes.NewBuffer(make([]byte, 0))
		v := view.New(buffer)

		flag.CommandLine = flag.NewFlagSet("with add subcommand", flag.ExitOnError)
		os.Args = []string{"with add subcommand", "add", "Hello World"}
		Main(v, userHomeDirWith("./", nil))

		want := "ID  TASK         STATUS\n-----------------------\n1   Hello World    ○\n"

		if buffer.String() != want {
			t.Errorf("want %#v, got %#v", want, buffer.String())
		}
	})

	t.Run("with delete subcommand", func(t *testing.T) {
		tearDown := setUpFile(t, defaultFileName, &model.TodoList{
			Filename: "",
			Items: []*model.Todo{
				model.NewTodo("should be deleted", 1),
			},
			TagList: make([]*model.Tag, 0),
		})
		oldArgs := os.Args
		defer func() {
			os.Args = oldArgs
		}()
		defer tearDown()
		buffer := bytes.NewBuffer(make([]byte, 0))
		v := view.New(buffer)

		flag.CommandLine = flag.NewFlagSet("with delete subcommand", flag.ExitOnError)
		os.Args = []string{"with delete subcommand", "rm", "1"}
		Main(v, userHomeDirWith("./", nil))

		want := "ID  TASK      STATUS\n--------------------\n"

		if buffer.String() != want {
			t.Errorf("want %#v, got %#v", want, buffer.String())
		}
	})

	t.Run("with toggle subcommand", func(t *testing.T) {
		tearDown := setUpFile(t, defaultFileName, &model.TodoList{
			Filename: "",
			Items: []*model.Todo{
				model.NewTodo("should be toggled", 1),
			},
			TagList: make([]*model.Tag, 0),
		})
		oldArgs := os.Args
		defer func() {
			os.Args = oldArgs
		}()
		defer tearDown()
		buffer := bytes.NewBuffer(make([]byte, 0))
		v := view.New(buffer)

		flag.CommandLine = flag.NewFlagSet("with toggle subcommand", flag.ExitOnError)
		os.Args = []string{"with toggle subcommand", "toggle", "1"}
		Main(v, userHomeDirWith("./", nil))

		want := "ID  TASK               STATUS\n-----------------------------\n\x1b[90m1   should be toggled    ✓\x1b[0m\n"

		if buffer.String() != want {
			t.Errorf("want %#v, got %#v", want, buffer.String())
		}
	})

	t.Run("with edit subcommand", func(t *testing.T) {
		tearDown := setUpFile(t, defaultFileName, &model.TodoList{
			Filename: "",
			Items: []*model.Todo{
				model.NewTodo("should be changed", 1),
			},
			TagList: make([]*model.Tag, 0),
		})
		oldArgs := os.Args
		defer func() {
			os.Args = oldArgs
		}()
		defer tearDown()
		buffer := bytes.NewBuffer(make([]byte, 0))
		v := view.New(buffer)

		flag.CommandLine = flag.NewFlagSet("with edit subcommand", flag.ExitOnError)
		os.Args = []string{"with edit subcommand", "edit", "1 Hello there!"}
		Main(v, userHomeDirWith("./", nil))

		want := "ID  TASK          STATUS\n------------------------\n1   Hello there!    ○\n"

		if buffer.String() != want {
			t.Errorf("want %#v, got %#v", want, buffer.String())
		}
	})

	t.Run("with clear subcommand", func(t *testing.T) {
		tearDown := setUpFile(t, defaultFileName, &model.TodoList{
			Filename: "",
			Items: []*model.Todo{
				model.NewTodo("should be deleted", 1),
				model.NewTodo("should also be deleted", 2),
			},
			TagList: make([]*model.Tag, 0),
		})
		oldArgs := os.Args
		defer func() {
			os.Args = oldArgs
		}()
		defer tearDown()
		buffer := bytes.NewBuffer(make([]byte, 0))
		v := view.New(buffer)

		flag.CommandLine = flag.NewFlagSet("with clear subcommand", flag.ExitOnError)
		os.Args = []string{"with clear subcommand", "clear"}
		Main(v, userHomeDirWith("./", nil))

		want := "ID  TASK      STATUS\n--------------------\n"

		if buffer.String() != want {
			t.Errorf("want %#v, got %#v", want, buffer.String())
		}
	})

	t.Run("with list subcommand", func(t *testing.T) {
		tearDown := setUpFile(t, defaultFileName, &model.TodoList{
			Filename: "",
			Items: []*model.Todo{
				model.NewTodo("Hi", 1),
			},
			TagList: make([]*model.Tag, 0),
		})
		oldArgs := os.Args
		defer func() {
			os.Args = oldArgs
		}()
		defer tearDown()
		buffer := bytes.NewBuffer(make([]byte, 0))
		v := view.New(buffer)

		flag.CommandLine = flag.NewFlagSet("with list subcommand", flag.ExitOnError)
		os.Args = []string{"with list subcommand", "list"}
		Main(v, userHomeDirWith("./", nil))

		want := "ID  TASK    STATUS\n------------------\n1   Hi        ○\n"

		if buffer.String() != want {
			t.Errorf("want %#v, got %#v", want, buffer.String())
		}
	})

	t.Run("with file flag", func(t *testing.T) {
		fileToUse := "another.json"
		tearDownFile := setUpFile(t, fileToUse, &model.TodoList{
			Filename: "",
			Items: []*model.Todo{
				model.NewTodo("I should be printed", 1),
			},
			TagList: make([]*model.Tag, 0),
		})
		defer tearDownFile()
		oldArgs := os.Args
		defer func() {
			os.Args = oldArgs
		}()
		buffer := bytes.NewBuffer(make([]byte, 0))
		v := view.New(buffer)

		flag.CommandLine = flag.NewFlagSet("with file flag", flag.ExitOnError)
		os.Args = []string{"with file flag", "-f", fileToUse, "list"}
		Main(v, userHomeDirWith("./", nil))

		want := "ID  TASK                 STATUS\n-------------------------------\n1   I should be printed    ○\n"

		if buffer.String() != want {
			t.Errorf("want %#v, got %#v", want, buffer.String())
		}
	})

	t.Run("print error when file creating fails", func(t *testing.T) {
		fileToUse := "/my-file.json"
		oldArgs := os.Args
		defer func() {
			os.Args = oldArgs
		}()
		buffer := bytes.NewBuffer(make([]byte, 0))
		v := view.New(buffer)

		os.Args = []string{"print error when file creating fails", "-f", fileToUse}
		flag.CommandLine = flag.NewFlagSet("print error when file creating fails", flag.ExitOnError)
		Main(v, userHomeDirWith("", nil))

		want := "open /my-file.json: permission denied\n"

		if buffer.String() != want {
			t.Errorf("want %#v, got %#v", want, buffer.String())
		}
	})

	t.Run("correct flag.Usage", func(t *testing.T) {
		filename = defaultFileName
		tearDown := setUpFile(t, defaultFileName, &model.TodoList{})
		oldArgs := os.Args
		defer func() {
			os.Args = oldArgs
		}()
		defer tearDown()
		buffer := bytes.NewBuffer(make([]byte, 0))
		v := view.New(buffer)

		os.Args = []string{"correct flag.Usage"}
		flag.CommandLine = flag.NewFlagSet("correct flag.Usage", flag.ExitOnError)
		Main(v, userHomeDirWith("./", nil))
		flag.Usage()

		want := "How to use goo\n  -f, --file\n    Path to a file to use (has to be json, if the file does not exist it gets created, has to be the first argument before the subcommands)\n    goo -f path/to/file.json list\n\n  list: List all todos\n    goo list\n\n  toggle: Toggle the state of a todo by its id\n    goo toggle <ID of todo:int>\n\n  rm: Delete a todo by its id\n    goo rm <ID of todo:int>\n\n  edit: Edit a todo by its id and a new label, use '{}' to insert the old value\n    goo edit <ID of todo:int> <new label:string>\n\n  add: Add a new todo\n    goo add <label:string>\n\n  clear: Clear the whole list (no confirmation required)\n    goo clear\n\n  swap: Swap the labels of two todos by their id\n    goo swap <ID of the first todo:int> <ID of the second todo:int>\n\n  tags: List all tags\n    goo tags\n       -tid <ID of todo:int> show all tags on this todo\n       -id <ID of tag:int> show all todos with this tag\n\n  tag: tag handling\n    goo tag <ID of tag:int> <ID of the todo:int>\n       adds the given tag to the todo\n\n    goo tag -c <Label of the tag:string>\n       creates a new tag\n\n    goo tag -rm\n       remove a tag or a tag from a todo\n       -rm <ID of tag:int> removes the tag from all todos and the tag itself\n       -rm <ID of tag:int> <ID of todo:int> removes the specific tag from the todo\n  \n  version: prints the version\n    goo version\n"

		if !strings.Contains(buffer.String(), want) {
			t.Errorf("want %#v, got %#v", want, buffer.String())
		}
	})

	t.Run("todolist creating fails", func(t *testing.T) {
		tearDown := setUpFile(t, defaultFileName, map[string]interface{}{
			"tagList": []int{1, 2, 3},
		})
		defer tearDown()
		oldArgs := os.Args
		defer func() {
			os.Args = oldArgs
		}()
		buffer := bytes.NewBuffer(make([]byte, 0))
		v := view.New(buffer)

		os.Args = []string{"todolist creating fails", "-f", defaultFileName}
		flag.CommandLine = flag.NewFlagSet("todolist creating fails", flag.ExitOnError)
		Main(v, userHomeDirWith("./", nil))

		want := "json: cannot unmarshal number into Go struct field TodoList.tagList of type model.Tag\n"

		if buffer.String() != want {
			t.Errorf("want %#v, got %#v", want, buffer.String())
		}
	})

	t.Run("print error when user home dir returns error", func(t *testing.T) {
		filename = ""
		oldArgs := os.Args
		defer func() {
			os.Args = oldArgs
		}()
		buffer := bytes.NewBuffer(make([]byte, 0))
		v := view.New(buffer)
		wantError := "directory not found\n"

		os.Args = []string{"print error when user home dir fails"}
		flag.CommandLine = flag.NewFlagSet("print error when user home dir fails", flag.ExitOnError)
		Main(v, userHomeDirWith("./", errors.New("directory not found")))

		if buffer.String() != wantError {
			t.Errorf("want %#v, got %#v", wantError, buffer.String())
		}
	})

	t.Run("non existent subcommand should print error", func(t *testing.T) {
		tearDown := setUpFile(t, defaultFileName, &model.TodoList{})
		defer tearDown()
		oldArgs := os.Args
		defer func() {
			os.Args = oldArgs
		}()
		buffer := bytes.NewBuffer(make([]byte, 0))
		v := view.New(buffer)

		os.Args = []string{"non existent subcommand should print error", "abc"}
		flag.CommandLine = flag.NewFlagSet("non existent subcommand should print error", flag.ExitOnError)
		Main(v, userHomeDirWith("./", nil))

		want := fmt.Sprintf("%s\n", command.ErrCommandNotFound)

		if buffer.String() != want {
			t.Errorf("want %#v, got %#v", want, buffer.String())
		}
	})
}
