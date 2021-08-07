package rockpaperscissors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetermineWinner(t *testing.T) {

	testCases := []struct {
		Name       string
		HaveP1     player
		HaveP2     player
		WantResult string
		WantWinner player
	}{
		{
			Name: "p1: rock, p2: rock",
			HaveP1: player{
				Name:   "p1",
				Choice: "rock",
			},
			HaveP2: player{
				Name:   "p2",
				Choice: "rock",
			},
			WantResult: winningExpressions["tiegame"],
			WantWinner: player{
				Name: "Nobody",
			},
		},
		{
			Name: "p1: rock, p2: paper",
			HaveP1: player{
				Name:   "p1",
				Choice: "rock",
			},
			HaveP2: player{
				Name:   "p2",
				Choice: "paper",
			},
			WantResult: winningExpressions["paperrock"],
			WantWinner: player{
				Name:   "p2",
				Choice: "paper",
			},
		},
		{
			Name: "p1: rock, p2: scissors",
			HaveP1: player{
				Name:   "p1",
				Choice: "rock",
			},
			HaveP2: player{
				Name:   "p2",
				Choice: "scissors",
			},
			WantResult: winningExpressions["rockscissors"],
			WantWinner: player{
				Name:   "p1",
				Choice: "rock",
			},
		},
		// paper, paper
		// paper, rock
		// paper, scissors
		// scissors, scissors
		// scissors, rock
		// scissors, paper
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// test
			gotWinner, gotResult := determineWinner(tc.HaveP1, tc.HaveP2)

			// assert
			assert.Equal(t, tc.WantResult, gotResult, "result mismatch")
			assert.Equal(t, tc.WantWinner.Name, gotWinner.Name, "result mismatch")
			assert.Equal(t, tc.WantWinner.Choice, gotWinner.Choice, "result mismatch")
		})
	}
}
