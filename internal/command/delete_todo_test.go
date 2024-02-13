package command

import (
	"errors"
	"github.com/Hydoc/goo/internal"
	"reflect"
	"testing"
)

func TestNewDeleteTodo(t *testing.T) {
	tests := []struct {
		name     string
		todoList *internal.TodoList
		id       int
		err      error
		want     *DeleteTodo
	}{
		{
			name: "create normally",
			todoList: &internal.TodoList{
				Filename: "",
				Items: []*internal.Todo{
					{
						Id:     1,
						Label:  "Test",
						IsDone: false,
					},
				},
			},
			id:  1,
			err: nil,
			want: &DeleteTodo{
				todoList: &internal.TodoList{
					Filename: "",
					Items: []*internal.Todo{
						{
							Id:     1,
							Label:  "Test",
							IsDone: false,
						},
					},
				},
				idToDelete: 1,
			},
		},
		{
			name: "not create when todo list does not have id in payload",
			todoList: &internal.TodoList{
				Filename: "",
				Items: []*internal.Todo{
					{
						Id:     1,
						Label:  "Test",
						IsDone: false,
					},
				},
			},
			id:   56,
			err:  errors.New("there is no todo with id 56"),
			want: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := newDeleteTodo(test.todoList, test.id)

			if err != nil && !reflect.DeepEqual(test.err, err) {
				t.Errorf("want error %v, got %v", test.err, err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}

func TestDeleteTodo_Execute(t *testing.T) {
	todoList := &internal.TodoList{
		Filename: "",
		Items: []*internal.Todo{
			{
				Id:     1,
				Label:  "",
				IsDone: false,
			},
		},
	}

	cmd, _ := newDeleteTodo(todoList, 1)
	cmd.Execute()

	if todoList.Has(1) {
		t.Errorf("expected to delete the item with id %d", cmd.idToDelete)
	}
}
