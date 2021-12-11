package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Pos struct {
	x int
	y int
}

type Grid [][]int

func (g Grid) String() string {
	var total string
	for ind, row := range g {
		total += fmt.Sprintf("%v", row)
		if ind < len(g)-1 {
			total += "\n"
		}
	}
	return total
}

func (g Grid) FlashCount() int {
	var count int

	for i := range g {
		for j := range g[i] {
			if g[i][j] > 9 {
				count++
			}
		}
	}

	return count
}

// Gets a list of positions to update on a flash
func (g Grid) Flash(row, col int) {
	for y := row - 1; y < row+2; y++ {
		for x := col - 1; x < col+2; x++ {
			// Not the original position
			if y == row && x == col {
				continue
			}
			if y >= 0 && y < len(g) && x >= 0 && x < len(g[row]) {
				g[y][x]++
			}
		}
	}
}

func (g Grid) Step() int {
	var flashes int
	// Increment everything
	for i := range g {
		for j := range g[i] {
			g[i][j]++
		}
	}

	hasFlashed := []Pos{}

	for g.FlashCount() != 0 {
		for i := range g {
			for j := range g[i] {
				if g[i][j] > 9 {
					hasFlashed = append(hasFlashed, Pos{i, j})
					flashes++
					g.Flash(i, j)
				}
			}
		}
		for _, pos := range hasFlashed {
			g[pos.x][pos.y] = 0
		}
	}

	return flashes
}

func main() {
	rawInput, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(rawInput), "\n")
	input = input[:len(input)-1]

	// Make a grid of numbers
	nums := make(Grid, len(input))
	for li, line := range input {
		nums[li] = make([]int, len(line))
		for ri, char := range line {
			n, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			nums[li][ri] = n
		}
	}

	part1(nums)
	part2(nums)
}

func part1(nums Grid) {
	var totalFlashes int

	for i := 0; i < 100; i++ {
		totalFlashes += nums.Step()
	}

	fmt.Println(totalFlashes)
}

func part2(nums Grid) {
	blankGrid := make(Grid, len(input))
	for ind, line := range input {
		blankGrid[ind] = make([]int, len(line))
	}

	steps := 0
	for ; !reflect.DeepEqual(nums, blankGrid); steps++ {
		nums.Step()
	}

	fmt.Println(steps)
}
