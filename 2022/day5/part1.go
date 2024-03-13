package day4

func Part1() {
	//list, err := os.ReadFile("./day4/input")
	//if err != nil {
	//	panic(err)
	//}
	//
	//input := parseInput(string(list))

	//count := findOverlaps(input)

	//fmt.Printf("There are %d overlapping pairs\n", count)
}

//func parseInput(raw string) []Pair {
//	lines := strings.Split(raw, "\n")
//	pairs := make([]Pair, len(lines))
//
//	for k, l := range lines {
//		pair := Pair{}
//
//		sections := strings.Split(l, ",")
//		if len(sections) != 2 {
//			panic(fmt.Sprintf("expected two sections on line %s", l))
//		}
//
//		for elfNo, s := range sections {
//			numbers := strings.Split(s, "-")
//			if len(numbers) != 2 {
//				panic(fmt.Sprintf("expected two numbers for a section %s", s))
//			}
//
//			start, err := strconv.Atoi(numbers[0])
//			if err != nil {
//				panic(err)
//			}
//			end, err := strconv.Atoi(numbers[1])
//			if err != nil {
//				panic(err)
//			}
//
//			section := Section{
//				Start: start,
//				End:   end,
//			}
//
//			switch elfNo {
//			case 0:
//				pair.Elf1 = section
//			case 1:
//				pair.Elf2 = section
//			}
//		}
//
//		pairs[k] = pair
//	}
//
//	return pairs
//}
