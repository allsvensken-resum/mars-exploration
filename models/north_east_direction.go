package models

type NORTHEAST struct {
	name string
}

func NewNorthEastDirection() Direction {
	return &NORTHEAST{name: "NE"}
}

func (ne *NORTHEAST) MoveFollowDirection(x, y, val int) (int, int) {
	return x + val, y + val
}

func (ne *NORTHEAST) GetName() string {
	return ne.name
}
