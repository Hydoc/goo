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
