package main

import (
	"time"

	"github.com/DeanBrunt/conways-game-of-life/pkg/grid"
	"github.com/DeanBrunt/conways-game-of-life/pkg/render"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Conway's Game of Life",
		Bounds: pixel.R(0, 0, 1024, 1024),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	curGrid := grid.NewGrid(50, 50)
	renderer := render.New(win)
	curGrid[0][1] = true
	curGrid[1][2] = true
	curGrid[2][0] = true
	curGrid[2][1] = true
	curGrid[2][2] = true

	for !win.Closed() {
		renderer.RenderGridFrame(curGrid)
		curGrid = grid.CalculateNextGrid(curGrid)
		time.Sleep(100 * time.Millisecond)
	}
}
