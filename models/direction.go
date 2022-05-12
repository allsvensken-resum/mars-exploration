package models

import "fmt"

var (
	NORTH, _ = NewDirection("N")
	SOUTH, _ = NewDirection("S")
	WEST, _  = NewDirection("W")
	EAST, _  = NewDirection("E")
)

var (
	directions = []direction{WEST, NORTH, EAST, SOUTH}
)

type direction string

func NewDirection(dir string) (direction, error) {
	switch dir {
	case "N":
	case "E":
	case "S":
	case "W":
	default:
		return "WRONG DIRECTION", fmt.Errorf("Wrong value for direction type.")
	}

	return (direction)(dir), nil
}
