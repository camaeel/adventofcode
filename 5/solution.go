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
	fmt.Println("Resulting program part 1: ", result)

	tmpProg2 := intcode.CloneProgram(program)

	input2 := []int{5}

	result2 := intcode.Intcode(tmpProg2, input2)
	fmt.Println("Resulting program part 2: ", result2)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
