package models

type SOUTH struct {
	name string
}

func NewSouthDirection() *SOUTH {
	return &SOUTH{name: "S"}
}

func (s *SOUTH) MoveFollowDirection(x, y, val int) (int, int) {
	return x, y - val
}

func (s *SOUTH) GetName() string {
	return s.name
}
