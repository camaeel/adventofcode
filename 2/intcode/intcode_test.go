package intcode

import (
	"reflect"
	"testing"
)

func TestIntcode1(t *testing.T) {
	input := []int{1, 0, 0, 0, 99}
	expected := []int{2, 0, 0, 0, 99}
	output := Intcode(input)
	if !reflect.DeepEqual(output, expected) {
		t.Error("Test error. Input: ", input, "\nExpected: ", expected, "\nGot: ", output)
	}
}

func TestIntcode2(t *testing.T) {
	input := []int{2, 3, 0, 3, 99}
	expected := []int{2, 3, 0, 6, 99}
	output := Intcode(input)
	if !reflect.DeepEqual(output, expected) {
		t.Error("Test error. Input: ", input, "\nExpected: ", expected, "\nGot: ", output)
	}
}

func TestIntcode3(t *testing.T) {
	input := []int{2, 4, 4, 5, 99, 0}
	expected := []int{2, 4, 4, 5, 99, 9801}
	output := Intcode(input)
	if !reflect.DeepEqual(output, expected) {
		t.Error("Test error. Input: ", input, "\nExpected: ", expected, "\nGot: ", output)
	}
}

func TestIntcode4(t *testing.T) {
	input := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	expected := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
	output := Intcode(input)
	if !reflect.DeepEqual(output, expected) {
		t.Error("Test error. Input: ", input, "\nExpected: ", expected, "\nGot: ", output)
	}
}
