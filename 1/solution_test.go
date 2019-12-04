package main

import "testing"

func Test12(t *testing.T) {
	if CalcFuel(12) != 2 {
		t.Error()
	}
}

func Test14(t *testing.T) {
	if CalcFuel(14) != 2 {
		t.Error()
	}
}
func Test3(t *testing.T) {
	if CalcFuel(1969) != 654 {
		t.Error()
	}
}
func Test4(t *testing.T) {
	if CalcFuel(100756) != 33583 {
		t.Error()
	}
}
