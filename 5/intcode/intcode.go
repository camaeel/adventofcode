package intcode

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Intcode(program []int, input []int) []int {

	output := make([]int, 0)
	counter := 0
	inputCounter := 0

	for {
		if Opcode(program, &counter, input, &inputCounter, output) == false {
			return output
		}
	}

}

// returns false if program finished
func Opcode(program []int, counter *int, input []int, inputCounter *int, output []int) bool {
	opcode, modes := DecodeOpcode(program[*counter])
	*counter++
	switch opcode {
	case 99:
		return false
	case 1:
		op1 := GetOperand(program, counter, modes[0])
		*counter++
		op2 := GetOperand(program, counter, modes[1])
		*counter++
		SetOperand(program, counter, op1+op2)
		*counter++
	case 2:
		op1 := GetOperand(program, counter, modes[0])
		*counter++
		op2 := GetOperand(program, counter, modes[1])
		*counter++
		SetOperand(program, counter, op1*op2)
		*counter++
	case 3:
		SetOperand(program, counter, input[*inputCounter])
		*inputCounter++
		*counter++
	case 4:
		op1 := GetOperand(program, counter, modes[0])
		*counter++
		output = append(output, op1)
	default:
		panic("Unknown opcode")
	}
	return true
}

func DecodeOpcode(oc int) (int, []int) {
	opcode := oc % 100
	mode := []int{(oc / 100) % 10, (oc / 1000) % 10, (oc / 10000) % 10}
	return opcode, mode
}

func SetOperand(program []int, counter *int, value int) {
	program[program[*counter]] = value
}

func GetOperand(program []int, counter *int, mode int) int {
	if mode == 0 {
		return program[program[*counter]]
	} else if mode == 1 {
		return program[*counter]
	} else {
		panic("Unknown mode")
	}
}

func ReadProgram(fileName string) []int {
	program := make([]int, 0)
	f, err := os.Open(fileName)
	defer f.Close()
	check(err)

	scanner := bufio.NewScanner(f)
	line := ""
	for scanner.Scan() {
		line += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	arr := strings.Split(line, ",")
	for _, v := range arr {
		num, err := strconv.Atoi(v)
		check(err)
		program = append(program, num)
	}

	return program
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CloneProgram(program []int) []int {
	tmpProg := make([]int, len(program))
	copy(tmpProg, program)
	return tmpProg
}
