package main

import (
	"../intcode"
	"fmt"
)

type TileType int

const (
	Empty TileType = iota
	Wall
	Block
	Paddle
	Ball
)

type Tile struct {
	X    int
	Y    int
	Type TileType
}

func main() {
	program := intcode.ReadProgram("input.txt")
	computer := intcode.CreateIntcode(program, 10000)
	computer.Run(false)

	tiles := make([]Tile, 0)
	working := true

	for working == true {
		var type1, x, y int
		x, working = <-computer.Output
		y, working = <-computer.Output
		type1, working = <-computer.Output

		found := false
		fmt.Println("X:", x, " Y:", y, " Type:", type1)
		for i, v := range tiles {
			if x == v.X && y == v.Y {
				tiles[i].Type = TileType(type1)
				found = true
				break
			}
		}
		if found == false {
			tiles = append(tiles, Tile{X: x, Y: y, Type: TileType(type1)})
		}

	}
	fmt.Println(tiles)

	blocks := 0
	for _, v := range tiles {
		if v.Type == Block {
			blocks++
		}
	}
	fmt.Println("Blocks:", blocks)
}
