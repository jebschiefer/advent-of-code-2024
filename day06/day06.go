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

func Day06(grid Grid) int {
	position := getStartingposition(grid)

	for {
		done, currentposition := moveAndMark(&grid, position)
		position = currentposition

		if done {
			break
		}
	}

	return countVisited(grid)
}

func moveAndMark(gridPointer *Grid, currentPosition Position) (done bool, position Position) {
	grid := *gridPointer
	xMin, yMin := 0, 0
	xMax, yMax := len(grid[0])-1, len(grid)-1

	newPosition := currentPosition.move(grid)

	if newPosition.x >= xMin && newPosition.x <= xMax && newPosition.y >= yMin && newPosition.y <= yMax {
		if grid[newPosition.y][newPosition.x] == "#" {
			newPosition = currentPosition.rotate()
		} else {
			grid[currentPosition.y][currentPosition.x] = "X"
		}

		grid[newPosition.y][newPosition.x] = newPosition.direction.getValue()

		return false, newPosition
	}

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
	count := 1

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == "X" {
				count++
			}
		}
	}

	return count
}
