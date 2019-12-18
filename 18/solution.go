package main

import (
	"strings"
)

func main() {

}

type Obj struct {
	Distance int
	IsKey    bool
	Letter   string
}

type Path struct {
	Objects []Obj
}

func ParseMap(input string) []string {
	return strings.Split(strings.Trim(input, "\n "), "\n")
}

func MapToPaths(mapa []string) []Path {
	res := make([]Path, 0)

	return res
}

func MinSteps(mapa []string) int {

	return 0
}
