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
	currIdx := r.findCurrIdx(r.direction)
	nextIdx := currIdx + val
	directionsLength := len(directions)

	if nextIdx < 0 {
		nextIdx = directionsLength - 1
	}

	if nextIdx >= directionsLength {
		nextIdx = nextIdx % directionsLength
	}

	r.direction = directions[nextIdx]
	fmt.Println(r.CurrentPosition())
}

func (r *Rover) TryMove(val int) (int, int) {
	newX, newY := r.x, r.y
	switch r.direction {
	case NORTH:
		newY += val
	case SOUTH:
		newY -= val
	case EAST:
		newX += val
	case WEST:
		newX -= val
	default:
	}

	return newX, newY
}

func (r *Rover) CurrentPosition() string {
	return fmt.Sprintf("%v:%v,%v", r.direction, r.x, r.y)
}

func (r *Rover) findCurrIdx(currDirection direction) int {
	for idx, direction := range directions {
		if direction == currDirection {
			return idx
		}
	}
	return 0
}
