package services

import "mars/models"

type IRover interface {
	Explore([]string) []string
}

type rover struct {
	land       models.Grid
	moveMapper map[string][2]int
	rover      models.Rover
}

func NewRover(rov models.Rover, land models.Grid, moveMapper map[string][2]int) IRover {
	return &rover{rover: rov, land: land, moveMapper: moveMapper}
}

func (r *rover) Explore(instructions []string) []string {
	histories := make([]string, 0)
	mappedInstructions := r.mapInstructionToNumber(instructions)

	for _, instruction := range mappedInstructions {
		move := instruction[0]
		rotate := instruction[1]

		if isMoveInstruction(move) {
			newX, newY := r.rover.TryMove(move)
			if !r.land.IsOutOfBound(newX, newY) {
				histories = append(histories, r.rover.Move(newX, newY))
			}
			continue
		}

		if isRotateInstruction(rotate) {
			histories = append(histories, r.rover.Rotate(rotate))
		}
	}

	return histories
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

func isRotateInstruction(instruction int) bool {
	return instruction != 0
}

func isMoveInstruction(instruction int) bool {
	return instruction != 0
}
