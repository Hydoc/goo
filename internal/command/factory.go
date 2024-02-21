package command

import (
	"flag"
	"strconv"
	"strings"
)

func Fabricate(flagArgs []string, list *flag.FlagSet, add *flag.FlagSet, doDelete *flag.FlagSet, toggle *flag.FlagSet, edit *flag.FlagSet, doClear *flag.FlagSet, swap *flag.FlagSet, tags *flag.FlagSet, showTagsOnTodo *int, showTodosForTag *int, tag *flag.FlagSet, tagRm *bool, tagAdd *bool) (FabricateCommand, string, error) {
	baseArgs := strings.TrimSpace(strings.Join(flagArgs[1:], " "))
	switch flagArgs[0] {
	case list.Name():
		return NewListTodos, baseArgs, nil
	case add.Name():
		return NewAddTodo, baseArgs, nil
	case doDelete.Name():
		return NewDeleteTodo, baseArgs, nil
	case toggle.Name():
		return NewToggleTodo, baseArgs, nil
	case edit.Name():
		return NewEditTodo, baseArgs, nil
	case doClear.Name():
		return NewClear, baseArgs, nil
	case swap.Name():
		return NewSwap, baseArgs, nil
	case tags.Name():
		tags.Parse(flagArgs[1:])
		switch {
		case *showTagsOnTodo > 0:
			args := strconv.Itoa(*showTagsOnTodo)
			return NewListTagsOnTodo, args, nil
		case *showTodosForTag > 0:
			args := strconv.Itoa(*showTodosForTag)
			return NewListTodosForTag, args, nil
		default:
			return NewListTags, baseArgs, nil
		}
	case tag.Name():
		tag.Parse(flagArgs[1:])
		args := strings.TrimSpace(strings.Join(tag.Args(), " "))
		switch {
		case *tagRm && len(tag.Args()) > 1:
			return NewRemoveTagFromTodo, args, nil
		case *tagRm && len(tag.Args()) == 1:
			return NewRemoveTag, args, nil
		case *tagAdd:
			return NewAddTag, args, nil
		default:
			return NewTagTodo, args, nil
		}
	default:
		return nil, "", ErrNoCommandFound
	}
}
