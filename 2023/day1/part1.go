package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// The newly-improved calibration document consists of lines of text;
// each line originally contained a specific calibration value that the Elves now need to recover.
// On each line, the calibration value can be found by combining the first digit and the last digit
// (in that order) to form a single two-digit number.
//
// For example:
//
// 1abc2
// pqr3stu8vwx
// a1b2c3d4e5f
// treb7uchet
//
// In this example, the calibration values of these four lines are 12, 38, 15, and 77.
// Adding these together produces 142.

func Part1() {
	input, err := os.ReadFile("./day1/input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	result := calculateCalibrationValue(lines)

	fmt.Println(result)
}

func calculateCalibrationValue(input []string) int {
	sum := 0

	// read input line by line
	for _, line := range input {
		// for each line, figure out the 1st and last number
		firstNum, lastNum := "", ""
		chars := strings.Split(line, "")

		for _, char := range chars {
			_, err := strconv.Atoi(char)
			if err != nil {
				continue // not a number
			}
			if firstNum == "" {
				firstNum = char
			} else {
				lastNum = char
			}
		}

		if lastNum == "" {
			lastNum = firstNum // for cases where there's just one number in the string
		}

		// joining the chars together to get a number
		numStr := firstNum + lastNum
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}

		// add all numbers together and return that as result
		sum += num
	}

	return sum
}
