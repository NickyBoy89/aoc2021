package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

func (p Point) Mul(f int) Point {
	return Point{p.X * f, p.Y * f}
}

func (p Point) Neighbors(grid map[Point]int) []Point {
	nei := []Point{
		Point{p.X - 1, p.Y},
		Point{p.X + 1, p.Y},
		Point{p.X, p.Y - 1},
		Point{p.X, p.Y + 1},
	}

	valid := []Point{}

	for _, item := range nei {
		if _, in := grid[item]; in {
			valid = append(valid, item)
		}
	}

	return valid
}

func d(source, target Point, grid map[Point]int) int {
	frontier := PriorityQueue{&Item{
		value:    source,
		priority: 0,
		index:    0,
	}}
	heap.Init(&frontier)
	came_from := make(map[Point]Point)
	cost_so_far := make(map[Point]int)
	came_from[source] = Point{-1, -1}
	cost_so_far[source] = 0

	for frontier.Len() != 0 {
		current := heap.Pop(&frontier).(*Item)
		if current.value == target {
			return current.priority
		}

		for _, next := range current.value.Neighbors(grid) {
			new_cost := cost_so_far[current.value] + grid[next]
			if _, in := cost_so_far[next]; !in || new_cost < cost_so_far[next] {
				cost_so_far[next] = new_cost
				item := &Item{
					value:    next,
					priority: new_cost,
				}
				heap.Push(&frontier, item)
				came_from[next] = current.value
			}
		}
	}
	panic("No path")
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
	graph := make(map[Point]int)
	for ri := range input {
		for ci := range input[ri] {
			n, err := strconv.Atoi(string(input[ri][ci]))
			if err != nil {
				panic(err)
			}
			graph[Point{X: ci, Y: ri}] = n
		}
	}

	fmt.Println(d(Point{0, 0}, Point{len(input) - 1, len(input[0]) - 1}, graph))
}

func printMap(graph map[Point]int, width, height int) {
	result := make([][]int, height)
	for ri := range result {
		result[ri] = make([]int, width)
	}
	for node, weight := range graph {
		result[node.Y][node.X] = weight
	}

	for _, line := range result {
		for _, char := range line {
			fmt.Printf("%v", char)
		}
		fmt.Println("")
	}
}

func expandMap(graph map[Point]int, factor, width, height int) map[Point]int {
	expanded := make(map[Point]int)
	for point, weight := range graph {
		for f := 0; f < factor; f++ {
			for g := 0; g < factor; g++ {
				newWeight := weight + f + g
				for newWeight > 9 {
					newWeight -= 9
				}
				expanded[Point{point.X + width*f, point.Y + height*g}] = newWeight
			}
		}
	}
	return expanded
}

func part2(input []string) {
	graph := make(map[Point]int)
	for ri := range input {
		for ci := range input[ri] {
			n, err := strconv.Atoi(string(input[ri][ci]))
			if err != nil {
				panic(err)
			}
			graph[Point{X: ci, Y: ri}] = n
		}
	}

	factor := 5

	expanded := expandMap(graph, factor, len(input), len(input[0]))

	fmt.Println(d(Point{0, 0}, Point{len(input)*factor - 1, len(input[0])*factor - 1}, expanded))
}
