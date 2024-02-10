package internal

import "fmt"

type TodoList struct {
	items []*Todo
}

func (list *TodoList) Add(todo *Todo) {
	list.items = append(list.items, todo)
}

func (list *TodoList) HasItems() bool {
	return len(list.items) > 0
}

func (list *TodoList) String() string {
	out := ""
	for i, todo := range list.items {
		if todo.IsDone {
			out += fmt.Sprintf("[x] %s", todo.Label)
		} else {
			out += fmt.Sprintf("[ ] %s", todo.Label)
		}

		if i != len(list.items)-1 {
			out += "\r\n"
		}
	}

	return out
}

func NewTodoList() *TodoList {
	return &TodoList{
		items: make([]*Todo, 0),
	}
}
