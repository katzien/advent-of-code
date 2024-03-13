package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExampleInput(t *testing.T) {
	input := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	result := calculateCalibrationValue(input)

	assert.Equal(t, 142, result)
}
