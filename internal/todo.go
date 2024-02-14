package internal

type Todo struct {
	Id     int    `json:"id"`
	Label  string `json:"label"`
	IsDone bool   `json:"isDone"`
}

func (t *Todo) DoneAsString() string {
	if t.IsDone {
		return "✓"
	}
	return "○"
}

func NewTodo(label string, id int) *Todo {
	return &Todo{
		Id:     id,
		Label:  label,
		IsDone: false,
	}
}
