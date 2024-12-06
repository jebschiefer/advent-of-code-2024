package day04

import "aoc2024/utils/grid"

type CountFunc func(grid grid.StringGrid) int

func Day04(grid grid.StringGrid, doCount CountFunc) int {
	count := 0

	for i := 0; i < 2; i++ {
		count += doCount(grid)
		grid = grid.Rotate()
	}

	return count
}

func countXmas(grid grid.StringGrid) int {
	count := 0
	xMax, yMax := len(grid[0])-1, len(grid)-1

	for y := range grid {
		for x := range grid[y] {
			if x+3 <= xMax {
				if grid[y][x] == "X" && grid[y][x+1] == "M" && grid[y][x+2] == "A" && grid[y][x+3] == "S" {
					count++
				} else if grid[y][x] == "S" && grid[y][x+1] == "A" && grid[y][x+2] == "M" && grid[y][x+3] == "X" {
					count++
				}

				if y+3 <= yMax {
					if grid[y][x] == "X" && grid[y+1][x+1] == "M" && grid[y+2][x+2] == "A" && grid[y+3][x+3] == "S" {
						count++
					} else if grid[y][x] == "S" && grid[y+1][x+1] == "A" && grid[y+2][x+2] == "M" && grid[y+3][x+3] == "X" {
						count++
					}
				}
			}
		}
	}

	return count
}

func countCrossMas(grid grid.StringGrid) int {
	count := 0
	xMax, yMax := len(grid[0])-1, len(grid)-1

	for y := range grid {
		for x := range grid[y] {
			if x+2 <= xMax && y+2 <= yMax {
				if grid[y][x] == "M" && grid[y+1][x+1] == "A" && grid[y+2][x+2] == "S" && grid[y+2][x] == "M" && grid[y][x+2] == "S" {
					count++
				} else if grid[y][x] == "S" && grid[y+1][x+1] == "A" && grid[y+2][x+2] == "M" && grid[y+2][x] == "S" && grid[y][x+2] == "M" {
					count++
				}
			}
		}
	}

	return count
}
