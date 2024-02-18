package application

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"github.com/Hydoc/goo/internal/model"
	"github.com/Hydoc/goo/internal/view"
	"os"
	"path/filepath"
	"strings"
	"testing"
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
		v := view.New(buffer)
		userHomeDir := "./"
		expectedFile := filepath.Join(userHomeDir, defaultFileName)
		Main(v, userHomeDirWith(userHomeDir, nil))

		if _, err := os.Stat(expectedFile); errors.Is(err, os.ErrNotExist) {
			t.Errorf("expected file %s to be created", expectedFile)
		}
		os.Remove(filename)
	})

	t.Run("without arguments should print the list", func(t *testing.T) {
		tearDown := setUpFile(t, "./.goo.json", make([]*model.Todo, 0))
		oldArgs := os.Args
		defer func() {
			os.Args = oldArgs
		}()
		defer tearDown()
		buffer := bytes.NewBuffer(make([]byte, 0))
		v := view.New(buffer)

		flag.CommandLine = flag.NewFlagSet("with add flag", flag.ExitOnError)
		os.Args = []string{"without arguments should print the list"}
		Main(v, userHomeDirWith("./", nil))

		want := "ID  TASK      STATUS\n--------------------\n"

		if buffer.String() != want {
			t.Errorf("want %#v, got %#v", want, buffer.String())
		}
	})

	t.Run("with add flag", func(t *testing.T) {
		tearDown := setUpFile(t, "./.goo.json", make([]*model.Todo, 0))
		oldArgs := os.Args
		defer func() {
			os.Args = oldArgs
		}()
		defer tearDown()
		buffer := bytes.NewBuffer(make([]byte, 0))
		v := view.New(buffer)

		flag.CommandLine = flag.NewFlagSet("with add flag", flag.ExitOnError)
		os.Args = []string{"with add flag", "-a", "Hello World"}
		Main(v, userHomeDirWith("./", nil))

		want := "ID  TASK         STATUS\n-----------------------\n1   Hello World    ○\n"

		if buffer.String() != want {
			t.Errorf("want %#v, got %#v", want, buffer.String())
		}
	})

	t.Run("with delete flag", func(t *testing.T) {
		tearDown := setUpFile(t, "./.goo.json", []*model.Todo{
			model.NewTodo("should be deleted", 1),
		})
		oldArgs := os.Args
		defer func() {
			os.Args = oldArgs
		}()
		defer tearDown()
		buffer := bytes.NewBuffer(make([]byte, 0))
		v := view.New(buffer)

		flag.CommandLine = flag.NewFlagSet("with delete flag", flag.ExitOnError)
		os.Args = []string{"with delete flag", "-d", "1"}
		Main(v, userHomeDirWith("./", nil))

		want := "ID  TASK      STATUS\n--------------------\n"

		if buffer.String() != want {
			t.Errorf("want %#v, got %#v", want, buffer.String())
		}
	})

	t.Run("with toggle flag", func(t *testing.T) {
		tearDown := setUpFile(t, "./.goo.json", []*model.Todo{
			model.NewTodo("should be toggled", 1),
		})
		oldArgs := os.Args
		defer func() {
			os.Args = oldArgs
		}()
		defer tearDown()
		buffer := bytes.NewBuffer(make([]byte, 0))
		v := view.New(buffer)

		flag.CommandLine = flag.NewFlagSet("with toggle flag", flag.ExitOnError)
		os.Args = []string{"with toggle flag", "-t", "1"}
		Main(v, userHomeDirWith("./", nil))

		want := "ID  TASK               STATUS\n-----------------------------\n\x1b[90m1   should be toggled    ✓\x1b[0m\n"

		if buffer.String() != want {
			t.Errorf("want %#v, got %#v", want, buffer.String())
		}
	})

	t.Run("with edit flag", func(t *testing.T) {
		tearDown := setUpFile(t, "./.goo.json", []*model.Todo{
			model.NewTodo("should be changed", 1),
		})
		oldArgs := os.Args
		defer func() {
			os.Args = oldArgs
		}()
		defer tearDown()
		buffer := bytes.NewBuffer(make([]byte, 0))
		v := view.New(buffer)

		flag.CommandLine = flag.NewFlagSet("with edit flag", flag.ExitOnError)
		os.Args = []string{"with edit flag", "-e", "1 Hello there!"}
		Main(v, userHomeDirWith("./", nil))

		want := "ID  TASK          STATUS\n------------------------\n1   Hello there!    ○\n"

		if buffer.String() != want {
			t.Errorf("want %#v, got %#v", want, buffer.String())
		}
	})

	t.Run("with clear flag", func(t *testing.T) {
		tearDown := setUpFile(t, "./.goo.json", []*model.Todo{
			model.NewTodo("should be deleted", 1),
			model.NewTodo("should also be deleted", 2),
		})
		oldArgs := os.Args
		defer func() {
			os.Args = oldArgs
		}()
		defer tearDown()
		buffer := bytes.NewBuffer(make([]byte, 0))
		v := view.New(buffer)

		flag.CommandLine = flag.NewFlagSet("with clear flag", flag.ExitOnError)
		os.Args = []string{"with clear flag", "-clear"}
		Main(v, userHomeDirWith("./", nil))

		want := "ID  TASK      STATUS\n--------------------\n"

		if buffer.String() != want {
			t.Errorf("want %#v, got %#v", want, buffer.String())
		}
	})

	t.Run("with list flag", func(t *testing.T) {
		tearDown := setUpFile(t, "./.goo.json", []*model.Todo{
			model.NewTodo("Hi", 1),
		})
		oldArgs := os.Args
		defer func() {
			os.Args = oldArgs
		}()
		defer tearDown()
		buffer := bytes.NewBuffer(make([]byte, 0))
		v := view.New(buffer)

		flag.CommandLine = flag.NewFlagSet("with list flag", flag.ExitOnError)
		os.Args = []string{"with list flag", "-l"}
		Main(v, userHomeDirWith("./", nil))

		want := "ID  TASK    STATUS\n------------------\n1   Hi        ○\n"

		if buffer.String() != want {
			t.Errorf("want %#v, got %#v", want, buffer.String())
		}
	})

	t.Run("with file flag", func(t *testing.T) {
		fileToUse := "another.json"
		tearDownFile := setUpFile(t, fileToUse, []*model.Todo{
			model.NewTodo("I should be printed", 1),
		})
		defer tearDownFile()
		oldArgs := os.Args
		defer func() {
			os.Args = oldArgs
		}()
		buffer := bytes.NewBuffer(make([]byte, 0))
		v := view.New(buffer)

		flag.CommandLine = flag.NewFlagSet("with file flag", flag.ExitOnError)
		os.Args = []string{"with file flag", "-f", fileToUse, "-l"}
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
		tearDown := setUpFile(t, "./.goo.json", make([]*model.Todo, 0))
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

		want := "How to use goo\n  -h, --help\n    \tPrints this information\n\n  -f, --file\n    \tPath to a file to use (has to be json, if the file does not exist it gets created)\n\n  -l, --list\n        List all todos\n\n  -t, --toggle\n        Toggle the state of a todo by its id\n\n  -d, --delete\n        Delete a todo by its id\n\n  -e, --edit\n        Edit a todo by its id and a new label, use '{}' to insert the old value\n             e.g goo --edit 1 {} World!\n\n  -a, --add\n        Add a new todo\n\n  --clear\n        Clear the whole list\n"

		if !strings.Contains(buffer.String(), want) {
			t.Errorf("want %#v, got %#v", want, buffer.String())
		}
	})

	t.Run("todolist creating fails", func(t *testing.T) {
		file := "./.goo.json"
		tearDown := setUpFile(t, file, map[string]int{})
		defer tearDown()
		oldArgs := os.Args
		defer func() {
			os.Args = oldArgs
		}()
		buffer := bytes.NewBuffer(make([]byte, 0))
		v := view.New(buffer)

		os.Args = []string{"todolist creating fails", "-f", file}
		flag.CommandLine = flag.NewFlagSet("todolist creating fails", flag.ExitOnError)
		Main(v, userHomeDirWith("./", nil))

		want := "json: cannot unmarshal object into Go value of type []*model.Todo\n"

		if buffer.String() != want {
			t.Errorf("want %#v, got %#v", want, buffer.String())
		}
	})

	t.Run("print error when controller returns error", func(t *testing.T) {
		file := "./.goo.json"
		tearDown := setUpFile(t, file, make([]*model.Todo, 0))
		defer tearDown()
		oldArgs := os.Args
		defer func() {
			os.Args = oldArgs
		}()
		buffer := bytes.NewBuffer(make([]byte, 0))
		v := view.New(buffer)

		os.Args = []string{"todolist creating fails", "-a"}
		flag.CommandLine = flag.NewFlagSet("todolist creating fails", flag.ExitOnError)
		Main(v, userHomeDirWith("./", nil))

		want := "empty todo is not allowed\n"

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
}
