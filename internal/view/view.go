package view

import (
	"fmt"
	"github.com/Hydoc/goo/internal/model"
	"io"
	"strconv"
	"strings"
)

const (
	idMarginRight = 4
)

type View interface {
	RenderList(todoList *model.TodoList)
	RenderLine(str string)
	RenderTags(todoList *model.TodoList)
}

type StdoutView struct {
	writer io.Writer
}

func (v *StdoutView) RenderLine(str string) {
	fmt.Fprintln(v.writer, str)
}

func (v *StdoutView) RenderTags(todoList *model.TodoList) {
	longestEntry := todoList.LenOfLongestTag()

	idStr := v.addMargin(0, idMarginRight, "ID")
	labelStr := v.addMargin(idMarginRight, longestEntry, "TAG")
	headline := fmt.Sprintf("%s%s", idStr, labelStr)
	v.RenderLine(headline)
	v.RenderLine(strings.Repeat("-", len(headline)))
	for _, tag := range todoList.TagList {
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

		var todoLabelStr string
		if todo.HasTags() {
			todoLabelStr = v.addMargin(0, longestEntry-1, todo.LabelAsString())
		} else {
			todoLabelStr = v.addMargin(0, longestEntry, todo.LabelAsString())
		}

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

func New(writer io.Writer) *StdoutView {
	return &StdoutView{writer}
}
