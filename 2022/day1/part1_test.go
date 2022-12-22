package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	t.Parallel()

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

	assert.Len(t, input, 5)
	assert.Equal(t, 6000, input[0])
	assert.Equal(t, 4000, input[1])
	assert.Equal(t, 11000, input[2])
	assert.Equal(t, 24000, input[3])
	assert.Equal(t, 10000, input[4])

	elf, calories := findMostCalories(input)

	assert.Equal(t, 4, elf)
	assert.Equal(t, 24000, calories)
}
