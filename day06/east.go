package day06

type East struct {
	value string
}

func (d East) getValue() string {
	return d.value
}

func (d East) move(grid Grid, position Position) Position {
	return Position{
		x:         position.x + 1,
		y:         position.y,
		direction: east,
	}
}

func (d East) rotate(position Position) Position {
	position.direction = south
	return position
}
