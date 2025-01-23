package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func CreateMyFyneApplication() {
	a := app.NewWithID("com.example.myfyneapplication")
	w := a.NewWindow("My Fyne Application")
	w.Resize(fyne.NewSize(1200, 900))
	w.CenterOnScreen()

	// Create the Gui
	gui := NewGui(&w)
	w.SetContent(gui.makeGui())

	w.ShowAndRun()
}
