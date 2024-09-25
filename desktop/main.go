package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create the app
	myApp := app.New()
	myWindow := myApp.NewWindow("OlympicBET")

	// Create the header content
	headerContent := container.NewHBox(
		canvas.NewText("OlympicBET", color.White),
		layout.NewSpacer(),
		widget.NewHyperlink("Jogos", nil),
		widget.NewHyperlink("Suas Apostas", nil),
		layout.NewSpacer(),
		widget.NewLabelWithStyle("Nome de usuário", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewLabel("R$ 1000"),
	)

	// Create a background rectangle for the header
	headerBackground := canvas.NewRectangle(color.RGBA{50, 50, 50, 255}) // Dark gray background

	// Combine background and header content using container.NewMax
	header := container.NewStack(
		headerBackground,
		container.NewPadded(headerContent),
	)

	// Create the main content area for the games list
	mainContent := container.NewVBox(
		widget.NewLabel("Jogos"),
		widget.NewButton("Jogo 1", nil),
		widget.NewButton("Jogo 2", nil),
		widget.NewButton("Jogo 3", nil),
		widget.NewButton("Jogo 4", nil),
		widget.NewButton("Jogo 5", nil),
	)

	// Create the main page container
	mainPage := container.NewBorder(header, nil, nil, nil, mainContent)

	// Create the Esportes page content with a back button
	esportesContent := widget.NewLabel("Esportes Page")

	backButton := widget.NewButton("Back", func() {
		// Set the content back to the main page when clicking "Back"
		myWindow.SetContent(mainPage)
	})
	esportesPage := container.NewVBox(esportesContent, backButton)
	// Create the side menu with a button for "Esportes"
	sideMenu := container.NewVBox(
		widget.NewButton("Esportes", func() {
			// Navigate to the Esportes page when clicking "Esportes"
			myWindow.SetContent(esportesPage)
		}),
	)

	// Arrange the side menu and the main content
	content := container.NewHSplit(
		sideMenu,
		mainContent,
	)
	content.SetOffset(0.2) // Set ratio for the split

	// Create the final layout with the header and the content
	finalLayout := container.NewBorder(header, nil, nil, nil, content)

	// Set the initial content of the window to the main page
	myWindow.SetContent(finalLayout)
	myWindow.Resize(fyne.NewSize(800, 600)) // Set window size
	myWindow.ShowAndRun()
}
