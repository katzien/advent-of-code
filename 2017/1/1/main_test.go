package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGivenInputsAndOutputs(t *testing.T) {
	cases := map[string]int{
		"1122": 3,
		"1111": 4,
		"1234": 0,
		"91212129": 9,
	}

	for input, output := range cases {

		result := findSum(input)

		assert.Equal(t, output, result)
	}
}
