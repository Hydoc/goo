package command

import (
	"github.com/Hydoc/goo/internal/model"
	"reflect"
	"testing"
)

func TestFactory_Fabricate(t *testing.T) {
	tests := []struct {
		name     string
		want     Command
		err      error
		todoList *model.TodoList
		toggle   int
		add      bool
		doDelete int
		edit     bool
		doClear  bool
		args     string
	}{
		{
			name: "edit",
			want: &EditTodo{
				todoList: &model.TodoList{
					Items: []*model.Todo{
						{
							Id:     1,
							Label:  "Test",
							IsDone: false,
						},
					},
				},
				idToEdit: 1,
				newLabel: "ABC",
			},
			err: nil,
			todoList: &model.TodoList{
				Items: []*model.Todo{
					{
						Id:     1,
						Label:  "Test",
						IsDone: false,
					},
				},
			},
			toggle:   0,
			add:      false,
			doDelete: 0,
			edit:     true,
			doClear:  false,
			args:     "1 ABC",
		},
		{
			name: "delete",
			want: &DeleteTodo{
				idToDelete: 1,
				todoList: &model.TodoList{
					Items: []*model.Todo{
						{
							Id:     1,
							Label:  "Test",
							IsDone: false,
						},
					},
				},
			},
			err: nil,
			todoList: &model.TodoList{
				Items: []*model.Todo{
					{
						Id:     1,
						Label:  "Test",
						IsDone: false,
					},
				},
			},
			toggle:   0,
			add:      false,
			doDelete: 1,
			edit:     false,
			doClear:  false,
			args:     "",
		},
		{
			name: "toggle",
			want: &ToggleTodo{
				idToToggle: 1,
				todoList: &model.TodoList{
					Items: []*model.Todo{
						{
							Id:     1,
							Label:  "Test",
							IsDone: false,
						},
					},
				},
			},
			err: nil,
			todoList: &model.TodoList{
				Items: []*model.Todo{
					{
						Id:     1,
						Label:  "Test",
						IsDone: false,
					},
				},
			},
			toggle:   1,
			add:      false,
			doDelete: 0,
			edit:     false,
			doClear:  false,
			args:     "",
		},
		{
			name: "add",
			want: &AddTodo{
				todoToAdd: "Hello!",
				todoList:  &model.TodoList{Items: make([]*model.Todo, 0)},
			},
			err:      nil,
			todoList: &model.TodoList{Items: make([]*model.Todo, 0)},
			toggle:   0,
			add:      true,
			doDelete: 0,
			edit:     false,
			doClear:  false,
			args:     "Hello!",
		},
		{
			name:     "clear",
			want:     &Clear{&model.TodoList{}},
			err:      nil,
			todoList: &model.TodoList{},
			toggle:   0,
			add:      false,
			doDelete: 0,
			edit:     false,
			doClear:  true,
			args:     "",
		},
		{
			name:     "invalid",
			want:     nil,
			err:      nil,
			todoList: nil,
			toggle:   0,
			add:      false,
			doDelete: 0,
			edit:     false,
			doClear:  false,
			args:     "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := NewFactory().Fabricate(test.todoList, test.toggle, test.add, test.doDelete, test.edit, test.doClear, test.args)

			if err != nil && !reflect.DeepEqual(err, test.err) {
				t.Errorf("expected err %v, got %v", test.err, err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}
