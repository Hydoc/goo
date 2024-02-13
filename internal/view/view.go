package view

import (
	"bufio"
	"fmt"
	"github.com/Hydoc/goo/internal"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

const (
	colorGray = "gray"
)

type Argument struct {
	RawCommand string
	Payload    string
}

type StdoutView struct {
	reader *bufio.Reader
	writer io.Writer
}

func (v *StdoutView) ClearScreen() {
	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
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

func (v *StdoutView) moveToBottom() {
	fmt.Printf("\x1B[%d;0f", v.getTerminalWidth())
}

func (v *StdoutView) bold(str string) string {
	return fmt.Sprintf("\033[1m%s\033[0m", str)
}

func (v *StdoutView) toColor(str, color string) string {
	switch color {
	case colorGray:
		return fmt.Sprintf("\033[90m%s\033[0m", str)
	default:
		return str
	}
}

func (v *StdoutView) Prompt() *Argument {
	v.moveToBottom()
	v.Render("> ")
	choice, _ := v.reader.ReadString('\n')
	choiceSplit := strings.Split(strings.TrimSuffix(choice, "\n"), " ")
	var payload string
	if len(choiceSplit) > 1 {
		payload = strings.TrimSpace(strings.Join(choiceSplit[1:], " "))
	}

	return &Argument{
		RawCommand: choiceSplit[0],
		Payload:    payload,
	}
}

func (v *StdoutView) getTerminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	width, err := strconv.Atoi(strings.TrimSuffix(strings.Split(string(out), " ")[1], "\n"))
	if err != nil {
		log.Fatal(err)
	}
	return width
}

func New(reader *bufio.Reader, writer io.Writer) *StdoutView {
	return &StdoutView{
		reader: reader,
		writer: writer,
	}
}
