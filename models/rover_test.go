package models_test

import (
	"mars/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMove(t *testing.T) {
	type testCase struct {
		name     string
		xVal     int
		yVal     int
		expected string
	}

	cases := []testCase{
		{
			name:     "should move follow new x and new y value",
			xVal:     1,
			yVal:     1,
			expected: "N:1,1",
		},
		{
			name:     "should stay the same direction",
			xVal:     2,
			yVal:     2,
			expected: "N:2,2",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			newX, newY := c.xVal, c.yVal
			direction := models.NewDirection("N")
			rover := models.NewRover(0, 0, direction, []models.Direction{})
			actual := rover.Move(newX, newY)
			assert.Equal(t, c.expected, actual)
		})
	}
}

func TestTryMove(t *testing.T) {
	type testCase struct {
		name      string
		val       int
		direction models.Direction
		startX    int
		startY    int
		expectedX int
		expectedY int
	}

	cases := []testCase{
		{
			name:      "should increase x value.",
			val:       1,
			direction: models.NewDirection("E"),
			startX:    0,
			startY:    0,
			expectedX: 1,
			expectedY: 0,
		},
		{
			name:      "should decrease x value.",
			val:       1,
			direction: models.NewDirection("W"),
			startX:    1,
			startY:    0,
			expectedX: 0,
			expectedY: 0,
		},
		{
			name:      "should increase y value.",
			val:       1,
			direction: models.NewDirection("N"),
			startX:    1,
			startY:    0,
			expectedX: 1,
			expectedY: 1,
		},
		{
			name:      "should decrease y value.",
			val:       1,
			direction: models.NewDirection("S"),
			startX:    1,
			startY:    1,
			expectedX: 1,
			expectedY: 0,
		},
		{
			name:      "should stay the same position if something wrong with direction.",
			val:       1,
			direction: models.NewDirection("WRONG"),
			startX:    1,
			startY:    1,
			expectedX: 1,
			expectedY: 1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			rover := models.NewRover(c.startX, c.startY, c.direction, []models.Direction{})
			actualX, actualY := rover.TryMove(c.val)
			assert.Equal(t, c.expectedX, actualX)
			assert.Equal(t, c.expectedY, actualY)
		})
	}
}
