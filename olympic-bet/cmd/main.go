package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"olympic-bet/internal/start"
)

func main() {
	fyneApp := app.New()

	window := fyneApp.NewWindow("OlympicBET")

	window.Resize(fyne.NewSize(800, 600))

	window.SetContent(start.InitPage())

	window.ShowAndRun()

}
