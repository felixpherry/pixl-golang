package pxcanvas

import (
	"math"

	"fyne.io/fyne/v2"
)

func (pxCanvas *PxCanvas) Pan(previousCoord, currentCoord fyne.PointEvent) {
	xDiff := currentCoord.Position.X - previousCoord.Position.X
	yDiff := currentCoord.Position.Y - previousCoord.Position.Y

	pxCanvas.CanvasOffset.X += xDiff
	pxCanvas.CanvasOffset.Y += yDiff
	pxCanvas.Refresh()
}

func (pxCanvas *PxCanvas) scale(direction int) {
	switch {
	case direction > 0:
		pxCanvas.PxSize += 1
	case direction < 0:
		pxCanvas.PxSize = int(math.Max(float64(pxCanvas.PxSize-1), 0.0))
	default:
		pxCanvas.PxSize = 10
	}
}
