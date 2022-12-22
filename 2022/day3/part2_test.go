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

	input := parseInput(list)

	assert.Len(t, input, 6)
	assert.Equal(t, Rucksack{
		LeftCompartment:  []rune("vJrwpWtwJgWr"),
		RightCompartment: []rune("hcsFMMfFFhFp"),
	}, *input[0])
	assert.Equal(t, Rucksack{
		LeftCompartment:  []rune("jqHRNqRjqzjGDLGL"),
		RightCompartment: []rune("rsFMfFZSrLrFZsSL"),
	}, *input[1])
	assert.Equal(t, Rucksack{
		LeftCompartment:  []rune("PmmdzqPrV"),
		RightCompartment: []rune("vPwwTWBwg"),
	}, *input[2])
	assert.Equal(t, Rucksack{
		LeftCompartment:  []rune("wMqvLMZHhHMvwLH"),
		RightCompartment: []rune("jbvcjnnSBnvTQFn"),
	}, *input[3])
	assert.Equal(t, Rucksack{
		LeftCompartment:  []rune("ttgJtRGJ"),
		RightCompartment: []rune("QctTZtZT"),
	}, *input[4])
	assert.Equal(t, Rucksack{
		LeftCompartment:  []rune("CrZsJsPPZsGz"),
		RightCompartment: []rune("wwsLwLmpwMDw"),
	}, *input[5])

	findCommonItems(input)

	assert.Equal(t, 'p', input[0].CommonItem)
	assert.Equal(t, 'L', input[1].CommonItem)
	assert.Equal(t, 'P', input[2].CommonItem)
	assert.Equal(t, 'v', input[3].CommonItem)
	assert.Equal(t, 't', input[4].CommonItem)
	assert.Equal(t, 's', input[5].CommonItem)

	//sum := scorePriority(input)
	//
	//assert.Equal(t, int32(157), sum)
}
