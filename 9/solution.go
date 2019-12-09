package main

import (
	"fmt"

	"../intcode"
)

func main() {
	program := intcode.ReadProgram("input.txt")
	comp := intcode.CreateIntcode(program, 10000)
	comp.Run(false)

	comp.Input <- (1)
	result := <-comp.Output

	fmt.Println("Part 1:", result)
}
