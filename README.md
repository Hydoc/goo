[![Tests](https://github.com/Hydoc/goo/actions/workflows/test.yaml/badge.svg)](https://github.com/Hydoc/goo/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/Hydoc/goo/graph/badge.svg?token=5TWYKUEG84)](https://codecov.io/gh/Hydoc/goo)
[![Go Report Card](https://goreportcard.com/badge/github.com/Hydoc/goo)](https://goreportcard.com/report/github.com/Hydoc/goo)

# Goo

A simple CLI todo list written in Go that supports different files, editing and tags!

Example

```shell
goo -f path/to/a/file.json list
```

Or using the default file (`~/.goo.json`)

```shell
goo add Hello World!
```

## Installation

**Using go**

```shell
go install github.com/Hydoc/goo@latest
```

**Note** Don't forget to add the $HOME/go/bin to your $PATH

```shell
export PATH=$PATH:$HOME/go/bin && goo add I did it!
```

## Commands and flags

### Flags

* `-f, --file`: Path to the file to use (defaults to `~/.goo.json`, if it does not exist it gets created (**has to be
  json**). The flag should always be **before** the subcommands
    * `goo -f path/to/my-file.json add Hello World!`

### Commands

| Command   | Description                                                                                                                                                                                                                                                                                                                                                                  | Example                                                                            |
|-----------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------|
| `add`     | Adds a new todo to the given list                                                                                                                                                                                                                                                                                                                                            | `goo add Hello World!`                                                             |
| `rm`      | Removes a todo by its id                                                                                                                                                                                                                                                                                                                                                     | `goo rm 1`                                                                         |
| `edit`    | Edits a todo by its id and a new label. Use curly braces (`{}`) to insert old value                                                                                                                                                                                                                                                                                          | `goo edit 1 {} World`                                                              |
| `toggle`  | Toggle the done state of a todo by its id                                                                                                                                                                                                                                                                                                                                    | `goo toggle 1`                                                                     |
| `list`    | List all todos in the file                                                                                                                                                                                                                                                                                                                                                   | `goo list`                                                                         |
| `clear`   | Clear the whole list                                                                                                                                                                                                                                                                                                                                                         | `goo clear`                                                                        |
| `swap`    | Swap the label of two todos by their id                                                                                                                                                                                                                                                                                                                                      | `goo swap 1 2`                                                                     |
| `tags`    | Shows all tags<br/>Has two optional flags `-tid` and `-id`<br/>Providing `-tid <Id of todo>` to show all tags on that todo<br/>Providing `-id <Id of tag>` to show all todos for that tag                                                                                                                                                                                    | `goo tags`<br/>`goo tags -tid 1`<br/>`goo tags -id 1`                              |
| `tag`     | Tag a todo with the provided tag and todo id. The first argument is the tag id, the second the todo id<br/>Has two optional flags `-c` and `-rm`<br/>Providing `-c <Label of a tag to create>` to create a new tag with that label<br/>Providing `-rm <Id of tag>` to remove the tag completely<br/>Providing `-rm <Id of tag> <Id of todo>` to remove the tag from the todo | `goo tag 1 1`<br/>`goo tag -c Hello Tag`<br/>`goo tag -rm 1`<br/>`goo tag -rm 1 1` |                                                                                                                                              
| `version` | Prints the current goo version                                                                                                                                                                                                                                                                                                                                               | `goo version`                                                                      |

