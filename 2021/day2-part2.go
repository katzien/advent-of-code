package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instr struct {
	dir string
	num int
}

func Day2P2() {
	input, err := os.ReadFile("./2021/input/day_2")
	if err != nil {
		panic(err)
	}

	dataArr := strings.Split(string(input), "\n")

	data := []instr{}
	for _, dd := range dataArr {
		parts := strings.Split(dd, " ")
		intVal, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		data = append(data, instr{dir: parts[0], num: intVal})
	}

	horiz, depth, aim := calculatePath(data)

	fmt.Printf("🐠 Final position: %d depth: %d aim: %d\n", horiz, depth, aim)
	fmt.Printf("🐡 Horiz x Depth: %d\n", horiz*depth)
}

func calculatePath(data []instr) (horiz int, depth int, aim int) {
	fmt.Printf("🧮 Total numbers: %d\n", len(data))

	for _, dd := range data {
		//fmt.Printf("🍤 Next instruction: %s %d\n", dd.dir, dd.num)
		switch dd.dir {
		case "forward":
			horiz += dd.num
			depth += aim * dd.num
		case "up":
			aim -= dd.num
		case "down":
			aim += dd.num
		}
		//fmt.Printf("🍤 Result: H: %d D: %d A: %d\n", horiz, depth, aim)
	}

	return horiz, depth, aim
}
