package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var xMin, xMax, yMin, yMax int

type Point struct {
	X, Y int
}

func (p Point) Valid() bool {
	return p.X >= xMin && p.X <= xMax && p.Y >= yMin && p.Y <= yMax
}

func main() {
	rawInput, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(rawInput), "\n")
	input = input[:len(input)-1]

	rawX := input[0][len("target area: x="):strings.IndexRune(input[0], ',')]

	rawY := input[0][strings.Index(input[0], "y=")+len("y="):]

	xMin, err = strconv.Atoi(rawX[:strings.IndexRune(rawX, '.')])
	if err != nil {
		panic(err)
	}

	xMax, err = strconv.Atoi(rawX[strings.LastIndex(rawX, ".")+1:])
	if err != nil {
		panic(err)
	}

	yMin, err = strconv.Atoi(rawY[:strings.IndexRune(rawY, '.')])
	if err != nil {
		panic(err)
	}

	yMax, err = strconv.Atoi(rawY[strings.LastIndex(rawY, ".")+1:])
	if err != nil {
		panic(err)
	}

	part1()
	part2()
}

func trajectory(xVel, yVel int) []Point {
	cur := Point{0, 0}
	points := []Point{cur}
	for cur.Y >= yMin {
		cur.X += xVel
		cur.Y += yVel
		points = append(points, Point{cur.X, cur.Y})
		if xVel > 0 {
			xVel--
		} else if xVel < 0 {
			xVel++
		}
		yVel--
	}
	return points
}

func part1() {
	valid := make(map[Point]int)
	iter := 300
	for i := 0; i < iter; i++ {
		for j := 0; j < iter; j++ {
			var maxHeight int
			for _, point := range trajectory(i, j) {
				if point.Y > maxHeight {
					maxHeight = point.Y
				}
				if point.Valid() {
					valid[Point{i, j}] = maxHeight
					break
				}
			}
		}
	}

	var max int
	for _, val := range valid {
		if val > max {
			max = val
		}
	}

	fmt.Println(max)
}

func part2() {
	valid := make(map[Point]int)
	iter := 1_000
	for i := -iter; i < iter; i++ {
		for j := -iter; j < iter; j++ {
			var maxHeight int
			for _, point := range trajectory(i, j) {
				if point.Y > maxHeight {
					maxHeight = point.Y
				}
				if point.Valid() {
					valid[Point{i, j}] = maxHeight
					break
				}
			}
		}
	}

	fmt.Println(len(valid))
}
