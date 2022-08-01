package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day2P1() {
	input, err := os.ReadFile("./2021/input/day_2")
	if err != nil {
		panic(err)
	}

	dataArr := strings.Split(string(input), "\n")

	type instr struct {
		dir string
		num int
	}
	data := []instr{}
	for _, dd := range dataArr {
		parts := strings.Split(dd, " ")
		intVal, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		data = append(data, instr{dir: parts[0], num: intVal})
	}

	fmt.Printf("🧮 Total numbers: %d\n", len(data))

	horiz := 0
	depth := 0

	for _, dd := range data {
		switch dd.dir {
		case "forward":
			horiz += dd.num
		case "up":
			depth -= dd.num
		case "down":
			depth += dd.num
		}
	}

	fmt.Printf("🐠 Final position: %d depth: %d\n", horiz, depth)
}
