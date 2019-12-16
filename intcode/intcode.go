package intcode

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Intcode struct {
	Program     []int
	Stopped     bool
	Input       chan int
	Output      chan int
	InputNeeded chan int
	WaitGroup   *sync.WaitGroup
	OutputArray []int //duplicate output
	baseOffset  int
	InputSync   bool
}

func CreateIntcode(program []int, memoryLimit int, inputSync bool) *Intcode {
	var inputChan chan int
	if inputSync {
		inputChan = make(chan int)
	} else {
		inputChan = make(chan int, memoryLimit)
	}

	wg := sync.WaitGroup{}
	ret := Intcode{Program: CloneProgram(program, memoryLimit), Input: inputChan, InputNeeded: make(chan int), Output: make(chan int, memoryLimit), Stopped: true, OutputArray: make([]int, 0), WaitGroup: &wg, InputSync: inputSync}

	return &ret
}

func (comp *Intcode) Run(debug bool) {
	comp.WaitGroup.Add(1)
	go comp.runAsync(debug)
}

func (comp *Intcode) runAsync(debug bool) {
	counter := 0
	for {
		if comp.Opcode(&counter, debug) == false {
			// close(comp.Output)
			comp.WaitGroup.Done()
			return
		}
	}

}

// returns false if program finished
func (comp *Intcode) Opcode(counter *int, debug bool) bool {
	opcode, modes := DecodeOpcode(comp.Program[*counter])
	opCounter := *counter
	*counter++
	switch opcode {
	case 99:
		if debug {
			fmt.Println("Counter: ", opCounter, "Next opcode: ", opcode)
		}
		close(comp.Output)
		return false
	case 1:
		op1 := comp.GetOperand(counter, modes[0])
		*counter++
		op2 := comp.GetOperand(counter, modes[1])
		*counter++
		comp.SetOperand(counter, modes[2], op1+op2)
		if debug {
			fmt.Println("Counter:", opCounter, "Opcode=", opcode, ", modes:", modes, ", program line: ", comp.Program[opCounter:opCounter+4], ", op1=", op1, " op2=", op2, ", result: ", op1+op2)
		}
		*counter++
	case 2:
		op1 := comp.GetOperand(counter, modes[0])
		*counter++
		op2 := comp.GetOperand(counter, modes[1])
		*counter++
		comp.SetOperand(counter, modes[2], op1*op2)
		if debug {
			fmt.Println("Counter:", opCounter, "Opcode=", opcode, ", modes:", modes, ", program line: ", comp.Program[opCounter:opCounter+4], ", op1=", op1, " op2=", op2, ", result: ", op1*op2)
		}
		*counter++
	case 3:
		var input int
		if debug {
			fmt.Println("Waiting for input")
		}
		if comp.InputSync {
			comp.InputNeeded <- (1)
		}
		input = <-comp.Input
		if debug {
			fmt.Println("Got input")
		}
		comp.SetOperand(counter, modes[0], input)
		if debug {
			fmt.Println("Counter:", opCounter, "Opcode=", opcode, ", modes:", modes, ", program line: ", comp.Program[opCounter:opCounter+2], ", input=", input)
		}

		*counter++
	case 4:
		op1 := comp.GetOperand(counter, modes[0])
		if debug {
			fmt.Println("Counter:", opCounter, "Opcode=", opcode, ", modes:", modes, ", program line: ", comp.Program[opCounter:opCounter+2], ", output=", op1)
		}
		*counter++
		comp.OutputArray = append(comp.OutputArray, op1)
		if debug {
			fmt.Println("Writing output")
		}
		comp.Output <- op1
		if debug {
			fmt.Println("Output written")
		}

	case 5:
		op1 := comp.GetOperand(counter, modes[0])
		*counter++
		op2 := comp.GetOperand(counter, modes[1])
		*counter++
		if debug {
			fmt.Println("Counter:", opCounter, "Opcode=", opcode, ", modes:", modes, ", program line: ", comp.Program[opCounter:opCounter+3], ", op1=", op1, " op2=", op2)
		}
		if op1 != 0 {
			*counter = op2
		}
	case 6:
		op1 := comp.GetOperand(counter, modes[0])
		*counter++
		op2 := comp.GetOperand(counter, modes[1])
		*counter++
		if debug {
			fmt.Println("Counter:", opCounter, "Opcode=", opcode, ", modes:", modes, ", program line: ", comp.Program[opCounter:opCounter+3], ", op1=", op1, " op2=", op2)
		}
		if op1 == 0 {
			*counter = op2
		}
	case 7:
		// is less than: if the first parameter is less than the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
		op1 := comp.GetOperand(counter, modes[0])
		*counter++
		op2 := comp.GetOperand(counter, modes[1])
		*counter++
		val := 0
		if op1 < op2 {
			val = 1
		}
		if debug {
			fmt.Println("Counter:", opCounter, "Opcode=", opcode, ", modes:", modes, ", program line: ", comp.Program[opCounter:opCounter+4], ", op1=", op1, " op2=", op2, " result=", val)
		}
		comp.SetOperand(counter, modes[2], val)
		*counter++
	case 8:
		// if the first parameter is equal to the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
		op1 := comp.GetOperand(counter, modes[0])
		*counter++
		op2 := comp.GetOperand(counter, modes[1])
		*counter++
		val := 0
		if op1 == op2 {
			val = 1
		}
		if debug {
			fmt.Println("Counter:", opCounter, "Opcode=", opcode, ", modes:", modes, ", program line: ", comp.Program[opCounter:opCounter+4], ", op1=", op1, " op2=", op2, " result=", val)
		}
		comp.SetOperand(counter, modes[2], val)
		*counter++
	case 9:
		op1 := comp.GetOperand(counter, modes[0])
		*counter++
		comp.baseOffset += op1
		if debug {
			fmt.Println("Counter:", opCounter, "Opcode=", opcode, ", modes:", modes, ", program line: ", comp.Program[opCounter:opCounter+2], ", op1=", op1, " Moving baseOffset to ", comp.baseOffset)
		}
	default:
		panic("Unknown opcode: " + strconv.Itoa(opcode) + ", at counter: " + strconv.Itoa(*counter-1))
	}

	return true
}

func DecodeOpcode(oc int) (int, []int) {
	opcode := oc % 100
	mode := []int{(oc / 100) % 10, (oc / 1000) % 10, (oc / 10000) % 10}
	return opcode, mode
}

func (comp *Intcode) SetOperand(counter *int, mode int, value int) {
	comp.Program[comp.AddressDecoder(mode, *counter)] = value
}

func (comp *Intcode) GetOperand(counter *int, mode int) int {
	return comp.Program[comp.AddressDecoder(mode, *counter)]
}

func ReadProgram(fileName string) []int {
	program := make([]int, 0)
	f, err := os.Open(fileName)
	defer f.Close()
	check(err)

	scanner := bufio.NewScanner(f)
	line := ""
	for scanner.Scan() {
		line += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	arr := strings.Split(line, ",")
	for _, v := range arr {
		num, err := strconv.Atoi(v)
		check(err)
		program = append(program, num)
	}

	return program
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CloneProgram(program []int, memorySize int) []int {
	tmpProg := make([]int, memorySize)
	copy(tmpProg, program)
	return tmpProg
}

//Reutrns address of target cell
func (comp *Intcode) AddressDecoder(mode int, counter int) int {
	if mode == 0 {
		return comp.Program[counter]
	} else if mode == 1 {
		return counter
	} else if mode == 2 {
		return comp.Program[counter] + comp.baseOffset
	} else {
		panic("Unknown address mode")
	}
}
