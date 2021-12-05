package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rawInput, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(rawInput), "\n")
	input = input[:len(input)-1]

	part1(input)
	part2(input)
}

func part1(input []string) {
	points := [][]string{}
	for _, line := range input {
		point := strings.Split(line, " -> ")
		points = append(points, []string{point[0], point[1]})
	}

	overlapping := make(map[string]int)

	for _, point := range points {
		start := point[0]
		end := point[1]
		startPos := strings.Split(start, ",")
		endPos := strings.Split(end, ",")
		if startPos[0] != endPos[0] && startPos[1] != endPos[1] {
			continue
		}
		startX, _ := strconv.Atoi(startPos[0])
		startY, _ := strconv.Atoi(startPos[1])

		endX, _ := strconv.Atoi(endPos[0])
		endY, _ := strconv.Atoi(endPos[1])

		slopeX := clamp(endX - startX)
		slopeY := clamp(endY - startY)

		curX, curY := startX, startY
		// For the off-by-one error
		var valid bool
		for curX != endX || curY != endY || valid {
			if _, present := overlapping[fmt.Sprintf("%v,%v", curX, curY)]; present {
				overlapping[fmt.Sprintf("%v,%v", curX, curY)]++
			} else {
				overlapping[fmt.Sprintf("%v,%v", curX, curY)] = 1
			}
			curX += slopeX
			curY += slopeY
			if valid {
				break
			}
			if curX == endX && curY == endY {
				valid = true
			}
		}
	}

	var overlaps int

	for _, count := range overlapping {
		if count > 1 {
			overlaps++
		}
	}

	fmt.Println(overlaps)
}

func clamp(num int) int {
	if num > 0 {
		return 1
	} else if num < 0 {
		return -1
	}
	return 0
}

func part2(input []string) {
	points := [][]string{}
	for _, line := range input {
		point := strings.Split(line, " -> ")
		points = append(points, []string{point[0], point[1]})
	}

	overlapping := make(map[string]int)

	for _, point := range points {
		start := point[0]
		end := point[1]
		startPos := strings.Split(start, ",")
		endPos := strings.Split(end, ",")
		startX, _ := strconv.Atoi(startPos[0])
		startY, _ := strconv.Atoi(startPos[1])

		endX, _ := strconv.Atoi(endPos[0])
		endY, _ := strconv.Atoi(endPos[1])

		slopeX := clamp(endX - startX)
		slopeY := clamp(endY - startY)

		curX, curY := startX, startY
		// For the off-by-one error
		var valid bool
		for curX != endX || curY != endY || valid {
			if _, present := overlapping[fmt.Sprintf("%v,%v", curX, curY)]; present {
				overlapping[fmt.Sprintf("%v,%v", curX, curY)]++
			} else {
				overlapping[fmt.Sprintf("%v,%v", curX, curY)] = 1
			}
			curX += slopeX
			curY += slopeY
			if valid {
				break
			}
			if curX == endX && curY == endY {
				valid = true
			}
		}
	}

	var overlaps int

	for _, count := range overlapping {
		if count > 1 {
			overlaps++
		}
	}

	fmt.Println(overlaps)
}
