package command

import (
	"github.com/Hydoc/goo/internal/model"
	"reflect"
	"testing"
)

func TestNewListTagsOnTodo(t *testing.T) {
	tests := []struct {
		name     string
		payload  string
		todoList *model.TodoList
		want     Command
		err      error
	}{
		{
			name:    "create normally",
			payload: "1",
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
					{
						Id:   1,
						Name: "blub",
					},
				},
			},
			want: &ListTagsOnTodo{
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
						{
							Id:   1,
							Name: "blub",
						},
					},
				},
				view:     newDummyView(),
				idOfTodo: 1,
			},
			err: nil,
		},
		{
			name:     "not create due to invalid id",
			payload:  "1ab",
			todoList: &model.TodoList{},
			want:     nil,
			err:      errInvalidId("1ab"),
		},
		{
			name:    "not create due to todo not in list",
			payload: "1",
			todoList: &model.TodoList{
				Filename: "",
				Items:    make([]*model.Todo, 0),
				TagList:  make([]*model.Tag, 0),
			},
			want: nil,
			err:  errNoTodoWithId(1),
		},
		{
			name:    "not create due to todo has no tags",
			payload: "1",
			todoList: &model.TodoList{
				Filename: "",
				Items: []*model.Todo{
					{
						Id:     1,
						Label:  "Test",
						IsDone: false,
						Tags:   make([]model.TagId, 0),
					},
				},
				TagList: make([]*model.Tag, 0),
			},
			want: nil,
			err:  ErrTodoHasNoTags,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := NewListTagsOnTodo(test.todoList, newDummyView(), test.payload)

			if err != nil && !reflect.DeepEqual(err, test.err) {
				t.Errorf("want error %v, got %v", test.err, err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}

func TestListTagsOnTodo_Execute(t *testing.T) {
	view := newDummyView()
	todoList := &model.TodoList{
		Filename: "",
		Items: []*model.Todo{
			{
				Id:     1,
				Label:  "Bla",
				IsDone: false,
				Tags:   []model.TagId{1},
			},
		},
		TagList: []*model.Tag{
			{
				Id:   1,
				Name: "my test tag",
			},
		},
	}
	cmd, _ := NewListTagsOnTodo(todoList, view, "1")
	cmd.Execute()

	if view.RenderTagsCalls == 0 {
		t.Error("expected view.RenderTags to have been called")
	}
}
