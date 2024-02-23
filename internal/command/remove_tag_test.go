package command

import (
	"errors"
	"github.com/Hydoc/goo/internal/model"
	"os"
	"reflect"
	"testing"
)

func TestNewRemoveTag(t *testing.T) {
	tests := []struct {
		name     string
		want     Command
		err      error
		payload  string
		todoList *model.TodoList
	}{
		{
			name: "create normally",
			want: &RemoveTag{
				todoList: &model.TodoList{
					Filename: "",
					Items:    make([]*model.Todo, 0),
					TagList: []*model.Tag{
						model.NewTag(1, "test tag"),
					},
				},
				view:          newDummyView(),
				tagIdToRemove: 1,
			},
			err:     nil,
			payload: "1",
			todoList: &model.TodoList{
				Filename: "",
				Items:    make([]*model.Todo, 0),
				TagList: []*model.Tag{
					model.NewTag(1, "test tag"),
				},
			},
		},
		{
			name:     "not create due to invalid id",
			want:     nil,
			err:      errInvalidId("1ab"),
			payload:  "1ab",
			todoList: &model.TodoList{},
		},
		{
			name:    "not create because of missing tag in todolist",
			want:    nil,
			err:     errNoTagWithId(1),
			payload: "1",
			todoList: &model.TodoList{
				Filename: "test.json",
				Items:    make([]*model.Todo, 0),
				TagList:  make([]*model.Tag, 0),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := NewRemoveTag(test.todoList, newDummyView(), test.payload)

			if err != nil && !reflect.DeepEqual(err, test.err) {
				t.Errorf("want error %v, got %v", test.err, err)
			}

			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}

func TestRemoveTag_Execute(t *testing.T) {
	file := "./test.json"
	defer os.Remove(file)

	view := newDummyView()
	todoList := &model.TodoList{
		Filename: file,
		Items: []*model.Todo{
			{
				Id:     1,
				Label:  "Test",
				IsDone: false,
				Tags:   []model.TagId{1},
			},
		},
		TagList: []*model.Tag{
			model.NewTag(1, "test tag"),
		},
	}
	cmd, err := NewRemoveTag(todoList, view, "1")
	cmd.Execute()

	if err != nil {
		t.Errorf("expected not an error, got %v", err)
	}

	if view.RenderTagsCalls == 0 {
		t.Errorf("expected a call to view.RenderTags")
	}

	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		t.Errorf("expected file %v to exist", file)
	}
}
