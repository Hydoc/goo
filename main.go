package main

import (
	"github.com/Hydoc/goo/internal/application"
	"github.com/Hydoc/goo/internal/view"
	"log"
	"os"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	application.Main(view.New(os.Stdout), home)
}
