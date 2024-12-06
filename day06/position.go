package day06

type Position struct {
	x         int
	y         int
	direction Direction
}

func (position Position) move(grid Grid) Position {
	return position.direction.move(grid, position)
}

func (position Position) rotate() Position {
	return position.direction.rotate(position)
}
