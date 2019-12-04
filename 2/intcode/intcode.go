package intcode

import "strconv"

func Intcode(program []int) []int {

	for counter := 0; program[counter] != 99; {
		if program[counter] != 1 && program[counter] != 2 {
			panic("Illegal operand: " + strconv.Itoa(program[counter]) + ", at counter: " + strconv.Itoa(counter))
		}
		srcIndex1 := program[counter+1]
		srcIndex2 := program[counter+2]
		tgtIndex := program[counter+3]

		if program[counter] == 1 {
			program[tgtIndex] = program[srcIndex1] + program[srcIndex2]
		} else if program[counter] == 2 {
			program[tgtIndex] = program[srcIndex1] * program[srcIndex2]
		}

		counter += 4
	}
	return program

}
