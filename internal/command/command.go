package command

import (
	"github.com/Hydoc/goo/internal/model"
	"github.com/Hydoc/goo/internal/view"
)

type Command interface {
	Execute()
}

type FabricateCommand func(todoList *model.TodoList, view view.View, payload string) (Command, error)
