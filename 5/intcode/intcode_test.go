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
		t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
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
		t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
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
		t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
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
		t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	}
}

func TestIntcode5(t *testing.T) {
	program := []int{3, 0, 4, 0, 99}
	fmt.Println()
	fmt.Println("Program: ", program)
	tmpProg := CloneProgram(program)
	input := []int{200}
	expected := []int{200, 0, 4, 0, 99}
	expectedOutput := make([]int, 0)
	output := Intcode(tmpProg, input)
	if !reflect.DeepEqual(tmpProg, expected) {
		t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	}
	if !reflect.DeepEqual(expectedOutput, output) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", output)
	}
}

func TestDecodeOpcode(t *testing.T) {
	opcode := 1002
	result1, result2 := DecodeOpcode(opcode)
	expectedResult1 := 2
	expectedResult2 := []int{0, 1, 0}
	if !reflect.DeepEqual(result1, expectedResult1) {
		t.Error("Test error. Input: ", opcode, "\nExpected: ", expectedResult1, "\nGot: ", result1)
	}
	if !reflect.DeepEqual(result2, expectedResult2) {
		t.Error("Test error - wrong output. Input: ", opcode, "\nExpected: ", expectedResult2, "\nGot: ", result2)
	}
}

func TestIntcode6(t *testing.T) {
	program := []int{1002, 4, 3, 4, 33}
	fmt.Println()
	fmt.Println("Program: ", program)
	tmpProg := CloneProgram(program)
	input := []int{}
	expected := []int{1002, 4, 3, 4, 99}
	expectedOutput := make([]int, 0)
	output := Intcode(tmpProg, input)
	if !reflect.DeepEqual(tmpProg, expected) {
		t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	}
	if !reflect.DeepEqual(expectedOutput, output) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", output)
	}
}
