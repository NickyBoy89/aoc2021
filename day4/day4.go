package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board [][]string

func (b Board) DrawNumber(num string) {
	for ci, col := range b {
		for ri, row := range col {
			if row == num {
				b[ci][ri] = "*" + b[ci][ri]
			}
		}
	}
}

func (b Board) HasMatch() bool {
	for _, col := range b {
		match := true
		for _, item := range col {
			if string(item[0]) != "*" {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}

	for i := range b[0] {
		match := true
		for _, col := range b {
			if string(col[i][0]) != "*" {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}

	return false
}

func (b Board) SumUnmarked() int {
	var sum int
	for _, col := range b {
		for _, row := range col {
			if string(row[0]) == "*" {
				continue
			}
			num, err := strconv.Atoi(row)
			if err != nil {
				panic(err)
			}
			sum += num
		}
	}
	return sum
}

func main() {
	rawInput, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	guessLine := string(rawInput)[:strings.IndexRune(string(rawInput), '\n')]

	guesses := strings.Split(guessLine, ",")

	rawBoards := strings.Split(string(rawInput)[len(guessLine)+1:], "\n\n")

	boards := make([]Board, len(rawBoards))
	for ind, raw := range rawBoards {
		temp := make(Board, 5)
		for ri, row := range dups(strings.Split(raw, "\n")) {
			for _, col := range dups(strings.Split(row, " ")) {
				temp[ri] = append(temp[ri], col)
			}
		}
		boards[ind] = temp
	}

	part1(boards, guesses)
	part2(boards, guesses)
}

func dups(lines []string) []string {
	result := []string{}
	for _, inp := range lines {
		if inp != "" {
			result = append(result, inp)
		}
	}
	return result
}

func part1(boards []Board, guesses []string) {
	for _, guess := range guesses {
		for _, board := range boards {
			board.DrawNumber(guess)
			if board.HasMatch() {
				parsedGuess, err := strconv.Atoi(guess)
				if err != nil {
					panic(err)
				}
				fmt.Println(board.SumUnmarked() * parsedGuess)
				return
			}
		}
	}
}

func part2(boards []Board, guesses []string) {
	for _, guess := range guesses {
		failed := []Board{}
		for _, board := range boards {
			board.DrawNumber(guess)
			if !board.HasMatch() {
				failed = append(failed, board)
			}
		}
		if len(boards) == 1 && len(failed) == 0 {
			num, err := strconv.Atoi(guess)
			if err != nil {
				panic(err)
			}
			fmt.Println(boards[0].SumUnmarked() * num)
			return
		}
		boards = failed
	}
}
