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
}

func NewRoverHandler() Rover {
	return &rover{}
}

func (r *rover) ExploreMars(c *gin.Context) {
	var form Form
	err := c.ShouldBind(&form)
	const ALLOW_FILE_TYPE = "text/plain"

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "can't bind file"})
		return
	}

	fileType := form.File.Header["Content-Type"][0]
	if !strings.Contains(ALLOW_FILE_TYPE, fileType) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "accept only .txt file."})
		return
	}

	openedFile, err := form.File.Open()

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "can't open file."})
		return
	}

	defer openedFile.Close()

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

	NORTH := models.NewNorthDirection()

	directions := models.NewDirections()
	land := models.NewGrid(maxX, maxY)
	rover := models.NewRover(0, 0, NORTH, directions)
	roverService := services.NewRover(rover, land)
	histories := roverService.Explore(instructions[1:])

	c.JSON(http.StatusOK, gin.H{"instructions": instructions, "result": histories})
}
