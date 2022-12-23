package day4

import (
	"fmt"
	"os"
)

func Part2() {
	list, err := os.ReadFile("./day4/input")
	if err != nil {
		panic(err)
	}

	input := parseInput(string(list))

	count := findAnyOverlaps(input)

	fmt.Printf("There are %d overlapping pairs\n", count)
}

func findAnyOverlaps(pairs []Pair) int {
	count := 0
	for _, p := range pairs {
		// the only time there's no overlap is if either the first section finishes before the second,
		// or the second finishes before the first - e.g.
		// .23456...  2-6
		// ......78.  4-8

		// ......78.  4-8
		// .23456...  2-6
		if p.Elf1.End < p.Elf2.Start || p.Elf1.Start > p.Elf2.End {
			continue
		}
		count++
	}

	return count
}
