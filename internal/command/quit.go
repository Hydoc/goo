package command

import (
	"os"
)

var QuitAliases = []string{"q"}

const (
	QuitAbbr = "quit"
	QuitDesc = "Quit the app"
)

type Quit struct{}

func (q *Quit) Execute() {
	os.Exit(0)
}

func NewQuit() *Quit {
	return &Quit{}
}
