package services

import "mars/models"

type IRover interface {
	Action([]string) string
}

type rover struct {
	land       models.Grid
	moveMapper map[string][2]int
	rover      models.Rover
}

func NewRover(land models.Grid, moveMapper map[string][2]int) IRover {
	return &rover{land: land, moveMapper: moveMapper}
}

func (r *rover) Action(instructions []string) string {
	mappedInstructions := r.mapInstructionToNumber(instructions)

	for _, instruction := range mappedInstructions {
		move := instruction[0]
		rotate := instruction[1]

		if move != 0 {
			newX, newY := r.rover.TryMove(move)
			if !r.land.IsOutOfBound(newX, newY) {
				r.rover.Move(newX, newY)
			}
			continue
		}

		if rotate != 0 {

		}
	}

	return r.rover.CurrentPosition()
}

func (r *rover) mapInstructionToNumber(instructions []string) [][2]int {
	mappedInstructions := make([][2]int, 0)
	for _, instruction := range instructions {
		mappedInstructions = append(mappedInstructions, r.moveMapper[instruction])
	}
	return mappedInstructions
}
