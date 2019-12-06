package intcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCloneProgram(t *testing.T) {
	program := []int{1, 0, 0, 0, 99}
	tmpProg := CloneProgram(program)
	if !reflect.DeepEqual(tmpProg, program) {
		t.Error("Test error. \nExpected: ", tmpProg, "\nGot: ", program)
	}

}

func TestIntcode1(t *testing.T) {
	program := []int{1, 0, 0, 0, 99}
	fmt.Println()
	fmt.Println("Program: ", program)
	tmpProg := CloneProgram(program)
	input := make([]int, 0)
	expected := []int{2, 0, 0, 0, 99}
	Intcode(tmpProg, input)
	if !reflect.DeepEqual(tmpProg, expected) {
		t.Error("Test error. Input: ", input, "\nExpected: ", expected, "\nGot: ", tmpProg)
	}
}

func TestIntcode2(t *testing.T) {
	program := []int{2, 3, 0, 3, 99}
	fmt.Println()
	fmt.Println("Program: ", program)
	tmpProg := CloneProgram(program)
	input := make([]int, 0)
	expected := []int{2, 3, 0, 6, 99}
	Intcode(tmpProg, input)
	if !reflect.DeepEqual(tmpProg, expected) {
		t.Error("Test error. Input: ", input, "\nExpected: ", expected, "\nGot: ", tmpProg)
	}
}

func TestIntcode3(t *testing.T) {
	program := []int{2, 4, 4, 5, 99, 0}
	fmt.Println()
	fmt.Println("Program: ", program)
	tmpProg := CloneProgram(program)
	input := make([]int, 0)
	expected := []int{2, 4, 4, 5, 99, 9801}
	Intcode(tmpProg, input)
	if !reflect.DeepEqual(tmpProg, expected) {
		t.Error("Test error. Input: ", input, "\nExpected: ", expected, "\nGot: ", tmpProg)
	}
}

func TestIntcode4(t *testing.T) {
	program := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	fmt.Println()
	fmt.Println("Program: ", program)
	tmpProg := CloneProgram(program)
	input := make([]int, 0)
	expected := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
	Intcode(tmpProg, input)
	if !reflect.DeepEqual(tmpProg, expected) {
		t.Error("Test error. Input: ", input, "\nExpected: ", expected, "\nGot: ", tmpProg)
	}
}
