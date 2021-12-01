package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	rawInput, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(rawInput), "\n")

	var numPrev int

	var prev int
	for _, depth := range input {
		num, _ := strconv.Atoi(depth)
		if num > prev {
			numPrev++
		}
		prev = num
	}

	fmt.Println(numPrev)
}

func part2() {
	rawInput, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	strInput := strings.Split(string(rawInput), "\n")

	input := make([]int, len(strInput))
	for ind, val := range strInput {
		if val == "" {
			continue
		}
		num, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		input[ind] = num
	}

	var numPrev int

	var prev int
	for ind, _ := range input {
		if ind-2 < 0 {
			continue
		}
		if sum(input[ind-2:ind+1]) > prev {
			numPrev++
		}
		prev = sum(input[ind-2 : ind+1])
	}

	fmt.Println(numPrev)
}

func sum(nums []int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	//part1()
	part2()
}
