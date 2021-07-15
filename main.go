package main

import (
	"fmt"

	"github.com/willsem/long-chain-of-ones/solve"
)

func main() {
	var n int
	fmt.Scan(&n)

	sequence := make([]byte, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&sequence[i])
	}

	for i := 0; i < n; i++ {
		if sequence[i] != 0 && sequence[i] != 1 {
			fmt.Println("All values should be 0 or 1")
			return
		}
	}

	fmt.Println("Sequence: ", sequence)
	fmt.Println("Answer: ", solve.MaxOnesAfterRemoveItem(sequence))
}
