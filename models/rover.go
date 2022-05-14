package models

import "fmt"

type IRover interface {
	Move(newX, newY int) string
	Rotate(val int) string
	TryMove(val int) (int, int)
	CurrentPosition() string
}

type rover struct {
	x         int
	y         int
	direction Direction
}

func NewRover(x, y int, direction Direction) IRover {
	return &rover{x: x, y: y, direction: direction}
}

func (r *rover) Move(newX, newY int) string {
	r.x, r.y = newX, newY
	return r.CurrentPosition()
}

func (r *rover) Rotate(val int) string {
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
	return r.CurrentPosition()
}

func (r *rover) TryMove(val int) (int, int) {
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
	}

	return newX, newY
}

func (r *rover) CurrentPosition() string {
	return fmt.Sprintf("%v:%v,%v", r.direction, r.x, r.y)
}

func (r *rover) findCurrIdx(currDirection Direction) int {
	for idx, direction := range directions {
		if direction == currDirection {
			return idx
		}
	}
	return 0
}
