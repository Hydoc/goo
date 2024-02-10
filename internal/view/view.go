package view

import (
	"bufio"
	"fmt"
	"strings"
)

type Argument struct {
	RawCommand string
	Payload    string
}

type StdoutView struct {
	reader *bufio.Reader
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
		payload = choiceSplit[1]
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
