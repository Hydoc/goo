package internal

import "testing"

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
