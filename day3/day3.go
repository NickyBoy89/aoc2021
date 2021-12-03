package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	rawInput, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(rawInput), "\n")
	input = input[:len(input)-1]

	//part1(input)
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

	n1 := toDecimal(string(gamma))
	n2 := toDecimal(invert(string(gamma)))

	fmt.Println(string(gamma), n1, string(gamma), n2, n1*n2)
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

	gamma := pool[0]

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

	omega := pool[0]

	fmt.Println(gamma, omega)

	n1 := toDecimal(gamma)
	n2 := toDecimal(string(omega))

	fmt.Println(n1 * n2)
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

func toDecimal(input string) int {
	var result int
	for ind, char := range input {
		if char == '1' {
			result |= 1 << (len(input) - ind - 1)
		}
	}
	return result
}

func invert(input string) string {
	result := make([]rune, len(input))
	for ind, char := range input {
		if char == '0' {
			result[ind] = '1'
		}
	}
	return string(result)
}
