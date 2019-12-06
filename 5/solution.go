package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"./intcode"
)

func main() {
	program := ReadProgram("input.txt")
	fmt.Println("Input program: ", program)
	//debug settings

	tmpProg := CloneProgram(program)
	tmpProg[1] = 12
	tmpProg[2] = 2

	result := intcode.Intcode(tmpProg)
	fmt.Println("Input program: ", program)
	fmt.Println("Resulting program part 1: ", result)

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			if GetResult(program, noun, verb) == 19690720 {
				fmt.Println("Got: ", 19690720, ", verb= ", verb, ", noun=", noun)
				fmt.Println("Rsult : ", noun*100+verb)
			}
		}
	}
}

func GetResult(program []int, noun int, verb int) int {
	tmpProg := make([]int, len(program))
	copy(tmpProg, program)
	tmpProg[1] = noun
	tmpProg[2] = verb
	intcode.Intcode(tmpProg)

	return tmpProg[0]
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
