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

func NewNorthDirection() *NORTH {
	return &NORTH{name: "N"}
}

func NewSouthDirection() *SOUTH {
	return &SOUTH{name: "S"}
}

func NewWestDirection() *WEST {
	return &WEST{name: "W"}
}

func NewEastDirection() *EAST {
	return &EAST{name: "E"}
}

type NORTH struct {
	name string
}

func (n *NORTH) MoveFollowDirection(x, y, val int) (int, int) {
	return x, y + val
}

func (n *NORTH) GetName() string {
	return n.name
}

type SOUTH struct {
	name string
}

func (s *SOUTH) MoveFollowDirection(x, y, val int) (int, int) {
	return x, y - val
}

func (s *SOUTH) GetName() string {
	return s.name
}

type WEST struct {
	name string
}

func (w *WEST) MoveFollowDirection(x, y, val int) (int, int) {
	return x - val, y
}

func (w *WEST) GetName() string {
	return w.name
}

type EAST struct {
	name string
}

func (e *EAST) MoveFollowDirection(x, y, val int) (int, int) {
	return x + val, y
}

func (e *EAST) GetName() string {
	return e.name
}
