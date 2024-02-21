package model

import (
	"reflect"
	"testing"
)

func TestNewTag(t *testing.T) {
	want := &Tag{1, "hello"}
	got := NewTag(1, "  HelLo ")

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
