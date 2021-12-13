package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Path struct {
	Start string
	End   string
}

func (p Path) String() string {
	return fmt.Sprintf("%v->%v", p.Start, p.End)
}

func main() {
	rawInput, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(rawInput), "\n")
	input = input[:len(input)-1]

	nodes := make(map[string][]string)

	for _, line := range input {
		parts := strings.Split(line, "-")
		nodes[parts[0]] = append(nodes[parts[0]], parts[1])
		nodes[parts[1]] = append(nodes[parts[1]], parts[0])
	}

	part1(nodes)
	part2(nodes)
}

func part1(nodes map[string][]string) {
	fmt.Println(Search(nodes, "start", "end", make(map[string]int)))
}

func Search(nodes map[string][]string, node, target string, visited map[string]int) int {
	visited[node]++
	var total int
	for _, other := range nodes[node] {
		if other == target {
			total++
			continue
		}
		if unicode.IsLower(rune(other[0])) && visited[other] > 0 {
			continue
		}
		total += Search(nodes, other, target, DeepCopy(visited))
	}
	return total
}

func DeepCopy(m map[string]int) map[string]int {
	result := make(map[string]int, len(m))
	for str, num := range m {
		result[str] = num
	}
	return result
}

func part2(nodes map[string][]string) {
	found := Search2(nodes, "start", "end", make(map[string]int), "")
	fmt.Println(Filter(found))
	fmt.Println(len(Filter(found)))
}

func Filter(paths []string) []string {
	valid := []string{}
	for _, path := range paths {
		nodeCounts := make(map[string]int)
		nodes := strings.Split(path, "->")
		var hasDuplicates bool
		for _, node := range nodes {
			nodeCounts[node]++
			if nodeCounts[node] > 1 && unicode.IsLower(rune(node[0])) {
				hasDuplicates = true
				break
			}
		}
		if !hasDuplicates {
			valid = append(valid, path)
		}
	}
	return valid
}

func Search2(nodes map[string][]string, node, target string, visited map[string]int, path string) []string {
	path += node + "->"
	visited[node]++
	total := []string{}
	for _, other := range nodes[node] {
		if other == target {
			total = append(total, path+target)
			continue
		}
		if other == "start" {
			continue
		}
		if unicode.IsLower(rune(other[0])) && visited[other] > 1 {
			continue
		}
		total = append(total, Search2(nodes, other, target, DeepCopy(visited), path)...)
	}
	return total
}
