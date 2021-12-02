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

	//part1(input)
	part2(input)
}

func part1(lines []string) {
	var position, depth int

	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		num, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		switch parts[0] {
		case "forward":
			position += num
		case "up":
			depth -= num
		case "down":
			depth += num
		default:
			panic("Unknown" + parts[0])
		}
	}

	fmt.Println(position * depth)

}

func part2(lines []string) {
	var position, depth, aim int

	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		num, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		switch parts[0] {
		case "forward":
			position += num
			depth += aim * num
		case "up":
			aim -= num
		case "down":
			aim += num
		default:
			panic("Unknown" + parts[0])
		}
	}

	fmt.Println(position * depth)
}
