package main

import (
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

	part1(input)
	part2(input)
}

func part1(input []string) {
	vals := make(map[rune]int)
	for _, instr := range input {
		switch instr[:strings.IndexRune(instr, ' ')] {
		case "inp":
			vals[rune(instr[len(instr)-1])] = 0
		case "add":
			locations := strings.Split(instr[len("add")+1:], " ")
			first := rune(locations[0][0])
			second := rune(locations[1][0])
			vals[first] = vals[first] * vals[second]
		case "div":
			locations := strings.Split(instr[len("div")+1:], " ")
			first := rune(locations[0][0])
			second := rune(locations[1][0])
			vals[first] = vals[first] / vals[second]
		case "mod":
			locations := strings.Split(instr[len("mod")+1:], " ")
			first := rune(locations[0][0])
			second := rune(locations[1][0])
			vals[first] = vals[first] % vals[second]
		case "eql":
			locations := strings.Split(instr[len("eql")+1:] " ")
			first := rune(locations[0][0])
			second := rune(locations[1][0])
			if vals[first] == vals[second] {
				vals[first] = 1
			} else {
				vals[first] = 0
			}
		}
	}
}

func part2(input []string) {
}
