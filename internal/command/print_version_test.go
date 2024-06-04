package command

import "testing"

func TestPrintVersion_Execute(t *testing.T) {
	view := newDummyView()
	cmd, _ := NewPrintVersion(view, "1.4.3")
	cmd.Execute()

	if view.RenderLineCalls != 1 {
		t.Errorf("want one call to view.RenderLine, got %d", view.RenderLineCalls)
	}
}
