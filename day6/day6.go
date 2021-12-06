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

	input := strings.Split(string(rawInput), ",")

	part1(input)
	part2(input)
}

func part1(input []string) {
	fishes := toInts(input)
	for gen := 0; gen < 80; gen++ {
		for ind, fish := range fishes {
			if fish == 0 {
				fishes[ind] = 6
				fishes = append(fishes, 8)
			} else {
				fishes[ind]--
			}
		}
	}

	fmt.Println(len(fishes))
}

func toInts(strs []string) []int {
	result := make([]int, len(strs))
	for ind, str := range strs {
		n, err := strconv.Atoi(strings.Trim(str, "\n"))
		if err != nil {
			panic(err)
		}
		result[ind] = n
	}
	return result
}

func part2(input []string) {
	intInput := toInts(input)
	fishes := make(map[int]int64, len(intInput))
	for _, inp := range intInput {
		fishes[inp] += 1
	}
	for gen := 0; gen < 256; gen++ {
		for fish, count := range deepCopy(fishes) {
			if fish == 0 {
				fishes[0] -= count
				fishes[6] += count
				fishes[8] += count
				continue
			}
			fishes[fish-1] += count
			fishes[fish] -= count
		}
	}

	var total int64
	for _, count := range fishes {
		total += count
	}
	fmt.Println(total)
}

func deepCopy(src map[int]int64) map[int]int64 {
	result := make(map[int]int64, len(src))
	for key, val := range src {
		result[key] = val
	}
	return result
}
