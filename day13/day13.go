package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Fold struct {
	axis  string
	value int
}

func main() {
	rawInput, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(rawInput), "\n")
	input = input[:len(input)-1]

	points := []Point{}
	folds := []Fold{}

	for _, line := range input {
		if strings.ContainsRune(line, ',') {
			parts := strings.Split(line, ",")
			n1, _ := strconv.Atoi(parts[0])
			n2, _ := strconv.Atoi(parts[1])
			points = append(points, Point{n1, n2})
		} else if strings.ContainsRune(line, '=') {
			eq := strings.IndexRune(line, '=')
			n, _ := strconv.Atoi(line[eq+1:])
			folds = append(folds, Fold{axis: line[eq-1 : eq], value: n})
		}
	}

	part1(points, folds)
	part2(points, folds)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Display(points []Point) {
	var maxX, maxY int
	for _, point := range points {
		if point.x > maxX {
			maxX = point.x
		}
		if point.y > maxY {
			maxY = point.y
		}
	}

	grid := make([][]rune, maxY+1)
	for i := 0; i <= maxY; i++ {
		grid[i] = make([]rune, maxX+1)
		for j := 0; j <= maxX; j++ {
			grid[i][j] = ' '
		}
	}

	for _, point := range points {
		grid[point.y][point.x] = '#'
	}

	for _, line := range grid {
		fmt.Printf("%v\n", string(line))
	}
}

func part1(points []Point, folds []Fold) {
	for _, fold := range folds {
		if fold.axis == "x" {
			for ind, point := range points {
				if point.x > fold.value {
					points[ind].x = fold.value - abs(point.x-fold.value)
				}
			}
		} else { // y
			for ind, point := range points {
				if point.y > fold.value {
					points[ind].y = fold.value - abs(point.y-fold.value)
				}
			}
		}
		break // Just the first instruction
	}

	uniq := make(map[Point]struct{})
	for _, point := range points {
		uniq[point] = struct{}{}
	}

	fmt.Println(len(uniq))
}

func part2(points []Point, folds []Fold) {
	for _, fold := range folds {
		if fold.axis == "x" {
			for ind, point := range points {
				if point.x > fold.value {
					points[ind].x = fold.value - abs(point.x-fold.value)
				}
			}
		} else { // y
			for ind, point := range points {
				if point.y > fold.value {
					points[ind].y = fold.value - abs(point.y-fold.value)
				}
			}
		}
	}

	Display(points)
}
