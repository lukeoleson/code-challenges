package rockpaperscissors

import (
	"fmt"
	"math/rand"
	"time"
)

var choices = []string{"rock", "paper", "scissors"}

var winningExpressions = map[string]string{
	"paperrock":     "Paper covers rock!",
	"rockscissors":  "Rock smashes scissors!",
	"scissorspaper": "Scissors cut paper!",
	"tiegame":       "TIE GAME!",
}

var winDeterminer = map[string]map[string]bool{
	"rock":     {"paper": false, "scissors": true},
	"paper":    {"rock": true, "scissors": false},
	"scissors": {"rock": false, "paper": true},
}

func StartGame() {
	prompt()
	p1, p2 := makeChoices()
	winner, result := determineWinner(p1, p2)
	printResults(result, p1, p2, winner)
}

func printResults(result string, p1, p2, winner player) {
	fmt.Printf("%v chose %v\n", p1.Name, p1.Choice)
	fmt.Printf("%v chose %v\n", p2.Name, p2.Choice)
	fmt.Printf("\n\n%v -- %v is the winner!\n", result, winner.Name)
}

func makeChoices() (player, player) {
	src := rand.NewSource(time.Now().Unix())
	chooser := rand.New(src)

	choice := chooser.Intn(len(choices))
	p1 := player{
		Name:   "Luke",
		Choice: choices[choice],
	}

	choice = chooser.Intn(len(choices))
	p2 := player{
		Name:   "Kim",
		Choice: choices[choice],
	}

	return p1, p2
}

func prompt() {
	for i := 3; i > 0; i-- {
		for j := 0; j < 2; j++ {
			time.Sleep(time.Millisecond * 250)
			fmt.Print(".")
		}
		time.Sleep(time.Millisecond * 250)
		fmt.Printf(" %v ", i)
	}
	time.Sleep(time.Millisecond * 250)
	fmt.Print(".. GO!\n\n")
	time.Sleep(time.Millisecond * 250)

}

func determineWinner(p1, p2 player) (player, string) {

	if p1.Choice == p2.Choice {
		return player{Name: "Nobody"}, winningExpressions["tiegame"]
	}

	p1Options := winDeterminer[p1.Choice]
	p1Won := p1Options[p2.Choice]

	if p1Won {
		return p1, winningExpressions[p1.Choice+p2.Choice]
	}
	return p2, winningExpressions[p2.Choice+p1.Choice]
}

type player struct {
	Name   string
	Choice string
}
