package day06

type West struct {
	value string
}

func (d West) getValue() string {
	return d.value
}

func (d West) move(grid Grid, position Position) Position {
	return Position{
		x:         position.x - 1,
		y:         position.y,
		direction: west,
	}
}

func (d West) rotate(position Position) Position {
	position.direction = north
	return position
}
