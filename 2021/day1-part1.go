package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day1P1() {
	data, err := os.ReadFile("./2021/input/day_1")
	if err != nil {
		panic(err)
	}

	increasedCount := 0
	prevVal := 0

	dataArr := strings.Split(string(data), "\n")

	for k, v := range dataArr {
		k := k
		v := v
		intVal, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		if k == 0 {
			// first value
			prevVal = intVal
			fmt.Printf("📌 Set first value to %d\n\n", prevVal)
			continue
		}

		if intVal > prevVal {
			//fmt.Printf("🟢 Current value %d is higher than the previous value %d\n", intVal, prevVal)
			increasedCount++
			//fmt.Printf("🟢 Increased count: %d\n", increasedCount)
		} else {
			//fmt.Printf("🔴 Current value %d is lower than the previous value %d\n", intVal, prevVal)
		}

		prevVal = intVal
	}

	fmt.Printf("🎯 Increased count: %d out of %d values", increasedCount, len(dataArr))
}
