package models

import "fmt"

type IRover interface {
	Move(newX, newY int) string
	Rotate(val int) string
	TryMove(val int) (int, int)
	CurrentPosition() string
}

type rover struct {
	x          int
	y          int
	direction  Direction
	directions []Direction
}

func NewRover(x, y int, direction Direction, directions []Direction) IRover {
	return &rover{x: x, y: y, direction: direction, directions: directions}
}

func (r *rover) Move(newX, newY int) string {
	r.x, r.y = newX, newY
	return r.CurrentPosition()
}

func (r *rover) Rotate(val int) string {
	currIdx := r.findCurrIdx(r.direction)
	nextIdx := currIdx + val
	directionsLength := len(r.directions)

	if nextIdx < 0 {
		nextIdx = directionsLength - 1
	}

	if nextIdx >= directionsLength {
		nextIdx = nextIdx % directionsLength
	}

	r.direction = r.directions[nextIdx]
	return r.CurrentPosition()
}

func (r *rover) TryMove(val int) (int, int) {
	return r.direction.MoveFollowDirection(r.x, r.y, val)
}

func (r *rover) CurrentPosition() string {
	return fmt.Sprintf("%v:%v,%v", r.direction.GetName(), r.x, r.y)
}

func (r *rover) findCurrIdx(currDirection Direction) int {
	for idx, direction := range r.directions {
		if direction.GetName() == currDirection.GetName() {
			return idx
		}
	}
	return 0
}
