package command

import (
	"errors"
	"fmt"
)

const (
	ErrInvalidId           = "%s is an invalid id"
	ErrNoTodoWithId        = "there is no todo with id %d"
	ErrNoTagWithId         = "there is no tag with id %d"
	ErrEmptyTodoNotAllowed = "empty todo is not allowed"
	ErrTagAlreadyExists    = "the tag %s already exists"
)

func errEmptyTodoNotAllowed() error { return errors.New(ErrEmptyTodoNotAllowed) }

func errTagAlreadyExists(tagName string) error {
	return errors.New(fmt.Sprintf(ErrTagAlreadyExists, tagName))
}

func errInvalidId(id string) error {
	return errors.New(fmt.Sprintf(ErrInvalidId, id))
}
