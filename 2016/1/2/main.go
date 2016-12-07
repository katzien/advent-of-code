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

	if d := distanceToFirstRevisitedLocation(input); d > 0 {
		log.Printf("The first location visited twice is %d blocks away.", distanceToFirstRevisitedLocation(input))
	} else {
		log.Println("It looks like we didn't visit any location twice.")
	}

}

func distanceToFirstRevisitedLocation(input string) int {
	direction := "N"
	location := position{x: 0, y: 0}

	moves := parseInput(input)

	list := locations{visited: make(map[position]struct{}, 0)}

	if verbose {
		log.Printf("Moves: %v", moves)
	}

	for _, move := range moves {

		if verbose {
			log.Printf("Currently at %v", location)
			log.Printf("Next move: %v", move)
		}

		if direction == "N" && move.direction == "L" || direction == "S" && move.direction == "R" {
			direction = "W"
			for i := 0; i < move.distance; i++ {
				location.x--
				if list.seen(location) {
					return distance(location)
				}
				list.add(location)
			}
		} else if direction == "N" && move.direction == "R" || direction == "S" && move.direction == "L" {
			direction = "E"
			for i := 0; i < move.distance; i++ {
				location.x++
				if list.seen(location) {
					return distance(location)
				}
				list.add(location)
			}
		} else if direction == "E" && move.direction == "L" || direction == "W" && move.direction == "R" {
			direction = "N"
			for i := 0; i < move.distance; i++ {
				location.y++
				if list.seen(location) {
					return distance(location)
				}
				list.add(location)
			}
		} else if direction == "W" && move.direction == "L" || direction == "E" && move.direction == "R" {
			direction = "S"
			for i := 0; i < move.distance; i++ {
				location.y--
				if list.seen(location) {
					return distance(location)
				}
				list.add(location)
			}
		}
	}

	if verbose {
		log.Printf("Finished at [%d, %d] facing %s", location.x, location.y, direction)
	}

	return 0
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

func distance(location position) int {
	return int(math.Abs(float64(location.x)) + math.Abs(float64(location.y)))
}

type move struct {
	direction string
	distance  int
}

type position struct {
	x, y int
}

type locations struct {
	visited map[position]struct{}
}

func (l *locations) add(place position) {
	l.visited[place] = struct{}{}
}

func (l *locations) seen(place position) bool {
	if verbose {
		log.Printf("Checking if we've been to %v already", place)
	}

	_, exists := l.visited[place]
	return exists
}
