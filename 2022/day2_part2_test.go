package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay2P2(t *testing.T) {
	list := `A Y
B X
C Z`

	input := parseTournamentInputP2(list)

	assert.Len(t, input, 3)
	assert.Equal(t, input[0], RoundP2{OpponentMove: OpponentRock, EndResult: ShouldDraw})
	assert.Equal(t, input[1], RoundP2{OpponentMove: OpponentPaper, EndResult: ShouldLose})
	assert.Equal(t, input[2], RoundP2{OpponentMove: OpponentScissors, EndResult: ShouldWin})

	totalScore := playTournamentP2(input)

	assert.Equal(t, 12, totalScore)
}

// The example above now goes like this:
// In the first round, your opponent will choose Rock (A), and you need the round to end in a draw (Y), so you also choose Rock. This gives you a score of 1 + 3 = 4.
// In the second round, your opponent will choose Paper (B), and you choose Rock so you lose (X) with a score of 1 + 0 = 1.
// In the third round, you will defeat your opponent's Scissors with Rock for a score of 1 + 6 = 7.
// Now that you're correctly decrypting the ultra top secret strategy guide, you would get a total score of 12.
