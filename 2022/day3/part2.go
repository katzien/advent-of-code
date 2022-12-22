package day3

import (
	"fmt"
	"os"
)

func Part2() {
	list, err := os.ReadFile("./day3/input")
	if err != nil {
		panic(err)
	}

	input := parseInput(string(list))

	findCommonItems(input)

	totalScore := scorePriority(input)

	fmt.Printf("Common items total priority score: %d\n", totalScore)
}
