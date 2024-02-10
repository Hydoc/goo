package command

type Command interface {
	Execute()
}

type UndoableCommand interface {
	Undo()
}

type StringCommand struct {
	Abbreviation string
	Description  string
	Aliases      []string
}

func NewStringCommand(abbreviation, description string, aliases []string) *StringCommand {
	return &StringCommand{
		Abbreviation: abbreviation,
		Description:  description,
		Aliases:      aliases,
	}
}
