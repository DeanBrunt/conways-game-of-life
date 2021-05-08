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

	verticalCursor := verticalGridSpacing + float64(lineThickness)
	for i := 0; i < horizontalLinesRequired; i++ {
		r.imd.Color = colornames.White
		r.imd.EndShape = imdraw.RoundEndShape
		r.imd.Push(pixel.V(0, verticalCursor), pixel.V(r.win.Bounds().W(), verticalCursor))
		r.imd.Line(float64(lineThickness))
		verticalCursor += verticalGridSpacing + float64(lineThickness)
	}

	verticalLinesRequired := len(gridFrame[0]) - 1
	heightRequiredForVerticLines := verticalLinesRequired * lineThickness

	remainingWidth := r.win.Bounds().W() - float64(heightRequiredForVerticLines)
	horizontalGridSpacing := remainingWidth / float64(len(gridFrame[0]))

	horizontalCursor := horizontalGridSpacing + float64(lineThickness)
	for i := 0; i < horizontalLinesRequired; i++ {
		r.imd.Color = colornames.White
		r.imd.EndShape = imdraw.RoundEndShape
		r.imd.Push(pixel.V(horizontalCursor, 0), pixel.V(horizontalCursor, r.win.Bounds().H()))
		r.imd.Line(float64(lineThickness))
		horizontalCursor += horizontalGridSpacing + float64(lineThickness)
	}

	r.win.Clear(colornames.Black)
	r.imd.Draw(r.win)
	r.win.Update()
}
