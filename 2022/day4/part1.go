package day4

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Section struct {
	Start int
	End   int
}

type Pair struct {
	Elf1 Section
	Elf2 Section
}

func Part1() {
	list, err := os.ReadFile("./day4/input")
	if err != nil {
		panic(err)
	}

	input := parseInput(string(list))

	count := findOverlaps(input)

	fmt.Printf("There are %d overlapping pairs\n", count)
}

func parseInput(raw string) []Pair {
	lines := strings.Split(raw, "\n")
	pairs := make([]Pair, len(lines))

	for k, l := range lines {
		pair := Pair{}

		sections := strings.Split(l, ",")
		if len(sections) != 2 {
			panic(fmt.Sprintf("expected two sections on line %s", l))
		}

		for elfNo, s := range sections {
			numbers := strings.Split(s, "-")
			if len(numbers) != 2 {
				panic(fmt.Sprintf("expected two numbers for a section %s", s))
			}

			start, err := strconv.Atoi(numbers[0])
			if err != nil {
				panic(err)
			}
			end, err := strconv.Atoi(numbers[1])
			if err != nil {
				panic(err)
			}

			section := Section{
				Start: start,
				End:   end,
			}

			switch elfNo {
			case 0:
				pair.Elf1 = section
			case 1:
				pair.Elf2 = section
			}
		}

		pairs[k] = pair
	}

	return pairs
}

// .234.....  2-4
// .....678.  6-8
//
// .23......  2-3
// ...45....  4-5
//
// ....567..  5-7
// ......789  7-9
//
// .2345678.  2-8
// ..34567..  3-7
//
// .....6...  6-6
// ...456...  4-6
//
// .23456...  2-6
// ...45678.  4-8
func findOverlaps(pairs []Pair) int {
	count := 0

	for _, p := range pairs {
		lengthSection1 := p.Elf1.End - p.Elf1.Start
		lengthSection2 := p.Elf2.End - p.Elf2.Start

		switch {
		case lengthSection1 >= lengthSection2:
			// .2345678.  2-8
			// ..34567..  3-7
			if p.Elf1.Start <= p.Elf2.Start && p.Elf1.End >= p.Elf2.End {
				count++
				continue
			}
			// .....6...  6-6
			// ...456...  4-6
		case lengthSection1 < lengthSection2:
			if p.Elf1.Start >= p.Elf2.Start && p.Elf1.End <= p.Elf2.End {
				count++
				continue
			}
		}
	}

	return count
}
