package main

import (
	"math"
)

func main() {

}

func DealIntoNewStack(deck []int) []int {
	res := make([]int, len(deck))

	for i := 0; i <= int(math.Ceil(float64(len(deck))/2.0)); i++ {
		res[i], res[len(deck)-1-i] = deck[len(deck)-1-i], deck[i]
	}
	return res
}
