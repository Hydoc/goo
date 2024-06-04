package command

import (
	"flag"
	"reflect"
	"testing"

	"github.com/Hydoc/goo/internal/model"
)

func TestFabricate(t *testing.T) {
	file := flag.String("file", "", "Path to a file to use (has to be json, if the file does not exist it gets created)")
	flag.StringVar(file, "f", "", "Path to a file to use (has to be json, if the file does not exist it gets created)")

	list := flag.NewFlagSet("list", flag.ExitOnError)
	add := flag.NewFlagSet("add", flag.ExitOnError)
	doDelete := flag.NewFlagSet("rm", flag.ExitOnError)
	toggle := flag.NewFlagSet("toggle", flag.ExitOnError)
	edit := flag.NewFlagSet("edit", flag.ExitOnError)
	doClear := flag.NewFlagSet("clear", flag.ExitOnError)
	swap := flag.NewFlagSet("swap", flag.ExitOnError)
	tags := flag.NewFlagSet("tags", flag.ExitOnError)
	showTagsOnTodo := tags.Int("tid", 0, "<ID of todo> show all tags on this todo")
	showTodosForTag := tags.Int("id", 0, "<ID of tag> show all todos with this tag")
	tag := flag.NewFlagSet("tag", flag.ExitOnError)
	tagRm := tag.Bool("rm", false, "delete a tag")
	tagAdd := tag.Bool("c", false, "add a tag")
	versionFlag := flag.NewFlagSet("version", flag.ExitOnError)
	version := "1.4.3"

	tests := []struct {
		name     string
		want     reflect.Type
		flagArgs []string
		todoList *model.TodoList
		err      error
	}{
		{
			name:     "ListTodos",
			want:     reflect.TypeOf(&ListTodos{}),
			flagArgs: []string{"list"},
			todoList: &model.TodoList{},
			err:      nil,
		},
		{
			name:     "AddTodo",
			want:     reflect.TypeOf(&AddTodo{}),
			flagArgs: []string{"add", "Hi"},
			todoList: &model.TodoList{},
			err:      nil,
		},
		{
			name:     "DeleteTodo",
			want:     reflect.TypeOf(&DeleteTodo{}),
			flagArgs: []string{"rm", "1"},
			todoList: &model.TodoList{
				Filename: "",
				Items:    []*model.Todo{model.NewTodo("test", 1)},
				TagList:  make([]*model.Tag, 0),
			},
			err: nil,
		},
		{
			name:     "ToggleTodo",
			want:     reflect.TypeOf(&ToggleTodo{}),
			flagArgs: []string{"toggle", "1"},
			todoList: &model.TodoList{
				Filename: "",
				Items:    []*model.Todo{model.NewTodo("test", 1)},
				TagList:  make([]*model.Tag, 0),
			},
			err: nil,
		},
		{
			name:     "EditTodo",
			want:     reflect.TypeOf(&EditTodo{}),
			flagArgs: []string{"edit", "1 Hello World"},
			todoList: &model.TodoList{
				Filename: "",
				Items:    []*model.Todo{model.NewTodo("test", 1)},
				TagList:  make([]*model.Tag, 0),
			},
			err: nil,
		},
		{
			name:     "Clear",
			want:     reflect.TypeOf(&Clear{}),
			flagArgs: []string{"clear"},
			todoList: nil,
			err:      nil,
		},
		{
			name:     "Swap",
			want:     reflect.TypeOf(&Swap{}),
			flagArgs: []string{"swap", "1 2"},
			todoList: &model.TodoList{
				Filename: "",
				Items:    []*model.Todo{model.NewTodo("test", 1), model.NewTodo("ABC", 2)},
				TagList:  make([]*model.Tag, 0),
			},
			err: nil,
		},
		{
			name:     "ListTags",
			want:     reflect.TypeOf(&ListTags{}),
			flagArgs: []string{"tags"},
			todoList: &model.TodoList{
				Filename: "",
				Items:    make([]*model.Todo, 0),
				TagList:  make([]*model.Tag, 0),
			},
			err: nil,
		},
		{
			name:     "TagTodo",
			want:     reflect.TypeOf(&TagTodo{}),
			flagArgs: []string{"tag", "1 1"},
			todoList: &model.TodoList{
				Filename: "",
				Items: []*model.Todo{
					model.NewTodo("Test", 1),
				},
				TagList: []*model.Tag{
					model.NewTag(1, "test tag"),
				},
			},
			err: nil,
		},
		{
			name:     "ListTodosForTag",
			want:     reflect.TypeOf(&ListTodosForTag{}),
			flagArgs: []string{"tags", "-id", "1"},
			todoList: &model.TodoList{
				Filename: "",
				Items: []*model.Todo{
					{1, "Test", false, []model.TagId{1}},
				},
				TagList: []*model.Tag{
					model.NewTag(1, "test"),
				},
			},
			err: nil,
		},
		{
			name:     "ListTagsOnTodo",
			want:     reflect.TypeOf(&ListTagsOnTodo{}),
			flagArgs: []string{"tags", "-tid", "1"},
			todoList: &model.TodoList{
				Filename: "",
				Items: []*model.Todo{
					{1, "Test", false, []model.TagId{1}},
				},
				TagList: make([]*model.Tag, 0),
			},
			err: nil,
		},
		{
			name:     "AddTag",
			want:     reflect.TypeOf(&AddTag{}),
			flagArgs: []string{"tag", "-c", "Hello World"},
			todoList: &model.TodoList{},
			err:      nil,
		},
		{
			name:     "RemoveTagFromTodo",
			want:     reflect.TypeOf(&RemoveTagFromTodo{}),
			flagArgs: []string{"tag", "-rm", "1", "1"},
			todoList: &model.TodoList{
				Filename: "",
				Items: []*model.Todo{
					{1, "Test", false, []model.TagId{1}},
				},
				TagList: []*model.Tag{
					model.NewTag(1, "test tag"),
				},
			},
			err: nil,
		},
		{
			name:     "RemoveTag",
			want:     reflect.TypeOf(&RemoveTag{}),
			flagArgs: []string{"tag", "-rm", "1"},
			todoList: &model.TodoList{
				Filename: "",
				Items:    make([]*model.Todo, 0),
				TagList: []*model.Tag{
					model.NewTag(1, "test tag"),
				},
			},
			err: nil,
		},
		{
			name:     "error when command was not found",
			want:     nil,
			flagArgs: []string{"shouldnotwork"},
			todoList: &model.TodoList{},
			err:      ErrCommandNotFound,
		},
		{
			name:     "PrintVersion",
			want:     reflect.TypeOf(&PrintVersion{}),
			flagArgs: []string{"version"},
			todoList: &model.TodoList{},
			err:      nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Fabricate(
				test.todoList,
				newDummyView(),
				test.flagArgs,
				list,
				add,
				doDelete,
				toggle,
				edit,
				doClear,
				swap,
				tags,
				showTagsOnTodo,
				showTodosForTag,
				tag,
				tagRm,
				tagAdd,
				versionFlag,
				version,
			)

			if err != nil && !reflect.DeepEqual(err, test.err) {
				t.Errorf("want err %v, got %v", test.err, err)
			}

			if reflect.TypeOf(got) != test.want {
				t.Errorf("want %#v, got %#v", test.want, got)
			}
		})
	}
}
