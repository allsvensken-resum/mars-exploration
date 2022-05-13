package main

import (
	"mars/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20

	moveInstructionMapper := getMoveInstructionMapper()
	roverHandler := handlers.NewRoverHandler(moveInstructionMapper)

	router.POST("/mars/explore", roverHandler.ExploreMars)
	err := router.Run(":8080")

	if err != nil {
		panic(err)
	}
}

func getMoveInstructionMapper() map[string][2]int {
	moveMap := make(map[string][2]int)
	moveMap["L"] = [2]int{0, -1}
	moveMap["R"] = [2]int{0, 1}
	moveMap["F"] = [2]int{1, 0}
	return moveMap
}
