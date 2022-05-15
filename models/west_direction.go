package models

type WEST struct {
	name string
}

func NewWestDirection() *WEST {
	return &WEST{name: "W"}
}

func (w *WEST) MoveFollowDirection(x, y, val int) (int, int) {
	return x - val, y
}

func (w *WEST) GetName() string {
	return w.name
}
