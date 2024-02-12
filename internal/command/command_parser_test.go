package command

import (
	"errors"
	"github.com/Hydoc/goo/internal/view"
	"reflect"
	"testing"
)

func TestNewParser(t *testing.T) {
	validCommands := []*StringCommand{
		{
			Abbreviation: "test",
			Description:  "Test desc",
			Aliases:      []string{"t"},
		},
	}

	parser := NewParser(validCommands)

	if !reflect.DeepEqual(parser.validCommands, validCommands) {
		t.Errorf("want valid commands %v, got %v", validCommands, parser.validCommands)
	}
}

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		name          string
		arg           *view.Argument
		validCommands []*StringCommand
		err           error
		want          *ParsedCommand
	}{
		{
			name: "find command in valid commands by abbreviation",
			arg:  &view.Argument{RawCommand: "test", Payload: "any payload"},
			validCommands: []*StringCommand{
				{
					Abbreviation: "test",
					Description:  "a test cmd",
					Aliases:      []string{"t"},
				},
			},
			err: nil,
			want: &ParsedCommand{
				abbreviation: "test",
				payload:      "any payload",
			},
		},
		{
			name: "find command in valid commands by alias",
			arg:  &view.Argument{RawCommand: "t", Payload: "any payload"},
			validCommands: []*StringCommand{
				{
					Abbreviation: "test",
					Description:  "a test cmd",
					Aliases:      []string{"t"},
				},
			},
			err: nil,
			want: &ParsedCommand{
				abbreviation: "test",
				payload:      "any payload",
			},
		},
		{
			name:          "not find command",
			arg:           &view.Argument{RawCommand: "blub", Payload: "any payload"},
			validCommands: []*StringCommand{},
			err:           errors.New("could not find command 'blub'"),
			want:          nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			parser := NewParser(test.validCommands)
			got, err := parser.Parse(test.arg)

			if err != nil && !reflect.DeepEqual(err, test.err) {
				t.Errorf("want error %v, got %v", test.err, err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}
