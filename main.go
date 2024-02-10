package main

import (
	"bufio"
	"fmt"
	"github.com/Hydoc/goo/internal"
	"github.com/Hydoc/goo/internal/command"
	"os"
	"strings"
)

type Argument struct {
	rawCommand string
	payload    string
}

func main() {
	quit := command.NewStringCommand(command.QuitAbbr, command.QuitDesc, command.QuitAliases)
	help := command.NewStringCommand(command.HelpAbbr, command.HelpDesc, command.HelpAliases)
	addTodo := command.NewStringCommand(command.AddTodoAbbr, command.AddTodoDesc, command.AddTodoAliases)
	todoList := internal.NewTodoList()
	parser := command.NewParser([]*command.StringCommand{quit, help, addTodo}, todoList)
	reader := bufio.NewReader(os.Stdin)
	for {
		if todoList.HasItems() {
			fmt.Println(todoList)
		}
		fmt.Print("> ")
		argument := readline(reader)
		cmd, err := parser.Parse(argument.rawCommand, argument.payload)
		if err != nil {
			fmt.Println(err)
			continue
		}
		cmd.Execute()
	}
}

func readline(reader *bufio.Reader) *Argument {
	choice, _ := reader.ReadString('\n')
	choiceSplit := strings.Split(strings.TrimSuffix(choice, "\n"), " ")
	var payload string
	if len(choiceSplit) > 1 {
		payload = choiceSplit[1]
	}

	return &Argument{
		rawCommand: choiceSplit[0],
		payload:    payload,
	}
}
