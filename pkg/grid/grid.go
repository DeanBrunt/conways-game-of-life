package grid

type (
	Row  = []bool
	Grid = []Row
)

func NewGrid(height, width int) Grid {
	grid := make(Grid, height)
	for i := 0; i < height; i++ {
		grid[i] = make(Row, width)
	}

	return grid
}
