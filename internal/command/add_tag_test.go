package command

import (
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/Hydoc/goo/internal/model"
)

func TestNewAddTag(t *testing.T) {
	tests := []struct {
		name     string
		payload  string
		err      error
		todoList *model.TodoList
		want     Command
	}{
		{
			name:    "add normally",
			payload: "HELLO   ",
			err:     nil,
			todoList: &model.TodoList{
				Filename: "",
				Items:    make([]*model.Todo, 0),
				TagList:  make([]*model.Tag, 0),
			},
			want: &AddTag{
				todoList: &model.TodoList{
					Filename: "",
					Items:    make([]*model.Todo, 0),
					TagList:  make([]*model.Tag, 0),
				},
				view:            newDummyView(),
				tagNameToCreate: "hello",
			},
		},
		{
			name:    "not add due to existing tag",
			payload: "HELLO ",
			err:     errTagAlreadyExists("hello"),
			todoList: &model.TodoList{
				Filename: "",
				Items:    nil,
				TagList: []*model.Tag{
					{
						Id:   1,
						Name: "hello",
					},
				},
			},
			want: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := NewAddTag(test.todoList, newDummyView(), test.payload)

			if err != nil && !reflect.DeepEqual(err, test.err) {
				t.Errorf("want err %v, got %v", test.err, err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}

func TestAddTag_Execute(t *testing.T) {
	file := "./test.json"
	defer os.Remove(file)
	todoList := &model.TodoList{
		Filename: file,
		Items:    make([]*model.Todo, 0),
		TagList:  make([]*model.Tag, 0),
	}

	payload := "My COOOL tag"
	view := newDummyView()
	cmd, _ := NewAddTag(todoList, view, payload)

	cmd.Execute()

	if !todoList.HasTag(1) {
		t.Error("expected todo list to have tag with id 1")
	}

	if !todoList.HasTagWith("my coool tag") {
		t.Error("expected todo list to have tag my coool tag")
	}

	if view.RenderTagsCalls == 0 {
		t.Error("expected view.RenderTags to have been called")
	}

	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		t.Errorf("expected file %v to exist", file)
	}
}
