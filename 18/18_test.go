package main

import (
	"testing"
)

func TestTC1(t *testing.T) {
	input := TC1Puzzle()
	rows := ParseMap(input)
	result := MinSteps(rows)
	if result != TC1Expected() {
		t.Error("ERRoR: Expected:", TC1Expected, ", got: ", result)
	}
}

func TestTC2(t *testing.T) {
	input := TC2Puzzle()
	rows := ParseMap(input)
	result := MinSteps(rows)
	if result != TC2Expected() {
		t.Error("ERRoR: Expected:", TC2Expected, ", got: ", result)
	}
}

func TestTC3(t *testing.T) {
	input := TC3Puzzle()
	rows := ParseMap(input)
	result := MinSteps(rows)
	if result != TC3Expected() {
		t.Error("ERRoR: Expected:", TC3Expected, ", got: ", result)
	}
}

func TestPhase1(t *testing.T) {
	input := PuzzleInput()
	rows := ParseMap(input)
	result := MinSteps(rows)
	if result != 1 {
		t.Error("ERRoR: Expected:", 1, ", got: ", result)
	}
}
