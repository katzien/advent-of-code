package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day1P2() {
	input, err := os.ReadFile("./2021/input/day_1")
	if err != nil {
		panic(err)
	}

	dataArr := strings.Split(string(input), "\n")
	data := make([]int, len(dataArr))
	for k, v := range dataArr {
		intVal, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		data[k] = intVal
	}

	fmt.Printf("🧮 Total numbers: %d\n", len(data))

	sums := tripleUp(data)

	countIncreases(sums)
}

func tripleUp(data []int) []int {
	var triples []int

	for i := 0; i < len(data)-2; i++ {
		sum := data[i] + data[i+1] + data[i+2]
		triples = append(triples, sum)
	}

	fmt.Printf("🧮 Total triples: %d\n", len(triples))

	return triples
}

func countIncreases(data []int) {
	increasedCount := 0
	prevVal := 0

	for k, v := range data {
		if k == 0 {
			// first value
			prevVal = v
			fmt.Printf("📌 Set first value to %d\n", prevVal)
			continue
		}

		if v > prevVal {
			//fmt.Printf("🟢 Current value %d is higher than the previous value %d\n", intVal, prevVal)
			increasedCount++
			//fmt.Printf("🟢 Increased count: %d\n", increasedCount)
		} else {
			//fmt.Printf("🔴 Current value %d is lower than the previous value %d\n", intVal, prevVal)
		}

		prevVal = v
	}

	fmt.Printf("🎯 Increased count: %d out of %d values", increasedCount, len(data))
}
