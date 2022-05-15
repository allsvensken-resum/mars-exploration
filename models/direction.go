package models

type Direction interface {
	MoveFollowDirection(x, y, val int) (int, int)
	GetName() string
}

func NewDirections() []Direction {
	return []Direction{
		NewNorthDirection(), NewEastDirection(), NewSouthDirection(), NewWestDirection(),
	}
}
