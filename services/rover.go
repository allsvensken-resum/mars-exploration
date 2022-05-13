package services

import "mars/models"

type IRover interface {
	Explore([]string) string
}

type rover struct {
	land       models.Grid
	moveMapper map[string][2]int
	rover      models.Rover
}

func NewRover(rov models.Rover, land models.Grid, moveMapper map[string][2]int) IRover {
	return &rover{rover: rov, land: land, moveMapper: moveMapper}
}

func (r *rover) Explore(instructions []string) string {
	const MOVE_INSTRUCTION = 0
	const ROTATE_INSTRUCTION = 1
	mappedInstructions := r.mapInstructionToNumber(instructions)

	for _, instruction := range mappedInstructions {
		move := instruction[0]
		rotate := instruction[1]

		if move != MOVE_INSTRUCTION {
			newX, newY := r.rover.TryMove(move)
			if !r.land.IsOutOfBound(newX, newY) {
				r.rover.Move(newX, newY)
			}
			continue
		}

		if rotate != ROTATE_INSTRUCTION {
			r.rover.Rotate(rotate)
		}
	}

	return r.rover.CurrentPosition()
}

func (r *rover) mapInstructionToNumber(instructions []string) [][2]int {
	mappedInstructions := make([][2]int, 0)
	for _, instruction := range instructions {
		mapped, ok := r.moveMapper[instruction]

		if ok {
			mappedInstructions = append(mappedInstructions, mapped)
		}
	}
	return mappedInstructions
}
