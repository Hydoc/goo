package command

import (
	"errors"
	"github.com/Hydoc/goo/internal/model"
	"os"
	"testing"
)

func TestClear_Execute(t *testing.T) {
	file := "./test.json"
	defer os.Remove(file)
	todoList := &model.TodoList{Filename: file, Items: []*model.Todo{
		{
			Id:     1,
			Label:  "should be removed",
			IsDone: false,
		},
		{
			Id:     2,
			Label:  "should also be removed",
			IsDone: true,
		},
	}}
	view := newDummyView()
	cmd, _ := NewClear(todoList, view, "")
	cmd.Execute()

	if view.RenderListCalls == 0 {
		t.Errorf("expected view.RenderList to have been called")
	}

	if todoList.HasItems() {
		t.Error("expected todo list to be cleared")
	}

	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		t.Errorf("expected file %v to exist", file)
	}
}
