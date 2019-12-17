package main

import (
	"fmt"
	"strings"

	"../intcode"
)

func main() {
	program := intcode.ReadProgram("input.txt")
	computer := intcode.CreateIntcode(program, 10000, false)

	computer.Run(false)
	mapa := ""
	for data, isOpen := <-computer.Output; isOpen == true; data, isOpen = <-computer.Output {
		mapa += string(rune(data))
	}
	// fmt.Println(mapa)

	alignement := 0
	rows := strings.Split(strings.Trim(mapa, "\n "), "\n")
	w, h := len(rows[0]), len(rows)
	fmt.Println(w, h)
	for y := 1; y < (h - 1); y++ {
		fmt.Println(rows[y])
		fmt.Println(rows[y+1])
		for x := 1; x < (w - 1); x++ {
			fmt.Println(x, y)
			if rows[y][x] == '#' {
				if rows[y-1][x] == '#' && rows[y][x-1] == '#' && rows[y+1][x] == '#' && rows[y][x+1] == '#' {
					//intersection
					alignement += x * y
				}
			}
		}
	}

	fmt.Println("part1:", alignement)
}
