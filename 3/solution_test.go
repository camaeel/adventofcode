package main

import "testing"

func TestSample1(t *testing.T) {
	set1 := []string{"R8", "U5", "L5", "D3"}
	set2 := []string{"U7", "R6", "D4", "L4"}
	result := GetNearestIntersect(set1, set2)
	expected := 6
	if result != expected {
		t.Error("wrong result expected=", expected, " got= ", result)
	}
}

func TestSample2(t *testing.T) {
	set1 := []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}
	set2 := []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}
	result := GetNearestIntersect(set1, set2)
	expected := 159
	if result != expected {
		t.Error("wrong result expected=", expected, " got= ", result)
	}
}

func TestSample3(t *testing.T) {
	set1 := []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}
	set2 := []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}
	result := GetNearestIntersect(set1, set2)
	expected := 135
	if result != expected {
		t.Error("wrong result expected=", expected, " got= ", result)
	}
}

func TestDetectIntersect(t *testing.T) {
	s1 := Section{A: Point{X: 0, Y: 0}, B: Point{X: 10, Y: 0}}
	s2 := Section{A: Point{X: 5, Y: 0}, B: Point{X: 5, Y: 10}}
	p, found := DetectIntersect(s1, s2)
	if found != true {
		t.Error("Intersection not found")
	}
	if p.X != 5 && p.Y != 0 {
		t.Error("Found wrong intersection: ", p)
	}

}
