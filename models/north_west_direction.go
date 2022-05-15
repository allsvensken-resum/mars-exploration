package models

type NORTHWEST struct {
	name string
}

func NewNorthWestDirection() Direction {
	return &NORTHWEST{name: "NW"}
}

func (nw *NORTHWEST) MoveFollowDirection(x, y, val int) (int, int) {
	return x - val, y + val
}

func (nw *NORTHWEST) GetName() string {
	return nw.name
}
