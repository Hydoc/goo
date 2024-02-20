package model

import (
	"fmt"
	"slices"
)

type Todo struct {
	Id     int     `json:"id"`
	Label  string  `json:"label"`
	IsDone bool    `json:"isDone"`
	Tags   []TagId `json:"tags"`
}

func (t *Todo) AddTag(tagId TagId) {
	if !slices.Contains(t.Tags, tagId) {
		t.Tags = append(t.Tags, tagId)
	}
}

func (t *Todo) HasTag(id TagId) bool {
	for _, tagId := range t.Tags {
		if tagId == id {
			return true
		}
	}
	return false
}

func (t *Todo) RemoveTag(tagId TagId) {
	indexOfTag := slices.Index(t.Tags, tagId)
	t.Tags = append(t.Tags[:indexOfTag], t.Tags[indexOfTag+1:]...)
}

func (t *Todo) LabelAsString() string {
	if len(t.Tags) > 0 {
		return fmt.Sprintf("%s 🏷", t.Label)
	}
	return t.Label
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
		make([]TagId, 0),
	}
}
