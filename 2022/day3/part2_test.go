package day3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart2(t *testing.T) {
	t.Parallel()

	list := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

	input := parseInputP2(list)

	assert.Len(t, input, 2)
	assert.Equal(t, Group{
		Rucksacks: []string{
			"vJrwpWtwJgWrhcsFMMfFFhFp",
			"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			"PmmdzqPrVvPwwTWBwg",
		},
	}, *input[0])
	assert.Equal(t, Group{
		Rucksacks: []string{
			"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			"ttgJtRGJQctTZtZT",
			"CrZsJsPPZsGzwwsLwLmpwMDw",
		},
	}, *input[1])

	findCommonItem(input)

	assert.Equal(t, 'r', input[0].CommonItem)
	assert.Equal(t, 'Z', input[1].CommonItem)

	sum := scorePriorityP2(input)

	assert.Equal(t, int32(70), sum)
}
