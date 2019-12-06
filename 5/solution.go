package main

import (
	"./intcode"
	"fmt"
)

func main() {
	program := intcode.ReadProgram("input.txt")
	fmt.Println("Input program: ", program)
	//debug settings

	tmpProg := intcode.CloneProgram(program)

	input := []int{1}

	result := intcode.Intcode(tmpProg, input)
	fmt.Println("Input program: ", program)
	fmt.Println("Resulting program part 1: ", result)

	output := intcode.Intcode(tmpProg, input)

	fmt.Println("Output: ", output)

	// for noun := 0; noun < 100; noun++ {
	// 	for verb := 0; verb < 100; verb++ {
	// 		if GetResult(program, noun, verb) == 19690720 {
	// 			fmt.Println("Got: ", 19690720, ", verb= ", verb, ", noun=", noun)
	// 			fmt.Println("Rsult : ", noun*100+verb)
	// 		}
	// 	}
	// }
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
