package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println()
}

func StringToArr(str string) []int {
	res := make([]int, 0)
	for _, v := range str {
		v1, _ := strconv.Atoi(string(v))
		res = append(res, v1)
	}

	return res
}

func ArrToString(arr []int) string {
	res := ""
	for _, v := range arr {
		res += strconv.Itoa(v)
	}

	return res
}

func FFTPhase(arr []int, pattern []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		value := 0
		for j, v := range arr {
			mult := pattern[((j+1)/(i+1))%len(pattern)]
			// fmt.Print(v, " * ", mult, " + ")
			value += v * mult
		}
		z := int(math.Abs(float64(value))) % 10
		// fmt.Println("=", z)
		res = append(res, z)
	}
	return res
}

func FFT(arr []int, pattern []int, iterations int) []int {
	for i := 0; i < iterations; i++ {
		arr = FFTPhase(arr, pattern)
	}
	return arr
}

func Input() string {
	return "59713137269801099632654181286233935219811755500455380934770765569131734596763695509279561685788856471420060118738307712184666979727705799202164390635688439701763288535574113283975613430058332890215685102656193056939765590473237031584326028162831872694742473094498692690926378560215065112055277042957192884484736885085776095601258138827407479864966595805684283736114104361200511149403415264005242802552220930514486188661282691447267079869746222193563352374541269431531666903127492467446100184447658357579189070698707540721959527692466414290626633017164810627099243281653139996025661993610763947987942741831185002756364249992028050315704531567916821944"
}
