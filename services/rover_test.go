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

	t.Run("should have 3 result with backward result at last", func(t *testing.T) {
		instructions := []string{"F", "B"}
		expected := 3
		expectedPos := "N:0,0"
		roverModel := models.NewRover(x, y, startDirection, directions)
		rover := NewRover(roverModel, land)

		actual := rover.Explore(instructions)

		assert.Equal(t, expected, len(actual))
		assert.Equal(t, expectedPos, actual[len(actual)-1])
	})

	type testReachingEdge struct {
		instructions []string
		expectedPos  string
	}

	reachingEdgeCases := []testReachingEdge{
		{
			instructions: []string{"B", "B", "B"},
			expectedPos:  "0,0",
		},
		{
			instructions: []string{"F", "F", "F", "F", "F", "F", "F"},
			expectedPos:  "0,5",
		},
		{
			instructions: []string{"R", "F", "F", "F", "F", "F", "F"},
			expectedPos:  "5,0",
		},
	}

	for _, rc := range reachingEdgeCases {
		t.Run("should not move when reaching edge", func(t *testing.T) {
			roverModel := models.NewRover(x, y, startDirection, directions)
			rover := NewRover(roverModel, land)

			actual := rover.Explore(rc.instructions)

			assert.Contains(t, actual[len(actual)-1], rc.expectedPos)
		})
	}

}
