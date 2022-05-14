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
			direction := models.NewNorthDirection()
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

	EAST := models.NewEastDirection()
	NORTH := models.NewNorthDirection()
	SOUTH := models.NewSouthDirection()
	WEST := models.NewWestDirection()

	cases := []testCase{
		{
			name:      "should increase x value.",
			val:       1,
			direction: EAST,
			startX:    0,
			startY:    0,
			expectedX: 1,
			expectedY: 0,
		},
		{
			name:      "should decrease x value.",
			val:       1,
			direction: WEST,
			startX:    1,
			startY:    0,
			expectedX: 0,
			expectedY: 0,
		},
		{
			name:      "should increase y value.",
			val:       1,
			direction: NORTH,
			startX:    1,
			startY:    0,
			expectedX: 1,
			expectedY: 1,
		},
		{
			name:      "should decrease y value.",
			val:       1,
			direction: SOUTH,
			startX:    1,
			startY:    1,
			expectedX: 1,
			expectedY: 0,
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

func TestRotate(t *testing.T) {
	type testCase struct {
		name           string
		val            int
		startDirection models.Direction
		expected       string
	}

	NORTH := models.NewNorthDirection()
	SOUTH := models.NewSouthDirection()
	WEST := models.NewWestDirection()
	EAST := models.NewEastDirection()

	directions := []models.Direction{
		NORTH,
		EAST,
		SOUTH,
		WEST,
	}

	cases := []testCase{
		{
			name:           "should change from NORTH to WEST",
			val:            -1,
			startDirection: NORTH,
			expected:       WEST.GetName(),
		},
		{
			name:           "should change from NORTH to EAST",
			val:            1,
			startDirection: NORTH,
			expected:       EAST.GetName(),
		},
		{
			name:           "should change from WEST to SOUTH",
			val:            -1,
			startDirection: WEST,
			expected:       SOUTH.GetName(),
		},
		{
			name:           "should change from SOUTH to WEST",
			val:            1,
			startDirection: SOUTH,
			expected:       WEST.GetName(),
		},
		{
			name:           "should change from NORTH TO NORTH",
			val:            8,
			startDirection: NORTH,
			expected:       NORTH.GetName(),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			rover := models.NewRover(0, 0, c.startDirection, directions)
			actual := rover.Rotate(c.val)
			assert.Contains(t, actual, c.expected)
		})
	}
}
