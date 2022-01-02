package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Num struct {
	L, R     *Num
	Val      int
	exploded bool
}

func (n Num) String() string {
	if n.L == nil && n.R == nil {
		return strconv.Itoa(n.Val)
	}
	if n.exploded {
		return fmt.Sprintf("[%v,%v]: EXPLODED", n.L, n.R)
	}
	return fmt.Sprintf("[%v,%v]", n.L, n.R)
}

func (n Num) IsNum() bool {
	return n.L == nil && n.R == nil
}

func (n *Num) Add(other *Num) {
	n.L = &Num{L: n.L, R: n.R}
	n.R = other
}

func (n *Num) MarkExplodable(depth int) {
	// Since we always go to the leftmost branch first, this should give the
	// first pair to be exploded
	if depth == 4 {
		n.exploded = true
		return
	}
	if n.L != nil {
		n.L.MarkExplodable(depth + 1)
	}
	if n.R != nil {
		n.R.MarkExplodable(depth + 1)
	}
}

func (n *Num) Reduce() {
	changed := true
	for changed {
		_, _, changed = n.reduce(0, 0, 0, false)
	}
}

func (n *Num) reduce(d, ladd, radd int, changed bool) (l, r int, change bool) {
	// If any pair is nested 4 deep, then explode it and start returning down
	// the values to add to other pairs
	if !changed && d == 4 && !n.IsNum() {
		fmt.Println(n)
		n.exploded = true
		if n.L != nil {
			l = n.L.Val
		}
		if n.R != nil {
			r = n.R.Val
		}
		return l, r, true
	}

	if n.L != nil {
		// If the left branch is a number greater than 10, split it into a pair
		if !changed && n.L.IsNum() && n.L.Val > 10 {
			n.L = &Num{L: &Num{Val: n.L.Val / 2}, R: &Num{Val: (n.L.Val / 2) + 1}}
			return 0, 0, true
		}

		if n.L.IsNum() {
			n.L.Val += radd
			radd = 0
		}

		// Reduce the left nested pair first
		l, r, change = n.L.reduce(d+1, ladd, radd, changed)
		if change {
			changed = true
		}
		// If the right is a number, then add to it and reset the number
		if n.R.IsNum() {
			n.R.Val += r
			r = 0
		}

		// If that same pair was exploded, then remove it and replace it with a zero
		if n.L.exploded {
			n.L = &Num{Val: 0}
		}
	}

	if n.R != nil {
		// If the left branch is a number greater than 10, split it into a pair
		if !changed && n.R.IsNum() && n.R.Val > 10 {
			n.R = &Num{L: &Num{Val: n.R.Val / 2}, R: &Num{Val: (n.R.Val / 2) + 1}}
			return 0, 0, true
		}

		if n.R.IsNum() {
			n.R.Val += ladd
			ladd = 0
		}

		l, r, change = n.R.reduce(d+1, l, r, changed)
		if change {
			changed = true
		}
		if n.L.IsNum() {
			n.L.Val += l
			l = 0
		}

		if n.R.exploded {
			n.R = &Num{Val: 0}
		}
	}

	return l, r, changed
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

func part1(input []string) {
	total := Parse(input[0])
	for _, inp := range input[1:] {
		fmt.Printf("%v + %v = ", total, Parse(inp))
		total.Add(Parse(inp))
		total.MarkExplodable(0)
		fmt.Printf("%v\n", total)
	}
}

func part2(input []string) {
}
