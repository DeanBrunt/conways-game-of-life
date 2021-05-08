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

	i := 0
	j := 0
	for !win.Closed() {
		renderer.RenderNextGridFrame(grid)
		grid[i][j] = !grid[i][j]
		if j+1 == len(grid[0]) {
			j = 0
			if i+1 == len(grid) {
				i = 0
				continue
			}
			i++
			continue
		}
		j++
	}
}
