package main

import (
	"fmt"

	"../intcode"
)

type Direction int

const (
	Up    Direction = 0
	Right Direction = 1
	Down  Direction = 2
	Left  Direction = 3
)

type Point struct {
	X     int
	Y     int
	Color int
}

func main() {
	program := intcode.ReadProgram("input.txt")
	computer := intcode.CreateIntcode(program, 10000)
	computer.Run(false)

	painted := make([]*Point, 0)
	position := Point{0, 0, 0}
	direction := Up

	finished := false

	for finished == false {
		computer.Input <- GetColor(painted, position)
		color, result1 := <-computer.Output
		dirCmd, result2 := <-computer.Output

		if result1 == true && result2 == true {

			position.Color = color

			found := false
			for _, v := range painted {
				if v.X == position.X && v.Y == position.Y {
					v.Color = color
					found = true
					break
				}
			}
			if !found {
				painted = append(painted, &Point{X: position.X, Y: position.Y, Color: color})
			}
			position, direction = Move(position, direction, dirCmd)
		} else {
			finished = true
		}
	}

	// fmt.Println("result:", computer.OutputArray)
	fmt.Println("len result:", len(painted))
}

func GetColor(painted []*Point, position Point) int {
	for _, p := range painted {
		if p.X == position.X && p.Y == position.Y {
			return p.Color
		}
	}
	return 0
}

func Move(position Point, direction Direction, dirCmd int) (Point, Direction) {
	if dirCmd == 1 {
		direction = (direction + 1) % 4
	} else if dirCmd == 0 {
		direction = (direction + 3) % 4
	} else {
		panic("Unknown direction")
	}
	if direction == Up {
		position.Y++
	} else if direction == Down {
		position.Y--
	} else if direction == Left {
		position.X--
	} else if direction == Right {
		position.X++
	}
	return position, direction
}
