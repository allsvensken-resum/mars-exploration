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
			expected: "1,1",
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
	NORTHEAST := models.NewNorthEastDirection()
	NORTHWEST := models.NewNorthWestDirection()
	SOUTHEAST := models.NewSouthEastDirection()
	SOUTHWEST := models.NewSouthWestDirection()

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
		{
			name:      "should increase both x and y value.",
			val:       1,
			direction: NORTHEAST,
			startX:    1,
			startY:    1,
			expectedX: 2,
			expectedY: 2,
		},
		{
			name:      "should decrease x and increase y value.",
			val:       -1,
			direction: SOUTHEAST,
			startX:    1,
			startY:    1,
			expectedX: 0,
			expectedY: 2,
		},
		{
			name:      "should decrease x and decrease y value.",
			direction: SOUTHWEST,
			val:       1,
			startX:    1,
			startY:    1,
			expectedX: 0,
			expectedY: 0,
		},
		{
			name:      "should decrease x and increase y value.",
			direction: NORTHWEST,
			val:       1,
			startX:    1,
			startY:    1,
			expectedX: 0,
			expectedY: 2,
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
		expected       models.Direction
	}

	NORTH := models.NewNorthDirection()
	SOUTH := models.NewSouthDirection()
	WEST := models.NewWestDirection()
	EAST := models.NewEastDirection()
	NORTHEAST := models.NewNorthEastDirection()
	NORTHWEST := models.NewNorthWestDirection()
	SOUTHEAST := models.NewSouthEastDirection()
	SOUTHWEST := models.NewSouthWestDirection()

	directions := models.NewDirections()

	cases := []testCase{
		{
			name:           "should change from NORTH to WEST",
			val:            -2,
			startDirection: NORTH,
			expected:       WEST,
		},
		{
			name:           "should change from NORTH to EAST",
			val:            2,
			startDirection: NORTH,
			expected:       EAST,
		},
		{
			name:           "should change from WEST to SOUTH",
			val:            -2,
			startDirection: WEST,
			expected:       SOUTH,
		},
		{
			name:           "should change from SOUTH to WEST",
			val:            2,
			startDirection: SOUTH,
			expected:       WEST,
		},
		{
			name:           "should change from NORTH TO NORTH",
			val:            16,
			startDirection: NORTH,
			expected:       NORTH,
		},
		{
			name:           "should change from NORTH TO NORTH EAST",
			val:            1,
			startDirection: NORTH,
			expected:       NORTHEAST,
		},
		{
			name:           "should change from NORTH EAST TO EAST",
			val:            1,
			startDirection: NORTHEAST,
			expected:       EAST,
		},
		{
			name:           "should change from NORTH EAST TO SOUTH EAST",
			val:            2,
			startDirection: NORTHEAST,
			expected:       SOUTHEAST,
		},
		{
			name:           "should change from SOUTH EAST TO SOUTH WEST",
			val:            2,
			startDirection: SOUTHEAST,
			expected:       SOUTHWEST,
		},
		{
			name:           "should change from SOUTH WEST TO NORTH WEST",
			val:            2,
			startDirection: SOUTHWEST,
			expected:       NORTHWEST,
		},
		{
			name:           "should change from NORTH to NORTH WEST",
			val:            -1,
			startDirection: NORTH,
			expected:       NORTHWEST,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			rover := models.NewRover(0, 0, c.startDirection, directions)
			actual := rover.Rotate(c.val)
			assert.Equal(t, c.expected.GetName(), actual)
		})
	}
}
