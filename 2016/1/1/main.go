package main

import (
	"log"
	"math"
	"strconv"
	"strings"
)

var verbose = false

func main() {
	input := "L4, L1, R4, R1, R1, L3, R5, L5, L2, L3, R2, R1, L4, R5, R4, L2, R1, R3, L5, R1, L3, L2, R5, L4, L5, R1, R2, L1, R5, L3, R2, R2, L1, R5, R2, L1, L1, R2, L1, R1, L2, L2, R4, R3, R2, L3, L188, L3, R2, R54, R1, R1, L2, L4, L3, L2, R3, L1, L1, R3, R5, L1, R5, L1, L1, R2, R4, R4, L5, L4, L1, R2, R4, R5, L2, L3, R5, L5, R1, R5, L2, R4, L2, L1, R4, R3, R4, L4, R3, L4, R78, R2, L3, R188, R2, R3, L2, R2, R3, R1, R5, R1, L1, L1, R4, R2, R1, R5, L1, R4, L4, R2, R5, L2, L5, R4, L3, L2, R1, R1, L5, L4, R1, L5, L1, L5, L1, L4, L3, L5, R4, R5, R2, L5, R5, R5, R4, R2, L1, L2, R3, R5, R5, R5, L2, L1, R4, R3, R1, L4, L2, L3, R2, L3, L5, L2, L2, L1, L2, R5, L2, L2, L3, L1, R1, L4, R2, L4, R3, R5, R3, R4, R1, R5, L3, L5, L5, L3, L2, L1, R3, L4, R3, R2, L1, R3, R1, L2, R4, L3, L3, L3, L1, L2"

	log.Printf("The destination is %d blocks away.", distanceToDestination(input))
}

type move struct {
	direction string
	distance  int
}

func parseInput(input string) []move {
	moves := []move{}

	if len(input) == 0 {
		return moves
	}

	steps := strings.Split(input, ", ")

	for _, s := range steps {
		d, e := strconv.Atoi(s[1:])
		if e != nil {
			log.Fatalf("Error converting %s to int: %s", s[1:], e.Error())
		}

		moves = append(moves, move{direction: s[0:1], distance: d})
	}

	return moves
}

func distanceToDestination(input string) int {

	var direction = "N"
	var x, y int

	moves := parseInput(input)

	if verbose {
		log.Printf("Moves: %v", moves)
	}

	for _, move := range moves {

		if verbose {
			log.Printf("Currently at [%d, %d] facing %s", x, y, direction)
			log.Printf("Next move: %v", move)
		}

		switch direction {
		case "N":
			switch move.direction {
			case "L":
				direction = "W"
				x = x - move.distance
			case "R":
				direction = "E"
				x = x + move.distance
			}
		case "S":
			switch move.direction {
			case "L":
				direction = "E"
				x = x + move.distance
			case "R":
				direction = "W"
				x = x - move.distance
			}
		case "W":
			switch move.direction {
			case "L":
				direction = "S"
				y = y - move.distance
			case "R":
				direction = "N"
				y = y + move.distance
			}
		case "E":
			switch move.direction {
			case "L":
				direction = "N"
				y = y + move.distance
			case "R":
				direction = "S"
				y = y - move.distance
			}
		}
	}

	if verbose {
		log.Printf("Finished at [%d, %d] facing %s", x, y, direction)
	}

	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}
