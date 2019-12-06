package main

import (
	"fmt"
)

func main() {
	min := 272091
	max := 815432
	count := 0
	for i := min; i <= max; i++ {
		if IsPossiblePassword(i) {
			count++
			fmt.Println(i)
		}
	}
	fmt.Println("Possible passwords number= ", count)

}

func IsPossiblePassword(i int) bool {
	digits := make([]int, 0)
	for num := i; num > 0; num = num / 10 {
		digits = append(digits, num%10)
	}

	// invert
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	// fmt.Println(i, " => ", digits)
	same := 1
	sameSatisfied := false
	for i := 1; i < len(digits); i++ {
		// fmt.Println("digit[", i-1, "]=", digits[i-1], " digit[", i, "]=", digits[i])
		// fmt.Println("same=", same)
		if digits[i] < digits[i-1] {
			// fmt.Println("  lower than prev")
			return false
		} else if digits[i] == digits[i-1] {
			// fmt.Println("  same")
			same++

		} else {
			// fmt.Println("  higher than prev")
			if same == 2 {
				sameSatisfied = true
			}
			// } else {
			// 	// fmt.Println("  reset same")
			// 	same = 1
			// }
			same = 1
		}

	}
	if same == 2 {
		sameSatisfied = true

	}

	return sameSatisfied
}
