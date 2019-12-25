package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	bugMaps := make([]string, 0)
	bugs := PuzzleInput()
	bugMaps = append(bugMaps, strings.Trim(bugs, "\n"))

	found := false
	for found == false {
		bugs = strings.Trim(ProgressTime(bugs), "\n")
		for _, v := range bugMaps {
			if bugs == v {
				found = true
				fmt.Println("cases: ", len(bugMaps))
				break
			}
		}
		bugMaps = append(bugMaps, strings.Trim(bugs, "\n"))
	}
	fmt.Println("Dupliaated map: ", bugs)
	fmt.Println("biodiversity: ", CalculateBiodiversity(bugs))
}

func PuzzleInput() string {
	return `##.##
#.##.
#...#
##.#.
####.`
}

func CalculateBiodiversity(bugs string) int {
	result := 0
	bugs = strings.ReplaceAll(bugs, " ", "")
	bugs = strings.ReplaceAll(bugs, "\n", "")
	for i, v := range bugs {
		if v == '#' {
			result += int(math.Pow(2.0, float64(i)))
		}
	}
	return result
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
