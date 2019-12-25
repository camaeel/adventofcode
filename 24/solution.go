package main

import (
	"strings"
)

func main() {

}

func PuzzleInput() string {
	return `##.##
#.##.
#...#
##.#.
####.`
}

func ProgressTime(input string) string {
	result := ""
	rows := strings.Split(strings.Trim(input, "\n"), "\n")
	for y := 0; y < len(rows); y++ {
		for x := 0; x < len(rows[0]); x++ {
			neighbourBugs := CountNeighbourBugs(rows, x, y)
			if rows[y][x] == '#' {
				if neighbourBugs == 1 {
					result += "#"
				} else {
					result += "."
				}
			} else {
				if neighbourBugs == 1 || neighbourBugs == 2 {
					result += "#"
				} else {
					result += "."
				}
			}
		}
		result += "\n"
	}
	return result

}

func CountNeighbourBugs(rows []string, x int, y int) int {
	result := 0
	if x > 0 {
		if rows[y][x-1] == '#' {
			result++
		}
	}
	if y > 0 {
		if rows[y-1][x] == '#' {
			result++
		}
	}
	if x < (len(rows[0]) - 1) {
		if rows[y][x+1] == '#' {
			result++
		}
	}
	if y < (len(rows) - 1) {
		if rows[y+1][x] == '#' {
			result++
		}
	}

	return result
}
