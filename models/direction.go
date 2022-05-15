package models

type Direction interface {
	MoveFollowDirection(x, y, val int) (int, int)
	GetName() string
}

func NewDirections() []Direction {
	return []Direction{
		NewNorthDirection(),
		NewNorthEastDirection(),
		NewEastDirection(),
		NewSouthEastDirection(),
		NewSouthDirection(),
		NewSouthWestDirection(),
		NewWestDirection(),
		NewNorthWestDirection(),
	}
}
