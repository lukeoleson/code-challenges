package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBalance(t *testing.T) {

	testCases := []struct{
		Name string
		Have []float64
		Want []Transaction
		WantTemp []float64
	}{
		//{
		//	Name: "Simple",
		//	Have: []float64 {
		//		-100.00,
		//		-200.00,
		//	},
		//	Want: []Transaction{
		//		{
		//			From: "Luke",
		//			To: "Justin",
		//			Amount: 50.00,
		//		},
		//	},
		//	WantTemp: []float64{0, 0},
		//},
		{
			Name: "Simple++",
			Have: []float64 {
				-100.00,
				-200.00,
				-50.00,
			},
			Want: []Transaction{
				{
					From: "Luke",
					To: "Justin",
					Amount: 50.00,
				},
			},
			WantTemp: []float64{0, 0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			got := balance(tc.Have)
			for i, transaction := range tc.Want {
				assert.Equal(t, transaction.From, got[i].From)
				assert.Equal(t, transaction.To, got[i].To)
				assert.Equal(t, transaction.Amount, got[i].Amount)
			}
		})
	}


}