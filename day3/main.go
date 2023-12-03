package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isDigit(char rune) bool {
	return char >= 48 && char <= 57
}

func isSymbol(char rune) bool {
	return !isDigit(char) && char != 46
}

func uniqueSymbols(lines []string) map[rune]int {
	symbols := map[rune]int{}
	for _, line := range lines {
		for _, char := range line {
			if !isDigit(char) {
				symbols[char] += 1
			}
		}
	}

	return symbols
}

type Direction int

const (
	right Direction = iota
	left
)

func findAllDigits(line string, startPos int, direction Direction) string {
	switch direction {
	case right:
		rightSide := line[startPos+1:]
		endPos := strings.Index(rightSide, ".")
		if endPos == -1 {
			endPos = len(line) - 1
		} else {
			endPos += startPos
		}

		number := line[startPos+1 : endPos+1]
		return number
	case left:
		leftSide := line[:startPos]
		endPos := strings.LastIndex(leftSide, ".")
		if endPos == -1 {
			endPos = 0
		}

		number := line[endPos+1 : startPos]
		return number
	default:
		panic("unknown direction")
	}
}

func checkLeft(line string, pos, row int) (*Position, int) {
	if pos > 0 {
		if isDigit(rune(line[pos-1])) {
			numberRaw := findAllDigits(line, pos, left)
			number, err := strconv.Atoi(numberRaw)
			if err != nil {
				panic(err)
			}

			pos := &Position{row, pos - len(numberRaw), pos}
			return pos, number
		}
	}

	return nil, 0
}

func checkRight(line string, pos, row int) (*Position, int) {
	if pos < len(line) {
		if isDigit(rune(line[pos+1])) {
			numberRaw := findAllDigits(line, pos, right)
			number, err := strconv.Atoi(numberRaw)
			if err != nil {
				panic(err)
			}

			pos := &Position{row, pos + len(numberRaw), pos}
			return pos, number
		}
	}

	return nil, 0
}

// Position stores information about where a number is located (for deduplication).
type Position struct {
	row   int
	start int
	end   int
}

func partOne(lines []string) int {
	found := map[Position]int{}

	for row, line := range lines {
		for pos, char := range line {
			if isSymbol(char) {
				foundPosition, foundNumber := checkLeft(line, pos, row)
				if foundPosition != nil {
					found[*foundPosition] = foundNumber
				}

				// TODO(alecmerdler): Check right...
				foundPosition, foundNumber = checkRight(line, pos, row)
				if foundPosition != nil {
					found[*foundPosition] = foundNumber
				}
				// TODO(alecmerdler): Check above...
				// TODO(alecmerdler): Check below...
			}
		}
	}

	total := 0
	for _, number := range found {
		total += number
	}

	return total
}

func main() {
	inputRaw, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	file := string(inputRaw)
	lines := strings.Split(file, "\n")

	partOneAnswer := partOne(lines)
	fmt.Printf("\nPart one answer: %d\n", partOneAnswer)
}
