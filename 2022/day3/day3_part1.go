package day3

import (
	"fmt"
	"os"
	"strings"
)

type Rucksack struct {
	LeftCompartment  []rune
	RightCompartment []rune
	CommonItem       rune
}

func Part1() {
	list, err := os.ReadFile("./day3/input")
	if err != nil {
		panic(err)
	}

	input := parseInput(string(list))

	findCommonItems(input)

	totalScore := scorePriority(input)

	fmt.Printf("Common items total priority score: %d\n", totalScore)
}

func parseInput(raw string) []*Rucksack {
	lines := strings.Split(raw, "\n")

	input := make([]*Rucksack, len(lines))

	for k, l := range lines {
		halfway := len(l) / 2

		left := l[:halfway]
		right := l[halfway:]

		if len(left) != len(right) {
			panic(fmt.Sprintf("uneven ruckack: %s | %s", left, right))
		}

		input[k] = &Rucksack{
			LeftCompartment:  []rune(left),
			RightCompartment: []rune(right),
		}
	}

	return input
}

func findCommonItems(input []*Rucksack) {
OUTER:
	for _, r := range input {
		for _, i := range r.LeftCompartment {
			for _, j := range r.RightCompartment {
				if j == i {
					r.CommonItem = j
					continue OUTER
				}
			}
		}
	}
}

func scorePriority(input []*Rucksack) int32 {
	var totalScore int32

	for _, r := range input {
		// ascii lowercase starts at 97, so if "a" is to score 1 the difference is 96
		if r.CommonItem > 97 {
			// lowercase item types a through z have priorities 1 through 26
			totalScore += r.CommonItem - 96
			// ascii uppercase starts at 65, so if "A" is to score 27 the difference is 38
		} else {
			// uppercase item types A through Z have priorities 27 through 52
			totalScore += r.CommonItem - 38
		}
	}

	return totalScore
}
