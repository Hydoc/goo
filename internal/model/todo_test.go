package model

import (
	"reflect"
	"testing"
)

func TestNewTodo(t *testing.T) {
	todo := NewTodo("Test", 1)

	if todo.Id != 1 {
		t.Errorf("expected id to be 1, got %d", todo.Id)
	}

	if todo.Label != "Test" {
		t.Errorf("expected label to be Test, got %s", todo.Label)
	}

	if todo.IsDone {
		t.Error("expected todo not to be done")
	}
}

func TestTodo_DoneAsString(t *testing.T) {
	tests := []struct {
		name string
		want string
		todo *Todo
	}{
		{
			name: "done = false",
			want: "○",
			todo: &Todo{
				Id:     1,
				Label:  "Test",
				IsDone: false,
			},
		},
		{
			name: "done = true",
			want: "✓",
			todo: &Todo{
				Id:     1,
				Label:  "Test",
				IsDone: true,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.todo.DoneAsString()
			if got != test.want {
				t.Errorf("want %s, got %s", test.want, got)
			}
		})
	}
}

func TestTodo_compare(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		todo  *Todo
		other *Todo
	}{
		{
			name:  "a is done, b is not",
			want:  -1,
			todo:  &Todo{1, "Test", true, make([]TagId, 0)},
			other: &Todo{2, "Test", false, make([]TagId, 0)},
		},
		{
			name:  "both are done",
			want:  1,
			todo:  &Todo{2, "Test", true, make([]TagId, 0)},
			other: &Todo{1, "Test", true, make([]TagId, 0)},
		},
		{
			name:  "b is done, a is not",
			want:  -1,
			todo:  &Todo{1, "Test", false, make([]TagId, 0)},
			other: &Todo{2, "Test", true, make([]TagId, 0)},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.todo.compare(test.other)

			if got != test.want {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}

func TestTodo_AddTag(t *testing.T) {
	tests := []struct {
		name     string
		want     []TagId
		tagToAdd int
		todo     *Todo
	}{
		{
			name:     "add tag",
			want:     []TagId{1},
			tagToAdd: 1,
			todo: &Todo{
				Id:     1,
				Label:  "Test",
				IsDone: false,
				Tags:   make([]TagId, 0),
			},
		},
		{
			name:     "not add tag when tag already in todo",
			want:     []TagId{1},
			tagToAdd: 1,
			todo: &Todo{
				Id:     1,
				Label:  "Test",
				IsDone: false,
				Tags:   []TagId{1},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.todo.AddTag(TagId(test.tagToAdd))
			got := test.todo.Tags

			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}
