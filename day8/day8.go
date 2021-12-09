package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set"
)

// Display is a list of maps from the characters, and the possible digits
type Display []map[string][]int

func main() {
	rawInput, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(rawInput), "\n")
	input = input[:len(input)-1]

	inputs := make([][]string, len(input))
	outputs := make([][]string, len(input))
	for ind, inp := range input {
		parts := strings.Split(inp, "|")
		inputs[ind] = strings.Split(parts[0], " ")
		outputs[ind] = strings.Split(parts[1], " ")
	}

	part1(inputs, outputs)
	part2(inputs, outputs)
}

func part1(inputs, outputs [][]string) {
	var overlapped int
	// For every line
	for lineNumber := range outputs {
		display := make(Display, len(outputs[lineNumber]))
		for ind, group := range outputs[lineNumber] {
			display[ind] = make(map[string][]int)
			switch len(group) {
			case 2, 4, 3, 7:
				overlapped++
			}
		}
	}
	fmt.Println(overlapped)
}

func part2(inputs, outputs [][]string) {
	var total int
	// For every line
	for lineNumber := range inputs {
		// Every digit and its known possibilities
		known := make(map[int]mapset.Set)

		for _, digits := range inputs[lineNumber] {
			digitSet := mapset.NewSet()
			for _, d := range digits {
				digitSet.Add(string(d))
			}
			switch len(digits) {
			case 2: // 1
				known[1] = digitSet
			case 3: // 7
				known[7] = digitSet
			case 7: // 8
				known[8] = digitSet
			case 4: // 4
				known[4] = digitSet
			}
		}

		for _, digits := range inputs[lineNumber] {
			digitSet := mapset.NewSet()
			for _, d := range digits {
				digitSet.Add(string(d))
			}
			if len(digits) == 5 && known[1].Difference(digitSet).Cardinality() == 0 {
				known[3] = digitSet
			} else if len(digits) == 6 {
				if known[1].Difference(digitSet).Cardinality() == 1 {
					known[6] = digitSet
				} else if known[4].Difference(digitSet).Cardinality() == 0 {
					known[9] = digitSet
				} else if known[4].Difference(digitSet).Cardinality() == 1 {
					known[0] = digitSet
				}
			}
		}

		for _, digits := range inputs[lineNumber] {
			digitSet := mapset.NewSet()
			for _, d := range digits {
				digitSet.Add(string(d))
			}
			if len(digits) != 5 || digitSet.Equal(known[3]) {
				continue
			}

			if digitSet.Contains(known[1].Difference(known[6]).ToSlice()[0]) {
				known[2] = digitSet
			} else {
				known[5] = digitSet
			}
		}

		var result string
		for _, digits := range outputs[lineNumber] {
			digitSet := mapset.NewSet()
			for _, d := range digits {
				digitSet.Add(string(d))
			}
			for k, v := range known {
				if v.Equal(digitSet) {
					result += fmt.Sprintf("%v", k)
				}
			}
		}

		num, err := strconv.Atoi(result)
		if err != nil {
			panic(err)
		}
		total += num
	}
	fmt.Println(total)
}
