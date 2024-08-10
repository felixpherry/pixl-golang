package main

import (
	"image/color"
	"pixl/apptype"
	"pixl/pxcanvas"
	"pixl/swatch"
	"pixl/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	pixlApp := app.New()
	pixlWindow := pixlApp.NewWindow("pixl")

	state := apptype.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	pxCanvasConfig := &apptype.PxCanvasConfig{
		DrawingArea:  fyne.NewSize(20, 20),
		CanvasOffset: fyne.NewPos(0, 0),
		PxRows:       10,
		PxCols:       10,
		PxSize:       20,
	}
	pxCanvas := pxcanvas.NewPxCanvas(&state, *pxCanvasConfig)

	appInit := ui.AppInit{
		PixlCanvas: pxCanvas,
		PixlWindow: pixlWindow,
		State:      &state,
		Swatches:   make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)
	appInit.PixlWindow.ShowAndRun()
}
