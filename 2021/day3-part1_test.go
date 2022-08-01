package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountGammaAndEpsilon(t *testing.T) {
	data := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	gamma, epsilon := countGammaAndEpsilon(data)

	assert.Equal(t, "10110", gamma)
	assert.Equal(t, "01001", epsilon)
}

func TestCalculatePower(t *testing.T) {
	assert.Equal(t, int64(198), calculatePowerConsumption("10110", "01001"))
}
