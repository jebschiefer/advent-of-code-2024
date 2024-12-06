package day06

type Grid [][]string

type Direction interface {
	move(grid Grid, position Position) Position
	rotate(position Position) Position
	getValue() string
}

var defaultPosition = Position{-1, -1, north}

var north = North{value: "^"}
var south = South{value: "v"}
var east = East{value: ">"}
var west = West{value: "<"}

var directions = map[string]Direction{
	"^": north,
	"v": south,
	">": east,
	"<": west,
}

func Day06Part1(grid Grid) int {
	markedGrid := markGrid(grid)
	return countVisited(markedGrid)
}

func Day06Part2(grid Grid) int {
	markedGrid := markGrid(grid)
	position := getStartingposition(grid)
	cycleCount := 0

	for y := range grid {
		for x := range grid[0] {
			if grid[y][x] == "." && markedGrid[y][x] == "X" {
				grid[y][x] = "O"

				if hasCycle(grid, position) {
					cycleCount++
				}

				grid[y][x] = "."
			}
		}

	}

	return cycleCount
}

func markGrid(grid Grid) Grid {
	markedGrid := copyGrid(grid)
	position := getStartingposition(markedGrid)

	for {
		done, currentposition := moveAndMark(&markedGrid, position)
		position = currentposition

		if done {
			break
		}
	}

	return markedGrid
}

func moveAndMark(gridPointer *Grid, currentPosition Position) (done bool, position Position) {
	grid := *gridPointer
	xMin, yMin := 0, 0
	xMax, yMax := len(grid[0])-1, len(grid)-1

	newPosition := currentPosition.move(grid)

	if newPosition.x >= xMin && newPosition.x <= xMax && newPosition.y >= yMin && newPosition.y <= yMax {
		if grid[newPosition.y][newPosition.x] == "#" || grid[newPosition.y][newPosition.x] == "O" {
			newPosition = currentPosition.rotate()
		} else {
			grid[currentPosition.y][currentPosition.x] = "X"
		}

		grid[newPosition.y][newPosition.x] = newPosition.direction.getValue()

		return false, newPosition
	}

	grid[currentPosition.y][currentPosition.x] = "X"

	return true, defaultPosition
}

func getStartingposition(grid Grid) Position {
	position := defaultPosition

	for y := range grid {
		for x := range grid[y] {
			direction, ok := directions[grid[y][x]]

			if ok {
				position = Position{x, y, direction}
			}
		}
	}

	return position
}

func countVisited(grid Grid) int {
	count := 0

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == "X" {
				count++
			}
		}
	}

	return count
}

func isCycle(seenPositions []Position, position Position) bool {
	for _, seen := range seenPositions {
		if seen.equals(position) {
			return true
		}
	}

	return false
}

func hasCycle(grid Grid, position Position) bool {
	grid = copyGrid(grid)
	seenPositions := []Position{}

	for {
		if isCycle(seenPositions, position) {
			return true
		}

		seenPositions = append(seenPositions, position)

		done, currentposition := moveAndMark(&grid, position)
		position = currentposition

		if done {
			return false
		}
	}
}

func copyGrid(grid Grid) Grid {
	copy := Grid{}

	for y := range grid {
		row := []string{}
		row = append(row, grid[y]...)
		copy = append(copy, row)
	}

	return copy
}
