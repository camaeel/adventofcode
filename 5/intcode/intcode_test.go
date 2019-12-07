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
	Intcode(tmpProg, input, false)
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
	Intcode(tmpProg, input, false)
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
	Intcode(tmpProg, input, false)
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
	Intcode(tmpProg, input, false)
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
	expectedOutput := []int{200}
	output := Intcode(tmpProg, input, false)
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
	output := Intcode(tmpProg, input, false)
	if !reflect.DeepEqual(tmpProg, expected) {
		t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	}
	if !reflect.DeepEqual(expectedOutput, output) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", output)
	}
}

func TestIntcode7(t *testing.T) {
	program := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	fmt.Println()
	fmt.Println("Program: ", program)
	tmpProg := CloneProgram(program)
	input := []int{8}
	// expected := []int{1002, 4, 3, 4, 99}
	expectedOutput := []int{1}
	output := Intcode(tmpProg, input, false)
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, output) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", output)
	}
}

func TestIntcode8(t *testing.T) {
	program := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	fmt.Println()
	fmt.Println("Program: ", program)
	tmpProg := CloneProgram(program)
	input := []int{10000}
	// expected := []int{1002, 4, 3, 4, 99}
	expectedOutput := []int{0}
	output := Intcode(tmpProg, input, false)
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, output) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", output)
	}
}

func TestIntcode9(t *testing.T) {
	program := []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	fmt.Println()
	fmt.Println("Program: ", program)
	tmpProg := CloneProgram(program)
	input := []int{7}
	// expected := []int{1002, 4, 3, 4, 99}
	expectedOutput := []int{1}
	output := Intcode(tmpProg, input, false)
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, output) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", output)
	}
}

func TestIntcode10(t *testing.T) {
	program := []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	fmt.Println()
	fmt.Println("Program: ", program)
	tmpProg := CloneProgram(program)
	input := []int{8}
	// expected := []int{1002, 4, 3, 4, 99}
	expectedOutput := []int{0}
	output := Intcode(tmpProg, input, false)
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, output) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", output)
	}
}

func TestIntcode11(t *testing.T) {
	program := []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}
	fmt.Println()
	fmt.Println("Program: ", program)
	tmpProg := CloneProgram(program)
	input := []int{8}
	// expected := []int{1002, 4, 3, 4, 99}
	expectedOutput := []int{1}
	output := Intcode(tmpProg, input, false)
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, output) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", output)
	}
}

func TestIntcode12(t *testing.T) {
	program := []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}
	fmt.Println()
	fmt.Println("Program: ", program)
	tmpProg := CloneProgram(program)
	input := []int{7}
	// expected := []int{1002, 4, 3, 4, 99}
	expectedOutput := []int{1}
	output := Intcode(tmpProg, input, false)
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, output) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", output)
	}
}

func TestIntcode13(t *testing.T) {
	program := []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}
	fmt.Println()
	fmt.Println("Program: ", program)
	tmpProg := CloneProgram(program)
	input := []int{8}
	// expected := []int{1002, 4, 3, 4, 99}
	expectedOutput := []int{0}
	output := Intcode(tmpProg, input, false)
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, output) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", output)
	}
}

func TestIntcode14(t *testing.T) {
	program := []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}
	fmt.Println()
	fmt.Println("Program: ", program)
	tmpProg := CloneProgram(program)
	input := []int{7}
	// expected := []int{1002, 4, 3, 4, 99}
	expectedOutput := []int{0}
	output := Intcode(tmpProg, input, false)
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, output) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", output)
	}
}

func TestIntcodeJump1(t *testing.T) {
	program := []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	fmt.Println()
	fmt.Println("Program: ", program)
	tmpProg := CloneProgram(program)
	input := []int{8}
	// expected := []int{1002, 4, 3, 4, 99}
	expectedOutput := []int{1}
	output := Intcode(tmpProg, input, false)
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, output) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", output)
	}
}

func TestIntcodeJump2(t *testing.T) {
	program := []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	fmt.Println()
	fmt.Println("Program: ", program)
	tmpProg := CloneProgram(program)
	input := []int{0}
	// expected := []int{1002, 4, 3, 4, 99}
	expectedOutput := []int{0}
	output := Intcode(tmpProg, input, false)
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, output) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", output)
	}
}

func TestIntcodeJump3(t *testing.T) {
	program := []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	fmt.Println()
	fmt.Println("Program: ", program)
	tmpProg := CloneProgram(program)
	input := []int{8}
	// expected := []int{1002, 4, 3, 4, 99}
	expectedOutput := []int{1}
	output := Intcode(tmpProg, input, false)
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, output) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", output)
	}
}

func TestIntcodeJump4(t *testing.T) {
	program := []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	fmt.Println()
	fmt.Println("Program: ", program)
	tmpProg := CloneProgram(program)
	input := []int{0}
	// expected := []int{1002, 4, 3, 4, 99}
	expectedOutput := []int{0}
	output := Intcode(tmpProg, input, false)
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, output) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", output)
	}
}

func TestIntcodeLarge1(t *testing.T) {
	program := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	fmt.Println()
	fmt.Println("Program: ", program)
	tmpProg := CloneProgram(program)
	input := []int{3}
	// expected := []int{1002, 4, 3, 4, 99}
	expectedOutput := []int{999}
	output := Intcode(tmpProg, input, false)
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, output) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", output)
	}
}

func TestIntcodeLarge2(t *testing.T) {
	program := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	fmt.Println()
	fmt.Println("Program: ", program)
	tmpProg := CloneProgram(program)
	input := []int{8}
	// expected := []int{1002, 4, 3, 4, 99}
	expectedOutput := []int{1000}
	output := Intcode(tmpProg, input, false)
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, output) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", output)
	}
}

func TestIntcodeLarge3(t *testing.T) {
	program := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	fmt.Println()
	fmt.Println("Program: ", program)
	tmpProg := CloneProgram(program)
	input := []int{13}
	// expected := []int{1002, 4, 3, 4, 99}
	expectedOutput := []int{1001}
	output := Intcode(tmpProg, input, false)
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, output) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", output)
	}
}
