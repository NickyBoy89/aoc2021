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
	for i := 0; i < 10; i++ {
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
	pairs := make(map[string]int)
	for i := 0; i < len(template); i++ {
		if i+1 < len(template) {
			pairs[string(template[i:i+2])]++
		}
	}

	charCounts := make(map[rune]int)

	for _, char := range template {
		charCounts[char]++
	}

	for step := 0; step < 40; step++ {
		newCounts := make(map[string]int)
		for item, count := range pairs {
			charCounts[rules[item]] += count
			newChar := string(rules[item])
			newCounts[string(item[0])+newChar] += count
			newCounts[newChar+string(item[1])] += count
		}

		pairs = newCounts
	}

	charCounts[template[0]]++

	var max int
	min := -1
	for _, count := range charCounts {
		if count > max {
			max = count
		}
		if min == -1 || count < min {
			min = count
		}
	}
	fmt.Println(max - min)
}
