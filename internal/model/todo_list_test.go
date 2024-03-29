package model

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
		todoList   interface{}
		err        error
		want       *TodoList
		createFile bool
	}{
		{
			name: "create from empty file",
			todoList: &TodoList{
				Filename: filename,
				TagList:  nil,
				Items:    []*Todo{},
			},
			err:        nil,
			createFile: true,
			want: &TodoList{
				Filename: filename,
				Items:    []*Todo{},
				TagList:  make([]*Tag, 0),
			},
		},
		{
			name: "create from file with todos",
			todoList: &TodoList{
				Filename: filename,
				Items: []*Todo{
					NewTodo("Test", 1),
					NewTodo("Another Test", 2),
				},
				TagList: []*Tag{
					{
						Id:   1,
						Name: "hi",
					},
				},
			},
			err:        nil,
			createFile: true,
			want: &TodoList{
				Filename: filename,
				Items: []*Todo{
					NewTodo("Test", 1),
					NewTodo("Another Test", 2),
				},
				TagList: []*Tag{
					{
						Id:   1,
						Name: "hi",
					},
				},
			},
		},
		{
			name:       "not create due to not existing file",
			err:        &os.PathError{Op: "open", Path: "test.json", Err: syscall.ENOENT},
			want:       nil,
			createFile: false,
		},
		{
			name: "not create due to invalid json",
			todoList: map[string]interface{}{
				"tagList": []int{1, 2, 3},
			},
			err:        &json.UnmarshalTypeError{Value: "number", Type: reflect.TypeOf(Tag{}), Offset: 13, Struct: "TodoList", Field: "tagList"},
			want:       nil,
			createFile: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.createFile {
				tearDown := setUpFile(t, filename, test.todoList)
				defer tearDown()
			}
			todoList, err := NewTodoListFromFile(filename)

			if err != nil && test.err.Error() != err.Error() {
				t.Errorf("expected error %#v, but got %#v", test.err, err)
			}

			if !reflect.DeepEqual(todoList, test.want) {
				t.Errorf("want todo list %v, got %v", test.want, todoList)
			}
		})
	}
}

func TestTodoList_Swap(t *testing.T) {
	list := &TodoList{
		Filename: "",
		Items: []*Todo{
			{
				Id:     1,
				Label:  "Hello",
				IsDone: true,
			},
			{
				Id:     2,
				Label:  "World",
				IsDone: false,
			},
		},
	}

	list.Swap(1, 2)

	firstTodo := list.Items[0]
	if firstTodo.Label != "World" {
		t.Errorf("want first todo label World, got %s", firstTodo.Label)
	}

	if firstTodo.IsDone != true {
		t.Errorf("expected done state not to be swapped")
	}

	secondTodo := list.Items[1]
	if secondTodo.Label != "Hello" {
		t.Errorf("want second todo label Hello, got %s", secondTodo.Label)
	}

	if secondTodo.IsDone != false {
		t.Errorf("expected done state not to be swapped")
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
					Tags:   make([]TagId, 0),
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

func TestTodoList_NextTodoId(t *testing.T) {
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
				TagList:  make([]*Tag, 0),
			},
			want: 1,
		},
		{
			name: "with entries",
			todoList: &TodoList{
				Filename: "",
				Items: []*Todo{
					NewTodo("Hello", 1),
					NewTodo("Hello World", 3),
				},
				TagList: make([]*Tag, 0),
			},
			want: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			nextId := test.todoList.NextTodoId()
			if !reflect.DeepEqual(nextId, test.want) {
				t.Errorf("want %v, got %v", test.want, nextId)
			}
		})
	}
}

func TestTodoList_SaveToFile(t *testing.T) {
	filename := "test.json"
	tearDown := setUpFile(t, filename, &TodoList{})
	defer tearDown()

	todoList, _ := NewTodoListFromFile(filename)
	todoList.Add(NewTodo("A", todoList.NextTodoId()))
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

func TestTodoList_SortedByIdAndState(t *testing.T) {
	tests := []struct {
		name     string
		todoList *TodoList
		want     *TodoList
	}{
		{
			name: "should be sorted by id and done",
			todoList: &TodoList{
				Items: []*Todo{
					{
						Id:     1,
						Label:  "",
						IsDone: true,
					},
					{
						Id:     2,
						Label:  "",
						IsDone: false,
					},
					{
						Id:     3,
						Label:  "",
						IsDone: true,
					},
				},
			},
			want: &TodoList{
				Items: []*Todo{
					{
						Id:     2,
						Label:  "",
						IsDone: false,
					},
					{
						Id:     1,
						Label:  "",
						IsDone: true,
					},
					{
						Id:     3,
						Label:  "",
						IsDone: true,
					},
				},
			},
		},
		{
			name: "should be sorted by id when no todo is done",
			todoList: &TodoList{
				Items: []*Todo{
					{
						Id:     3,
						Label:  "",
						IsDone: false,
					},
					{
						Id:     2,
						Label:  "",
						IsDone: false,
					},
					{
						Id:     1,
						Label:  "",
						IsDone: false,
					},
				},
			},
			want: &TodoList{
				Items: []*Todo{
					{
						Id:     1,
						Label:  "",
						IsDone: false,
					},
					{
						Id:     2,
						Label:  "",
						IsDone: false,
					},
					{
						Id:     3,
						Label:  "",
						IsDone: false,
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.todoList.SortedByIdAndState()

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("want order %v, got %v", test.want, got)
			}
		})
	}
}

func TestTodoList_Clear(t *testing.T) {
	todoList := &TodoList{
		Filename: "test.json",
		Items: []*Todo{
			NewTodo("Test", 1),
			NewTodo("Another", 2),
			NewTodo("Third", 3),
		},
	}

	todoList.Clear()

	if todoList.HasItems() {
		t.Error("expected todolist to be empty")
	}
}

func TestTodoList_AddTag(t *testing.T) {
	todoList := &TodoList{
		Filename: "",
		Items:    nil,
		TagList:  nil,
	}
	wantTagList := []*Tag{
		{
			Id:   1,
			Name: "test tag",
		},
	}
	todoList.AddTag(NewTag(1, "Test Tag"))

	if !reflect.DeepEqual(todoList.TagList, wantTagList) {
		t.Errorf("want %v, got %v", wantTagList, todoList.TagList)
	}
}

func TestTodoList_FindTag(t *testing.T) {
	todoList := &TodoList{
		TagList: []*Tag{
			{
				Id:   1,
				Name: "hi",
			},
		},
	}
	tests := []struct {
		name        string
		tagIdToFind int
		want        *Tag
	}{
		{
			name:        "find when todo list has tag",
			tagIdToFind: 1,
			want: &Tag{
				Id:   1,
				Name: "hi",
			},
		},
		{
			name:        "not find when id not found in todo list tags",
			tagIdToFind: 21,
			want:        nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := todoList.FindTag(test.tagIdToFind)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}

func TestTodoList_TagTodo(t *testing.T) {
	todoList := TodoList{
		Filename: "",
		Items: []*Todo{
			{
				Id:     1,
				Label:  "Test",
				IsDone: false,
				Tags:   make([]TagId, 0),
			},
		},
		TagList: nil,
	}
	todoList.TagTodo(1, 12)

	if !todoList.Find(1).HasTag(12) {
		t.Error("expected todo with id 1 to have tag with id 12")
	}
}

func TestTodoList_RemoveTagFromTodo(t *testing.T) {
	todoList := TodoList{
		Filename: "",
		Items: []*Todo{
			{
				Id:     1,
				Label:  "Test",
				IsDone: false,
				Tags:   []TagId{12},
			},
		},
		TagList: nil,
	}
	todoList.RemoveTagFromTodo(12, 1)

	if todoList.Find(1).HasTag(12) {
		t.Error("expected todo with id 1 to not have tag with id 12")
	}
}

func TestTodoList_LenOfLongestTag(t *testing.T) {
	tests := []struct {
		name     string
		todoList *TodoList
		want     int
	}{
		{
			name: "one entry",
			todoList: &TodoList{
				Filename: "",
				TagList: []*Tag{
					NewTag(1, "BLA"),
				},
			},
			want: len("BLA"),
		},
		{
			name: "no entry",
			todoList: &TodoList{
				Filename: "",
				TagList:  make([]*Tag, 0),
			},
			want: 0,
		},
		{
			name: "multiple entries",
			todoList: &TodoList{
				Filename: "",
				TagList: []*Tag{
					NewTag(1, "Hello"),
					NewTag(2, "Hello World123"),
					NewTag(3, "Hello World!  "),
				},
			},
			want: len("Hello World123"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.todoList.LenOfLongestTag()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("want todo list %v, got %v", test.want, got)
			}
		})
	}
}

func TestTodoList_HasTag(t *testing.T) {
	todoList := &TodoList{
		Filename: "",
		Items:    make([]*Todo, 0),
		TagList: []*Tag{
			NewTag(1, "bla"),
		},
	}

	tests := []struct {
		name      string
		want      bool
		tagToFind TagId
	}{
		{
			name:      "true if todo list has tag",
			want:      true,
			tagToFind: 1,
		},
		{
			name:      "false if todo list does not have tag",
			want:      false,
			tagToFind: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := todoList.HasTag(test.tagToFind)

			if got != test.want {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}

func TestTodoList_HasTagWith(t *testing.T) {
	todoList := &TodoList{
		Filename: "",
		Items:    make([]*Todo, 0),
		TagList: []*Tag{
			NewTag(1, "bla"),
		},
	}

	tests := []struct {
		name      string
		want      bool
		tagToFind string
	}{
		{
			name:      "true if todo list has tag",
			want:      true,
			tagToFind: "bla",
		},
		{
			name:      "false if todo list does not have tag",
			want:      false,
			tagToFind: "BLUB",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := todoList.HasTagWith(test.tagToFind)

			if got != test.want {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}

func TestTodoList_NextTagId(t *testing.T) {
	tests := []struct {
		name     string
		todoList *TodoList
		want     TagId
	}{
		{
			name: "for empty list",
			todoList: &TodoList{
				Filename: "",
				Items:    make([]*Todo, 0),
				TagList:  make([]*Tag, 0),
			},
			want: 1,
		},
		{
			name: "with entries",
			todoList: &TodoList{
				Filename: "",
				Items:    make([]*Todo, 0),
				TagList: []*Tag{
					NewTag(1, "bla"),
					NewTag(3, "123"),
				},
			},
			want: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			nextId := test.todoList.NextTagId()
			if nextId != test.want {
				t.Errorf("want %v, got %v", test.want, nextId)
			}
		})
	}
}

func TestTodoList_TagsForTodo(t *testing.T) {
	tests := []struct {
		name       string
		want       []*Tag
		idToSearch int
		todoList   *TodoList
	}{
		{
			name: "do find",
			want: []*Tag{
				NewTag(1, "bla"),
			},
			idToSearch: 12,
			todoList: &TodoList{
				Filename: "",
				Items: []*Todo{
					{
						Id:     12,
						Label:  "My Todo",
						IsDone: false,
						Tags:   []TagId{1},
					},
				},
				TagList: []*Tag{
					NewTag(1, "bla"),
					NewTag(2, "123"),
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.todoList.TagsForTodo(test.idToSearch)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}

func TestTodoList_RemoveTag(t *testing.T) {
	tests := []struct {
		name        string
		wantItems   []*Todo
		wantTagList []*Tag
		todoList    *TodoList
		idToRemove  TagId
	}{
		{
			name:       "remove a tag",
			idToRemove: 25,
			wantItems: []*Todo{
				{
					Id:     1,
					Label:  "Test",
					IsDone: false,
					Tags:   make([]TagId, 0),
				},
			},
			wantTagList: make([]*Tag, 0),
			todoList: &TodoList{
				Filename: "",
				Items: []*Todo{
					{
						Id:     1,
						Label:  "Test",
						IsDone: false,
						Tags:   []TagId{25},
					},
				},
				TagList: []*Tag{
					NewTag(25, "test tag"),
				},
			},
		},
		{
			name: "remove a tag but not from todo if it does not have the id",
			wantItems: []*Todo{
				{
					Id:     1,
					Label:  "Test",
					IsDone: false,
					Tags:   []TagId{25},
				},
			},
			wantTagList: []*Tag{
				NewTag(25, "test tag"),
			},
			todoList: &TodoList{
				Filename: "",
				Items: []*Todo{
					{
						Id:     1,
						Label:  "Test",
						IsDone: false,
						Tags:   []TagId{25},
					},
				},
				TagList: []*Tag{
					NewTag(25, "test tag"),
					NewTag(1, "test tag #2"),
				},
			},
			idToRemove: 1,
		},
		{
			name:       "not remove a tag when tag can not be found",
			idToRemove: 800,
			wantItems: []*Todo{
				{
					Id:     1,
					Label:  "Test",
					IsDone: false,
					Tags:   []TagId{25},
				},
			},
			wantTagList: []*Tag{
				NewTag(25, "test tag"),
			},
			todoList: &TodoList{
				Filename: "",
				Items: []*Todo{
					{
						Id:     1,
						Label:  "Test",
						IsDone: false,
						Tags:   []TagId{25},
					},
				},
				TagList: []*Tag{
					NewTag(25, "test tag"),
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.todoList.RemoveTag(test.idToRemove)
			gotItems := test.todoList.Items
			gotTagList := test.todoList.TagList

			if !reflect.DeepEqual(gotItems, test.wantItems) {
				t.Errorf("want items %#v, got %#v", test.wantItems, gotItems)
			}

			if !reflect.DeepEqual(gotTagList, test.wantTagList) {
				t.Errorf("want tag list %#v, got %#v", test.wantTagList, gotTagList)
			}
		})
	}
}

func TestTodoList_TodosForTag(t *testing.T) {
	tests := []struct {
		name          string
		tagIdToSearch TagId
		wantTodoItems []*Todo
		todoList      *TodoList
	}{
		{
			name:          "find todos for tag",
			tagIdToSearch: 1,
			wantTodoItems: []*Todo{
				{123, "i should be in", false, []TagId{1}},
				{789, "i also should be in", false, []TagId{1}},
			},
			todoList: &TodoList{
				Filename: "test.json",
				Items: []*Todo{
					NewTodo("irrelevant #1", 1),
					{123, "i should be in", false, []TagId{1}},
					NewTodo("irrelevant #2", 1),
					{789, "i also should be in", false, []TagId{1}},
				},
				TagList: []*Tag{
					NewTag(1, "test tag"),
				},
			},
		},
		{
			name:          "empty slice for not found todos",
			tagIdToSearch: 1,
			wantTodoItems: make([]*Todo, 0),
			todoList: &TodoList{
				Filename: "",
				Items: []*Todo{
					NewTodo("test", 123),
				},
				TagList: []*Tag{
					NewTag(1, "test tag"),
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.todoList.TodosForTag(test.tagIdToSearch)

			if !reflect.DeepEqual(got.Items, test.wantTodoItems) {
				t.Errorf("want %#v, got %#v", test.wantTodoItems, got)
			}
		})
	}
}
