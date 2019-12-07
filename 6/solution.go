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

	result2 := GetOrbitTransfers(orbits, com, Object("YOU"), Object("SAN"))
	fmt.Println("Result part 2: ", result2)
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

func ChildrenSearch(children []Object, needle Object) bool {
	for _, v := range children {
		if needle == v {
			return true
		}
	}
	return false
}

func GetOrbitTransfers(orbits []Orbit, com Object, needle1 Object, needle2 Object) int {
	path1 := FindObjectPath(orbits, com, needle1)
	path2 := FindObjectPath(orbits, com, needle2)
	fmt.Println("Path1:", path1)
	fmt.Println("Path2:", path2)
	i1 := len(path1) - 1
	i2 := len(path2) - 1
	for i1 > 0 && i2 > 0 && path1[i1-1] == path2[i2-1] {
		i1--
		i2--
	}
	return i1 + i2
}

func FindObjectPath(orbits []Orbit, parent Object, needle Object) []Object {
	children := FindChildren(orbits, parent)
	if ChildrenSearch(children, needle) {
		return []Object{parent}
	}
	for _, v := range children {
		path := FindObjectPath(orbits, v, needle)
		if len(path) > 0 {
			return append(path, parent)
		}
	}
	return make([]Object, 0)
}
