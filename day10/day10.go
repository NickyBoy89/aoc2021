package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

var syntaxScores = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var autoScores = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

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
	var totalScore int

	open := []rune{}
	for _, line := range input {
		for _, char := range line {
			keepGoing := true
			switch rune(char) {
			case '[', '(', '<', '{':
				open = append(open, rune(char))
			case ']':
				if open[len(open)-1] != '[' {
					totalScore += syntaxScores[rune(char)]
					keepGoing = false
				} else {
					open = open[:len(open)-1]
				}
			case ')':
				if open[len(open)-1] != '(' {
					totalScore += syntaxScores[rune(char)]
					keepGoing = false
				} else {
					open = open[:len(open)-1]
				}
			case '>':
				if open[len(open)-1] != '<' {
					totalScore += syntaxScores[rune(char)]
					keepGoing = false
				} else {
					open = open[:len(open)-1]
				}
			case '}':
				if open[len(open)-1] != '{' {
					totalScore += syntaxScores[rune(char)]
					keepGoing = false
				} else {
					open = open[:len(open)-1]
				}
			}
			if !keepGoing {
				break
			}
		}
	}

	fmt.Println(totalScore)
}

func part2(input []string) {
	scores := []int{}

	for _, line := range input {
		open := []rune{}
		for ind, char := range line {
			keepGoing := true
			switch rune(char) {
			case '[', '(', '<', '{':
				open = append(open, rune(char))
			case ']':
				if open[len(open)-1] != '[' {
					keepGoing = false
				} else {
					open = open[:len(open)-1]
				}
			case ')':
				if open[len(open)-1] != '(' {
					keepGoing = false
				} else {
					open = open[:len(open)-1]
				}
			case '>':
				if open[len(open)-1] != '<' {
					keepGoing = false
				} else {
					open = open[:len(open)-1]
				}
			case '}':
				if open[len(open)-1] != '{' {
					keepGoing = false
				} else {
					open = open[:len(open)-1]
				}
			}
			if !keepGoing {
				break
			}

			if ind == len(line)-1 && len(open) != 0 {
				var totalScore int
				for i := len(open) - 1; i >= 0; i-- {
					remaining := open[i]
					totalScore *= 5
					totalScore += autoScores[rune(remaining)]
				}
				scores = append(scores, totalScore)
			}
		}
	}

	sort.Ints(scores)

	fmt.Println(scores[len(scores)/2])
}
