package command

import (
	"reflect"
	"testing"
)

func TestNewStringCommand(t *testing.T) {
	abbr := "Test"
	desc := "Test desc"
	aliases := []string{"t"}
	cmd := NewStringCommand(abbr, desc, aliases)

	if cmd.Abbreviation != abbr {
		t.Errorf("want abbreviation %v, got %v", abbr, cmd.Abbreviation)
	}

	if cmd.Description != desc {
		t.Errorf("want description %v, got %v", desc, cmd.Description)
	}

	if !reflect.DeepEqual(cmd.Aliases, aliases) {
		t.Errorf("want aliases %v, got %v", aliases, cmd.Aliases)
	}
}
