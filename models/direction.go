package models

var (
	NORTH = NewDirection("N")
	SOUTH = NewDirection("S")
	WEST  = NewDirection("W")
	EAST  = NewDirection("E")
)

var (
	directions = []Direction{WEST, NORTH, EAST, SOUTH}
)

type Direction string

func NewDirection(dir string) Direction {
	switch dir {
	case "N":
	case "E":
	case "S":
	case "W":
	default:
		return "WRONG DIRECTION"
	}

	return (Direction)(dir)
}
