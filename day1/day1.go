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

	ints := make([]int, len(input))
	for ind, item := range input {
		num, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		ints[ind] = num
	}

	part1(ints)
	part2(ints)
}

func part1(input []int) {
	var numPrev int

	for ind, num := range input {
		if ind > 0 && num > input[ind-1] {
			numPrev++
		}
	}

	fmt.Println(numPrev)
}

func part2(input []int) {
	var numPrev int

	for ind := range input {
		if ind > 2 && sum(input[ind-2:ind+1]) > sum(input[ind-3:ind]) {
			numPrev++
		}
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
