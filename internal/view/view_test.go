package view

import (
	"bytes"
	"github.com/Hydoc/goo/internal/model"
	"testing"
)

func TestStdoutView_RenderLine(t *testing.T) {
	want := "Hello World!\n"
	buffer := bytes.NewBuffer(make([]byte, 0))
	v := New(buffer)

	v.RenderLine("Hello World!")

	if want != buffer.String() {
		t.Errorf("want %#v, got %#v", want, buffer.String())
	}
}

func TestStdoutView_RenderList(t *testing.T) {
	tests := []struct {
		name     string
		want     string
		buffer   *bytes.Buffer
		todoList *model.TodoList
	}{
		{
			name:   "without items",
			want:   "ID  TASK      STATUS\n--------------------\n",
			buffer: bytes.NewBuffer(make([]byte, 0)),
			todoList: &model.TodoList{
				Items: make([]*model.Todo, 0),
			},
		},
		{
			name:   "with one item",
			want:   "ID  TASK  STATUS\n----------------\n1   Test    ○\n",
			buffer: bytes.NewBuffer(make([]byte, 0)),
			todoList: &model.TodoList{
				Items: []*model.Todo{
					model.NewTodo("Test", 1),
				},
			},
		},
		{
			name:   "multiple items with one done",
			want:   "ID  TASK              STATUS\n----------------------------\n2   should be first     ○\n3   should be second    ○\n\x1b[90m1   should be last      ✓\x1b[0m\n",
			buffer: bytes.NewBuffer(make([]byte, 0)),
			todoList: &model.TodoList{
				Items: []*model.Todo{
					{
						Id:     1,
						Label:  "should be last",
						IsDone: true,
					},
					model.NewTodo("should be first", 2),
					model.NewTodo("should be second", 3),
				},
			},
		},
		{
			name:   "multiple items with one tag",
			want:   "ID  TASK                  STATUS\n--------------------------------\n1   should be first 🏷       ○\n",
			buffer: bytes.NewBuffer(make([]byte, 0)),
			todoList: &model.TodoList{
				Items: []*model.Todo{
					{1, "should be first", false, []model.TagId{1}},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := New(test.buffer)
			v.RenderList(test.todoList)

			if test.buffer.String() != test.want {
				t.Errorf("want %#v, got %#v", test.want, test.buffer.String())
			}
		})
	}
}

func TestStdoutView_RenderTags(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		buffer  *bytes.Buffer
		tagList []*model.Tag
	}{
		{
			name:   "with multiple tags",
			want:   "ID  TAG        \n---------------\n1   test tag   \n2   another tag\n",
			buffer: bytes.NewBuffer(make([]byte, 0)),
			tagList: []*model.Tag{
				model.NewTag(1, "test tag"),
				model.NewTag(2, "another tag"),
			},
		},
		{
			name:    "without tags",
			want:    "ID   TAG\n--------\n",
			buffer:  bytes.NewBuffer(make([]byte, 0)),
			tagList: make([]*model.Tag, 0),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := New(test.buffer)
			v.RenderTags(test.tagList)

			if test.buffer.String() != test.want {
				t.Errorf("want %#v, got %#v", test.want, test.buffer.String())
			}
		})
	}
}
