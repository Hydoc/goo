package internal

type Todo struct {
	Id     int    `json:"id"`
	Label  string `json:"label"`
	IsDone bool   `json:"isDone"`
}

func NewTodo(label string, id int) *Todo {
	return &Todo{
		Id:     id,
		Label:  label,
		IsDone: false,
	}
}
