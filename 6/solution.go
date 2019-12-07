package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Object string
type Orbit struct {
	Center    Object
	Satellite Object
}

func main() {
	f, err := os.Open("input.txt")
	check(err)

	orbits := make([]Orbit, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		orbitData := strings.Split(line, ")")
		orbits = append(orbits, Orbit{Center: Object(orbitData[0]), Satellite: Object(orbitData[1])})
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Println(orbits)
	com := FindCenterOfMass(orbits)
	if com == "" {
		panic("No COM")
	}
	fmt.Println("COM is:", com)
	result := GetChildrenOrbits(orbits, com, 0)
	fmt.Println("Result: ", result)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func FindCenterOfMass(orbits []Orbit) Object {
	for _, o := range orbits {
		possibleCenter := o.Center
		found := false
		for _, o2 := range orbits {
			if o2.Satellite == possibleCenter {
				found = true
			}
		}
		if found == false {
			return possibleCenter
		}

	}

	return ""
}

func GetChildrenOrbits(orbits []Orbit, parent Object, height int) int {
	children := FindChildren(orbits, parent)
	result := height
	for _, v := range children {
		result += GetChildrenOrbits(orbits, v, height+1)
	}
	return result
}

func FindChildren(orbits []Orbit, parent Object) []Object {
	result := make([]Object, 0)
	for _, v := range orbits {
		if parent == v.Center {
			result = append(result, v.Satellite)
		}
	}

	return result
}
