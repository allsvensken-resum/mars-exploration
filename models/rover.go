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

func (r *Rover) Move(newX, newY int) {
	r.x, r.y = newX, newY
	fmt.Println(r.CurrentPosition())
}

func (r *Rover) Rotate(val int) {

}

func (r *Rover) TryMove(val int) (x, y int) {
	newX, newY := r.x, r.y
	switch r.direction {
	case NORTH, SOUTH:
		newY += val
	case EAST, WEST:
		newX += val
	default:
	}

	return newX, newY
}

func (r *Rover) CurrentPosition() string {
	return fmt.Sprintf("%v:%v,%v", r.direction, r.x, r.y)
}
