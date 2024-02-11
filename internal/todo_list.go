package internal

import (
	"fmt"
	"slices"
)

type TodoList struct {
	Items []*Todo
}

func (list *TodoList) Add(todo *Todo) {
	list.Items = append(list.Items, todo)
}

func (list *TodoList) Has(id int) bool {
	for _, todo := range list.Items {
		if todo.Id == id {
			return true
		}
	}
	return false
}

func (list *TodoList) Remove(id int) {
	i := slices.IndexFunc(list.Items, func(todo *Todo) bool {
		return todo.Id == id
	})
	if i != -1 {
		list.Items = append(list.Items[:i], list.Items[i+1:]...)
	}
}

func (list *TodoList) HasItems() bool {
	return len(list.Items) > 0
}

func (list *TodoList) Toggle(id int) {
	for _, todo := range list.Items {
		if todo.Id == id {
			todo.IsDone = !todo.IsDone
			return
		}
	}
}

func (list *TodoList) NextId() int {
	if len(list.Items) == 0 {
		return 1
	}

	return list.Items[len(list.Items)-1].Id + 1
}

func (list *TodoList) String() string {
	out := ""
	for i, todo := range list.Items {
		if todo.IsDone {
			out += fmt.Sprintf("%d [x] %s", todo.Id, todo.Label)
		} else {
			out += fmt.Sprintf("%d [ ] %s", todo.Id, todo.Label)
		}

		if i != len(list.Items)-1 {
			out += "\r\n"
		}
	}

	return out
}

func NewTodoList() *TodoList {
	return &TodoList{
		Items: make([]*Todo, 0),
	}
}
