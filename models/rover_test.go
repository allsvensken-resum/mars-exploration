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

	rover := models.NewRover()

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			newX, newY := c.xVal, c.yVal
			actual := rover.Move(newX, newY)
			assert.Equal(t, c.expected, actual)
		})
	}
}
