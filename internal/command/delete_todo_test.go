package command

import (
	"errors"
	"github.com/Hydoc/goo/internal/model"
	"reflect"
	"testing"
)

func TestNewDeleteTodo(t *testing.T) {
	tests := []struct {
		name     string
		todoList *model.TodoList
		payload  string
		err      error
		want     Command
	}{
		{
			name: "create normally",
			todoList: &model.TodoList{
				Filename: "",
				Items: []*model.Todo{
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
				view: newDummyView(),
				todoList: &model.TodoList{
					Filename: "",
					Items: []*model.Todo{
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
			todoList: &model.TodoList{
				Filename: "",
				Items: []*model.Todo{
					{
						Id:     1,
						Label:  "Test",
						IsDone: false,
					},
				},
			},
			payload: "56",
			err:     errors.New("there is no todo with id 56"),
			want:    nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := NewDeleteTodo(test.todoList, newDummyView(), test.payload)

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
	todoList := &model.TodoList{
		Filename: "",
		Items: []*model.Todo{
			{
				Id:     1,
				Label:  "",
				IsDone: false,
			},
		},
	}

	cmd, _ := NewDeleteTodo(todoList, newDummyView(), "1")
	cmd.Execute()

	if todoList.Has(1) {
		t.Errorf("expected to delete the item with id %d", cmd.(*DeleteTodo).idToDelete)
	}
}
