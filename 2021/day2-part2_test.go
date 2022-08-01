package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculatePath(t *testing.T) {

	//forward 5
	//down 5
	//forward 8
	//up 3
	//down 8
	//forward 2

	data := []instr{
		{dir: "forward", num: 5},
		{dir: "down", num: 5},
		{dir: "forward", num: 8},
		{dir: "up", num: 3},
		{dir: "down", num: 8},
		{dir: "forward", num: 2},
	}

	// After following these new instructions, you would have
	// a horizontal position of 15 and a depth of 60.
	// (Multiplying these produces 900.)
	horiz, depth, aim := calculatePath(data)

	assert.Equal(t, 15, horiz)
	assert.Equal(t, 60, depth)
	assert.Equal(t, 10, aim)
	assert.Equal(t, 900, horiz*depth)
}
