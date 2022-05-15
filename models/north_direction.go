package models

type NORTH struct {
	name string
}

func NewNorthDirection() Direction {
	return &NORTH{name: "N"}
}

func (n *NORTH) MoveFollowDirection(x, y, val int) (int, int) {
	return x, y + val
}

func (n *NORTH) GetName() string {
	return n.name
}
