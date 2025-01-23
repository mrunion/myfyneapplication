package main

//go:generate fyne bundle --package ui -output internal/ui/bundle.go assets/Icon.png

import (
	"log"

	"github.com/example/myfyneapplication/internal/ui"
)

func main() {
	// Set the logger prefix
	log.SetPrefix("My Fyne Application")
	ui.CreateMyFyneApplication()
}
