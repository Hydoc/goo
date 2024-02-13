package view

import (
	"fmt"
	"github.com/Hydoc/goo/internal"
	"io"
	"strconv"
	"strings"
)

const (
	colorGray = "gray"
)

type StdoutView struct {
	writer io.Writer
}

func (v *StdoutView) Render(str string) {
	fmt.Fprint(v.writer, str)
}

func (v *StdoutView) RenderLine(str string) {
	fmt.Fprintln(v.writer, str)
}

func (v *StdoutView) RenderList(todoList *internal.TodoList) {
	textWithMarginRight := func(margin int, str string) string {
		return fmt.Sprintf(fmt.Sprintf("%%-%ds", margin), str)
	}

	textWithMarginLeft := func(margin int, str string) string {
		return fmt.Sprintf(fmt.Sprintf("%%%ds", margin), str)
	}

	idMarginRight := 4
	offsetStatus := 8
	offsetCheck := 5
	longestEntry := todoList.LenOfLongestTodo()

	idStr := textWithMarginRight(idMarginRight, "ID")
	labelStr := textWithMarginLeft(idMarginRight, textWithMarginRight(longestEntry, "TASK"))
	statusStr := textWithMarginLeft(longestEntry-len(labelStr)+offsetStatus, "STATUS")

	headline := fmt.Sprintf("%s%s%s", idStr, labelStr, statusStr)
	v.RenderLine(headline)
	v.RenderLine(strings.Repeat("-", len(headline)))
	for _, todo := range todoList.Items {
		if todo.IsDone {
			v.RenderLine(v.toColor(fmt.Sprintf("%s%s%s", textWithMarginRight(idMarginRight, strconv.Itoa(todo.Id)), textWithMarginRight(longestEntry, todo.Label), textWithMarginLeft(offsetCheck, "✓")), colorGray))
		} else {
			v.RenderLine(fmt.Sprintf("%s%s%s", textWithMarginRight(idMarginRight, strconv.Itoa(todo.Id)), textWithMarginRight(longestEntry, todo.Label), textWithMarginLeft(offsetCheck, "○")))
		}
	}
}

func (v *StdoutView) toColor(str, color string) string {
	switch color {
	case colorGray:
		return fmt.Sprintf("\033[90m%s\033[0m", str)
	default:
		return str
	}
}

func New(writer io.Writer) *StdoutView {
	return &StdoutView{
		writer: writer,
	}
}
