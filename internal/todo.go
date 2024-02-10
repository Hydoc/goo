package internal

type Todo struct {
	Id     int
	Label  string
	IsDone bool
}

func NewTodo(label string, id int) *Todo {
	return &Todo{
		Id:     id,
		Label:  label,
		IsDone: false,
	}
}
