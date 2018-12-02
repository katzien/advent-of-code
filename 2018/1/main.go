package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := parseInput()
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}

func partOne(frequencies []int) int {
	freq := 0

	for _, ff := range frequencies {
		freq += ff
	}

	return freq
}

func partTwo(frequencies []int) int {
	return 0
}

func parseInput() ([]int, error) {
	var frequencies []int

	in, err := ioutil.ReadFile("input")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error reading the input file: %s", err))
	}

	input := strings.Split(string(in), "\n")
	for _, ff := range input {
		f, err := strconv.Atoi(ff)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error converting input %s to int: %s", ff, err))
		}
		frequencies = append(frequencies, f)
	}

	return frequencies, nil
}