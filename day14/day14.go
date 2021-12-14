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

	template := []rune(input[0])

	rules := make(map[string]rune, len(input[2:]))

	for _, rule := range input[2:] {
		parts := strings.Split(rule, " -> ")
		rules[parts[0]] = rune(parts[1][0])
	}

	part1(template, rules)
	part2(template, rules)
}

func part1(template []rune, rules map[string]rune) {
	for i := 0; i < 2; i++ {
		result := []rune{}
		for ind, item := range template {
			result = append(result, item)
			if ind < len(template)-1 {
				result = append(result, rules[string(template[ind:ind+2])])
			}
		}
		template = result
	}

	counts := make(map[rune]int)
	for _, poly := range template {
		counts[poly]++
	}

	var max int
	min := -1

	for _, count := range counts {
		if count > max {
			max = count
		}
		if min == -1 || count < min {
			min = count
		}
	}

	fmt.Println(max - min)
}

func part2(template []rune, rules map[string]rune) {
	counts := polyTree(template, rules, 2, make(map[string]string))
	fmt.Println(counts)
	var max int
	min := -1

	for _, count := range counts {
		if count > max {
			max = count
		}
		if min == -1 || count < min {
			min = count
		}
	}

	fmt.Println(max - min)
}

func polyTree(template []rune, rules map[string]rune, steps int, recipes map[string]string) map[rune]int {
	counts := make(map[rune]int)
	for ind := range template {
		if ind < len(template)-1 {
			pair := string(template[ind : ind+2])
			if _, in := recipes[pair]; !in {
				recipes[pair] = string(pair[0]) + string(rules[pair]) + string(pair[1])
			}
		}
	}
	return counts
}
