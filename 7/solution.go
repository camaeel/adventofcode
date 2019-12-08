package main

import (
	"../intcode"
	"fmt"
)

func main() {
	program := intcode.ReadProgram("input.txt")
	maxValue, maxSettings := MaxValue(program)
	fmt.Println("max value=", maxValue, " max settings: ", maxSettings)
	maxValuePart2, maxSettingsPart2 := MaxValuePart2(program)
	fmt.Println("Part2: max value=", maxValuePart2, " max settings: ", maxSettingsPart2)
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

func MaxValuePart2(program []int) (int, []int) {
	min, max := 5, 10

	maxValue := 0
	maxSettings := make([]int, 0)

	for a1 := min; a1 < max; a1++ {
		for a2 := min; a2 < max; a2++ {
			if a2 != a1 {
				for a3 := min; a3 < max; a3++ {
					if a3 != a2 && a3 != a1 {
						for a4 := min; a4 < max; a4++ {
							if a4 != a3 && a4 != a2 && a4 != a1 {
								for a5 := min; a5 < max; a5++ {
									if a5 != a4 && a5 != a3 && a5 != a2 && a5 != a1 {
										settings := []int{a1, a2, a3, a4, a5}
										value := AmplifiersPart2(program, settings)
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
	numOfComputers := 5
	computers := make([]*intcode.Intcode, numOfComputers)
	for i := 0; i < numOfComputers; i++ {
		computers[i] = intcode.CreateIntcode(program)
		if i > 0 {
			computers[i].Input = computers[i-1].Output
		}
	}
	for _, v := range computers {
		v.Run(false)
	}
	for i, v := range computers {
		v.Input <- settings[i]
	}
	computers[0].Input <- 0
	result = <-computers[len(computers)-1].Output
	return result
}

func AmplifiersPart2(program []int, settings []int) int {
	result := 0
	numOfComputers := 5
	computers := make([]*intcode.Intcode, numOfComputers)
	for i := 0; i < numOfComputers; i++ {
		computers[i] = intcode.CreateIntcode(program)
	}
	for i := 0; i < numOfComputers; i++ {
		computers[(i+1)%numOfComputers].Input = computers[i].Output
	}
	for _, v := range computers {
		v.Run(false)
	}
	for i, v := range computers {
		v.Input <- settings[i]
	}
	computers[0].Input <- 0
	computers[len(computers)-1].WaitGroup.Wait()
	// result = <-computers[len(computers)-1].Output
	result = computers[len(computers)-1].OutputArray[len(computers[len(computers)-1].OutputArray)-1]
	return result
}
