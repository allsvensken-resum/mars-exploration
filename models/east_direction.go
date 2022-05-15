package models

type EAST struct {
	name string
}

func NewEastDirection() *EAST {
	return &EAST{name: "E"}
}

func (e *EAST) MoveFollowDirection(x, y, val int) (int, int) {
	return x + val, y
}

func (e *EAST) GetName() string {
	return e.name
}
