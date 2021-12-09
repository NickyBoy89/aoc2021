package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
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

func num(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return n
}

func part1(input []string) {
	var risk int

	for ri, row := range input {
		for ci, col := range row {
			val := num(string(col))
			// Check left-right
			if ci+1 < len(row) && num(string(input[ri][ci+1])) <= val {
				continue
			}
			if ci-1 > -1 && num(string(input[ri][ci-1])) <= val {
				continue
			}
			// Up-down
			if ri+1 < len(input) && num(string(input[ri+1][ci])) <= val {
				continue
			}
			if ri-1 > -1 && num(string(input[ri-1][ci])) <= val {
				continue
			}
			risk += val + 1
		}
	}
	fmt.Println(risk)
}

type Pair struct {
	x, y int
}

func part2(input []string) {
	basins := []map[int][]Pair{}
	for ri, row := range input {
		for ci := range row {
			size, positions := floodFill(ci, ri, input)
			sort.Slice(positions, func(i, j int) bool {
				return fmt.Sprintf("%v%v", positions[i].x, positions[i].y) < fmt.Sprintf("%v%v", positions[j].x, positions[j].y)
			})
			var present bool
			for _, b := range basins {
				if reflect.DeepEqual(b, map[int][]Pair{size: positions}) {
					present = true
					break
				}
			}
			if !present {
				basins = append(basins, map[int][]Pair{size: positions})
			}
		}
	}

	sizes := make([]int, 0, len(basins))
	for _, s := range basins {
		for i := range s {
			sizes = append(sizes, i)
		}
	}
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})
	fmt.Println(mulSum(sizes[:3]))
}

func mulSum(n []int) int {
	total := 1
	for _, num := range n {
		total *= num
	}
	return total
}

func floodFill(x, y int, arr []string) (size int, positions []Pair) {
	q := []Pair{Pair{x, y}}
	positions = []Pair{}
	seen := make(map[Pair]struct{})
	for len(q) != 0 {
		n := q[0]
		q = q[1:]
		if _, in := seen[n]; in {
			continue
		}
		if n.y < len(arr) && n.x < len(arr[0]) && n.y > -1 && n.x > -1 && rune(arr[n.y][n.x]) != '9' {
			positions = append(positions, Pair{n.x, n.y})
			size++
			q = append(q, Pair{n.x + 1, n.y})
			q = append(q, Pair{n.x - 1, n.y})
			q = append(q, Pair{n.x, n.y + 1})
			q = append(q, Pair{n.x, n.y - 1})
		}
		seen[n] = struct{}{}
	}
	return size, positions
}
