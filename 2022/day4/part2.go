package day4

import (
	"os"
)

func Part2() {
	_, err := os.ReadFile("./day4/input")
	if err != nil {
		panic(err)
	}

}
