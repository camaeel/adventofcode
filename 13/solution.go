package main

import (
	"../intcode"
	"fmt"
	"strconv"
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
	program[0] = 2 // start game part 2
	computer := intcode.CreateIntcode(program, 10000, true)

	computer.Run(false)

	tiles := make([]Tile, 0)
	working := true

	score := 0

	buffer := make([]int, 0)

	for working == true {

		select {
		case c, working := <-computer.Output:
			if working == false {
				break
			}

			buffer = append(buffer, c)

			if len(buffer) == 3 {

				x, y, type1 := buffer[0], buffer[1], buffer[2]
				found := false
				if x == -1 && y == 0 {
					score = type1
				} else {
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
				//empty buffer
				buffer = make([]int, 0)
			}

		case <-computer.InputNeeded:
			ball := FindTileByType(tiles, Ball)
			paddle := FindTileByType(tiles, Paddle)
			decision := 0
			if ball.X < paddle.X {
				decision = -1
			} else if ball.X > paddle.X {
				decision = 1
			}
			fmt.Println("Draw")
			fmt.Println(Draw(tiles, score))
			computer.Input <- decision
			fmt.Println("Decision:", decision)
			//make decision
			//publish move
		}

	}

	blocks := 0
	for _, v := range tiles {
		if v.Type == Block {
			blocks++
		}
	}
	// fmt.Println("Blocks:", blocks)

	fmt.Println("Draw")
	fmt.Println(Draw(tiles, score))
}

func FindTileByType(tiles []Tile, ttype TileType) Tile {
	for _, v := range tiles {
		if v.Type == ttype {
			return v
		}
	}
	return Tile{X: -10000, Y: -10000}
}

func Draw(tiles []Tile, score int) string {
	ret := ""
	ret += "\n"
	ret += strconv.Itoa(score)
	ret += "\n"
	minX, minY := tiles[0].X, tiles[0].Y
	maxX, maxY := tiles[0].X, tiles[0].Y
	for _, v := range tiles {
		if minX > v.X {
			minX = v.X
		}
		if minY > v.Y {
			minY = v.Y
		}
		if maxX < v.X {
			maxX = v.X
		}
		if maxY < v.Y {
			maxY = v.Y
		}
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			for _, v := range tiles {
				if x == v.X && y == v.Y {
					ret += strconv.Itoa(int(v.Type))
				}
			}
		}
		ret += "\n"
	}

	return ret
}
