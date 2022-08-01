package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day3P1() {
	input, err := os.ReadFile("./2021/input/day_3")
	if err != nil {
		panic(err)
	}

	dataArr := strings.Split(string(input), "\n")

	gamma, epsilon := countGammaAndEpsilon(dataArr)

	power := calculatePowerConsumption(gamma, epsilon)

	fmt.Printf("🎏 Power consumption: %d\n", power)
}

func calculatePowerConsumption(gamma, epsilon string) int64 {
	gammaRate, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	epsilonRate, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	fmt.Printf("🐠 Gamma: %s (%d)\n", gamma, gammaRate)
	fmt.Printf("🐡 Epsilon: %s (%d)\n", epsilon, epsilonRate)

	return gammaRate * epsilonRate
}

func countGammaAndEpsilon(data []string) (gamma string, epsilon string) {
	length := 0
	dataRunes := [][]rune{}
	for _, line := range data {
		chars := []rune(line)
		dataRunes = append(dataRunes, chars)
		if length == 0 {
			length = len(chars)
		} else {
			if length != len(chars) {
				panic("🔴 numbers are of different length")
			}
		}
	}

	fmt.Printf("🔍 Length: %d\n", length)

	zeros := 0
	ones := 0

	for i := 0; i < length; i++ {
		for _, line := range dataRunes {
			char := string(line[i])
			if char == "1" {
				ones++
			} else {
				zeros++
			}
		}

		if ones > zeros {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		} else {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		}

		//fmt.Printf("🔍 Zeros: %d Ones: %d\n", zeros, ones)

		ones = 0
		zeros = 0
	}

	return
}
