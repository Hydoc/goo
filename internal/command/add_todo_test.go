package command

import (
	"errors"
	"github.com/Hydoc/goo/internal"
	"reflect"
	"testing"
)

func TestAddTodo_newAddTodo(t *testing.T) {
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
			err:     errors.New(addTodoUsage),
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
