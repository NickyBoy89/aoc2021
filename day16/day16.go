package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Packet struct {
	Version    byte
	Id         byte
	Data       int
	Subpackets []Packet
}

func (p *Packet) ReadFrom(input *bytes.Buffer) {
	packetVersion, err := strconv.ParseInt(string(input.Next(3)), 2, 64)
	if err != nil {
		panic(err)
	}
	p.Version = byte(packetVersion)

	packetId, err := strconv.ParseInt(string(input.Next(3)), 2, 64)
	if err != nil {
		panic(err)
	}
	p.Id = byte(packetId)

	switch packetId {
	case 4: // Literal value
		var packetData string
		for {
			// Read in chunks of 5 until the leading bit is 0
			chunk := input.Next(5)
			packetData += string(chunk[1:])
			if chunk[0] == '0' {
				break
			}
		}
		n, err := strconv.ParseInt(packetData, 2, 64)
		if err != nil {
			panic(err)
		}
		p.Data = int(n)
	default:
		lengthType := string(input.Next(1))
		switch lengthType {
		case "0": // Next 15 bits are total length of subpackets
			subLength, err := strconv.ParseInt(string(input.Next(15)), 2, 64)
			if err != nil {
				panic(err)
			}
			subpackets := bytes.NewBuffer(input.Next(int(subLength)))
			for subpackets.Len() != 0 {
				var newPack Packet
				newPack.ReadFrom(subpackets)
				p.Subpackets = append(p.Subpackets, newPack)
			}
		default: // Next 11 bits are number of subpackets within the packet
			numSub, err := strconv.ParseInt(string(input.Next(11)), 2, 64)
			if err != nil {
				panic(err)
			}
			for numSub > 0 {
				var newPack Packet
				newPack.ReadFrom(input)
				p.Subpackets = append(p.Subpackets, newPack)
				numSub--
			}
		}
	}
}

func (p Packet) Evaluate() int {
	switch p.Id {
	case 0: // Sum everything
		var sum int
		for _, p := range p.Subpackets {
			sum += p.Evaluate()
		}
		return sum
	case 1: // Multiply everything
		sum := 1
		for _, p := range p.Subpackets {
			sum *= p.Evaluate()
		}
		return sum
	case 2:
		min := p.Subpackets[0].Evaluate()
		for _, p := range p.Subpackets {
			if p.Evaluate() < min {
				min = p.Evaluate()
			}
		}
		return min
	case 3:
		max := p.Subpackets[0].Evaluate()
		for _, p := range p.Subpackets {
			if p.Evaluate() > max {
				max = p.Evaluate()
			}
		}
		return max
	case 4:
		return p.Data
	case 5:
		if p.Subpackets[0].Evaluate() > p.Subpackets[1].Evaluate() {
			return 1
		}
		return 0
	case 6:
		if p.Subpackets[0].Evaluate() < p.Subpackets[1].Evaluate() {
			return 1
		}
		return 0
	case 7:
		if p.Subpackets[0].Evaluate() == p.Subpackets[1].Evaluate() {
			return 1
		}
		return 0
	default:
		panic(fmt.Errorf("Unknown packet id: %v", p.Id))
	}
}

func (p Packet) SumVersions() int {
	total := int(p.Version)
	for _, sub := range p.Subpackets {
		total += sub.SumVersions()
	}
	return int(total)
}

func main() {
	rawInput, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(rawInput), "\n")
	input = input[:len(input)-1]

	part1(input[0])
	part2(input[0])
}

func part1(input string) {
	var root Packet
	n, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	var stringBits string
	for _, b := range n {
		stringBits += fmt.Sprintf("%.8b", b)
	}
	root.ReadFrom(bytes.NewBuffer([]byte(stringBits)))
	fmt.Println(root.SumVersions())
}

func part2(input string) {
	var root Packet
	n, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	var stringBits string
	for _, b := range n {
		stringBits += fmt.Sprintf("%.8b", b)
	}
	root.ReadFrom(bytes.NewBuffer([]byte(stringBits)))
	fmt.Println(root.Evaluate())
}
