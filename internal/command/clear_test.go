package command

import (
	"github.com/Hydoc/goo/internal"
	"testing"
)

func TestClear_Execute(t *testing.T) {
	todoList := &internal.TodoList{Items: []*internal.Todo{
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
	cmd := newClear(todoList)
	cmd.Execute()

	if todoList.HasItems() {
		t.Error("expected todo list to be cleared")
	}
}
