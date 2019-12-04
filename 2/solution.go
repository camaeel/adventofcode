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
	program := readProgram()
	fmt.Println("Input program: ", program)
	//debug settings
	program[1] = 12
	program[2] = 2

	result := intcode.Intcode(program)
	fmt.Println("Resulting program: ", result)
}

func readProgram() []int {
	program := make([]int, 0)
	f, err := os.Open("input.txt")
	defer f.Close()
	check(err)

	scanner := bufio.NewScanner(f)
	line := ""
	for scanner.Scan() {
		line += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
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
