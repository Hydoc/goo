[![Tests](https://github.com/Hydoc/goo/actions/workflows/test.yaml/badge.svg)](https://github.com/Hydoc/goo/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/Hydoc/goo/graph/badge.svg?token=5TWYKUEG84)](https://codecov.io/gh/Hydoc/goo)
[![Go Report Card](https://goreportcard.com/badge/github.com/Hydoc/goo)](https://goreportcard.com/report/github.com/Hydoc/goo)

# Goo
A simple CLI todo list written in Go that supports different files.

Example
```shell
goo -f path/to/a/file.json --list
```
Or using the default file (`~/.goo.json`)
```shell
goo --add Hello World!
```

**Notes**
1. The file has to be json
2. If the file does not exist it gets created

## Installation
**Using go**
```shell
go install github.com/Hydoc/goo@latest
```
**Note** Don't forget to add the $HOME/go/bin to your $PATH
```shell
export PATH=$PATH:$HOME/go/bin && goo -a I did it!
```