package main

import (
	"fmt"
	"reflect"
)

type Coords struct {
	X int
	Y int
	Z int
}

type Moon struct {
	Position Coords
	Velocity Coords
}

func main() {
	moons := GetPart1()
	Simulate(moons, 1000)
	energy := CalucalteEnergy(moons)
	fmt.Println("energy:", energy)

	//part2
	//X
	moons = GetPart1X()
	expected := GetPart1X()
	roundX := 0
	for ; reflect.DeepEqual(moons, expected) == false || roundX == 0; roundX++ {
		Round(moons)
		// fmt.Println()
		// fmt.Println("round: ", roundX)
		// fmt.Println(MoonsToStr(moons))
	}
	fmt.Println("Needed rounds X: ", roundX)

	//Y
	moons = GetPart1Y()
	expected = GetPart1Y()
	roundY := 0
	for ; reflect.DeepEqual(moons, expected) == false || roundY == 0; roundY++ {
		Round(moons)
		// fmt.Println()
		// fmt.Println("roundY: ", roundY)
		// fmt.Println(MoonsToStr(moons))
	}
	fmt.Println("Needed rounds Y: ", roundY)

	//X
	moons = GetPart1Z()
	expected = GetPart1Z()
	roundZ := 0
	for ; reflect.DeepEqual(moons, expected) == false || roundZ == 0; roundZ++ {
		Round(moons)
		// fmt.Println()
		// fmt.Println("roundZ: ", roundZ)
		// fmt.Println(MoonsToStr(moons))
	}
	fmt.Println("Needed rounds Z: ", roundZ)

	part2 := LCM(roundX, roundY, roundZ)
	fmt.Println("Result part 2: ", part2)
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func NWW(a int, b int) int {
	return a * b / NWD(a, b)
}

func NWD(a int, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func Round(moons []Moon) {
	ApplyGravity(moons)
	ApplyVelocity(moons)

}

func Simulate(moons []Moon, maxRounds int) {
	for round := 1; round <= maxRounds; round++ {
		Round(moons)
		fmt.Println()
		fmt.Println("round: ", round)
		fmt.Println(MoonsToStr(moons))
	}
}

func ApplyGravity(moons []Moon) {
	for i := 0; i < len(moons)-1; i++ {
		for j := i + 1; j < len(moons); j++ {
			ApplyGravityForPair(&moons[i], &moons[j])
		}
	}
}

func ApplyGravityForPair(a *Moon, b *Moon) {
	if a.Position.X < b.Position.X {
		a.Velocity.X++
		b.Velocity.X--
	} else if a.Position.X > b.Position.X {
		a.Velocity.X--
		b.Velocity.X++
	}

	if a.Position.Y < b.Position.Y {
		a.Velocity.Y++
		b.Velocity.Y--
	} else if a.Position.Y > b.Position.Y {
		a.Velocity.Y--
		b.Velocity.Y++
	}

	if a.Position.Z < b.Position.Z {
		a.Velocity.Z++
		b.Velocity.Z--
	} else if a.Position.Z > b.Position.Z {
		a.Velocity.Z--
		b.Velocity.Z++
	}
}

func ApplyVelocity(moons []Moon) {
	for i, v := range moons {
		moons[i].Position = AddCoords(v.Position, v.Velocity)
	}
}

func AddCoords(a Coords, b Coords) Coords {
	ret := a
	ret.X += b.X
	ret.Y += b.Y
	ret.Z += b.Z
	return ret
}

func MoonsToStr(moons []Moon) string {
	res := ""
	for _, v := range moons {
		res += "pos=" + CoordsToStr(v.Position)
		res += " vel=" + CoordsToStr(v.Velocity) + "\n"
	}
	return res
}

func CoordsToStr(c Coords) string {

	return fmt.Sprintf("<x=%3d, y=%3d, z=%3d>", c.X, c.Y, c.Z)
}

func GetPart1() []Moon {
	return []Moon{
		Moon{Position: Coords{X: 13, Y: -13, Z: -2}},
		Moon{Position: Coords{X: 16, Y: 2, Z: -15}},
		Moon{Position: Coords{X: 7, Y: -18, Z: -12}},
		Moon{Position: Coords{X: -3, Y: -8, Z: -8}}}
}

func GetPart1X() []Moon {
	return []Moon{
		Moon{Position: Coords{X: 13, Y: 0, Z: 0}},
		Moon{Position: Coords{X: 16, Y: 0, Z: 0}},
		Moon{Position: Coords{X: 7, Y: 0, Z: 0}},
		Moon{Position: Coords{X: -3, Y: 0, Z: 0}}}
}

func GetPart1Y() []Moon {
	return []Moon{
		Moon{Position: Coords{X: 0, Y: -13, Z: 0}},
		Moon{Position: Coords{X: 0, Y: 2, Z: 0}},
		Moon{Position: Coords{X: 0, Y: -18, Z: 0}},
		Moon{Position: Coords{X: 0, Y: -8, Z: 0}}}
}

func GetPart1Z() []Moon {
	return []Moon{
		Moon{Position: Coords{X: 0, Y: 0, Z: -2}},
		Moon{Position: Coords{X: 0, Y: 0, Z: -15}},
		Moon{Position: Coords{X: 0, Y: 0, Z: -12}},
		Moon{Position: Coords{X: 0, Y: 0, Z: -8}}}
}

func CalucalteEnergy(moons []Moon) int {
	ret := 0
	for _, v := range moons {
		pot := Abs(v.Position.X) + Abs(v.Position.Y) + Abs(v.Position.Z)
		kin := Abs(v.Velocity.X) + Abs(v.Velocity.Y) + Abs(v.Velocity.Z)
		ret += pot * kin
	}
	return ret
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func GetTc1() []Moon {
	return []Moon{
		Moon{Position: Coords{X: -1, Y: 0, Z: 2}},
		Moon{Position: Coords{X: 2, Y: -10, Z: -7}},
		Moon{Position: Coords{X: 4, Y: -8, Z: 8}},
		Moon{Position: Coords{X: 3, Y: 5, Z: -1}}}
}

func GetTc1X() []Moon {
	return []Moon{
		Moon{Position: Coords{X: -1, Y: 0, Z: 0}},
		Moon{Position: Coords{X: 2, Y: 0, Z: 0}},
		Moon{Position: Coords{X: 4, Y: 0, Z: 0}},
		Moon{Position: Coords{X: 3, Y: 0, Z: 0}}}
}

func GetTc1Y() []Moon {
	return []Moon{
		Moon{Position: Coords{X: 0, Y: 0, Z: 0}},
		Moon{Position: Coords{X: 0, Y: -10, Z: 0}},
		Moon{Position: Coords{X: 0, Y: -8, Z: 0}},
		Moon{Position: Coords{X: 0, Y: 5, Z: 0}}}
}

func GetTc1Z() []Moon {
	return []Moon{
		Moon{Position: Coords{X: 0, Y: 0, Z: 2}},
		Moon{Position: Coords{X: 0, Y: 0, Z: -7}},
		Moon{Position: Coords{X: 0, Y: 0, Z: 8}},
		Moon{Position: Coords{X: 0, Y: 0, Z: -1}}}
}
