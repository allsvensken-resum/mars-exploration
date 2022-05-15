package services

import (
	"mars/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExplore(t *testing.T) {
	type testCase struct {
		name         string
		instructions []string
		land         models.Grid
		expect       int
	}

	x, y := 0, 0
	maxX, maxY := 5, 5

	land := models.NewGrid(maxX, maxY)
	startDirection := models.NewNorthDirection()
	directions := models.NewDirections()

	cases := []testCase{
		{
			name:         "should have 6 result.",
			instructions: []string{"L", "R", "L", "F", "R"},
			land:         land,
			expect:       6,
		},
		{
			name:         "should have 3 result",
			instructions: []string{"L", "R"},
			land:         land,
			expect:       3,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			roverModel := models.NewRover(x, y, startDirection, directions)
			rover := NewRover(roverModel, land)

			actual := rover.Explore(c.instructions)
			assert.Equal(t, c.expect, len(actual))
		})
	}

}
