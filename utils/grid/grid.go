package grid

type Grid interface {
	rotate() Grid
}

type StringGrid [][]string

func (grid StringGrid) Rotate() StringGrid {
	rotated := StringGrid{}
	yMax := len(grid) - 1

	for x := range grid[0] {
		column := []string{}

		for y := yMax; y >= 0; y-- {
			column = append(column, grid[y][x])
		}

		rotated = append(rotated, column)
	}

	return rotated
}
