package command

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidId           = "%s is an invalid id"
	ErrNoTodoWithId        = "there is no todo with id %d"
	ErrNoTagWithId         = "there is no tag with id %d"
	ErrEmptyTodoNotAllowed = "empty todo is not allowed"
	ErrTagAlreadyExists    = "the tag %s already exists"
	ErrTodoDoesNotHaveTag  = "the todo %d does not have tag with id %d"
	ErrTodoHasNoTags       = "the todo has no tags"
	ErrNoCommandFound      = errors.New("the command could not be found")
)

func errTodoHasNoTags() error {
	return errors.New(ErrTodoHasNoTags)
}

func errEmptyTodoNotAllowed() error { return errors.New(ErrEmptyTodoNotAllowed) }

func errTagAlreadyExists(tagName string) error {
	return errors.New(fmt.Sprintf(ErrTagAlreadyExists, tagName))
}

func errInvalidId(id string) error {
	return errors.New(fmt.Sprintf(ErrInvalidId, id))
}
