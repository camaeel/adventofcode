package main

import (
	// "fmt"
	"fmt"
	"testing"
)

func TestCalculateAngle3Quarter(t *testing.T) {
	laser := Point{X: 4, Y: 3}
	target := Point{X: 0, Y: 0}
	result := CalculateAngle(laser, target)
	expectedMin := 180.0
	expectedMax := 270.0

	// fmt.Println("Expected angle between: ", expectedMin, " and ", expectedMax, " Got:", result)
	if !(result < expectedMax && result > expectedMin) {
		t.Error("Error. Expected angle between: ", expectedMin, " and ", expectedMax, " Got:", result)
	}
}

func TestCalculateAngle1Quarter(t *testing.T) {
	laser := Point{X: 4, Y: 3}
	target := Point{X: 10, Y: 10}
	result := CalculateAngle(laser, target)
	expectedMin := 0.0
	expectedMax := 90.0

	// fmt.Println("Expected angle between: ", expectedMin, " and ", expectedMax, " Got:", result)
	if !(result < expectedMax && result > expectedMin) {
		t.Error("Error. Expected angle between: ", expectedMin, " and ", expectedMax, " Got:", result)
	}
}

func TestCalculateAngle2Quarter(t *testing.T) {
	laser := Point{X: 4, Y: 3}
	target := Point{X: 6, Y: 0}
	result := CalculateAngle(laser, target)
	expectedMin := 90.0
	expectedMax := 180.0

	// fmt.Println("Expected angle between: ", expectedMin, " and ", expectedMax, " Got:", result)
	if !(result < expectedMax && result > expectedMin) {
		t.Error("Error. Expected angle between: ", expectedMin, " and ", expectedMax, " Got:", result)
	}
}

func TestCalculateAngle4Quarter(t *testing.T) {
	laser := Point{X: 4, Y: 3}
	target := Point{X: 1, Y: 6}
	result := CalculateAngle(laser, target)
	expectedMin := 270.0
	expectedMax := 360.0

	// fmt.Println("Expected angle between: ", expectedMin, " and ", expectedMax, " Got:", result)
	if !(result < expectedMax && result > expectedMin) {
		t.Error("Error. Expected angle between: ", expectedMin, " and ", expectedMax, " Got:", result)
	}
}

func TestCalculateAngleUp(t *testing.T) {
	laser := Point{X: 4, Y: 3}
	target := Point{X: 4, Y: 6}
	result := CalculateAngle(laser, target)
	expected := 0.0

	// fmt.Println("Expected angle between: ", expectedMin, " and ", expectedMax, " Got:", result)
	if !(result == expected) {
		t.Error("Error. Expected angle : ", expected, " Got:", result)
	}
}

func TestPart1(t *testing.T) {
	dataStr := `......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`
	mapa := ParseMap(dataStr)
	maxX, maxY, maxValue := FindBestSpot(mapa)
	expectedMaxX, expectedMaxY, expectedValue := 5, 8, 33
	if maxX != expectedMaxX || expectedMaxY != maxY || expectedValue != maxValue {
		t.Error("Error. Expected maxX,maxY  : ", expectedMaxX, expectedMaxY, " Got:", maxX, maxY)
		t.Error("Error. Expected mexpectedValue  : ", expectedValue, " Got: ", maxValue)
	}
	fmt.Println("maxValue=", maxValue)
}

func TestPart2(t *testing.T) {
	dataStr := `.#....#####...#..
##...##.#####..##
##...#...#.#####.
..#.....#...###..
..#.#.....#....##`
	mapa := ParseMap(dataStr)
	maxX, maxY, maxValue := FindBestSpot(mapa)
	expectedMaxX, expectedMaxY := 8, 3
	if maxX != expectedMaxX || expectedMaxY != maxY {
		t.Error("Error. Expected maxX,maxY  : ", expectedMaxX, expectedMaxY, " Got:", maxX, maxY)
	}
	fmt.Println("maxValue=", maxValue)
}
