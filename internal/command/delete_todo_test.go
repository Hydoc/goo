package command

import (
	"errors"
	"github.com/Hydoc/goo/internal"
	"reflect"
	"testing"
)

func TestDeleteTodo_newDeleteTodo(t *testing.T) {
	tests := []struct {
		name     string
		todoList *internal.TodoList
		payload  string
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
			payload: "1",
			err:     nil,
			want: &DeleteTodo{
				previousTodoListItems: nil,
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
			name: "not create due to invalid payload",
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
			payload: "1ab",
			err:     errors.New(deleteTodoUsage),
			want:    nil,
		},
		{
			name: "not create due to missing payload",
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
			payload: "",
			err:     errors.New(deleteTodoUsage),
			want:    nil,
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
			payload: "56",
			err:     errors.New(nothingToDelete),
			want:    nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := newDeleteTodo(test.todoList, test.payload)

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
	previousTodoListItems := []*internal.Todo{
		{
			Id:     1,
			Label:  "",
			IsDone: false,
		},
	}
	todoList := &internal.TodoList{
		Filename: "",
		Items:    previousTodoListItems,
	}

	cmd, _ := newDeleteTodo(todoList, "1")
	cmd.Execute()

	if todoList.Has(1) {
		t.Errorf("expected to delete the item with id %d", cmd.idToDelete)
	}

	if !reflect.DeepEqual(cmd.previousTodoListItems, previousTodoListItems) {
		t.Errorf("want previous todo list items in cmd %v, got %v", previousTodoListItems, cmd.previousTodoListItems)
	}
}

func TestDeleteTodo_Undo(t *testing.T) {
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

	cmd, _ := newDeleteTodo(todoList, "1")
	cmd.Execute()
	cmd.Undo()

	if !todoList.Has(1) {
		t.Errorf("expected to undo the cmd, todo with id %d is not set", 1)
	}
}
