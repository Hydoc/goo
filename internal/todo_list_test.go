package internal

import (
	"encoding/json"
	"os"
	"reflect"
	"syscall"
	"testing"
)

func setUpFile(t *testing.T, filename string, content interface{}) func() {
	jsonContent, err := json.Marshal(content)

	if err != nil {
		t.Errorf("there was an error marshaling the content %s", err)
	}

	err = os.WriteFile(filename, jsonContent, 0644)
	if err != nil {
		t.Errorf("there was an error creating the file %s", err)
	}

	return func() {
		err := os.Remove(filename)
		if err != nil {
			t.Errorf("there was an error removing the file %s", err)
		}
	}
}

func TestNewTodoListFromFile(t *testing.T) {
	filename := "test.json"
	tests := []struct {
		name       string
		todoItems  interface{}
		err        error
		want       *TodoList
		createFile bool
	}{
		{
			name:       "create from empty file",
			todoItems:  []*Todo{},
			err:        nil,
			createFile: true,
			want: &TodoList{
				Filename: filename,
				Items:    []*Todo{},
			},
		},
		{
			name: "create from file with todos",
			todoItems: []*Todo{
				NewTodo("Test", 1),
				NewTodo("Another Test", 2),
			},
			err:        nil,
			createFile: true,
			want: &TodoList{
				Filename: filename,
				Items: []*Todo{
					NewTodo("Test", 1),
					NewTodo("Another Test", 2),
				},
			},
		},
		{
			name:       "not create due to not existing file",
			todoItems:  make([]*Todo, 0),
			err:        &os.PathError{Op: "open", Path: "test.json", Err: syscall.ENOENT},
			want:       nil,
			createFile: false,
		},
		{
			name: "not create due to invalid json",
			todoItems: map[string]interface{}{
				"invalid": "no todo item",
			},
			err: &json.UnmarshalTypeError{
				Value:  "object",
				Type:   reflect.TypeOf([]*Todo{}),
				Offset: 1,
				Struct: "",
				Field:  "",
			},
			want:       nil,
			createFile: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.createFile {
				tearDown := setUpFile(t, filename, test.todoItems)
				defer tearDown()
			}
			todoList, err := NewTodoListFromFile(filename)

			if err != nil && !reflect.DeepEqual(err, test.err) {
				t.Errorf("expected error %#v, but got %#v", test.err, err)
			}

			if !reflect.DeepEqual(todoList, test.want) {
				t.Errorf("want todo list %v, got %v", test.want, todoList)
			}
		})
	}
}

func TestTodoList_Add(t *testing.T) {
	list := &TodoList{
		Filename: "",
		Items:    make([]*Todo, 0),
	}

	list.Add(NewTodo("Test", 1))
	firstTodo := list.Items[0]

	if firstTodo.Label != "Test" {
		t.Errorf("want first todo label %s, got %s", "Test", firstTodo.Label)
	}

	if firstTodo.Id != 1 {
		t.Errorf("want first todo id %d, got %d", 1, firstTodo.Id)
	}
}

func TestTodoList_Find(t *testing.T) {
	tests := []struct {
		name     string
		todoList *TodoList
		want     *Todo
		id       int
	}{
		{
			name: "do find",
			todoList: &TodoList{
				Filename: "",
				Items: []*Todo{
					NewTodo("Test", 1),
				},
			},
			want: NewTodo("Test", 1),
			id:   1,
		},
		{
			name: "not find",
			todoList: &TodoList{
				Filename: "",
				Items: []*Todo{
					NewTodo("Test", 1),
				},
			},
			want: nil,
			id:   2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.todoList.Find(test.id)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("want todo list %v, got %v", test.want, got)
			}
		})
	}
}

func TestTodoList_LenOfLongestTodo(t *testing.T) {
	tests := []struct {
		name     string
		todoList *TodoList
		want     int
	}{
		{
			name: "one entry",
			todoList: &TodoList{
				Filename: "",
				Items: []*Todo{
					NewTodo("Hello", 1),
					NewTodo("Hello World", 2),
				},
			},
			want: len("Hello World"),
		},
		{
			name: "no entry",
			todoList: &TodoList{
				Filename: "",
				Items:    make([]*Todo, 0),
			},
			want: 0,
		},
		{
			name: "multiple entries",
			todoList: &TodoList{
				Filename: "",
				Items: []*Todo{
					NewTodo("Hello", 1),
					NewTodo("Hello World", 2),
					NewTodo("Hello World123", 3),
					NewTodo("Hello World!", 4),
				},
			},
			want: len("Hello World123"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.todoList.LenOfLongestTodo()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("want todo list %v, got %v", test.want, got)
			}
		})
	}
}

func TestTodoList_Edit(t *testing.T) {
	tests := []struct {
		name     string
		todoList *TodoList
		id       int
		newLabel string
		want     *Todo
	}{
		{
			name: "do edit",
			todoList: &TodoList{
				Filename: "",
				Items: []*Todo{
					NewTodo("Hello", 1),
					NewTodo("Hello World", 2),
				},
			},
			id:       2,
			newLabel: "Bla",
			want:     NewTodo("Bla", 2),
		},
		{
			name: "not edit due to not found todo",
			todoList: &TodoList{
				Filename: "",
				Items:    make([]*Todo, 0),
			},
			id:       1,
			newLabel: "Hello",
			want:     nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.todoList.Edit(test.id, test.newLabel)
			if !reflect.DeepEqual(test.want, test.todoList.Find(test.id)) {
				t.Errorf("want todo %v, got %v", test.want, test.todoList.Find(test.id))
			}
		})
	}
}

func TestTodoList_Has(t *testing.T) {
	tests := []struct {
		name     string
		todoList *TodoList
		id       int
		want     bool
	}{
		{
			name: "true when todo list has item",
			todoList: &TodoList{
				Filename: "",
				Items: []*Todo{
					NewTodo("Hello", 1),
					NewTodo("Hello World", 2),
				},
			},
			id:   2,
			want: true,
		},
		{
			name: "false when todo list is missing item",
			todoList: &TodoList{
				Filename: "",
				Items:    make([]*Todo, 0),
			},
			id:   1,
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			has := test.todoList.Has(test.id)
			if has != test.want {
				t.Errorf("want %v, got %v", test.want, has)
			}
		})
	}
}

func TestTodoList_Remove(t *testing.T) {
	tests := []struct {
		name     string
		todoList *TodoList
		id       int
		want     []*Todo
	}{
		{
			name: "do delete",
			todoList: &TodoList{
				Filename: "",
				Items: []*Todo{
					NewTodo("Hello", 1),
					NewTodo("Hello World", 2),
				},
			},
			id: 2,
			want: []*Todo{
				NewTodo("Hello", 1),
			},
		},
		{
			name: "not delete when id not found",
			todoList: &TodoList{
				Filename: "",
				Items:    make([]*Todo, 0),
			},
			id:   1,
			want: make([]*Todo, 0),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.todoList.Remove(test.id)
			if !reflect.DeepEqual(test.todoList.Items, test.want) {
				t.Errorf("want todo list items %v, got %v", test.want, test.todoList.Items)
			}
		})
	}
}

func TestTodoList_Toggle(t *testing.T) {
	tests := []struct {
		name     string
		todoList *TodoList
		id       int
		want     []*Todo
	}{
		{
			name: "from false to true",
			todoList: &TodoList{
				Filename: "",
				Items: []*Todo{
					NewTodo("Hello", 1),
					NewTodo("Hello World!", 2),
					NewTodo("Hello World", 3),
				},
			},
			id: 2,
			want: []*Todo{
				NewTodo("Hello", 1),
				{
					Id:     2,
					Label:  "Hello World!",
					IsDone: true,
				},
				NewTodo("Hello World", 3),
			},
		},
		{
			name: "from true to false",
			todoList: &TodoList{
				Filename: "",
				Items: []*Todo{
					{
						Id:     1,
						Label:  "Hello",
						IsDone: true,
					},
					NewTodo("Hello World!", 2),
					NewTodo("Hello World", 3),
				},
			},
			id: 1,
			want: []*Todo{
				{
					Id:     1,
					Label:  "Hello",
					IsDone: false,
				},
				NewTodo("Hello World!", 2),
				NewTodo("Hello World", 3),
			},
		},
		{
			name: "not toggle due to not found todo",
			todoList: &TodoList{
				Filename: "",
				Items: []*Todo{
					NewTodo("Hello", 1),
					NewTodo("Hello World!", 2),
					NewTodo("Hello World", 3),
				},
			},
			id: 1212,
			want: []*Todo{
				NewTodo("Hello", 1),
				NewTodo("Hello World!", 2),
				NewTodo("Hello World", 3),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.todoList.Toggle(test.id)
			if !reflect.DeepEqual(test.todoList.Items, test.want) {
				t.Errorf("want todo list items %v, got %v", test.want, test.todoList.Items)
			}
		})
	}
}

func TestTodoList_NextId(t *testing.T) {
	tests := []struct {
		name     string
		todoList *TodoList
		want     int
	}{
		{
			name: "for empty list",
			todoList: &TodoList{
				Filename: "",
				Items:    make([]*Todo, 0),
			},
			want: 1,
		},
		{
			name: "with entries",
			todoList: &TodoList{
				Filename: "",
				Items: []*Todo{
					NewTodo("Hello", 1),
					NewTodo("Hello World!", 2),
					NewTodo("Hello World", 3),
				},
			},
			want: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			nextId := test.todoList.NextId()
			if !reflect.DeepEqual(nextId, test.want) {
				t.Errorf("want todo list items %v, got %v", test.want, nextId)
			}
		})
	}
}

func TestTodoList_SaveToFile(t *testing.T) {
	filename := "test.json"
	tearDown := setUpFile(t, filename, []*Todo{})
	defer tearDown()

	todoList, _ := NewTodoListFromFile(filename)
	todoList.Add(NewTodo("A", todoList.NextId()))
	todoList.SaveToFile()

	todoList, _ = NewTodoListFromFile(filename)

	if len(todoList.Items) == 0 {
		t.Errorf("expected saved todo list to have items")
	}

	if todoList.Items[0].Label != "A" {
		t.Errorf("expected saved todo list item at index 0 to have label 'A', got %v", todoList.Items[0].Label)
	}

	if todoList.Items[0].Id != 1 {
		t.Errorf("expected saved todo list item at index 0 to have id '1', got %v", todoList.Items[0].Id)
	}
}
