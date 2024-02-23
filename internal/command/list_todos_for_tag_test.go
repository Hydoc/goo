package command

import (
	"github.com/Hydoc/goo/internal/model"
	"reflect"
	"testing"
)

func TestNewListTodosForTag(t *testing.T) {
	tests := []struct {
		name     string
		want     Command
		err      error
		payload  string
		todoList *model.TodoList
	}{
		{
			name: "create normally",
			want: &ListTodosForTag{
				todoList: &model.TodoList{
					Filename: "",
					Items: []*model.Todo{
						{
							Id:     1,
							Label:  "Test todo",
							IsDone: false,
							Tags:   []model.TagId{1},
						},
					},
					TagList: []*model.Tag{
						model.NewTag(1, "test tag"),
					},
				},
				view:    newDummyView(),
				idOfTag: 1,
			},
			err:     nil,
			payload: "1",
			todoList: &model.TodoList{
				Filename: "",
				Items: []*model.Todo{
					{
						Id:     1,
						Label:  "Test todo",
						IsDone: false,
						Tags:   []model.TagId{1},
					},
				},
				TagList: []*model.Tag{
					model.NewTag(1, "test tag"),
				},
			},
		},
		{
			name:     "not create due to invalid",
			want:     nil,
			err:      errInvalidId("1ab"),
			payload:  "1ab",
			todoList: &model.TodoList{},
		},
		{
			name:    "not create due to missing tag",
			want:    nil,
			err:     errNoTagWithId(1),
			payload: "1",
			todoList: &model.TodoList{
				Filename: "",
				Items:    make([]*model.Todo, 0),
				TagList:  make([]*model.Tag, 0),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := NewListTodosForTag(test.todoList, newDummyView(), test.payload)

			if err != nil && !reflect.DeepEqual(err, test.err) {
				t.Errorf("want err %v, got %v", test.err, err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}

func TestListTodosForTag_Execute(t *testing.T) {
	todoList := &model.TodoList{
		Filename: "test.json",
		Items: []*model.Todo{
			model.NewTodo("irrelevant #1", 1),
			{123, "i should be in", false, []model.TagId{1}},
			model.NewTodo("irrelevant #2", 1),
			{789, "i also should be in", false, []model.TagId{1}},
		},
		TagList: []*model.Tag{
			model.NewTag(1, "test tag"),
		},
	}

	view := newDummyView()
	cmd, err := NewListTodosForTag(todoList, view, "1")
	cmd.Execute()

	if err != nil {
		t.Errorf("expected no err, got %v", err)
	}

	if view.RenderListCalls == 0 {
		t.Error("expected a view.RenderList call")
	}
}
