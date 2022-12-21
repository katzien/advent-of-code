package main

import (
	"fmt"
	"os"
	"sort"
)

func Day1P2() {
	list, err := os.ReadFile("./input/day_1")
	if err != nil {
		panic(err)
	}

	input := parseInput(string(list))

	elves, calories := findTopThreeElvesTotal(input)

	fmt.Printf("Elves %v carry %d calories in total\n", elves, calories)
}

func findTopThreeElvesTotal(input map[int]int) ([]int, int) {
	elves := make([]int, 3)
	totalCalories := 0

	type kv struct {
		Key   int
		Value int
	}

	var ss []kv
	for k, v := range input {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for k, s := range ss[:3] {
		elves[k] = s.Key + 1
		totalCalories += s.Value
	}

	return elves, totalCalories
}
