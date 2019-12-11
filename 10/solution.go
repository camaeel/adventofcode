package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	check(err)
	data, err1 := ioutil.ReadAll(f)
	check(err1)
	dataStr := string(data)
	fmt.Println(dataStr)
	mapa := ParseMap(dataStr)

	maxX, maxY, maxValue := FindBestSpot(mapa)

	fmt.Println("Max asteroids visible from: ", maxX, ",", maxY, ", visible:", maxValue)
	laser := Point{X: maxX, Y: maxY}

	targets := GetTargets(mapa, laser)

	fmt.Println("Found", len(targets), "targets")
	mapa2 := make([]string, len(mapa))
	copy(mapa2, mapa)

	canonDir := 180.0 + 0.000001
	maxTarget := 200
	// panic("aaa")
	var target *Target

	for i := 0; len(targets) > 0 && i < maxTarget; {
		if i == 0 {
			fmt.Println()
			// panic("aaa")
		}
		fmt.Println()
		target = FindNearestTarget(laser, canonDir, targets)
		if target != nil {
			canonDir = target.Angle
			DestroyTarget(target)
			fmt.Println("destroyed :", target)
			MakeTargetsVisible(targets, *target, laser)
			i++
			fmt.Println("Iteration:", i, ", Found target:", target)
		} else {
			canonDir = 360.0
			fmt.Println("Reset canon:", canonDir)
		}

	}

	fmt.Println("Target: ", target)
	fmt.Println("Part2: ", 100*target.P.X+target.P.Y)
}

type Point struct {
	X int
	Y int
}

type Target struct {
	P         Point
	Angle     float64
	Visible   bool
	Destroyed bool
}

func PointDistance(c Point, laser Point) int {
	return (c.X-laser.X)*(c.X-laser.X) + (c.Y-laser.Y)*(c.Y-laser.Y)
}

func MakeTargetsVisible(targets []*Target, target Target, laser Point) {
	candidates := make([]*Target, 0)
	for _, t := range targets {
		if t.Destroyed == false && t.Angle == target.Angle && (t.P.X != target.P.X || t.P.Y != target.P.Y) {
			candidates = append(candidates, t)
		}
	}
	minCandidateDistance := 0
	var minCandidate *Target
	for _, c := range candidates {
		dist := PointDistance(c.P, laser)
		if minCandidate == nil || minCandidateDistance > dist {
			minCandidate = c
			minCandidateDistance = dist
		}
	}
	if minCandidate != nil {
		minCandidate.Visible = true
		fmt.Println("making target visible:", minCandidate)
	}
}

func DestroyTarget(target *Target) {
	target.Destroyed = true
}

func FindNearestTarget(laser Point, canonDir float64, targets []*Target) *Target {
	maxAngle := -0.1
	var ret *Target
	for _, v := range targets {
		if v.Visible == true && v.Destroyed == false && v.Angle < canonDir {
			if maxAngle < v.Angle {
				ret = v
				maxAngle = ret.Angle
			}
		}
	}
	return ret
}

func GetTargets(mapa []string, laser Point) []*Target {
	result := make([]*Target, 0)
	w, h := MapDimensions(mapa)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if IsAsteroid(mapa, x, y) && (x != laser.X || y != laser.Y) {
				punkt := Point{X: x, Y: y}
				target := Target{P: punkt, Destroyed: false, Visible: true, Angle: CalculateAngle(laser, punkt)}
				result = append(result, &target)
			}
		}
	}

	for _, t1 := range result {
		for _, t2 := range result {
			if t1 != t2 && t1.Angle == t2.Angle {
				dist1 := PointDistance(t1.P, laser)
				dist2 := PointDistance(t2.P, laser)
				if dist1 < dist2 {
					// t1.Visible = true
					t2.Visible = false
				} else {
					t1.Visible = false
					// t2.Visible = true
				}
			}
		}
	}

	return result
}

func CalculateAngle(laser Point, punkt Point) float64 {

	angle := math.Atan2(float64(punkt.Y-laser.Y), float64(punkt.X-laser.X))*(180/math.Pi) + 90.0
	if angle < 0.0 {
		angle += 360.0
	}
	return angle
}

func FindBestSpot(mapa []string) (int, int, int) {

	// fmt.Println(mapa)
	w, h := MapDimensions(mapa)

	maxX, maxY := 0, 0
	maxValue := 0

	fmt.Println("w=", w, ", h=", h)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if IsAsteroid(mapa, x, y) {
				visibleAsteroids := CountVisibleAsteroids(mapa, x, y)
				if visibleAsteroids > maxValue {
					maxValue = visibleAsteroids
					maxX = x
					maxY = y
				}
			}
		}
	}
	return maxX, maxY, maxValue
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ParseMap(mapa string) []string {
	rows := strings.Split(mapa, "\n")
	return rows
}

func MapDimensions(mapa []string) (int, int) {
	return len(mapa[0]), len(mapa)
}

func IsAsteroid(mapa []string, x int, y int) bool {
	return (mapa[y][x] == '#')
}

func CountVisibleAsteroids(mapa []string, x int, y int) int {
	res := 0
	targets := GetTargets(mapa, Point{X: x, Y: y})
	for _, v := range targets {
		if v.Visible == true {
			res++
		}
	}
	// fmt.Println()
	// fmt.Println(Draw(mapa, targets, Point{X: x, Y: y}))
	// fmt.Println("visible tgts:", res)
	// fmt.Println()
	return res
}

func Draw(mapa []string, targets []*Target, laser Point) string {
	result := ""
	w, h := MapDimensions(mapa)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if x == laser.X && y == laser.Y {
				result += "X"
			} else {
				t := FindTarget(targets, x, y)
				if t == nil {
					result += "."
				} else if t.Visible == true {
					result += "#"
				} else if t.Visible == false {
					result += "+"
				}
			}
		}
		result += "\n"
	}

	return result
}

func FindTarget(targets []*Target, x int, y int) *Target {
	for _, v := range targets {
		if v.P.X == x && v.P.Y == y {
			return v
		}
	}
	return nil
}
