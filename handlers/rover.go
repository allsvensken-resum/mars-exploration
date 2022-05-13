package handlers

import (
	"io/ioutil"
	"mars/models"
	"mars/services"
	"mime/multipart"
	"net/http"
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
		c.JSON(http.StatusBadRequest, gin.H{"message": "can't open file."})
		return
	}

	file, err := ioutil.ReadAll(openedFile)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "can't read file."})
		return
	}

	instructions := strings.Fields(string(file))
	land := models.NewGrid(2, 4)
	rover := models.NewRover()
	roverService := services.NewRover(rover, land, r.moveInstructionMapper)
	roverService.Action(instructions)

	c.JSON(http.StatusOK, gin.H{"test": instructions})
}
