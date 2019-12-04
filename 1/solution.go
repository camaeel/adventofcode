package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	check(err)

	var totalFuel, totalFuelPart2 int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		weight, err1 := strconv.Atoi(line)
		check(err1)
		totalFuel += CalcFuel(weight)
		totalFuelPart2 += CalcFuelWithFuelWeight(weight)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println("Part 1:")
	fmt.Println(totalFuel)

	fmt.Println("Part 2:")
	fmt.Println(totalFuelPart2)
	defer f.Close()
}

func CalcFuel(w int) int {
	return w/3 - 2
}

func CalcFuelWithFuelWeight(w int) int {
	f := CalcFuel(w)
	if f <= 0 {
		return 0
	}
	return f + CalcFuelWithFuelWeight(f)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
