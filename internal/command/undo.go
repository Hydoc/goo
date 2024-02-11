package command

import "errors"

var UndoAliases = []string{"u"}

const (
	UndoAbbr      = "undo"
	UndoDesc      = "Undo the last action"
	nothingToUndo = "nothing to undo"
)

type Undo struct {
	cmd UndoableCommand
}

func (u *Undo) Execute() {
	u.cmd.Undo()
}

func newUndo(undoStack *UndoStack) (*Undo, error) {
	if !undoStack.HasItems() {
		return nil, errors.New(nothingToUndo)
	}

	return &Undo{
		cmd: undoStack.Pop(),
	}, nil
}
