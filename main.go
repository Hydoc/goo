package main

import (
	"github.com/Hydoc/goo/internal/application"
	"github.com/Hydoc/goo/internal/view"
	"os"
)

func main() {
	os.Exit(application.Main(view.New(os.Stdout), os.UserHomeDir))
}
