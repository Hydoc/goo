package internal

import (
	"encoding/json"
	"os"
	"slices"
)

type TodoList struct {
	Filename string
	Items    []*Todo
}

func (list *TodoList) Add(todo *Todo) {
	list.Items = append(list.Items, todo)
}

func (list *TodoList) Find(id int) *Todo {
	for _, todo := range list.Items {
		if todo.Id == id {
			return todo
		}
	}
	return nil
}

func (list *TodoList) LenOfLongestTodo() int {
	if !list.HasItems() {
		return 0
	}

	current := len(list.Items[0].Label)
	for _, todo := range list.Items {
		if len(todo.Label) > current {
			current = len(todo.Label)
		}
	}
	return current
}

func (list *TodoList) Edit(id int, label string) {
	for _, todo := range list.Items {
		if todo.Id == id {
			todo.Label = label
		}
	}
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

func (list *TodoList) SaveToFile() {
	encoded, _ := json.Marshal(list.Items)
	_ = os.WriteFile(list.Filename, encoded, 0644)
}

func NewTodoListFromFile(filename string) (*TodoList, error) {
	var items []*Todo
	jsonBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonBytes, &items)
	if err != nil {
		return nil, err
	}

	return &TodoList{
		Filename: filename,
		Items:    items,
	}, nil
}
