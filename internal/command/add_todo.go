package command

import "github.com/Hydoc/goo/internal"

var AddTodoAliases = []string{"a"}

const (
	AddTodoAbbr = "add"
	AddTodoDesc = "Add a new todo"
	AddTodoHelp = "use the command like so: add <label of the todo>"
)

type AddTodo struct {
	previousTodoListItems []*internal.Todo
	todoList              *internal.TodoList
	todoToAdd             string
}

func (cmd *AddTodo) Execute() {
	cmd.previousTodoListItems = cmd.todoList.Items
	cmd.todoList.Add(internal.NewTodo(cmd.todoToAdd, cmd.todoList.NextId()))
}

func (cmd *AddTodo) Undo() {
	cmd.todoList.Items = cmd.previousTodoListItems
}

func NewAddTodo(todoList *internal.TodoList, todoToAdd string) *AddTodo {
	return &AddTodo{
		todoList:  todoList,
		todoToAdd: todoToAdd,
	}
}

func CanCreateAddTodo(payload string) bool {
	return len(payload) > 0
}
