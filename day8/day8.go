package main

import (
	"fmt"
	"os"
	"strings"
)

var digitNumbers = map[int]int{
	0: 6,
	1: 2,
	2: 5,
	3: 5,
	4: 4,
	5: 5,
	6: 6,
	7: 3,
	8: 7,
	9: 6,
}

func digitsToNumber(digits int) int {
	for number, segments := range digitNumbers {
		if digits == segments {
			return number
		}
	}
	panic("asds")
}

func plausibleSegments(number int) []int {
	count := []int{}
	for num, segments := range digitNumbers {
		if segments == number {
			count = append(count, num)
		}
	}
	return count
}

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
		inputs[ind] = strings.Split(parts[0], " ")[1:]
		outputs[ind] = strings.Split(parts[1], " ")[1:]
	}

	part1(inputs, outputs)
	part2(inputs, outputs)
}

func numRange(n int) []int {
	result := make([]int, n)
	for ind := 0; ind < n; ind++ {
		result[ind] = ind
	}
	return result
}

func remove(nums []int, number int) []int {
	toRemove := -1
	for ind, n := range nums {
		if n == number {
			toRemove = ind
			break
		}
	}

	if toRemove == -1 {
		return nums
	}

	nums[toRemove] = nums[len(nums)-1]
	return nums[:len(nums)-1]
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
				/*
					case 2: // 1
						display[ind][group] = []int{1}
					case 4: // 4
						display[ind][group] = []int{4}
					case 3: // 7
						display[ind][group] = []int{7}
					case 7: // 8
						display[ind][group] = []int{8}
					default:
						//display[ind][group] = numRange(len(inputs[lineNumber]))
				*/
			}
		}

		/*
			fmt.Println(display)

			for _, group := range display {
				for name, possible := range group {
					if len(possible) != 1 {
						continue
					}
					fmt.Println(outputs[lineNumber])
					for _, inp := range outputs[lineNumber] {
						if name == inp {
							fmt.Println(name)
							overlapped++
						}
					}
				}
			}
		*/
	}
	fmt.Println(overlapped)
}

func filterStep(inputs Display) Display {
	for _, display := range inputs {
		for digits, possibleVals := range display {

			if len(possibleVals) == 1 {
				for ind := range inputs {
					for i, j := range inputs[ind] {
						if i == digits {
							continue
						}
						inputs[ind][i] = remove(inputs[ind][i], possibleVals[0])
						_ = j
					}
				}
			}
		}
	}
	return inputs
}

func union(n1, n2 []string) []string {
	fmt.Println(n1, n2)
	union := []string{}
	for _, i := range n1 {
		for _, j := range n2 {
			if j == i {
				union = append(union, i)
			}
		}
	}
	fmt.Println(union)
	return union
}

func part2(inputs, outputs [][]string) {
	// For every line
	for lineNumber := range inputs {
		possible := make(map[int][]string)

		for _, digits := range inputs[lineNumber] {
			positions := []int{}
			switch len(digits) {
			case 2:
				positions = []int{2, 5}
			case 4:
				positions = []int{1, 2, 3, 5}
			case 3:
				positions = []int{0, 2, 5}
			case 7:
				positions = []int{0, 1, 2, 3, 4, 5, 6, 7}
			case 5, 6:
				positions = numRange(8)
			}

			for _, pos := range positions {
				if len(possible[pos]) == 0 {
					possible[pos] = strings.Split(digits, "")
				} else {
					possible[pos] = union(possible[pos], strings.Split(digits, ""))
				}
			}
		}

		fmt.Println(possible)
	}
}
