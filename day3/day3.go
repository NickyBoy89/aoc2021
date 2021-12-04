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

func part1(lines []string) {
	zeroes := make([]int, len(lines[0]))
	ones := make([]int, len(lines[0]))
	for _, line := range lines {
		for ind, bit := range line {
			if bit == '0' {
				zeroes[ind] += 1
			} else if bit == '1' {
				ones[ind] += 1
			}
		}
	}

	gamma := make([]rune, len(lines[0]))

	for num := 0; num < len(zeroes); num++ {
		if zeroes[num] > ones[num] {
			gamma[num] = '0'
		} else {
			gamma[num] = '1'
		}
	}

	gammaNum, err := strconv.ParseUint(string(gamma), 2, 64)
	if err != nil {
		panic(err)
	}

	omega := ^gammaNum ^ (^uint64(0) << len(string(gamma)))

	fmt.Println(gammaNum * omega)
}

func part2(lines []string) {
	pool := make([]string, len(lines))
	copy(pool, lines)

	// For a tiebreaker, use 1
	for len(pool) > 1 {
		zeroes, ones := countBits(pool)
		for num := 0; num < len(zeroes); num++ {
			zeroes, ones = countBits(pool)
			selected := '1'
			if zeroes[num] > ones[num] {
				selected = '0'
			}

			passed := []string{}

			for _, item := range pool {
				if rune(item[num]) == selected {
					passed = append(passed, item)
				}
			}

			pool = passed
		}
	}

	gamma, err := strconv.ParseInt(pool[0], 2, 64)
	if err != nil {
		panic(err)
	}

	pool = make([]string, len(lines))
	copy(pool, lines)

	// For a tiebreaker, use 0
	for len(pool) > 1 {
		zeroes, ones := countBits(pool)
		for num := 0; num < len(zeroes); num++ {
			zeroes, ones = countBits(pool)
			selected := '0'
			if zeroes[num] > ones[num] {
				selected = '1'
			}

			passed := []string{}

			if len(pool) == 1 {
				break
			}
			for _, item := range pool {
				if rune(item[num]) == selected {
					passed = append(passed, item)
				}
			}

			pool = passed
		}
	}

	omega, err := strconv.ParseInt(pool[0], 2, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println(gamma * omega)
}

func countBits(lines []string) (zeroes, ones []int) {
	zeroes = make([]int, len(lines[0]))
	ones = make([]int, len(lines[0]))

	for _, line := range lines {
		for ind, bit := range line {
			if bit == '0' {
				zeroes[ind] += 1
			} else if bit == '1' {
				ones[ind] += 1
			}
		}
	}

	return zeroes, ones
}
