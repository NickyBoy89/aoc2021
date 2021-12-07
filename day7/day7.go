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

	input := strings.Split(string(rawInput[:len(rawInput)-1]), ",")

	vals := make([]int, len(input))
	for ind, val := range input {
		num, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		vals[ind] = num
	}

	maxValue := 0
	for _, val := range vals {
		if val > maxValue {
			maxValue = val
		}
	}

	part1(vals, maxValue)
	part2(vals, maxValue)
}

func part1(heights []int, maxValue int) {
	costs := make(map[int]int)
	for set := 0; set < maxValue; set++ {
		cost := 0
		for _, height := range heights {
			cost += abs(height - set)
		}
		costs[set] = cost
	}

	minCost := -1
	for _, cost := range costs {
		if minCost == -1 || cost < minCost {
			minCost = cost
		}
	}

	fmt.Println(minCost)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func gauss_sum(num int) int {
	return (num * (num + 1)) / 2
}

func part2(heights []int, maxValue int) {
	costs := make(map[int]int)
	for set := 0; set < maxValue; set++ {
		cost := 0
		for _, height := range heights {
			cost += gauss_sum(abs(height - set))
		}
		costs[set] = cost
	}

	minCost := -1
	for _, cost := range costs {
		if minCost == -1 || cost < minCost {
			minCost = cost
		}
	}

	fmt.Println(minCost)
}
