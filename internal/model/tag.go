package model

import "strings"

type TagId int

type Tag struct {
	Id   TagId  `json:"id"`
	Name string `json:"name"`
}

func NewTag(id TagId, name string) *Tag {
	return &Tag{id, strings.ToLower(strings.TrimSpace(name))}
}
