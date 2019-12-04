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

func TestCalcFuelWithFuelWeight(t *testing.T) {
	if CalcFuelWithFuelWeight(14) != 2 {
		t.Error()
	}

}

func TestCalcFuelWithFuelWeight2(t *testing.T) {
	if CalcFuelWithFuelWeight(1969) != 966 {
		t.Error(966)
	}

}

func TestCalcFuelWithFuelWeight3(t *testing.T) {
	if CalcFuelWithFuelWeight(100756) != 50346 {
		t.Error(50346)
	}

}
