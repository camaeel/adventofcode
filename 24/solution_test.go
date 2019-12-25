package main

import (
	"strings"
	"testing"
)

func TestProgressTime(t *testing.T) {
	input := `....#
#..#.
#..##
..#..
#....`
	expected := `#..#.
####.
###.#
##.##
.##..`
	result := ProgressTime(input)

	if strings.Trim(result, "\n") != strings.Trim(expected, "\n") {
		t.Error("Wrong result. Got:", result, ", expected: ", expected)
	}

}

func TestCalucalteBiodiv(t *testing.T) {
	input := `.....
.....
.....
#....
.#...`
	expected := 2129920
	result := CalculateBiodiversity(input)
	if result != expected {
		t.Error("Wrong result. Got:", result, ", expected: ", expected)
	}
}
