package main

import (
	"mars/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20

	roverHandler := handlers.NewRoverHandler()

	router.POST("/mars/explore", roverHandler.ExploreMars)
	err := router.Run(":8080")

	if err != nil {
		panic(err)
	}
}
