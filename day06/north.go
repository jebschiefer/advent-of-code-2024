package day06

type North struct {
	value string
}

func (d North) getValue() string {
	return d.value
}

func (d North) move(grid Grid, position Position) Position {
	return Position{
		x:         position.x,
		y:         position.y - 1,
		direction: north,
	}
}

func (d North) rotate(position Position) Position {
	position.direction = east
	return position
}
