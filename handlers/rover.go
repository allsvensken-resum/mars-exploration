package handlers

import (
	"io/ioutil"
	"mars/models"
	"mars/services"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Form struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

type Rover interface {
	ExploreMars(c *gin.Context)
}

type rover struct {
	moveInstructionMapper map[string][2]int
}

func NewRoverHandler(moveInstructionMapper map[string][2]int) Rover {
	return &rover{moveInstructionMapper: moveInstructionMapper}
}

func (r *rover) ExploreMars(c *gin.Context) {
	var form Form
	_ = c.ShouldBind(&form)

	openedFile, err := form.File.Open()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "can't open file."})
		return
	}

	file, err := ioutil.ReadAll(openedFile)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "can't read file."})
		return
	}

	instructions := strings.Fields(string(file))

	if len(instructions) < 2 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "you must give grid size and at least one instruction."})
	}

	gridSize := instructions[0]
	gridSizeArr := strings.Split(gridSize, "")

	maxX, err := strconv.Atoi(gridSizeArr[0])
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "grid size is invalid."})
		return
	}

	maxY, err := strconv.Atoi(gridSizeArr[1])
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "grid size is invalid."})
	}

	land := models.NewGrid(maxX, maxY)
	rover := models.NewRover()
	roverService := services.NewRover(rover, land, r.moveInstructionMapper)
	histories := roverService.Explore(instructions[1:])

	c.JSON(http.StatusOK, gin.H{"result": histories})
}
