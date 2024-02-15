package command

import "github.com/Hydoc/goo/internal"

type Factory struct{}

func (f *Factory) Fabricate(todoList *internal.TodoList, toggle int, add bool, doDelete int, edit bool, doClear bool, args string) (Command, error) {
	switch {
	case edit:
		return newEditTodo(todoList, args)
	case doDelete > 0:
		return newDeleteTodo(todoList, doDelete)
	case toggle > 0:
		return newToggleTodo(todoList, toggle)
	case add:
		return newAddTodo(todoList, args)
	case doClear:
		return newClear(todoList), nil
	default:
		return nil, nil
	}
}

func NewFactory() *Factory {
	return &Factory{}
}
