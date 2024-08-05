package ui

import "fyne.io/fyne/v2/container"

func Setup(app *AppInit) {
	swatchesContainer := BuildSwatches(app)
	colorpicker := SetupColorPicker(app)

	appContainer := container.NewBorder(nil, swatchesContainer, nil, colorpicker)
	app.PixlWindow.SetContent(appContainer)
}
