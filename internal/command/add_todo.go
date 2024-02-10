package command

import "github.com/Hydoc/goo/internal"

var AddTodoAliases = []string{"a"}

const (
	AddTodoAbbr = "add"
	AddTodoDesc = "Add a new todo"
)

type AddTodo struct {
	todoList  *internal.TodoList
	todoToAdd string
}

func (cmd *AddTodo) Execute() {
	cmd.todoList.Add(internal.NewTodo(cmd.todoToAdd))
}

func NewAddTodo(todoToAdd string, todoList *internal.TodoList) *AddTodo {
	return &AddTodo{
		todoList:  todoList,
		todoToAdd: todoToAdd,
	}
}
