package command

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidId           = "%s is an invalid id"
	ErrNoTodoWithId        = "there is no todo with id %d"
	ErrNoTagWithId         = "there is no tag with id %d"
	ErrEmptyTodoNotAllowed = errors.New("empty todo is not allowed")
	ErrTagAlreadyExists    = "the tag %s already exists"
	ErrTodoDoesNotHaveTag  = "the todo %d does not have tag with id %d"
	ErrTodoAlreadyHasTag   = "the todo %d already has the tag with id %d"
	ErrTodoHasNoTags       = errors.New("the todo has no tags")
	ErrCommandNotFound     = errors.New("the command could not be found")
)

func errTagAlreadyExists(tagName string) error {
	return errors.New(fmt.Sprintf(ErrTagAlreadyExists, tagName))
}

func errInvalidId(id string) error {
	return errors.New(fmt.Sprintf(ErrInvalidId, id))
}

func errTodoAlreadyHasTag(todoId, tagId int) error {
	return errors.New(fmt.Sprintf(ErrTodoAlreadyHasTag, todoId, tagId))
}
