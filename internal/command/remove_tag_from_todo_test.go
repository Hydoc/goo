package command

import (
	"errors"
	"fmt"
	"github.com/Hydoc/goo/internal/model"
	"os"
	"reflect"
	"testing"
)

func TestNewRemoveTagFromTodo(t *testing.T) {
	tests := []struct {
		name     string
		payload  string
		err      error
		todoList *model.TodoList
		want     Command
	}{
		{
			name:    "create normally",
			payload: "1 1",
			err:     nil,
			todoList: &model.TodoList{
				Filename: "",
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
			},
			want: &RemoveTagFromTodo{
				todoList: &model.TodoList{
					Filename: "",
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
				},
				view:          newDummyView(),
				idOfTodo:      1,
				tagIdToRemove: 1,
			},
		},
		{
			name:     "not create due to arguments > 2",
			payload:  "1 2 3",
			err:      fmt.Errorf("can not untag todo, need id of tag as first argument, the second has to be the id of the todo"),
			todoList: &model.TodoList{},
			want:     nil,
		},
		{
			name:     "not create due to arguments < 2",
			payload:  "1",
			err:      fmt.Errorf("can not untag todo, need id of tag as first argument, the second has to be the id of the todo"),
			todoList: &model.TodoList{},
			want:     nil,
		},
		{
			name:     "not create due to invalid tag id",
			payload:  "ab 2",
			err:      errInvalidId("ab"),
			todoList: &model.TodoList{},
			want:     nil,
		},
		{
			name:     "not create due to invalid todo id",
			payload:  "1 ab",
			err:      errInvalidId("ab"),
			todoList: &model.TodoList{},
			want:     nil,
		},
		{
			name:     "not create due to missing todo in todo list",
			payload:  "1 13",
			err:      errNoTodoWithId(13),
			todoList: &model.TodoList{},
			want:     nil,
		},
		{
			name:    "not create due to missing tag in tag list",
			payload: "1 13",
			err:     errNoTagWithId(1),
			todoList: &model.TodoList{
				Items: []*model.Todo{
					model.NewTodo("Test", 13),
				},
			},
			want: nil,
		},
		{
			name:    "not create when todo does not have tag",
			payload: "1 13",
			err:     fmt.Errorf(ErrTodoDoesNotHaveTag, 13, 1),
			todoList: &model.TodoList{
				Items: []*model.Todo{
					model.NewTodo("Test", 13),
				},
				TagList: []*model.Tag{
					model.NewTag(1, "test"),
				},
			},
			want: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := NewRemoveTagFromTodo(test.todoList, newDummyView(), test.payload)

			if err != nil && !reflect.DeepEqual(err, test.err) {
				t.Errorf("want err %v, got %v", test.err, err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}

func TestRemoveTagFromTodo_Execute(t *testing.T) {
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
	cmd, err := NewRemoveTagFromTodo(todoList, view, "1 1")
	cmd.Execute()

	if err != nil {
		t.Errorf("expected not an error, got %v", err)
	}

	if todoList.Items[0].HasTag(1) {
		t.Error("want todo at index 0 not to have tag 1")
	}

	if view.RenderListCalls == 0 {
		t.Errorf("expected a call to view.RenderList")
	}

	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		t.Errorf("expected file %v to exist", file)
	}
}
