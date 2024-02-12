package view

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Argument struct {
	RawCommand string
	Payload    string
}

type StdoutView struct {
	reader *bufio.Reader
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
	fmt.Print(str)
}

func (v *StdoutView) RenderLine(str string) {
	fmt.Println(str)
}

func (v *StdoutView) Prompt() *Argument {
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

func New(reader *bufio.Reader) *StdoutView {
	return &StdoutView{
		reader: reader,
	}
}
