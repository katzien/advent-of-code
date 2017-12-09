package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGivenInputsAndOutputs(t *testing.T) {
	cases := map[string]int{
		"1212": 6,
		"1221": 0,
		"123425": 4,
		"123123": 12,
		"12131415": 4,
	}

	for input, output := range cases {

		result := findSum(input)

		assert.Equal(t, output, result)
	}
}
