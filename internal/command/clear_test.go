package command

import (
	"github.com/Hydoc/goo/internal/model"
	"testing"
)

func TestClear_Execute(t *testing.T) {
	todoList := &model.TodoList{Items: []*model.Todo{
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
	cmd, _ := NewClear(todoList, newDummyView(), "")
	cmd.Execute()

	if todoList.HasItems() {
		t.Error("expected todo list to be cleared")
	}
}
