package internal

type Todo struct {
	Label  string
	IsDone bool
}

func NewTodo(label string) *Todo {
	return &Todo{
		Label:  label,
		IsDone: false,
	}
}
