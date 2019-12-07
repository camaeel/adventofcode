package main

import (
	"../5/intcode"
	"fmt"
)

func main() {
	program := intcode.ReadProgram("input.txt")
	maxValue, maxSettings := MaxValue(program)
	fmt.Println("max value=", maxValue, " max settings: ", maxSettings)
}

func MaxValue(program []int) (int, []int) {
	maxValue := 0
	maxSettings := make([]int, 0)

	for a1 := 0; a1 < 5; a1++ {
		for a2 := 0; a2 < 5; a2++ {
			if a2 != a1 {
				for a3 := 0; a3 < 5; a3++ {
					if a3 != a2 && a3 != a1 {
						for a4 := 0; a4 < 5; a4++ {
							if a4 != a3 && a4 != a2 && a4 != a1 {
								for a5 := 0; a5 < 5; a5++ {
									if a5 != a4 && a5 != a3 && a5 != a2 && a5 != a1 {
										settings := []int{a1, a2, a3, a4, a5}
										value := Amplifiers(program, settings)
										if value > maxValue {
											maxValue = value
											maxSettings = settings
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return maxValue, maxSettings
}

func Amplifiers(program []int, settings []int) int {
	result := 0
	for i := 0; i < 5; i++ {
		progCpy := intcode.CloneProgram(program)
		input := []int{settings[i], result}
		output := intcode.Intcode(progCpy, input, false)
		result = output[0]
		// fmt.Println("settings:", settings, ", input:", input, ", output:", output, ", result:", result)
	}
	return result
}
