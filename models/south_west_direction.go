package models

type SOUTHWEST struct {
	name string
}

func NewSouthWestDirection() Direction {
	return &SOUTHWEST{name: "SW"}
}

func (sw *SOUTHWEST) MoveFollowDirection(x, y, val int) (int, int) {
	return x - val, y - val
}

func (sw *SOUTHWEST) GetName() string {
	return sw.name
}
