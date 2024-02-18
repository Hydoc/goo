package command

import (
	"errors"
	"github.com/Hydoc/goo/internal/model"
	"reflect"
	"testing"
)

func TestNewEditTodo(t *testing.T) {
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
			payload: "1 Bla {} bla",
			err:     nil,
			want: &EditTodo{
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
				idToEdit: 1,
				newLabel: "Bla {} bla",
			},
		},
		{
			name: "not create due to invalid payload (missing id)",
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
			payload: "Bla {} bla",
			err:     errors.New("can not edit todo, id is missing"),
			want:    nil,
		},
		{
			name: "not create due to invalid payload (missing label)",
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
			err:     errors.New("empty todo is not allowed"),
			want:    nil,
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
			payload: "53",
			err:     errors.New("there is no todo with id 53"),
			want:    nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := NewEditTodo(test.todoList, newDummyView(), test.payload)

			if err != nil && !reflect.DeepEqual(test.err, err) {
				t.Errorf("want error %v, got %v", test.err, err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("want %#v, got %#v", test.want, got)
			}
		})
	}
}

func TestEditTodo_Execute(t *testing.T) {
	todoList := &model.TodoList{
		Filename: "",
		Items: []*model.Todo{
			{
				Id:     1,
				Label:  "Test",
				IsDone: false,
			},
		},
	}
	payload := "1 Bla {} bla"
	wantLabel := "Bla Test bla"

	cmd, _ := NewEditTodo(todoList, newDummyView(), payload)
	cmd.Execute()

	editedItem := todoList.Items[0]
	if editedItem.Label != wantLabel {
		t.Errorf("want label %v, got %v", wantLabel, editedItem.Label)
	}
}
