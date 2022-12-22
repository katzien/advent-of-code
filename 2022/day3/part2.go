package day3

import (
	"fmt"
	"os"
	"strings"
)

type Group struct {
	Rucksacks  []string
	CommonItem rune
}

func Part2() {
	list, err := os.ReadFile("./day3/input")
	if err != nil {
		panic(err)
	}

	input := parseInputP2(string(list))

	findCommonItem(input)

	totalScore := scorePriorityP2(input)

	fmt.Printf("Common items total priority score: %d\n", totalScore)
}

func parseInputP2(raw string) []*Group {
	lines := strings.Split(raw, "\n")
	groups := []*Group{}

	var currentGroup *Group
	for k, l := range lines {
		if k%3 == 0 {
			if k != 0 {
				groups = append(groups, currentGroup)
			}
			g := &Group{
				Rucksacks: []string{l},
			}
			currentGroup = g
		} else {
			currentGroup.Rucksacks = append(currentGroup.Rucksacks, l)
		}
	}

	groups = append(groups, currentGroup)

	return groups
}

func findCommonItem(groups []*Group) {
	for _, g := range groups {
		if len(g.Rucksacks) != 3 {
			panic(fmt.Sprintf("invalid group: %v", g.Rucksacks))
		}

		for _, item := range g.Rucksacks[0] {
			if strings.Contains(g.Rucksacks[1], string(item)) && strings.Contains(g.Rucksacks[2], string(item)) {
				g.CommonItem = item
			}
		}

		if g.CommonItem == 0 {
			panic(fmt.Sprintf("no common item found for group %v", g.Rucksacks))
		}
	}
}

func scorePriorityP2(input []*Group) int32 {
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
