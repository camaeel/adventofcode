package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	check(err)
	data, err1 := ioutil.ReadAll(f)
	if err1 != nil {
		log.Fatal(err)
	}

	width, height := 25, 6
	numberOfLayers := len(data) / (width * height)
	fmt.Println("numberOfLayers:", numberOfLayers)
	// strconv.Atoi(string(byteNumber))

	// parse layers
	layers := make([][]int, 100)
	for i, _ := range layers {
		layers[i] = make([]int, width*height)
	}
	for i, v := range data {
		var err2 error
		layers[i/(width*height)][i%(width*height)], err2 = strconv.Atoi(string(v))
		check(err2)
	}

	zeros, ones, twos := make([]int, 100), make([]int, 100), make([]int, 100)

	minZerosLayer := 0

	for i, layer := range layers {
		for _, digit := range layer {
			switch digit {
			case 0:
				zeros[i]++
			case 1:
				ones[i]++
			case 2:
				twos[i]++
			default:
				panic("unknown digit")
			}
		}
		if zeros[minZerosLayer] > zeros[i] {
			minZerosLayer = i
		}
	}

	fmt.Println("Part 1: Min zeros layer:", minZerosLayer, ", code:", ones[minZerosLayer]*twos[minZerosLayer])

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
