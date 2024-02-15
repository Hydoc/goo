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

func (t *Todo) compare(other *Todo) int {
	switch {
	case t.IsDone && other.IsDone:
		return t.Id - other.Id
	case other.IsDone:
		return -1
	default:
		return t.Id - other.Id
	}
}

func NewTodo(label string, id int) *Todo {
	return &Todo{
		id,
		label,
		false,
	}
}
