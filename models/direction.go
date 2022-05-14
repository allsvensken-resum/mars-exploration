package models

type Direction interface {
	MoveFollowDirection(x, y, val int) (int, int)
	GetName() string
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

type WRONG_DIRECTION struct {
	name string
}

func (w *WRONG_DIRECTION) MoveFollowDirection(x, y, val int) (int, int) {
	return x, y
}

func (w *WRONG_DIRECTION) GetName() string {
	return w.name
}

func NewDirection(dir string) Direction {
	switch dir {
	case "N":
		return &NORTH{name: dir}
	case "E":
		return &EAST{name: dir}
	case "S":
		return &SOUTH{name: dir}
	case "W":
		return &WEST{name: dir}
	}
	return &WRONG_DIRECTION{name: "WRONG DIRECTION"}
}
