package command

type UndoStack struct {
	items []UndoableCommand
}

func (stack *UndoStack) Pop() UndoableCommand {
	cmd, items := stack.items[len(stack.items)-1], stack.items[:len(stack.items)-1]
	stack.items = items
	return cmd
}

func (stack *UndoStack) Push(cmd UndoableCommand) {
	stack.items = append(stack.items, cmd)
}

func (stack *UndoStack) HasItems() bool {
	return len(stack.items) > 0
}

func NewUndoStack() *UndoStack {
	return &UndoStack{
		items: make([]UndoableCommand, 0),
	}
}
