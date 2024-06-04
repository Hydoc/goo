package command

import (
	"testing"

	"github.com/Hydoc/goo/internal/model"
)

func TestListTodos_Execute(t *testing.T) {
	view := newDummyView()
	todoList := &model.TodoList{}
	cmd, err := NewListTodos(todoList, view, "")
	cmd.Execute()

	if err != nil {
		t.Errorf("expected not an error, got %v", err)
	}

	if view.RenderListCalls == 0 {
		t.Errorf("expected a call to view.RenderList")
	}
}
