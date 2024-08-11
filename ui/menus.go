package ui

import (
	"errors"
	"image"
	"image/png"
	"os"
	"pixl/util"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func saveFileDialog(app *AppInit) {
	dialog.ShowFileSave(func(uri fyne.URIWriteCloser, e error) {
		if uri == nil {
			return
		} else {
			err := png.Encode(uri, app.PixlCanvas.PixelData)
			if err != nil {
				dialog.NewError(err, app.PixlWindow)
				return
			}
			app.State.SetFilePath(uri.URI().Path())
		}
	}, app.PixlWindow)
}

func BuildSaveAsMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Save As...", func() {
		saveFileDialog(app)
	})
}

func BuildSaveMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Save", func() {
		if app.State.FilePath == "" {
			saveFileDialog(app)
		} else {
			tryClose := func(fh *os.File) {
				err := fh.Close()
				if err != nil {
					dialog.NewError(err, app.PixlWindow)
				}
			}

			fh, err := os.Create(app.State.FilePath)
			defer tryClose(fh)
			if err != nil {
				dialog.NewError(err, app.PixlWindow)
				return
			}

			err = png.Encode(fh, app.PixlCanvas.PixelData)
			if err != nil {
				dialog.NewError(err, app.PixlWindow)
				return
			}
		}
	})
}

func BuildNewMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("New", func() {
		sizeValidator := func(s string) error {
			size, err := strconv.Atoi(s)
			if err != nil {
				return errors.New("must be positive integer")
			}

			if size <= 0 {
				return errors.New("must be > 0")
			}
			return nil
		}

		widthEntry := widget.NewEntry()
		widthEntry.Validator = sizeValidator

		heightEntry := widget.NewEntry()
		heightEntry.Validator = sizeValidator

		widthFormEntry := widget.NewFormItem("Width", widthEntry)
		heightFormEntry := widget.NewFormItem("Height", heightEntry)

		formItems := []*widget.FormItem{widthFormEntry, heightFormEntry}
		dialog.ShowForm("New Image", "Create", "Cancel", formItems, func(ok bool) {
			if ok {
				pixelWidth := 0
				pixelHeight := 0
				if widthEntry.Validate() != nil {
					dialog.ShowError(errors.New("invalid width"), app.PixlWindow)
				} else {
					pixelWidth, _ = strconv.Atoi(widthEntry.Text)
				}

				if heightEntry.Validate() != nil {
					dialog.ShowError(errors.New("invalid height"), app.PixlWindow)
				} else {
					pixelHeight, _ = strconv.Atoi(heightEntry.Text)
				}
				app.PixlCanvas.NewDrawing(pixelWidth, pixelHeight)
			}
		}, app.PixlWindow)
	})
}

func BuildOpenMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Open...", func() {
		dialog.ShowFileOpen(func(uri fyne.URIReadCloser, e error) {
			if uri == nil {
				return
			} else {
				img, _, err := image.Decode(uri)
				if err != nil {
					dialog.ShowError(err, app.PixlWindow)
					return
				}
				app.PixlCanvas.LoadImage(img)
				app.State.SetFilePath(uri.URI().Path())

				colors := util.GetImageColors(img)
				i := 0
				for color := range colors {
					if i == len(app.Swatches) {
						break
					}
					app.Swatches[i].SetColor(color)
					i++
				}
			}
		}, app.PixlWindow)
	})
}

func BuildMenus(app *AppInit) *fyne.Menu {
	return fyne.NewMenu(
		"File",
		BuildNewMenu(app),
		BuildOpenMenu(app),
		BuildSaveMenu(app),
		BuildSaveAsMenu(app),
	)
}

func SetupMenus(app *AppInit) {
	menus := BuildMenus(app)
	mainMenu := fyne.NewMainMenu(menus)
	app.PixlWindow.SetMainMenu(mainMenu)
}
