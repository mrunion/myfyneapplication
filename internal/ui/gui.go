package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/example/myfyneapplication/internal/ui/dialogs"
)

type gui struct {
	win *fyne.Window
}

func NewGui(win *fyne.Window) *gui {
	return &gui{
		win: win,
	}
}

func (g *gui) makeGui() fyne.CanvasObject {
	// TODO: make the ui

	return container.NewBorder(
		nil, //top
		widget.NewButton("About", func() { g.showAboutDialog() }), //bottom
		nil, //left
		nil, //right
		container.NewCenter(widget.NewLabel("We're just getting started...")))
}

func (g *gui) showAboutDialog() {
	about := dialogs.NewAbout(g.win, "About", resourceIconPng)
	about.Show()
}
