package command

import (
	"github.com/Hydoc/goo/internal/model"
	"reflect"
	"testing"
)

func TestNewListTags(t *testing.T) {
	todoList := &model.TodoList{}
	view := newDummyView()
	want := &ListTags{
		todoList: todoList,
		view:     view,
	}

	got, err := NewListTags(todoList, view, "")

	if err != nil {
		t.Errorf("want err to be nil, got %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestListTags_Execute(t *testing.T) {
	todoList := &model.TodoList{}
	view := newDummyView()
	cmd, _ := NewListTags(todoList, view, "")
	cmd.Execute()

	if view.RenderTagsCalls == 0 {
		t.Error("expected view.RenderTags to have been called")
	}
}
