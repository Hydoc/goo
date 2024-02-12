package command

import (
	"github.com/Hydoc/goo/internal"
	"reflect"
	"testing"
)

func TestFactory_Fabricate(t *testing.T) {
	todoList := &internal.TodoList{
		Filename: "",
		Items: []*internal.Todo{
			{
				Id:     1,
				Label:  "Test",
				IsDone: false,
			},
		},
	}

	validCommands := []*StringCommand{
		{
			Abbreviation: QuitAbbr,
			Description:  "",
			Aliases:      nil,
		},
		{
			Abbreviation: HelpAbbr,
			Description:  "",
			Aliases:      nil,
		},
		{
			Abbreviation: AddTodoAbbr,
			Description:  "",
			Aliases:      nil,
		},
		{
			Abbreviation: ToggleTodoAbbr,
			Description:  "",
			Aliases:      nil,
		},
		{
			Abbreviation: UndoAbbr,
			Description:  "",
			Aliases:      nil,
		},
		{
			Abbreviation: DeleteTodoAbbr,
			Description:  "",
			Aliases:      nil,
		},
		{
			Abbreviation: EditTodoAbbr,
			Description:  "",
			Aliases:      nil,
		},
	}

	tests := []struct {
		name          string
		parsedCommand *ParsedCommand
		want          Command
		todoList      *internal.TodoList
	}{
		{
			name: QuitAbbr,
			parsedCommand: &ParsedCommand{
				abbreviation: QuitAbbr,
				payload:      "",
			},
			want:     &Quit{todoList: todoList},
			todoList: todoList,
		},
		{
			name: HelpAbbr,
			parsedCommand: &ParsedCommand{
				abbreviation: HelpAbbr,
				payload:      "",
			},
			want:     &Help{validCommands: validCommands},
			todoList: todoList,
		},
		{
			name: AddTodoAbbr,
			parsedCommand: &ParsedCommand{
				abbreviation: AddTodoAbbr,
				payload:      "Another one",
			},
			want: &AddTodo{
				todoList:  todoList,
				todoToAdd: "Another one",
			},
			todoList: todoList,
		},
		{
			name: ToggleTodoAbbr,
			parsedCommand: &ParsedCommand{
				abbreviation: ToggleTodoAbbr,
				payload:      "1",
			},
			want: &ToggleTodo{
				todoList:       todoList,
				todoIdToToggle: 1,
			},
			todoList: todoList,
		},
		{
			name: UndoAbbr,
			parsedCommand: &ParsedCommand{
				abbreviation: UndoAbbr,
				payload:      "",
			},
			want:     &Undo{&DummyUndoableCommand{}},
			todoList: todoList,
		},
		{
			name: DeleteTodoAbbr,
			parsedCommand: &ParsedCommand{
				abbreviation: DeleteTodoAbbr,
				payload:      "1",
			},
			want: &DeleteTodo{
				todoList:   todoList,
				idToDelete: 1,
			},
			todoList: todoList,
		},
		{
			name: EditTodoAbbr,
			parsedCommand: &ParsedCommand{
				abbreviation: EditTodoAbbr,
				payload:      "1 Bla {}",
			},
			want: &EditTodo{
				todoList: todoList,
				idToEdit: 1,
				newLabel: "Bla {}",
			},
			todoList: todoList,
		},
		{
			name: "invalid",
			parsedCommand: &ParsedCommand{
				abbreviation: "invalid",
				payload:      "1 Bla {}",
			},
			want:     nil,
			todoList: todoList,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			factory := NewFactory(validCommands)
			got, _ := factory.Fabricate(test.parsedCommand, test.todoList, &UndoStack{items: []UndoableCommand{&DummyUndoableCommand{}}})

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}

type DummyUndoableCommand struct{}

func (d *DummyUndoableCommand) Undo() {}
