package command

var UndoAliases = []string{"u"}

const (
	UndoAbbr      = "undo"
	UndoDesc      = "Undo the last action"
	NothingToUndo = "nothing to undo"
)

type Undo struct {
	cmd UndoableCommand
}

func (u *Undo) Execute() {
	u.cmd.Undo()
}
