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
