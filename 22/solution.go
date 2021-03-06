package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, size := PuzzleInput()
	searchFor := 2019
	result := Shuffle(size, searchFor, input)
	fmt.Println(result)

	deck := 119315717514047
	iter := 101741582076661
	resultPos := 2020

	cmds := strings.Split(input, "\n")

	result2 := resultPos
	for i := 0; i < iter; i++ {
		if i%1000000000 == 0 {
			fmt.Println("iter: ", i)
		}
		for cmdI := len(cmds) - 1; cmdI >= 0; cmdI-- {
			result2 = ExecuteOrder(deck, result2, cmds[cmdI])
		}

	}

	fmt.Println("ph2:", result2)
}

func DealIntoNewStack(length int, tracked int) int {
	return length - tracked - 1
}

func Cut(length int, tracked int, param int) int {
	return ((tracked - param + length) % length)
}

func CutReversed(length int, tracked int, param int) int {
	return ((tracked + param + length) % length)
}

func DealWithIncrement(length int, tracked int, param int) int {
	return (tracked * param) % length

}

// 0123456789 /3
// 0741852963

// 10 mod 3 = 1

// 0123456789X /4
// 0369147X258

// len%p > 11 mod 4 = 3
// 10 -> 8
// 1 -> 3
// 2 -> 6
// 3 -> 9
// 4 -> 1

// y= (x * param) % length
// y + a * len = x * param

func DealWithIncrementReversed(length int, tracked int, param int) int {
	for i := 0; i <= param; i++ {
		if ((tracked + i*length) % param) == 0 {
			return (tracked + i*length) / param
		}
	}

	panic("not found")

}

func OpenDeck(size int) []int {
	res := make([]int, size)
	for i := 0; i < size; i++ {
		res[i] = i
	}

	return res
}

func Shuffle(length int, tracked int, input string) int {
	cmds := strings.Split(input, "\n")

	for _, v := range cmds {
		tracked = ExecuteOrder(length, tracked, v)
	}
	return tracked
}

func ExecuteOrder(length int, tracked int, cmd string) int {
	if cmd == "deal into new stack" {
		return DealIntoNewStack(length, tracked)
	} else {
		matched, err := regexp.MatchString(`^cut -?\d+$`, cmd)
		if matched == true && err == nil {

			re := regexp.MustCompile(` (-?\d+)$`)
			paramStr := re.FindString(cmd)
			param, _ := strconv.Atoi(strings.Trim(paramStr, " "))
			return Cut(length, tracked, param)
		} else {
			matched, err := regexp.MatchString(`^deal with increment -?\d+$`, cmd)
			if matched == true && err == nil {
				re := regexp.MustCompile(` (-?\d+)$`)
				paramStr := re.FindString(cmd)
				param, _ := strconv.Atoi(strings.Trim(paramStr, " "))
				return DealWithIncrement(length, tracked, param)
			} else {
				panic("wrong input")
			}
		}

	}

}

func ExecuteOrderReversed(length int, tracked int, cmd string) int {
	if cmd == "deal into new stack" {
		return DealIntoNewStack(length, tracked)
	} else {
		matched, err := regexp.MatchString(`^cut -?\d+$`, cmd)
		if matched == true && err == nil {

			re := regexp.MustCompile(` (-?\d+)$`)
			paramStr := re.FindString(cmd)
			param, _ := strconv.Atoi(strings.Trim(paramStr, " "))
			return CutReversed(length, tracked, param)
		} else {
			matched, err := regexp.MatchString(`^deal with increment -?\d+$`, cmd)
			if matched == true && err == nil {
				re := regexp.MustCompile(` (-?\d+)$`)
				paramStr := re.FindString(cmd)
				param, _ := strconv.Atoi(strings.Trim(paramStr, " "))
				return DealWithIncrementReversed(length, tracked, param)
			} else {
				panic("wrong input")
			}
		}

	}

}

func PuzzleInput() (string, int) {
	return `cut -135
deal with increment 38
deal into new stack
deal with increment 29
cut 120
deal with increment 30
deal into new stack
cut -7198
deal into new stack
deal with increment 59
cut -8217
deal with increment 75
cut 4868
deal with increment 29
cut 4871
deal with increment 2
deal into new stack
deal with increment 54
cut 777
deal with increment 40
cut -8611
deal with increment 3
cut -5726
deal with increment 57
deal into new stack
deal with increment 41
deal into new stack
cut -5027
deal with increment 12
cut -5883
deal with increment 45
cut 9989
deal with increment 14
cut 6535
deal with increment 18
cut -5544
deal with increment 29
deal into new stack
deal with increment 64
deal into new stack
deal with increment 41
deal into new stack
deal with increment 6
cut 4752
deal with increment 8
deal into new stack
deal with increment 26
cut -6635
deal with increment 10
deal into new stack
cut -3830
deal with increment 48
deal into new stack
deal with increment 39
cut -4768
deal with increment 65
deal into new stack
cut -5417
deal with increment 15
cut -4647
deal into new stack
cut -3596
deal with increment 17
cut -3771
deal with increment 50
cut 1682
deal into new stack
deal with increment 20
deal into new stack
deal with increment 22
deal into new stack
deal with increment 3
cut 8780
deal with increment 52
cut 7478
deal with increment 9
cut -8313
deal into new stack
cut 742
deal with increment 19
cut 9982
deal into new stack
deal with increment 68
cut 9997
deal with increment 23
cut -240
deal with increment 54
cut -7643
deal into new stack
deal with increment 6
cut -3493
deal with increment 74
deal into new stack
deal with increment 75
deal into new stack
deal with increment 40
cut 596
deal with increment 6
cut -4957
deal into new stack`, 10007
}

func Tc1() (string, int) {
	return `deal with increment 7
deal into new stack
deal into new stack`, 10
}

func Tc2() (string, int) {
	return `cut 6
deal with increment 7
deal into new stack`, 10
}

func Tc3() (string, int) {
	return `deal with increment 7
deal with increment 9
cut -2`, 10
}

func Tc4() (string, int) {
	return `deal into new stack
cut -2
deal with increment 7
cut 8
cut -4
deal with increment 7
cut 3
deal with increment 9
deal with increment 3
cut -1`, 10
}
