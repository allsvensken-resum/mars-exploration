package models

type Grid struct {
	maxX int
	maxY int
}

func NewGrid(x, y int) Grid {
	return Grid{maxX: x, maxY: y}
}

func (g *Grid) IsOutOfBound(x, y int) bool {
	if x > g.maxX || y > g.maxY || x < 0 || y < 0 {
		return true
	}

	return false
}
