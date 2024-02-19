package command

import "errors"

const (
	ErrInvalidId       = "%s is an invalid id"
	ErrNoTodoWithId    = "there is no todo with id %d"
	ErrNoTagWithId     = "there is no tag with id %d"
	ErrEmptyNotAllowed = "empty todo is not allowed"
)

func errEmptyNotAllowed() error { return errors.New(ErrEmptyNotAllowed) }
