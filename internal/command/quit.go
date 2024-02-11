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

func (cmd *Quit) Execute() {
	os.Exit(0)
}

func newQuit() *Quit {
	return &Quit{}
}
