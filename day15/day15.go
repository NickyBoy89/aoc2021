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
}

func part2(input []string) {
}

type Point struct {
	X, Y int
}

func manhattanDistance(p1, p2 Point) int {
	return abs(p2.Y-p1.Y) + abs(p2.X-p1.X)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func reconstruct_path(cameFrom map[Point]Point, current Point) []Point {
	total_path := []Point{current}
	for _, current := range cameFrom {
		total_path = append([]Point{current}, total_path...)
	}
	return total_path
}

func minInSet(set map[Point]int) Point {
	min := -1
	var minItem Point
	for item, score := range set {
		if score == -1 || score < min {
			min = score
			minItem = item
		}
	}
	return minItem
}

// A* finds a path from start to goal.
// h is the heuristic function. h(n) estimates the cost to reach goal from node n.
func A_Star(start, goal Point, h func(p1, p2 Point) int) []Point {
	// The set of discovered nodes that may need to be (re-)expanded.
	// Initially, only the start node is known.
	// This is usually implemented as a min-heap or priority queue rather than a hash-set.
	openSet := map[Point]struct{}{start: struct{}{}}

	// For node n, cameFrom[n] is the node immediately preceding it on the cheapest path from start
	// to n currently known.
	cameFrom := make(map[Point]Point)

	// For node n, gScore[n] is the cost of the cheapest path from start to n currently known.
	gScore := make(map[Point]int)
	gScore[start] = 0

	// For node n, fScore[n] := gScore[n] + h(n). fScore[n] represents our current best guess as to
	// how short a path from start to finish can be if it goes through n.
	fScore := make(map[Point]int)
	fScore[start] = h(start, goal)

	for len(openSet) != 0 {
		// This operation can occur in O(1) time if openSet is a min-heap or a priority queue
		// current := the node in openSet having the lowest fScore[] value
		current := minInSet(fScore)
		if current == goal {
			return reconstruct_path(cameFrom, current)
		}

		delete(openSet, current)
		neighbors := []Point{}
		if current.X > 0 {
			neighbors = append(neighbors, Point{current.X + 1, current.Y})
		}
		if current.X > 0 {
			neighbors = append(neighbors, Point{current.X - 1, current.Y})
		}
		if current.X > 0 {
			neighbors = append(neighbors, Point{current.X, current.Y + 1})
		}
		if current.X > 0 {
			neighbors = append(neighbors, Point{current.X, current.Y - 1})
		}
		for _, neighbor := range neighbors {
			// d(current,neighbor) is the weight of the edge from current to neighbor
			// tentative_gScore is the distance from start to the neighbor through current
			tentative_gScore := gScore[current] + d(current, neighbor)
			if tentative_gScore < gScore[neighbor] {
				// This path to neighbor is better than any previous one. Record it!
				cameFrom[neighbor] = current
				gScore[neighbor] = tentative_gScore
				fScore[neighbor] = tentative_gScore + h(neighbor, goal)
				if _, in := openSet[neighbor]; !in {
					openSet[neighbor] = struct{}{}
				}
			}
		}
	}

	// Open set is empty but goal was never reached
	return nil
}
