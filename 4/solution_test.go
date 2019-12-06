package main

import (
	"fmt"
	"testing"
)

func TestIsPossiblePassword1(t *testing.T) {
	fmt.Println()
	fmt.Println("start test case")

	input := 566779
	expected := true
	output := IsPossiblePassword(input)
	if output != expected {
		t.Error("Test error. Input: ", input, "\nExpected: ", expected, "\nGot: ", output)
	}
}

func TestIsPossiblePassword2(t *testing.T) {
	fmt.Println()
	fmt.Println("start test case")

	input := 566679
	expected := false
	output := IsPossiblePassword(input)
	if output != expected {
		t.Error("Test error. Input: ", input, "\nExpected: ", expected, "\nGot: ", output)
	}
}

func TestIsPossiblePassword3(t *testing.T) {
	fmt.Println()
	fmt.Println("start test case")
	input := 345678
	expected := false
	output := IsPossiblePassword(input)
	if output != expected {
		t.Error("Test error. Input: ", input, "\nExpected: ", expected, "\nGot: ", output)
	}
}

func TestIsPossiblePassword4(t *testing.T) {
	fmt.Println()
	fmt.Println("start test case")
	input := 688999
	expected := true
	output := IsPossiblePassword(input)
	if output != expected {
		t.Error("Test error. Input: ", input, "\nExpected: ", expected, "\nGot: ", output)
	}
}

func TestIsPossiblePassword5(t *testing.T) {
	fmt.Println()
	fmt.Println("start test case")
	input := 112233
	expected := true
	output := IsPossiblePassword(input)
	if output != expected {
		t.Error("Test error. Input: ", input, "\nExpected: ", expected, "\nGot: ", output)
	}
}

func TestIsPossiblePassword6(t *testing.T) {
	fmt.Println()
	fmt.Println("start test case")
	input := 123444
	expected := false
	output := IsPossiblePassword(input)
	if output != expected {
		t.Error("Test error. Input: ", input, "\nExpected: ", expected, "\nGot: ", output)
	}
}

func TestIsPossiblePassword7(t *testing.T) {
	fmt.Println()
	fmt.Println("start test case")
	input := 111122
	expected := true
	output := IsPossiblePassword(input)
	if output != expected {
		t.Error("Test error. Input: ", input, "\nExpected: ", expected, "\nGot: ", output)
	}
}
