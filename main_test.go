package main

import (
	"testing"

	"github.com/willsem/long-chain-of-ones/solve"
)

var sequences = [][]byte{
	{0, 0},
	{0, 1},
	{1, 0},
	{1, 1},
	{1, 1, 0, 1, 1},
	{1, 1, 0, 1, 1, 0, 1, 1, 1},
	{1, 1, 0, 1, 1, 0, 1, 1, 1, 0},
	{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
}

var results = []uint{
	0,
	1,
	1,
	1,
	4,
	5,
	5,
	2,
}

func TestMaxOnes(t *testing.T) {
	for i := 0; i < len(sequences); i++ {
		sequence := sequences[i]
		result := solve.MaxOnesAfterRemoveItem(sequence)
		expected := results[i]

		if result != expected {
			t.Error("Test", i+1, "\nData: ", sequence, "\nAsnwer: ", result, "\nExpected: ", expected, "\n")
		}
	}
}
