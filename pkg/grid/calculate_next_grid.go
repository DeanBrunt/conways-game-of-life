package grid

func CalculateNextGrid(currentGrid Grid) Grid {
	newGrid := NewGrid(len(currentGrid), len(currentGrid[0]))
	for i := 0; i < len(currentGrid); i++ {
		for j := 0; j < len(currentGrid[i]); j++ {
			newGrid[i][j] = determineNextState(i, j, currentGrid)
		}
	}

	return newGrid
}

func determineNextState(i, j int, grid Grid) bool {
	var neighbourCoords [][2]int
	for k := -1; k <= 1; k++ {
		for l := -1; l <= 1; l++ {
			if k == 0 && l == 0 {
				continue
			}
			coords := [2]int{}

			switch {
			case i+k < 0:
				coords[0] = len(grid) - 1
			case i+k >= len(grid):
				coords[0] = 0
			default:
				coords[0] = i + k
			}

			switch {
			case j+l < 0:
				coords[1] = len(grid[0]) - 1
			case j+l >= len(grid[0]):
				coords[1] = 0
			default:
				coords[1] = j + l
			}

			neighbourCoords = append(neighbourCoords, coords)
		}
	}

	neighbourCount := 0
	for _, neighbourCoord := range neighbourCoords {
		if grid[neighbourCoord[0]][neighbourCoord[1]] {
			neighbourCount++
		}
	}

	thisCellState := grid[i][j]
	if thisCellState && (neighbourCount == 2 || neighbourCount == 3) {
		return true
	}

	if !thisCellState && neighbourCount == 3 {
		return true
	}

	return false
}
