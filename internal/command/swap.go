package command

import (
	"fmt"
	"github.com/Hydoc/goo/internal/model"
	"github.com/Hydoc/goo/internal/view"
	"strconv"
	"strings"
)

type Swap struct {
	todoList *model.TodoList
	view     view.View
	firstId  int
	secondId int
}

func (cmd *Swap) Execute() {
	cmd.todoList.Swap(cmd.firstId, cmd.secondId)
	cmd.view.RenderList(cmd.todoList)
	cmd.todoList.SaveToFile()
}

func NewSwap(todoList *model.TodoList, view view.View, payload string) (Command, error) {
	splitBySpace := strings.Split(payload, " ")

	if len(splitBySpace) < 2 || len(splitBySpace) > 2 {
		return nil, fmt.Errorf("can not swap, need two ids separated by space")
	}

	var ids []int
	for _, entry := range splitBySpace {
		id, err := strconv.Atoi(entry)
		if err != nil {
			return nil, errInvalidId(entry)
		}

		if !todoList.Has(id) {
			return nil, errNoTodoWithId(id)
		}
		ids = append(ids, id)
	}

	return &Swap{todoList, view, ids[0], ids[1]}, nil
}
