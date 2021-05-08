package main

import (
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

	grid := grid.NewGrid(25, 25)
	renderer := render.New(win)
	grid[0][1] = true
	grid[1][2] = true
	grid[2][0] = true
	grid[2][1] = true
	grid[2][2] = true

	for !win.Closed() {
		renderer.RenderNextGridFrame(grid)
	}
}
