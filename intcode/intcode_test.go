package intcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCloneProgram(t *testing.T) {
	memorySize := 10
	program := []int{1, 0, 0, 0, 99, 0, 0, 0, 0, 0}
	tmpProg := CloneProgram(program, memorySize)
	if !reflect.DeepEqual(cap(tmpProg), memorySize) {
		t.Error("Test error. \nExpected: ", cap(tmpProg), "\nGot: ", memorySize)
	}
	if !reflect.DeepEqual(tmpProg, program) {
		t.Error("Test error. \nExpected: ", tmpProg, "\nGot: ", program)
	}
	if !reflect.DeepEqual(len(tmpProg), memorySize) {
		t.Error("Test error. \nExpected: ", len(tmpProg), "\nGot: ", memorySize)
	}
}

func TestIntcode1(t *testing.T) {
	program := []int{1, 0, 0, 0, 99}
	fmt.Println()
	fmt.Println("Program: ", program)
	computer := CreateIntcode(program)
	computer.Run(false)
	computer.WaitGroup.Wait()
	expected := []int{2, 0, 0, 0, 99}
	if !reflect.DeepEqual(computer.Program[0:len(expected)], expected) {
		t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", computer.Program)
	}
}

func TestIntcode2(t *testing.T) {
	program := []int{2, 3, 0, 3, 99}
	fmt.Println()
	fmt.Println("Program: ", program)
	computer := CreateIntcode(program)
	computer.Run(false)
	computer.WaitGroup.Wait()
	expected := []int{2, 3, 0, 6, 99}
	if !reflect.DeepEqual(computer.Program[0:len(expected)], expected) {
		t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", computer.Program)
	}
}

func TestIntcode3(t *testing.T) {
	program := []int{2, 4, 4, 5, 99, 0}
	fmt.Println()
	fmt.Println("Program: ", program)
	computer := CreateIntcode(program)
	computer.Run(false)
	computer.WaitGroup.Wait()
	expected := []int{2, 4, 4, 5, 99, 9801}
	if !reflect.DeepEqual(computer.Program[0:len(expected)], expected) {
		t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", computer.Program)
	}
}

func TestIntcode4(t *testing.T) {
	program := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	fmt.Println()
	fmt.Println("Program: ", program)
	computer := CreateIntcode(program)
	computer.Run(false)
	computer.WaitGroup.Wait()
	expected := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
	if !reflect.DeepEqual(computer.Program[0:len(expected)], expected) {
		t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", computer.Program)
	}
}

func TestIntcode5(t *testing.T) {
	program := []int{3, 0, 4, 0, 99}
	fmt.Println()
	fmt.Println("Program: ", program)
	computer := CreateIntcode(program)
	fmt.Println("Input written: ", program)
	computer.Run(false)
	computer.Input <- 200
	output := <-computer.Output
	computer.WaitGroup.Wait()

	expected := []int{200, 0, 4, 0, 99}
	expectedOutput := []int{200}
	if !reflect.DeepEqual(computer.Program[0:len(expected)], expected) {
		t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", computer.Program)
	}

	if !reflect.DeepEqual(expectedOutput, computer.OutputArray) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", computer.OutputArray)
	}
	if output != expectedOutput[0] {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", output, "\nGot: ", expectedOutput[0])
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
	computer := CreateIntcode(program)
	computer.Run(false)
	computer.WaitGroup.Wait()

	expected := []int{1002, 4, 3, 4, 99}
	expectedOutput := make([]int, 0)

	if !reflect.DeepEqual(computer.Program[0:len(expected)], expected) {
		t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", computer.Program)
	}
	if !reflect.DeepEqual(expectedOutput, computer.OutputArray) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", computer.OutputArray)
	}
}

func TestIntcode7(t *testing.T) {
	program := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	fmt.Println()
	fmt.Println("Program: ", program)
	computer := CreateIntcode(program)

	computer.Run(false)
	computer.Input <- 8
	_ = <-computer.Output

	computer.WaitGroup.Wait()
	// expected := []int{1002, 4, 3, 4, 99}
	expectedOutput := []int{1}
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, computer.OutputArray) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", computer.OutputArray)
	}
}

func TestIntcode8(t *testing.T) {
	program := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	fmt.Println()
	fmt.Println("Program: ", program)
	computer := CreateIntcode(program)

	computer.Run(false)
	computer.Input <- 10000
	_ = <-computer.Output

	computer.WaitGroup.Wait()
	expectedOutput := []int{0}

	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, computer.OutputArray) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", computer.OutputArray)
	}
}

func TestIntcode9(t *testing.T) {
	program := []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	fmt.Println()
	fmt.Println("Program: ", program)
	computer := CreateIntcode(program)

	computer.Run(false)
	computer.Input <- 7
	_ = <-computer.Output

	computer.WaitGroup.Wait()
	expectedOutput := []int{1}
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, computer.OutputArray) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", computer.OutputArray)
	}
}

func TestIntcode10(t *testing.T) {
	program := []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	fmt.Println()
	fmt.Println("Program: ", program)
	computer := CreateIntcode(program)

	computer.Run(false)
	computer.Input <- 8
	_ = <-computer.Output

	computer.WaitGroup.Wait()
	expectedOutput := []int{0}
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, computer.OutputArray) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", computer.OutputArray)
	}
}

func TestIntcode11(t *testing.T) {
	program := []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}
	fmt.Println()
	fmt.Println("Program: ", program)
	computer := CreateIntcode(program)

	computer.Run(false)
	computer.Input <- 8
	_ = <-computer.Output

	computer.WaitGroup.Wait()
	expectedOutput := []int{1}
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, computer.OutputArray) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", computer.OutputArray)
	}
}

func TestIntcode12(t *testing.T) {
	program := []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}
	fmt.Println()
	fmt.Println("Program: ", program)
	computer := CreateIntcode(program)

	computer.Run(false)
	computer.Input <- 7
	_ = <-computer.Output

	computer.WaitGroup.Wait()
	expectedOutput := []int{1}
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, computer.OutputArray) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", computer.OutputArray)
	}
}

func TestIntcode13(t *testing.T) {
	program := []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}
	fmt.Println()
	fmt.Println("Program: ", program)
	computer := CreateIntcode(program)

	computer.Run(false)
	computer.Input <- 8
	_ = <-computer.Output

	computer.WaitGroup.Wait()
	expectedOutput := []int{0}
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, computer.OutputArray) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", computer.OutputArray)
	}
}

func TestIntcode14(t *testing.T) {
	program := []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}
	fmt.Println()
	fmt.Println("Program: ", program)
	computer := CreateIntcode(program)

	computer.Run(false)
	computer.Input <- 7
	_ = <-computer.Output

	computer.WaitGroup.Wait()
	expectedOutput := []int{0}
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, computer.OutputArray) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", computer.OutputArray)
	}
}

func TestIntcodeJump1(t *testing.T) {
	program := []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	fmt.Println()
	fmt.Println("Program: ", program)
	computer := CreateIntcode(program)

	computer.Run(false)
	computer.Input <- 8
	_ = <-computer.Output

	computer.WaitGroup.Wait()
	expectedOutput := []int{1}
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, computer.OutputArray) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", computer.OutputArray)
	}
}

func TestIntcodeJump2(t *testing.T) {
	program := []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	fmt.Println()
	fmt.Println("Program: ", program)
	computer := CreateIntcode(program)

	computer.Run(false)
	computer.Input <- 0
	_ = <-computer.Output

	computer.WaitGroup.Wait()
	expectedOutput := []int{0}
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, computer.OutputArray) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", computer.OutputArray)
	}
}

func TestIntcodeJump3(t *testing.T) {
	program := []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	fmt.Println()
	fmt.Println("Program: ", program)
	computer := CreateIntcode(program)

	computer.Run(false)
	computer.Input <- 8
	_ = <-computer.Output

	computer.WaitGroup.Wait()
	expectedOutput := []int{1}
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, computer.OutputArray) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", computer.OutputArray)
	}
}

func TestIntcodeJump4(t *testing.T) {
	program := []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	fmt.Println()
	fmt.Println("Program: ", program)
	computer := CreateIntcode(program)

	computer.Run(false)
	computer.Input <- 0
	_ = <-computer.Output

	computer.WaitGroup.Wait()
	expectedOutput := []int{0}
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, computer.OutputArray) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", computer.OutputArray)
	}
}

func TestIntcodeLarge1(t *testing.T) {
	program := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	fmt.Println()
	fmt.Println("Program: ", program)
	computer := CreateIntcode(program)

	computer.Run(false)
	computer.Input <- 3
	_ = <-computer.Output

	computer.WaitGroup.Wait()
	expectedOutput := []int{999}
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, computer.OutputArray) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", computer.OutputArray)
	}
}

func TestIntcodeLarge2(t *testing.T) {
	program := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	fmt.Println()
	fmt.Println("Program: ", program)
	computer := CreateIntcode(program)

	computer.Run(false)
	computer.Input <- 8
	_ = <-computer.Output

	computer.WaitGroup.Wait()
	expectedOutput := []int{1000}
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, computer.OutputArray) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", computer.OutputArray)
	}
}

func TestIntcodeLarge3(t *testing.T) {
	program := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	fmt.Println()
	fmt.Println("Program: ", program)
	computer := CreateIntcode(program)

	computer.Run(false)
	computer.Input <- 13
	_ = <-computer.Output

	computer.WaitGroup.Wait()
	expectedOutput := []int{1001}
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, computer.OutputArray) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", computer.OutputArray)
	}
}

func TestIntcodeDay9RelativeAddr1(t *testing.T) {
	program := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	fmt.Println()
	fmt.Println("Program: ", program)
	computer := CreateIntcode(program)

	computer.Run(true)
	// computer.Input <- 13
	for i := 0; i < 16; i++ {
		_ = <-computer.Output

	}

	computer.WaitGroup.Wait()
	expectedOutput := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	// if !reflect.DeepEqual(tmpProg, expected) {
	// 	t.Error("Test error. Input: ", program, "\nExpected: ", expected, "\nGot: ", tmpProg)
	// }
	if !reflect.DeepEqual(expectedOutput, computer.OutputArray) {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expectedOutput, "\nGot: ", computer.OutputArray)
	}
}

func TestIntcodeDay9LargeNumber(t *testing.T) {
	program := []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}
	fmt.Println()
	fmt.Println("Program: ", program)
	computer := CreateIntcode(program)

	computer.Run(false)
	// computer.Input <- 13
	_ = <-computer.Output

	computer.WaitGroup.Wait()

	if computer.OutputArray[0] < 1000000000000000 {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: > 1000000000000000, \nGot: ", computer.OutputArray)
	}
}

func TestIntcodeDay9LargeNumber2(t *testing.T) {
	program := []int{104, 1125899906842624, 99}
	fmt.Println()
	fmt.Println("Program: ", program)
	computer := CreateIntcode(program)
	computer.Run(false)
	// computer.Input <- 13
	_ = <-computer.Output

	computer.WaitGroup.Wait()

	if computer.OutputArray[0] != program[1] {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", program[1], ", \nGot: ", computer.OutputArray)
	}
}

func TestAddresDecoderMode0(t *testing.T) {
	program := []int{10, 20, 30, 1, 3, 99}
	comp := CreateIntcode(program)
	comp.baseOffset = 11

	mode := 0
	counter := 3
	expected := 1

	result := comp.AddressDecoder(mode, counter)
	if result != expected {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expected, "\nGot: ", result)
	}
}

func TestAddresDecoderMode1(t *testing.T) {
	program := []int{10, 20, 30, 1, 3, 99}
	comp := CreateIntcode(program)
	comp.baseOffset = 11

	mode := 1
	counter := 3
	expected := 3

	result := comp.AddressDecoder(mode, counter)
	if result != expected {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expected, "\nGot: ", result)
	}
}

func TestAddresDecoderMode2(t *testing.T) {
	program := []int{10, 20, 30, 1, 3, 99}
	comp := CreateIntcode(program)
	comp.baseOffset = 11

	mode := 2
	counter := 3
	expected := 12

	result := comp.AddressDecoder(mode, counter)
	if result != expected {
		t.Error("Test error - wrong output. Input: ", program, "\nExpected: ", expected, "\nGot: ", result)
	}
}
