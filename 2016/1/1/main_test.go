package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert.Equal(t, 5, distanceToDestination("R2, L3"))
}

func TestExample2(t *testing.T) {
	assert.Equal(t, 2, distanceToDestination("R2, R2, R2"))
}

func TestExample3(t *testing.T) {
	assert.Equal(t, 12, distanceToDestination("R5, L5, R5, R3"))
}

func TestFinalAnswer(t *testing.T) {
	input := "L4, L1, R4, R1, R1, L3, R5, L5, L2, L3, R2, R1, L4, R5, R4, L2, R1, R3, L5, R1, L3, L2, R5, L4, L5, R1, R2, L1, R5, L3, R2, R2, L1, R5, R2, L1, L1, R2, L1, R1, L2, L2, R4, R3, R2, L3, L188, L3, R2, R54, R1, R1, L2, L4, L3, L2, R3, L1, L1, R3, R5, L1, R5, L1, L1, R2, R4, R4, L5, L4, L1, R2, R4, R5, L2, L3, R5, L5, R1, R5, L2, R4, L2, L1, R4, R3, R4, L4, R3, L4, R78, R2, L3, R188, R2, R3, L2, R2, R3, R1, R5, R1, L1, L1, R4, R2, R1, R5, L1, R4, L4, R2, R5, L2, L5, R4, L3, L2, R1, R1, L5, L4, R1, L5, L1, L5, L1, L4, L3, L5, R4, R5, R2, L5, R5, R5, R4, R2, L1, L2, R3, R5, R5, R5, L2, L1, R4, R3, R1, L4, L2, L3, R2, L3, L5, L2, L2, L1, L2, R5, L2, L2, L3, L1, R1, L4, R2, L4, R3, R5, R3, R4, R1, R5, L3, L5, L5, L3, L2, L1, R3, L4, R3, R2, L1, R3, R1, L2, R4, L3, L3, L3, L1, L2"

	assert.Equal(t, 279, distanceToDestination(input))
}

func TestParseInput(t *testing.T) {
	expected := []move{
		move{direction: "R", distance: 565},
		move{direction: "L", distance: 5},
		move{direction: "R", distance: 23},
		move{direction: "R", distance: 3},
	}

	assert.Equal(t, expected, parseInput("R565, L5, R23, R3"))
}

func TestParseEmptyInput(t *testing.T) {
	assert.Equal(t, []move{}, parseInput(""))
}
