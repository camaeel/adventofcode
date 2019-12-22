package main

import (
	"reflect"
	"testing"
)

func TestDealIntNewStack1(t *testing.T) {
	input := []int{0, 140, 12, 3, 99}
	expected := []int{99, 3, 12, 140, 0}

	result := DealIntoNewStack(input)
	if !reflect.DeepEqual(expected, result) {
		t.Error("Test error - wrong output. Input: ", input, "\nExpected: ", expected, "\nGot: ", result)
	}

}

func TestDealIntNewStack2(t *testing.T) {
	input := []int{0, 140, 12, 3}
	expected := []int{3, 12, 140, 0}

	result := DealIntoNewStack(input)
	if !reflect.DeepEqual(expected, result) {
		t.Error("Test error - wrong output. Input: ", input, "\nExpected: ", expected, "\nGot: ", result)
	}

}
