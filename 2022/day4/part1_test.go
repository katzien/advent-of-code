package day4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	list := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

	input := parseInput(list)

	assert.Len(t, input, 6)
	assert.Equal(t, Pair{
		Elf1: Section{
			Start: 2,
			End:   4,
		},
		Elf2: Section{
			Start: 6,
			End:   8,
		},
	}, input[0])
	assert.Equal(t, Pair{
		Elf1: Section{
			Start: 2,
			End:   3,
		},
		Elf2: Section{
			Start: 4,
			End:   5,
		},
	}, input[1])
	assert.Equal(t, Pair{
		Elf1: Section{
			Start: 5,
			End:   7,
		},
		Elf2: Section{
			Start: 7,
			End:   9,
		},
	}, input[2])
	assert.Equal(t, Pair{
		Elf1: Section{
			Start: 2,
			End:   8,
		},
		Elf2: Section{
			Start: 3,
			End:   7,
		},
	}, input[3])
	assert.Equal(t, Pair{
		Elf1: Section{
			Start: 6,
			End:   6,
		},
		Elf2: Section{
			Start: 4,
			End:   6,
		},
	}, input[4])
	assert.Equal(t, Pair{
		Elf1: Section{
			Start: 2,
			End:   6,
		},
		Elf2: Section{
			Start: 4,
			End:   8,
		},
	}, input[5])

	count := findOverlaps(input)

	assert.Equal(t, 2, count)
}
