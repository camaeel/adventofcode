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

func TestDealIntNewStack3(t *testing.T) {
	input := []int{0, 140, 12, 3, 99}
	expected := []int{99, 3, 12, 140, 0}

	result := ExecuteOrder(input, "deal into new stack")
	if !reflect.DeepEqual(expected, result) {
		t.Error("Test error - wrong output. Input: ", input, "\nExpected: ", expected, "\nGot: ", result)
	}

}

func TestCut1(t *testing.T) {
	input := []int{0, 140, 12, 3, 99}
	expected := []int{12, 3, 99, 0, 140}

	result := ExecuteOrder(input, "cut 2")
	if !reflect.DeepEqual(expected, result) {
		t.Error("Test error - wrong output. Input: ", input, "\nExpected: ", expected, "\nGot: ", result)
	}

}

func TestCut2(t *testing.T) {
	input := []int{0, 140, 12, 3, 99, 65, 12}
	expected := []int{65, 12, 0, 140, 12, 3, 99}

	result := ExecuteOrder(input, "cut -2")
	if !reflect.DeepEqual(expected, result) {
		t.Error("Test error - wrong output. Input: ", input, "\nExpected: ", expected, "\nGot: ", result)
	}

}

func TestDealWithIncrement1(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := []int{0, 7, 4, 1, 8, 5, 2, 9, 6, 3}

	result := ExecuteOrder(input, "deal with increment 3")
	if !reflect.DeepEqual(expected, result) {
		t.Error("Test error - wrong output. Input: ", input, "\nExpected: ", expected, "\nGot: ", result)
	}

}

func TestTc1(t *testing.T) {
	input, size := Tc1()
	expected := []int{0, 3, 6, 9, 2, 5, 8, 1, 4, 7}

	deck := OpenDeck(size)
	result := Shuffle(deck, input)

	if !reflect.DeepEqual(expected, result) {
		t.Error("Test error - wrong output. Input: ", input, "\nExpected: ", expected, "\nGot: ", result)
	}
}
