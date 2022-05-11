package models

import "fmt"

type Rover struct {
	x         int
	y         int
	direction direction
}

func NewRover() Rover {
	direction, _ := NewDirection("N")
	return Rover{x: 0, y: 0, direction: direction}
}

func (r *Rover) Move(val int) {
	switch r.direction {
	case NORTH, SOUTH:
		r.y += 1
	case EAST, WEST:
		r.x += 1
	default:
	}
}

func (r *Rover) CurrentPosition() string {
	return fmt.Sprintf("%v:%v.%v", r.direction, r.x, r.y)
}
