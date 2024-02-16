package controller

import (
	"encoding/json"
	"errors"
	"github.com/Hydoc/goo/internal/command"
	"github.com/Hydoc/goo/internal/model"
	"os"
	"reflect"
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

type DummyView struct {
	RenderListCalls int
}

func (d *DummyView) RenderList(_ *model.TodoList) {
	d.RenderListCalls++
}

func newDummyView() *DummyView {
	return &DummyView{0}
}

func TestController_Handle(t *testing.T) {
	t.Run("call view.RenderList when list command triggers", func(t *testing.T) {
		todoList := &model.TodoList{
			Items: []*model.Todo{
				model.NewTodo("Hello", 1),
			},
		}
		dummyView := newDummyView()
		ctr := New(dummyView, todoList, command.NewFactory())

		code, err := ctr.Handle(true, 0, false, 0, false, false, "")

		if err != nil {
			t.Errorf("want no error, but got %v", err)
		}

		if code != 0 {
			t.Errorf("want code 0, got %v", code)
		}

		if dummyView.RenderListCalls == 0 {
			t.Error("want dummyView.RenderListCalls to be greater than 0")
		}
	})

	t.Run("do nothing when no command was passed", func(t *testing.T) {
		todoList := &model.TodoList{
			Items: []*model.Todo{
				model.NewTodo("Hello", 1),
			},
		}
		ctr := New(newDummyView(), todoList, command.NewFactory())

		code, err := ctr.Handle(false, 0, false, 0, false, false, "")

		if err != nil {
			t.Errorf("want no error, but got %v", err)
		}

		if code != 0 {
			t.Errorf("want code 0, got %v", code)
		}
	})

	tests := []struct {
		name               string
		toggle             int
		add                bool
		doDelete           int
		edit               bool
		doClear            bool
		args               string
		err                error
		wantCode           int
		wantAfterExecution []*model.Todo
		todoItems          []*model.Todo
	}{
		{
			name:     "handle toggle",
			toggle:   1,
			add:      false,
			doDelete: 0,
			edit:     false,
			doClear:  false,
			err:      nil,
			args:     "",
			wantCode: 0,
			wantAfterExecution: []*model.Todo{
				{
					Id:     1,
					Label:  "Test",
					IsDone: true,
				},
			},
			todoItems: []*model.Todo{
				{
					Id:     1,
					Label:  "Test",
					IsDone: false,
				},
			},
		},
		{
			name:     "handle add",
			toggle:   0,
			add:      true,
			doDelete: 0,
			edit:     false,
			doClear:  false,
			args:     "Hello World",
			err:      nil,
			wantCode: 0,
			wantAfterExecution: []*model.Todo{
				{
					Id:     1,
					Label:  "Hello World",
					IsDone: false,
				},
			},
			todoItems: make([]*model.Todo, 0),
		},
		{
			name:               "handle delete",
			toggle:             0,
			add:                false,
			doDelete:           1,
			edit:               false,
			doClear:            false,
			args:               "",
			err:                nil,
			wantCode:           0,
			wantAfterExecution: make([]*model.Todo, 0),
			todoItems: []*model.Todo{
				model.NewTodo("abc", 1),
			},
		},
		{
			name:     "handle edit",
			toggle:   0,
			add:      false,
			doDelete: 0,
			edit:     true,
			doClear:  false,
			args:     "1 Hello{}",
			err:      nil,
			wantCode: 0,
			wantAfterExecution: []*model.Todo{
				{
					Id:     1,
					Label:  "Hello!",
					IsDone: false,
				},
			},
			todoItems: []*model.Todo{
				model.NewTodo("!", 1),
			},
		},
		{
			name:               "handle clear",
			toggle:             0,
			add:                false,
			doDelete:           0,
			edit:               false,
			doClear:            true,
			args:               "",
			err:                nil,
			wantCode:           0,
			wantAfterExecution: make([]*model.Todo, 0),
			todoItems: []*model.Todo{
				model.NewTodo("ABC", 2),
				model.NewTodo("Hello", 3),
				model.NewTodo("World", 5),
			},
		},
		{
			name:               "return error when factory returns error",
			toggle:             0,
			add:                true,
			doDelete:           0,
			edit:               false,
			doClear:            false,
			args:               "",
			err:                errors.New("empty todo is not allowed"),
			wantCode:           1,
			wantAfterExecution: nil,
			todoItems:          nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			filename := "test.json"
			tearDown := setUpFile(t, filename, test.todoItems)
			defer tearDown()

			actualTodoList, _ := model.NewTodoListFromFile(filename)

			ctr := New(newDummyView(), actualTodoList, command.NewFactory())
			code, err := ctr.Handle(false, test.toggle, test.add, test.doDelete, test.edit, test.doClear, test.args)

			writtenTodoList, _ := model.NewTodoListFromFile(filename)

			if !reflect.DeepEqual(writtenTodoList.Items, test.wantAfterExecution) {
				t.Errorf("want written list %v, got %v", test.wantAfterExecution, writtenTodoList)
			}

			if err != nil && !reflect.DeepEqual(err, test.err) {
				t.Errorf("want err %v, got %v", test.err, err)
			}

			if test.wantCode != code {
				t.Errorf("want code %v, got %v", test.wantCode, code)
			}
		})
	}
}
