package models

type SOUTHEAST struct {
	name string
}

func NewSouthEastDirection() Direction {
	return &SOUTHEAST{name: "SE"}
}

func (se *SOUTHEAST) MoveFollowDirection(x, y, val int) (int, int) {
	return x + val, y - val
}

func (se *SOUTHEAST) GetName() string {
	return se.name
}
