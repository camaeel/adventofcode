package main

import (
	"../intcode"
	"fmt"
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
	position := Point{0, 0, 1}
	direction := Up

	minPosition := position
	maxPosition := position

	finished := false

	for finished == false {
		if len(painted) > 0 {
			computer.Input <- GetColor(painted, position)
		} else {
			computer.Input <- (1)
		}

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

			if position.X < minPosition.X {
				minPosition.X = position.X
			}
			if position.Y < minPosition.Y {
				minPosition.Y = position.Y
			}
			if position.X > maxPosition.X {
				maxPosition.X = position.X
			}
			if position.Y > maxPosition.Y {
				maxPosition.Y = position.Y
			}

		} else {
			finished = true
		}
	}

	// fmt.Println("result:", computer.OutputArray)
	fmt.Println("len result:", len(painted))
	fmt.Println(Draw(painted, minPosition, maxPosition))
}

func Draw(painted []*Point, min Point, max Point) string {
	result := ""
	for y := max.Y; y >= min.Y; y-- {
		for x := min.X; x <= max.X; x++ {
			color := GetColor(painted, Point{X: x, Y: y})
			if color == 1 {
				result += "#"
			} else {
				result += "."
			}
		}
		result += "\n"
	}
	return result
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
