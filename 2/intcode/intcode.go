package intcode

func Intcode(program []int) []int {

	for counter := 0; program[counter] != 99; {
		counter++
	}
	return program

}
