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

	paths := make([]Path, len(input))
	for ind, line := range input {
		parts := strings.Split(line, "-")
		paths[ind] = Path{parts[0], parts[1]}
	}

	part1(paths)
	part2(paths)
}

type Node struct {
	Name        string
	Connections map[string]*Node
}

func (n *Node) String() string {
	return fmt.Sprintf("Node: %v", n.Name)
}

func (n *Node) AddEdge(other *Node) {
	n.Connections[other.Name] = other
	other.Connections[n.Name] = n
}

func part1(paths []Path) {
	nodes := make(map[string]*Node)
	for _, path := range paths {
		if _, exist := nodes[path.Start]; !exist {
			nodes[path.Start] = &Node{Name: path.Start, Connections: make(map[string]*Node)}
		}
		if _, exist := nodes[path.End]; !exist {
			nodes[path.End] = &Node{Name: path.End, Connections: make(map[string]*Node)}
		}
		nodes[path.Start].AddEdge(nodes[path.End])
	}

	fmt.Println(Search(nodes["start"], "end", make(map[string]int)))
}

func Search(node *Node, target string, visited map[string]int) int {
	fmt.Println(node.Name, len(node.Connections), visited)
	var total int

	for name, connection := range node.Connections {
		if name == "end" {
			total++
		}
		if unicode.IsLower(rune(node.Name[0])) && visited[name] > 0 {
			continue
		}
		visited[name]++
		total += Search(connection, target, DeepCopy(visited))
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

func part2(paths []Path) {
}
