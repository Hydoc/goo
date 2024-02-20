package model

import (
	"encoding/json"
	"os"
	"slices"
)

type TodoList struct {
	Filename string  `json:"-"`
	Items    []*Todo `json:"items"`
	TagList  []*Tag  `json:"tagList"`
}

func (list *TodoList) Add(todo *Todo) {
	list.Items = append(list.Items, todo)
}

func (list *TodoList) AddTag(tag *Tag) {
	list.TagList = append(list.TagList, tag)
}

func (list *TodoList) Swap(firstId, secondId int) {
	firstTodo := list.Find(firstId)
	secondTodo := list.Find(secondId)
	tmp := firstTodo.Label
	firstTodo.Label = secondTodo.Label
	secondTodo.Label = tmp
}

func (list *TodoList) Find(id int) *Todo {
	for _, todo := range list.Items {
		if todo.Id == id {
			return todo
		}
	}
	return nil
}

func (list *TodoList) FindTag(id int) *Tag {
	for _, tag := range list.TagList {
		if tag.Id == TagId(id) {
			return tag
		}
	}
	return nil
}

func (list *TodoList) TagTodo(id int, tagId TagId) {
	todo := list.Find(id)
	todo.AddTag(tagId)
}

func (list *TodoList) LenOfLongestTodo() int {
	if !list.HasItems() {
		return 0
	}

	current := len(list.Items[0].LabelAsString())
	for _, todo := range list.Items {
		if len(todo.LabelAsString()) > current {
			current = len(todo.LabelAsString())
		}
	}
	return current
}

func (list *TodoList) LenOfLongestTag() int {
	if len(list.TagList) == 0 {
		return 0
	}

	current := len(list.TagList[0].Name)
	for _, tag := range list.TagList {
		if len(tag.Name) > current {
			current = len(tag.Name)
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

func (list *TodoList) HasTag(id TagId) bool {
	for _, tag := range list.TagList {
		if tag.Id == id {
			return true
		}
	}
	return false
}

func (list *TodoList) HasTagWith(name string) bool {
	for _, tag := range list.TagList {
		if tag.Name == name {
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

func (list *TodoList) SortedByIdAndState() *TodoList {
	itemsCopy := make([]*Todo, len(list.Items))
	copy(itemsCopy, list.Items)
	slices.SortFunc(itemsCopy, func(a, b *Todo) int {
		return a.compare(b)
	})
	return &TodoList{
		Filename: list.Filename,
		Items:    itemsCopy,
	}
}

func (list *TodoList) NextTodoId() int {
	if len(list.Items) == 0 {
		return 1
	}

	return list.Items[len(list.Items)-1].Id + 1
}

func (list *TodoList) NextTagId() TagId {
	if len(list.TagList) == 0 {
		return 1
	}

	return list.TagList[len(list.TagList)-1].Id + 1
}

func (list *TodoList) SaveToFile() {
	encoded, _ := json.Marshal(list)
	_ = os.WriteFile(list.Filename, encoded, 0644)
}

func (list *TodoList) Clear() {
	list.Items = make([]*Todo, 0)
}

func NewTodoListFromFile(filename string) (*TodoList, error) {
	var todoList *TodoList
	jsonBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonBytes, &todoList)
	if err != nil {
		return nil, err
	}

	if todoList.TagList == nil {
		todoList.TagList = make([]*Tag, 0)
	}

	return &TodoList{
		Filename: filename,
		Items:    todoList.Items,
		TagList:  todoList.TagList,
	}, nil
}
