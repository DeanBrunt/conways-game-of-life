package render

import (
	"github.com/DeanBrunt/conways-game-of-life/pkg/grid"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Renderer struct {
	imd *imdraw.IMDraw
	win *pixelgl.Window
}

func New(win *pixelgl.Window) *Renderer {
	return &Renderer{
		imd: imdraw.New(nil),
		win: win,
	}
}

func (r *Renderer) RenderNextGridFrame(gridFrame grid.Grid) {
	lineThickness := 1
	horizontalLinesRequired := len(gridFrame) - 1
	heightRequiredForHorizLines := horizontalLinesRequired * lineThickness

	remainingHeight := r.win.Bounds().H() - float64(heightRequiredForHorizLines)
	verticalGridSpacing := remainingHeight / float64(len(gridFrame))

	verticalCursor := verticalGridSpacing
	for i := 0; i < horizontalLinesRequired; i++ {
		r.imd.Color = colornames.White
		r.imd.EndShape = imdraw.RoundEndShape
		r.imd.Push(pixel.V(0, verticalCursor), pixel.V(600, verticalCursor))
		r.imd.Line(float64(lineThickness))
		verticalCursor += verticalGridSpacing
	}

	r.win.Clear(colornames.Black)
	r.imd.Draw(r.win)
	r.win.Update()
}
