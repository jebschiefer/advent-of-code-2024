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

func (position Position) equals(position2 Position) bool {
	return position.x == position2.x &&
		position.y == position2.y &&
		position.direction == position2.direction
}
