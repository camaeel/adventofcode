package main

import (
	"reflect"
	"testing"
)

// func TestDealIntNewStack1(t *testing.T) {
// 	input := []int{0, 140, 12, 3, 99}
// 	expected := []int{99, 3, 12, 140, 0}

// 	result := DealIntoNewStack(input)
// 	if !reflect.DeepEqual(expected, result) {
// 		t.Error("Test error - wrong output. Input: ", input, "\nExpected: ", expected, "\nGot: ", result)
// 	}

// }

// func TestDealIntNewStack2(t *testing.T) {
// 	input := []int{0, 140, 12, 3}
// 	expected := []int{3, 12, 140, 0}

// 	result := DealIntoNewStack(input)
// 	if !reflect.DeepEqual(expected, result) {
// 		t.Error("Test error - wrong output. Input: ", input, "\nExpected: ", expected, "\nGot: ", result)
// 	}

// }

func TestDealIntNewStack3(t *testing.T) {
	expected := 6

	result := ExecuteOrder(10, 3, "deal into new stack")
	if !reflect.DeepEqual(expected, result) {
		t.Error("Test error - wrong output. Expected: ", expected, "\nGot: ", result)
	}

}

func TestDealIntNewStackReversed(t *testing.T) {
	expected := 3

	result := ExecuteOrderReversed(10, 6, "deal into new stack")
	if !reflect.DeepEqual(expected, result) {
		t.Error("Test error - wrong output. Expected: ", expected, "\nGot: ", result)
	}

}

func TestCut1(t *testing.T) {
	expected := 2

	result := ExecuteOrder(10, 5, "cut 3")
	if !reflect.DeepEqual(expected, result) {
		t.Error("Test error - wrong output. Expected: ", expected, "\nGot: ", result)
	}

}

func TestCut1Reversed(t *testing.T) {
	expected := 5

	result := ExecuteOrderReversed(10, 2, "cut 3")
	if !reflect.DeepEqual(expected, result) {
		t.Error("Test error - wrong output. Expected: ", expected, "\nGot: ", result)
	}

}

func TestCut2(t *testing.T) {
	expected := 9

	result := ExecuteOrder(10, 5, "cut -4")
	if !reflect.DeepEqual(expected, result) {
		t.Error("Test error - wrong output. Expected: ", expected, "\nGot: ", result)
	}

}

func TestCut2Reversed(t *testing.T) {
	expected := 5

	result := ExecuteOrderReversed(10, 9, "cut -4")
	if !reflect.DeepEqual(expected, result) {
		t.Error("Test error - wrong output. Expected: ", expected, "\nGot: ", result)
	}

}

func TestDealWithIncrement1(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := 2

	result := ExecuteOrder(10, 4, "deal with increment 3")
	if !reflect.DeepEqual(expected, result) {
		t.Error("Test error - wrong output. Input: ", input, "\nExpected: ", expected, "\nGot: ", result)
	}

}

func TestDealWithIncrement1Reversed(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := 4

	result := ExecuteOrderReversed(10, 2, "deal with increment 3")
	if !reflect.DeepEqual(expected, result) {
		t.Error("Test error - wrong output. Input: ", input, "\nExpected: ", expected, "\nGot: ", result)
	}

}

func TestTc1V2(t *testing.T) {
	input, size := Tc1()
	// expected := []int{0, 3, 6, 9, 2, 5, 8, 1, 4, 7}

	result := Shuffle(size, 6, input)

	if !reflect.DeepEqual(2, result) {
		t.Error("Test error - wrong output. Input: ", input, "\nExpected: ", 3, "\nGot: ", result)
	}
}
