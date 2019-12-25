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
