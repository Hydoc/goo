package command

import (
	"errors"
	"github.com/Hydoc/goo/internal/model"
	"os"
	"reflect"
	"testing"
)

type dummyView struct {
	RenderListCalls int
	RenderLineCalls int
}

func (d *dummyView) RenderList(_ *model.TodoList) {
	d.RenderListCalls++
}

func (d *dummyView) RenderLine(_ string) {
	d.RenderLineCalls++
}

func newDummyView() *dummyView {
	return &dummyView{0, 0}
}

func TestNewAddTodo(t *testing.T) {
	todoList := &model.TodoList{}
	tests := []struct {
		name    string
		payload string
		want    Command
		err     error
	}{
		{
			name:    "create normally",
			payload: "test",
			want:    &AddTodo{view: newDummyView(), todoList: todoList, todoToAdd: "test"},
			err:     nil,
		},
		{
			name:    "not create due to missing payload",
			payload: "",
			want:    nil,
			err:     errors.New("empty todo is not allowed"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := NewAddTodo(todoList, newDummyView(), test.payload)

			if test.err != nil && !reflect.DeepEqual(err, test.err) {
				t.Errorf("want error %#v, got %#v", test.err, err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("want AddTodo %v, got %v", test.want, got)
			}
		})
	}
}

func TestAddTodo_Execute(t *testing.T) {
	file := "./test.json"
	defer os.Remove(file)
	previousTodoList := []*model.Todo{
		{
			Id:     1,
			Label:  "Test",
			IsDone: false,
		},
	}
	todoList := &model.TodoList{
		Filename: file,
		Items:    previousTodoList,
	}

	payload := "new task"
	view := newDummyView()
	cmd, _ := NewAddTodo(todoList, view, payload)

	cmd.Execute()

	if len(todoList.Items) == 1 {
		t.Errorf("expected to add a todo")
	}

	addedTodo := todoList.Items[1]
	if addedTodo.Label != payload {
		t.Errorf("want label %v, got %v", payload, addedTodo.Label)
	}

	if addedTodo.Id != 2 {
		t.Errorf("want id %d, got %d", 1, addedTodo.Id)
	}

	if addedTodo.IsDone {
		t.Errorf("expected todo not to be done")
	}

	if view.RenderListCalls == 0 {
		t.Errorf("expected view.RenderList to have been called")
	}

	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		t.Errorf("expected file %v to exist", file)
	}
}
