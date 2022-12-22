package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay1P2(t *testing.T) {
	list := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

	input := parseInput(list)

	elves, calories := findTopThreeElvesTotal(input)

	assert.Equal(t, []int{4, 3, 5}, elves)
	assert.Equal(t, 45000, calories)
}
