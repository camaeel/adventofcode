package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCase1Phase1(t *testing.T) {
	input := "12345678"
	pattern := []int{0, 1, 0, -1}
	expected := "48226158"
	result := ArrToString(FFT(StringToArr(input), pattern, 1))

	if result != expected {
		t.Error("Error. Expected:", expected, ", Got:", result)
	}
}

func TestPart1(t *testing.T) {
	input := Input()
	pattern := []int{0, 1, 0, -1}
	expected := "10189359"
	result := ArrToString(FFT(StringToArr(input), pattern, 100))
	fmt.Println("Resul:", result)
	if result[0:8] != expected {
		t.Error("Error. Expected:", expected, ", Got:", result)
	}
}

func TestPart2(t *testing.T) {
	input := ""
	for i := 0; i < 10000; i++ {
		input += Input()
	}
	pattern := []int{0, 1, 0, -1}
	expected := "10189359"
	offset, _ := strconv.Atoi(Input()[0:7])
	result := ArrToString(FFT2(StringToArr(input), pattern, 100, offset, 8))

	if result[offset:offset+8] != expected {
		t.Error("Error. Expected:", expected, ", Got:", result)
	}
}

func TestPart2Long(t *testing.T) {
	input := ""
	for i := 0; i < 10000; i++ {
		input += Input()
	}
	pattern := []int{0, 1, 0, -1}
	expected := "10189359"
	offset, _ := strconv.Atoi(Input()[0:7])
	result := ArrToString(FFT(StringToArr(input), pattern, 100))

	if result[offset:offset+8] != expected {
		t.Error("Error. Expected:", expected, ", Got:", result)
	}
}
