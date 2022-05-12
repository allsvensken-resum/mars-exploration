package main

import (
	"mars/models"
	"mars/services"
)

func main() {
	moveMap := make(map[string][2]int)
	moveMap["L"] = [2]int{0, 1}
	moveMap["R"] = [2]int{0, -1}
	moveMap["F"] = [2]int{1, 0}

	land := models.NewGrid(2, 4)
	rover := models.NewRover()

	roverService := services.NewRover(rover, land, moveMap)
	instructions := []string{"L", "F", "L", "F", "L", "L", "F"}
	roverService.Action(instructions)
}
