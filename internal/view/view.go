package view

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/Hydoc/goo/internal/model"
)

const (
	idMarginRight = 4
)

type View interface {
	RenderList(todoList *model.TodoList)
	RenderLine(str string)
	RenderTags(tagList []*model.Tag)
}

type StdoutView struct {
	writer io.Writer
}

func (v *StdoutView) RenderLine(str string) {
	fmt.Fprintln(v.writer, str)
}

func (v *StdoutView) RenderTags(tagList []*model.Tag) {
	longestEntry := v.findLongestEntry(tagList)

	idStr := v.addMargin(0, idMarginRight, "ID")
	labelStr := v.addMargin(idMarginRight, longestEntry, "TAG")
	headline := fmt.Sprintf("%s%s", idStr, labelStr)
	v.RenderLine(headline)
	v.RenderLine(strings.Repeat("-", len(headline)))
	for _, tag := range tagList {
		tagIdStr := v.addMargin(0, idMarginRight, strconv.Itoa(int(tag.Id)))
		tagNameStr := v.addMargin(0, longestEntry, tag.Name)
		todoStr := fmt.Sprintf("%s%s", tagIdStr, tagNameStr)

		v.RenderLine(todoStr)
	}
}

func (v *StdoutView) RenderList(todoList *model.TodoList) {
	todoList = todoList.SortedByIdAndState()
	offsetStatus := 8
	offsetCheck := 7
	longestEntry := todoList.LenOfLongestTodo()

	idStr := v.addMargin(0, idMarginRight, "ID")
	labelStr := v.addMargin(idMarginRight, longestEntry, "TASK")
	statusStr := v.addMargin(len(labelStr)-longestEntry+offsetStatus, 0, "STATUS")

	headline := fmt.Sprintf("%s%s%s", idStr, labelStr, statusStr)
	v.RenderLine(headline)
	v.RenderLine(strings.Repeat("-", len(headline)))
	for _, todo := range todoList.Items {
		todoIdStr := v.addMargin(0, idMarginRight, strconv.Itoa(todo.Id))
		todoLabelStr := v.addMargin(0, longestEntry, todo.LabelAsString())

		todoStateStr := v.addMargin(len(headline)-longestEntry-offsetCheck, 0, todo.DoneAsString())
		todoStr := fmt.Sprintf("%s%s%s", todoIdStr, todoLabelStr, todoStateStr)

		if todo.IsDone {
			v.RenderLine(v.addGray(todoStr))
		} else {
			v.RenderLine(todoStr)
		}
	}
}

func (v *StdoutView) addMargin(left, right int, str string) string {
	rightMargin := fmt.Sprintf(fmt.Sprintf("%%-%ds", right), str)
	return fmt.Sprintf(fmt.Sprintf("%%%ds", left), rightMargin)
}

func (v *StdoutView) addGray(str string) string {
	return fmt.Sprintf("\033[90m%s\033[0m", str)
}

func (v *StdoutView) findLongestEntry(tagList []*model.Tag) int {
	if len(tagList) == 0 {
		return 0
	}

	current := len(tagList[0].Name)
	for _, tag := range tagList {
		if len(tag.Name) > current {
			current = len(tag.Name)
		}
	}
	return current
}

func New(writer io.Writer) *StdoutView {
	return &StdoutView{writer}
}
