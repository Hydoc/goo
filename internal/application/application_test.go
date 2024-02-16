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

func Test_Main(t *testing.T) {
	t.Run("when file in home dir does not exist, it should get created", func(t *testing.T) {
		buffer := bytes.NewBuffer(make([]byte, 0))
		v := view.New(buffer)
		userHomeDir := "./"
		expectedFile := filepath.Join(userHomeDir, defaultFileName)
		Main(v, userHomeDir)

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
		Main(v, "./")

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
		Main(v, "./")

		want := "ID  TASK         STATUS\n-----------------------\n1   Hello World    ○\n"

		if buffer.String() != want {
			t.Errorf("want %#v, got %#v", want, buffer.String())
		}
	})
}
