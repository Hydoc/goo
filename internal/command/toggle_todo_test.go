package command

import (
	"errors"
	"github.com/Hydoc/goo/internal/model"
	"os"
	"reflect"
	"testing"
)

func TestNewToggleTodo(t *testing.T) {
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
			want: &ToggleTodo{
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
				idToToggle: 1,
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
		{
			name: "not create when invalid id is passed",
			todoList: &model.TodoList{
				Filename: "",
				Items:    make([]*model.Todo, 0),
			},
			payload: "56a",
			err:     errors.New("56a is an invalid id"),
			want:    nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := NewToggleTodo(test.todoList, newDummyView(), test.payload)

			if err != nil && !reflect.DeepEqual(test.err, err) {
				t.Errorf("want error %v, got %v", test.err, err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("want %#v, got %#v", test.want, got)
			}
		})
	}
}

func TestToggleTodo_Execute(t *testing.T) {
	file := "./test.json"
	defer os.Remove(file)
	todoList := &model.TodoList{
		Filename: file,
		Items: []*model.Todo{
			{
				Id:     1,
				Label:  "",
				IsDone: false,
			},
		},
	}

	view := newDummyView()
	cmd, _ := NewToggleTodo(todoList, view, "1")
	cmd.Execute()

	if view.RenderListCalls == 0 {
		t.Errorf("expected view.RenderList to have been called")
	}

	if !todoList.Items[0].IsDone {
		t.Error("expected todo at index 0 to be done")
	}

	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		t.Errorf("expected file %v to exist", file)
	}
}
