package command

import (
	"errors"
	"fmt"
	"github.com/Hydoc/goo/internal/model"
	"os"
	"reflect"
	"testing"
)

func TestNewSwap(t *testing.T) {
	tests := []struct {
		name     string
		payload  string
		err      error
		todoList *model.TodoList
		want     Command
	}{
		{
			name:    "create normally",
			payload: "1 2",
			err:     nil,
			todoList: &model.TodoList{
				Filename: "",
				Items: []*model.Todo{
					model.NewTodo("Hello", 1),
					model.NewTodo("World", 2),
				},
			},
			want: &Swap{
				todoList: &model.TodoList{
					Filename: "",
					Items: []*model.Todo{
						model.NewTodo("Hello", 1),
						model.NewTodo("World", 2),
					},
				},
				view:     newDummyView(),
				firstId:  1,
				secondId: 2,
			},
		},
		{
			name:     "not create when arguments smaller than 2",
			payload:  "1",
			err:      fmt.Errorf("can not swap, need two ids separated by space"),
			todoList: &model.TodoList{},
			want:     nil,
		},
		{
			name:     "not create when arguments greater than 2",
			payload:  "1 2 3",
			err:      fmt.Errorf("can not swap, need two ids separated by space"),
			todoList: &model.TodoList{},
			want:     nil,
		},
		{
			name:     "not create when first id is invalid",
			payload:  "1a 2",
			err:      errors.New("1a is an invalid id"),
			todoList: &model.TodoList{},
			want:     nil,
		},
		{
			name:    "not create when second id is invalid",
			payload: "1 2a",
			err:     errors.New("2a is an invalid id"),
			todoList: &model.TodoList{
				Items: []*model.Todo{
					model.NewTodo("Test", 1),
				},
			},
			want: nil,
		},
		{
			name:     "not create when todo list does not have todo with first id",
			payload:  "56 2",
			err:      errors.New("there is no todo with id 56"),
			todoList: &model.TodoList{},
			want:     nil,
		},
		{
			name:    "not create when todo list does not have todo with second id",
			payload: "1 56",
			err:     errors.New("there is no todo with id 56"),
			todoList: &model.TodoList{
				Items: []*model.Todo{
					model.NewTodo("Test", 1),
				},
			},
			want: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := NewSwap(test.todoList, newDummyView(), test.payload)

			if err != nil && !reflect.DeepEqual(err, test.err) {
				t.Errorf("want error %v, got %v", test.err, err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}

func TestSwap_Execute(t *testing.T) {
	file := "./test.json"
	defer os.Remove(file)
	todoList := &model.TodoList{
		Filename: file,
		Items: []*model.Todo{
			{
				Id:     1,
				Label:  "Hello",
				IsDone: false,
			},
			{
				Id:     2,
				Label:  "World",
				IsDone: false,
			},
		},
	}

	view := newDummyView()
	cmd, _ := NewSwap(todoList, view, "1 2")
	cmd.Execute()

	if view.RenderListCalls == 0 {
		t.Errorf("expected view.RenderList to have been called")
	}

	if todoList.Items[0].Label != "World" {
		t.Error("expected todo at index 0 to be World")
	}

	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		t.Errorf("expected file %v to exist", file)
	}
}
