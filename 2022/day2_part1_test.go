package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay2P1(t *testing.T) {
	list := `A Y
B X
C Z`

	input := parseTournamentInput(list)

	assert.Len(t, input, 3)
	assert.Equal(t, input[0], Round{OpponentMove: OpponentRock, MyMove: MePaper})
	assert.Equal(t, input[1], Round{OpponentMove: OpponentPaper, MyMove: MeRock})
	assert.Equal(t, input[2], Round{OpponentMove: OpponentScissors, MyMove: MeScissors})

	totalScore := playTournament(input)

	assert.Equal(t, 15, totalScore)
}

//In the first round, your opponent will choose Rock (A), and you should choose Paper (Y).
//This ends in a win for you with a score of 8 (2 because you chose Paper + 6 because you won).
//In the second round, your opponent will choose Paper (B), and you should choose Rock (X).
//This ends in a loss for you with a score of 1 (1 + 0).
//The third round is a draw with both players choosing Scissors, giving you a score of 3 + 3 = 6.
//In this example, if you were to follow the strategy guide, you would get a total score of 15 (8 + 1 + 6).
