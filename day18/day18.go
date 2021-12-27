package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Num struct {
	L, R interface{}
}

func (n Num) String() string {
	return fmt.Sprintf("[%v,%v]", n.L, n.R)
}

func (n Num) Add(other Num) Num {
	return Num{L: n, R: other}
}

func (n Num) Reduce() {
	n.reduce(0)
}

func (n Num) reduce(d int) {
	if d == 4 {
		fmt.Println(n)
	}
	switch n.L.(type) {
	case Num:
		n.L.(Num).reduce(d + 1)
	}

	switch n.R.(type) {
	case Num:
		n.R.(Num).reduce(d + 1)
	}
}

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

func Parse(line string) interface{} {
	var num Num
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
	return n
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
	for _, inp := range input {
		p := Parse(inp).(Num)
		orig := p
		p.Reduce()
		fmt.Printf("%v -> %v\n", orig, p)
	}
}

func part2(input []string) {
}
