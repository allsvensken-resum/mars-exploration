package services

import "mars/models"

type IRover interface {
	Explore([]string) []string
}

type rover struct {
	land  models.Grid
	rover models.IRover
}

func NewRover(rov models.IRover, land models.Grid) IRover {
	return &rover{rover: rov, land: land}
}

func (r *rover) Explore(instructions []string) []string {
	histories := make([]string, 0)
	histories = append(histories, r.rover.Status())
	mappedInstructions := r.mapInstructionToNumber(instructions)

	for _, instruction := range mappedInstructions {
		move := instruction[0]
		rotate := instruction[1]

		if isMoveInstruction(move) {
			newX, newY := r.rover.TryMove(move)
			if !r.land.IsOutOfBound(newX, newY) {
				r.rover.Move(newX, newY)
			}
		}

		if isRotateInstruction(rotate) {
			r.rover.Rotate(rotate)
		}

		histories = append(histories, r.rover.Status())
	}

	return histories
}

func (r *rover) mapInstructionToNumber(instructions []string) [][2]int {
	mappedInstructions := make([][2]int, 0)
	moveMap := getMoveInstructionMapper()
	for _, instruction := range instructions {
		mapped, ok := moveMap[instruction]

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

func getMoveInstructionMapper() map[string][2]int {
	moveMap := make(map[string][2]int)
	moveMap["L"] = [2]int{0, -2}
	moveMap["HL"] = [2]int{0, -1}
	moveMap["R"] = [2]int{0, 2}
	moveMap["HR"] = [2]int{0, 1}
	moveMap["F"] = [2]int{1, 0}
	moveMap["B"] = [2]int{-1, 0}
	return moveMap
}
