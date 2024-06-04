package main

import (
	"os"

	"github.com/Hydoc/goo/internal/application"
	"github.com/Hydoc/goo/internal/view"
)

func main() {
	os.Exit(application.Main(view.New(os.Stdout), os.UserHomeDir))
}
