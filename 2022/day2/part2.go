package day2

import (
	"fmt"
	"os"
	"strings"
)

const (
	ShouldLose = "X"
	ShouldDraw = "Y"
	ShouldWin  = "Z"
)

type RoundP2 struct {
	OpponentMove string
	EndResult    string
}

func Part2() {
	list, err := os.ReadFile("./day2/input")
	if err != nil {
		panic(err)
	}

	input := parseTournamentInputP2(string(list))

	totalScore := playTournamentP2(input)

	fmt.Printf("You would score %d\n", totalScore)
}

func parseTournamentInputP2(raw string) []RoundP2 {
	lines := strings.Split(raw, "\n")

	input := make([]RoundP2, len(lines))

	for k, l := range lines {
		moves := strings.Split(l, " ")
		if len(moves) != 2 {
			panic(fmt.Sprintf("invalid moves line: %s", l))
		}

		input[k] = RoundP2{OpponentMove: moves[0], EndResult: moves[1]}
	}

	return input
}

// "Anyway, the second column says how the round needs to end: X means you need to lose,
// Y means you need to end the round in a draw, and Z means you need to win. Good luck!"
// The total score is still calculated in the same way, but now you need to figure out
// what shape to choose so the round ends as indicated.
// Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock.
func playTournamentP2(input []RoundP2) int {
	totalScore := 0

	roundScore := 0
	for _, round := range input {
		switch round.OpponentMove {
		case OpponentRock:
			switch round.EndResult {
			case ShouldLose:
				roundScore = ScoreScissors + ScoreLoss
			case ShouldDraw:
				roundScore = ScoreRock + ScoreDraw
			case ShouldWin:
				roundScore = ScorePaper + ScoreWin
			}
		case OpponentPaper:
			switch round.EndResult {
			case ShouldLose:
				roundScore = ScoreRock + ScoreLoss
			case ShouldDraw:
				roundScore = ScorePaper + ScoreDraw
			case ShouldWin:
				roundScore = ScoreScissors + ScoreWin
			}
		case OpponentScissors:
			switch round.EndResult {
			case ShouldLose:
				roundScore = ScorePaper + ScoreLoss
			case ShouldDraw:
				roundScore = ScoreScissors + ScoreDraw
			case ShouldWin:
				roundScore = ScoreRock + ScoreWin
			}
		}

		totalScore += roundScore
		roundScore = 0
	}

	return totalScore
}
