package command

import (
	"flag"
	"strconv"
	"strings"

	"github.com/Hydoc/goo/internal/model"
	"github.com/Hydoc/goo/internal/view"
)

func Fabricate(todoList *model.TodoList, view view.View, flagArgs []string, list *flag.FlagSet, add *flag.FlagSet, doDelete *flag.FlagSet, toggle *flag.FlagSet, edit *flag.FlagSet, doClear *flag.FlagSet, swap *flag.FlagSet, tags *flag.FlagSet, showTagsOnTodo *int, showTodosForTag *int, tag *flag.FlagSet, tagRm *bool, tagAdd *bool, versionFlag *flag.FlagSet, version string) (Command, error) {
	baseArgs := strings.TrimSpace(strings.Join(flagArgs[1:], " "))
	switch flagArgs[0] {
	case list.Name():
		return NewListTodos(todoList, view, baseArgs)
	case add.Name():
		return NewAddTodo(todoList, view, baseArgs)
	case doDelete.Name():
		return NewDeleteTodo(todoList, view, baseArgs)
	case toggle.Name():
		return NewToggleTodo(todoList, view, baseArgs)
	case edit.Name():
		return NewEditTodo(todoList, view, baseArgs)
	case doClear.Name():
		return NewClear(todoList, view, baseArgs)
	case swap.Name():
		return NewSwap(todoList, view, baseArgs)
	case versionFlag.Name():
		return NewPrintVersion(view, version)
	case tags.Name():
		tags.Parse(flagArgs[1:])
		switch {
		case *showTagsOnTodo > 0:
			args := strconv.Itoa(*showTagsOnTodo)
			return NewListTagsOnTodo(todoList, view, args)
		case *showTodosForTag > 0:
			args := strconv.Itoa(*showTodosForTag)
			return NewListTodosForTag(todoList, view, args)
		default:
			return NewListTags(todoList, view, baseArgs)
		}
	case tag.Name():
		tag.Parse(flagArgs[1:])
		args := strings.TrimSpace(strings.Join(tag.Args(), " "))
		switch {
		case *tagRm && len(tag.Args()) > 1:
			return NewRemoveTagFromTodo(todoList, view, args)
		case *tagRm && len(tag.Args()) == 1:
			return NewRemoveTag(todoList, view, args)
		case *tagAdd:
			return NewAddTag(todoList, view, args)
		default:
			return NewTagTodo(todoList, view, args)
		}
	default:
		return nil, ErrCommandNotFound
	}
}
