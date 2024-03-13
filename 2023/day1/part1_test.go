package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration value that the Elves now need to recover. On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.
//
//For example:
//
//1abc2
//pqr3stu8vwx
//a1b2c3d4e5f
//treb7uchet

// In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.
func TestName(t *testing.T) {

	input := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	result := calculateCalibrationValue(input)

	if result != 142 {
		t.Error("result is not 142")
		t.Fail()
	}
	//fmt.Println(142 == result)
	assert.Equal(t, 142, result)
}
