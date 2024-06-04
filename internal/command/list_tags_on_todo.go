package command

import (
	"strconv"

	"github.com/Hydoc/goo/internal/model"
	"github.com/Hydoc/goo/internal/view"
)

type ListTagsOnTodo struct {
	todoList *model.TodoList
	view     view.View
	idOfTodo int
}

func (cmd *ListTagsOnTodo) Execute() {
	cmd.view.RenderTags(cmd.todoList.TagsForTodo(cmd.idOfTodo))
}

func NewListTagsOnTodo(todoList *model.TodoList, view view.View, payload string) (Command, error) {
	idOfTodo, err := strconv.Atoi(payload)
	if err != nil {
		return nil, errInvalidId(payload)
	}

	todo := todoList.Find(idOfTodo)
	if todo == nil {
		return nil, errNoTodoWithId(idOfTodo)
	}

	if !todo.HasTags() {
		return nil, ErrTodoHasNoTags
	}

	return &ListTagsOnTodo{todoList, view, idOfTodo}, nil
}
