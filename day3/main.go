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

func findAllDigits(line string, startPos int) string {
	rightSide := line[startPos:]
	endPos := strings.IndexFunc(rightSide, func(r rune) bool { return !isDigit(r) })
	if endPos == -1 {
		endPos = len(rightSide)
	}

	number := rightSide[:endPos]
	_, err := strconv.Atoi(number)
	if err != nil {
		panic(err)
	}

	return number
}

func partOne(lines []string) int {
	found := []rune{}

	for row, line := range lines {
		skipTo := 0

		for pos, char := range line {
			if pos < skipTo {
				continue
			}

			// FIXME(alecmerdler): New approach: find every number and check if it is touching a symbol...

			if isDigit(char) {
				number := findAllDigits(line, pos)
				endPos := pos + len(number)

				// TODO(alecmerdler): Check right...
				if endPos < len(line) {
					rightChar := rune(line[endPos+1])
					if isSymbol(rightChar) {
						found = append(found, rightChar)
						continue
					}
				}

				if pos > 0 {
					// TODO(alecmerdler): Check left...
				}

				if row > 0 {
					// TODO(alecmerdler): Check above...
				}

				if row < len(lines)-1 {
					// TODO(alecmerdler): Check below...
				}

				skipTo = endPos + 1
			}
		}
	}

	total := 0
	for _, char := range found {
		number, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}

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
