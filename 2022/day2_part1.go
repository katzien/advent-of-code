package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	OpponentRock     = "A"
	OpponentPaper    = "B"
	OpponentScissors = "C"

	MeRock     = "X"
	MePaper    = "Y"
	MeScissors = "Z"

	ScoreRock     = 1
	ScorePaper    = 2
	ScoreScissors = 3

	ScoreLoss = 0
	ScoreDraw = 3
	ScoreWin  = 6
)

type Round struct {
	OpponentMove string
	MyMove       string
}

func Day2P1() {
	list, err := os.ReadFile("./input/day_2")
	if err != nil {
		panic(err)
	}

	input := parseTournamentInput(string(list))

	totalScore := playTournament(input)

	fmt.Printf("You would score %d\n", totalScore)
}

func parseTournamentInput(raw string) []Round {
	lines := strings.Split(raw, "\n")

	input := make([]Round, len(lines))

	for k, l := range lines {
		moves := strings.Split(l, " ")
		if len(moves) != 2 {
			panic(fmt.Sprintf("invalid moves line: %s", l))
		}

		input[k] = Round{OpponentMove: moves[0], MyMove: moves[1]}
	}

	return input
}

// The winner of the whole tournament is the player with the highest score.
// Your total score is the sum of your scores for each round.
// The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors)
// plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).
// Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock.
// If both players choose the same shape, the round instead ends in a draw.
func playTournament(input []Round) int {
	totalScore := 0

	roundScore := 0
	for _, round := range input {
		switch round.OpponentMove {
		case OpponentRock:
			switch round.MyMove {
			case MeRock:
				roundScore = ScoreRock + ScoreDraw
			case MePaper:
				roundScore = ScorePaper + ScoreWin
			case MeScissors:
				roundScore = ScoreScissors + ScoreLoss
			}
		case OpponentPaper:
			switch round.MyMove {
			case MeRock:
				roundScore = ScoreRock + ScoreLoss
			case MePaper:
				roundScore = ScorePaper + ScoreDraw
			case MeScissors:
				roundScore = ScoreScissors + ScoreWin
			}
		case OpponentScissors:
			switch round.MyMove {
			case MeRock:
				roundScore = ScoreRock + ScoreWin
			case MePaper:
				roundScore = ScorePaper + ScoreLoss
			case MeScissors:
				roundScore = ScoreScissors + ScoreDraw
			}
		}

		totalScore += roundScore
		roundScore = 0
	}

	return totalScore
}
