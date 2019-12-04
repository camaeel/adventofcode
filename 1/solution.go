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

	var totalFuel int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		weight, err1 := strconv.Atoi(line)
		check(err1)
		totalFuel += CalcFuel(weight)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(totalFuel)
	defer f.Close()
}

func CalcFuel(w int) int {
	return w/3 - 2
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
