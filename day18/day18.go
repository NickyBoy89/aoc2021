package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func IndexWithSkip(input string, target rune) int {
	var balance int
	for ind, char := range input {
		switch rune(char) {
		case '[':
			balance++
		case ']':
			balance--
		case target:
			if balance == 0 {
				return ind
			}
		}
	}
	return -1
}

func MatchingBracket(input string, ind int) int {
	balance := 0
	for i := ind; i < len(input); i++ {
		switch rune(input[i]) {
		case '[':
			balance--
		case ']':
			balance++
			if balance == 0 {
				return i
			}
		}
	}
	return -1
}

func nextNum(input string, ind, increment int) []int {
	start := -1
	for i := ind; i < len(input) && i > 0; i += increment {
		if unicode.IsDigit(rune(input[i])) {
			if start == -1 {
				start = i
			}
		} else if start != -1 {
			return []int{start, i}
		}
	}
	return nil
}

func reduce(input string, split bool) (string, bool) {
	balance := -1
	lastDigit := 0
	for i := 1; i < len(input); i++ {
		switch rune(input[i]) {
		case '[':
			// Nested inside four pairs, explode it
			if balance == -4 {
				end := MatchingBracket(input, i)
				explodedPair := input[i+1 : end]
				explosionNumbers := strings.Split(explodedPair, ",")
				// Replace the pair with a "0"
				input = input[:i] + "0" + input[end+1:]
				// Increment the first number to the right
				ni := nextNum(input, i+1, 1)
				if ni != nil {
					n, err := strconv.Atoi(input[ni[0]:ni[1]])
					if err != nil {
						panic(err)
					}
					n1, err := strconv.Atoi(explosionNumbers[1])
					if err != nil {
						panic(err)
					}
					input = input[:ni[0]] + strconv.Itoa(n+n1) + input[ni[1]:]
				}
				ni = nextNum(input, i-1, -1)
				if ni != nil {
					n, err := strconv.Atoi(input[ni[1]+1 : ni[0]+1])
					if err != nil {
						panic(err)
					}
					n1, err := strconv.Atoi(explosionNumbers[0])
					if err != nil {
						panic(err)
					}
					input = input[:ni[1]+1] + strconv.Itoa(n+n1) + input[ni[0]+1:]
				}
				return input, true
			}
			balance--
			lastDigit = i + 1
		case ']', ',': // Signal a number literal
			if i-lastDigit > 0 && split {
				n, err := strconv.Atoi(input[lastDigit:i])
				if err != nil {
					panic(err)
				}
				// Split if 10 or greater
				if n >= 10 {
					var splitPair string
					if n%2 == 0 { // If odd, both pairs are equal
						splitPair = fmt.Sprintf("[%v,%v]", n/2, n/2)
					} else { // Even, the right pair should be one more
						splitPair = fmt.Sprintf("[%v,%v]", n/2, n/2+1)
					}
					input = input[:lastDigit] + splitPair + input[i:]
					return input, true
				}
			}
			lastDigit = i + 1
			if rune(input[i]) == ']' {
				balance++
			}
		}
	}
	return input, false
}

type Num struct {
	L, R *Num
	Val  int
}

func (n Num) Magnitude() int {
	if n.L == nil && n.R == nil {
		return n.Val
	}
	return 3*n.L.Magnitude() + 2*n.R.Magnitude()
}

func Parse(line string) *Num {
	num := new(Num)
	if len(line) > 0 && line[0] == '[' {
		comma := IndexWithSkip(line[1:len(line)-1], ',') + 1
		num.L = Parse(line[1:comma])
		num.R = Parse(line[comma+1 : len(line)-1])
		return num
	}
	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	return &Num{Val: n}
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

func Reduce(input string) string {
	changed := true
	for changed {
		// Repeatedly try and explode numbers until there are no more
		input, changed = reduce(input, false)
		if changed {
			continue
		}
		// If no explosions, start reducing
		input, changed = reduce(input, true)
	}
	return input
}

func part1(input []string) {
	total := input[0]
	for _, inp := range input[1:] {
		total = fmt.Sprintf("[%v,%v]", total, inp)
		total = Reduce(total)
	}

	parsedTotal := Parse(total)
	fmt.Println(parsedTotal.Magnitude())
}

func part2(input []string) {
	var max int
	for fi, first := range input {
		for si, second := range input {
			if fi == si {
				continue
			}
			mag := Parse(Reduce(fmt.Sprintf("[%v,%v]", first, second))).Magnitude()
			if mag > max {
				max = mag
			}
		}
	}
	fmt.Println(max)
}
