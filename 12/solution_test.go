package main

import (
	"fmt"
	"reflect"
	"testing"
)

func GetTc2() []Moon {

	return []Moon{
		Moon{Position: Coords{X: -8, Y: -10, Z: 0}},
		Moon{Position: Coords{X: 5, Y: 5, Z: 10}},
		Moon{Position: Coords{X: 2, Y: -7, Z: 3}},
		Moon{Position: Coords{X: 9, Y: -8, Z: -3}}}
}

func TestPart1Case(t *testing.T) {
	moons := GetTc1()

	Round(moons)
	round := 1
	expectedMoons1 := []Moon{
		Moon{Position: Coords{X: 2, Y: -1, Z: 1}, Velocity: Coords{X: 3, Y: -1, Z: -1}},
		Moon{Position: Coords{X: 3, Y: -7, Z: -4}, Velocity: Coords{X: 1, Y: 3, Z: 3}},
		Moon{Position: Coords{X: 1, Y: -7, Z: 5}, Velocity: Coords{X: -3, Y: 1, Z: -3}},
		Moon{Position: Coords{X: 2, Y: 2, Z: 0}, Velocity: Coords{X: -1, Y: -3, Z: 1}}}
	if !reflect.DeepEqual(moons, expectedMoons1) {
		t.Error("Round:", round)
		t.Error("Expected:\n" + MoonsToStr(expectedMoons1))
		t.Error("Got:\n" + MoonsToStr(moons))
	}

	Round(moons)
	round++
	expectedMoons2 := []Moon{
		Moon{Position: Coords{X: 5, Y: -3, Z: -1}, Velocity: Coords{X: 3, Y: -2, Z: -2}},
		Moon{Position: Coords{X: 1, Y: -2, Z: 2}, Velocity: Coords{X: -2, Y: 5, Z: 6}},
		Moon{Position: Coords{X: 1, Y: -4, Z: -1}, Velocity: Coords{X: 0, Y: 3, Z: -6}},
		Moon{Position: Coords{X: 1, Y: -4, Z: 2}, Velocity: Coords{X: -1, Y: -6, Z: 2}}}
	if !reflect.DeepEqual(moons, expectedMoons2) {
		t.Error("Round:", round)
		t.Error("Expected:\n" + MoonsToStr(expectedMoons2))
		t.Error("Got:\n" + MoonsToStr(moons))
	}
}

func TestPart1CaseWithEnergy(t *testing.T) {
	moons := GetTc1()
	rounds := 10
	Simulate(moons, rounds)

	expectedMoons1 := []Moon{
		Moon{Position: Coords{X: 2, Y: 1, Z: -3}, Velocity: Coords{X: -3, Y: -2, Z: 1}},
		Moon{Position: Coords{X: 1, Y: -8, Z: 0}, Velocity: Coords{X: -1, Y: 1, Z: 3}},
		Moon{Position: Coords{X: 3, Y: -6, Z: 1}, Velocity: Coords{X: 3, Y: 2, Z: -3}},
		Moon{Position: Coords{X: 2, Y: 0, Z: 4}, Velocity: Coords{X: 1, Y: -1, Z: -1}}}
	if !reflect.DeepEqual(moons, expectedMoons1) {
		t.Error("Round:", rounds)
		t.Error("Expected:\n" + MoonsToStr(expectedMoons1))
		t.Error("Got:\n" + MoonsToStr(moons))
	}
	energy := CalucalteEnergy(moons)
	exepctedEnergy := 179
	if exepctedEnergy != energy {
		t.Error("Wrong energy. Expected:", exepctedEnergy, " Got:", energy)
	}
}

func TestPart2All(t *testing.T) {
	moons := GetTc1()
	expected := GetTc1()
	round := 0
	for ; reflect.DeepEqual(moons, expected) == false || round == 0; round++ {
		Round(moons)
		// fmt.Println()
		// fmt.Println("round: ", roundX)
		// fmt.Println(MoonsToStr(moons))
	}
	fmt.Println("Needed rounds: ", round)
	if round != 2772 {
		t.Error("Expected:", 2772, ", got:", round)
	}
}

func TestPart2Split(t *testing.T) {

	moons := GetTc1X()
	expected := GetTc1X()
	roundX := 0
	for ; reflect.DeepEqual(moons, expected) == false || roundX == 0; roundX++ {
		Round(moons)
		// fmt.Println()
		// fmt.Println("round: ", roundX)
		// fmt.Println(MoonsToStr(moons))
	}
	fmt.Println("Needed rounds X: ", roundX)

	//Y
	moons = GetTc1Y()
	expected = GetTc1Y()
	roundY := 0
	for ; reflect.DeepEqual(moons, expected) == false || roundY == 0; roundY++ {
		Round(moons)
		// fmt.Println()
		// fmt.Println("roundY: ", roundY)
		// fmt.Println(MoonsToStr(moons))
	}
	fmt.Println("Needed rounds Y: ", roundY)

	//X
	moons = GetTc1Z()
	expected = GetTc1Z()
	roundZ := 0
	for ; reflect.DeepEqual(moons, expected) == false || roundZ == 0; roundZ++ {
		Round(moons)
		// fmt.Println()
		// fmt.Println("roundZ: ", roundZ)
		// fmt.Println(MoonsToStr(moons))
	}
	fmt.Println("Needed rounds Z: ", roundZ)

	part2 := LCM(roundX, roundY, roundZ)
	if part2 != 2772 {
		t.Error("Expected:", 2772, ", got:", part2)
	}
}
