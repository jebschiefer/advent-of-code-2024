package day06

type South struct {
	value string
}

func (d South) getValue() string {
	return d.value
}

func (d South) move(grid Grid, position Position) Position {
	return Position{
		x:         position.x,
		y:         position.y + 1,
		direction: south,
	}
}

func (d South) rotate(position Position) Position {
	position.direction = west
	return position
}
