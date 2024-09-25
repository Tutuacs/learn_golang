package start

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"olympic-bet/pkg/colors"
)

func InitPage() (init_container *fyne.Container) {

	// Header content
	headerText := canvas.NewText("OlympicBET", colors.Colors.White)
	headerText.TextSize = 24

	header := container.NewVBox(
		layout.NewSpacer(), // Spacer for pushing header up
		container.NewHBox(
			layout.NewSpacer(), // Adds space to the left of the header
			headerText,         // Centered header text
			layout.NewSpacer(), // Adds space to the right of the header
		),
		layout.NewSpacer(), // Spacer for pushing content down after header
	)

	minSize := fyne.NewSize(200, 40) // Width: 200, Height: 40

	// Form content (Login form)
	emailEntry := widget.NewEntry()
	emailEntry.MinSize()
	emailEntry.SetPlaceHolder("Email")

	emailField := container.New(layout.NewGridWrapLayout(minSize), emailEntry)

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Senha")

	passwordField := container.New(layout.NewGridWrapLayout(minSize), passwordEntry)

	loginBTN := widget.NewButtonWithIcon("Login", theme.LoginIcon(), func() {
		// Login action
	})
	loginBTN.Importance = widget.HighImportance // Highlights the button
	loginBG := canvas.NewRectangle(colors.Colors.Red_primary)

	loginButton := container.NewStack(
		loginBG,
		loginBTN,
	)

	form := container.NewVBox(
		emailField,
		passwordField,
		layout.NewSpacer(), // Adds space between form and button
		loginButton,
	)

	// Center content layout
	centerBG := canvas.NewRectangle(colors.Colors.Gray_secondary) // Background rectangle
	centerContent := container.NewCenter(
		container.NewVBox(
			layout.NewSpacer(), // Adds space at the top
			form,
			layout.NewSpacer(), // Adds space at the bottom
		),
	)

	center := container.NewStack(centerBG, centerContent)

	// Combine header and center
	init_container = container.NewBorder(header, nil, nil, nil, center)

	return
}
