[![Tests](https://github.com/Hydoc/goo/actions/workflows/test.yaml/badge.svg)](https://github.com/Hydoc/goo/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/Hydoc/goo/graph/badge.svg?token=5TWYKUEG84)](https://codecov.io/gh/Hydoc/goo)
[![Go Report Card](https://goreportcard.com/badge/github.com/Hydoc/goo)](https://goreportcard.com/report/github.com/Hydoc/goo)

# Goo
A simple CLI todo list written in Go that supports different files.

Example
```shell
goo -file path/to/a/file.json
```

**Notes**
1. The file has to be json
2. If the file does not exist it gets created


Saving only happens when using the `quit` command.

The following commands are currently supported:
1. `add`
2. `delete`
3. `edit`
4. `toggle`
5. `undo`
6. `help`
7. `quit`
