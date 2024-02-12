package command

import (
	"bufio"
	"github.com/Hydoc/goo/internal/view"
	"os"
	"testing"
)

func TestHelp_Execute(t *testing.T) {
	validCommands := []*StringCommand{
		{
			Abbreviation: "Test",
			Description:  "do smth",
			Aliases:      []string{"t"},
		},
		{
			Abbreviation: "die",
			Description:  "die",
			Aliases:      []string{"d"},
		},
	}

	want := "Here is a list of all possible commands:\r\nTest: do smth (aliases: t)\r\ndie: die (aliases: d)\n"

	writer := &DummyWriter{}
	cmd := newHelp(view.New(bufio.NewReader(os.Stdin), writer), validCommands)

	cmd.Execute()

	writtenBytes := writer.writtenBytes
	if want != string(writtenBytes) {
		t.Errorf("want %#v, got %#v", want, writtenBytes)
	}
}

type DummyWriter struct {
	writtenBytes []byte
}

func (w *DummyWriter) Write(p []byte) (n int, err error) {
	w.writtenBytes = append(w.writtenBytes, p...)
	return 0, nil
}
