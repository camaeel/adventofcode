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

	comp2 := intcode.CreateIntcode(program, 10000)
	comp2.Run(false)

	comp2.Input <- (2)
	result2 := <-comp2.Output

	fmt.Println("Part 2:", result2)
}
