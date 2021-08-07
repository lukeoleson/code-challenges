package main

import (
	"math"
	"sort"
)

func main() {
	//balance()
}

type Transaction struct {
	From string
	To string
	Amount float64
}

func balance(costsByPerson []float64) []Transaction {

	//
	// Get total cost of trip
	//
	var tripTotal float64
	people := len(costsByPerson)
	for _, cost := range costsByPerson {
		tripTotal += cost
	}

	//
	// Get amount each person should pay to be fair
	//
	split := math.Abs(tripTotal / float64(people))

	//
	// Calculate how much each person over/under paid
	//
	for i, cost := range costsByPerson {
		costsByPerson[i] = cost + split
	}
	sort.Float64s(costsByPerson)

	//
	// Redistribute funds from greatest underpaid to greatest overpaid
	//
	i := 0
	j := len(costsByPerson) - 1
	// TODO: Implement returning a slice of Transactions indicating how much goes from who to whom
	for i < j {
		if costsByPerson[i] == 0 {
			i++
		}
		if costsByPerson[j] == 0 {
			j--
		}
		if equalEnough(costsByPerson[i], costsByPerson[j]) {
			costsByPerson[i] = 0
			costsByPerson[j] = 0
			i++
			j--
		}
		if math.Abs(costsByPerson[j]) > math.Abs(costsByPerson[i]) {
			costsByPerson[j] = costsByPerson[j] + costsByPerson[i]
			costsByPerson[i] = 0
			i++
		}
		if math.Abs(costsByPerson[j]) < math.Abs(costsByPerson[i]) {
			costsByPerson[i] = costsByPerson[i] + costsByPerson[j]
			costsByPerson[j] = 0
			j--
		}
	}

	return []Transaction{
	{
		 From:   "Luke",
		 To:     "Justin",
		 Amount: 51.00,
	},
	}
}

func equalEnough(i, j float64) bool {
	if math.Abs(i) - math.Abs(j) > 0.001 {
		return false
	}
	return true
}

//     vv
// [4, 7, 9, 12, 16, 22]
// target: 16
// 26-20
// (4, 12), (7, 9)

//     vv
// [0, 4, 7, 9, 12, 16, 22]
// target: 16
// (0, 16), (4, 12)

// Find the total of the trip
// Find how much everyone should have spent
// Find out each person's +/-
// sort spent least to spent most
// get how much is available from the first person
// work your way back transfering it to each person until you run out.
// move to the next person (second spot) take as much as you can or as much as you need from the first, give to the last.
// repeat until you hit the middle

// {"Luke": -200, "Justin": -100}
// -300
// -300 / 2 = -150
// {"Luke": -50, "Justin": 50}
// {"Justin": 50, "Luke": -50}
// have 50 to distribute