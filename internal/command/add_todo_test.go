package command

import (
	"errors"
	"github.com/Hydoc/goo/internal"
	"reflect"
	"testing"
)

func TestNewAddTodo(t *testing.T) {
	todoList := &internal.TodoList{}
	tests := []struct {
		name    string
		payload string
		want    *AddTodo
		err     error
	}{
		{
			name:    "create normally",
			payload: "test",
			want:    &AddTodo{todoList: todoList, todoToAdd: "test"},
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
			got, err := newAddTodo(todoList, test.payload)

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
	previousTodoList := []*internal.Todo{
		{
			Id:     1,
			Label:  "Test",
			IsDone: false,
		},
	}
	todoList := &internal.TodoList{
		Filename: "",
		Items:    previousTodoList,
	}

	payload := "new task"
	cmd, _ := newAddTodo(todoList, payload)

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
}
