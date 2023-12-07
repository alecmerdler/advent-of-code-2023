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

type checkFunc func(line string, startPos, endPos int) bool

func checkRight(line string, startPos, endPos int) bool {
	// FIXME(alecmerdler): Debugging
	fmt.Println("checkRight()")

	if endPos >= len(line)-1 {
		return false
	}

	rightChar := rune(line[endPos+1])
	return isSymbol(rightChar)
}

func checkLeft(line string, startPos, endPos int) bool {
	// FIXME(alecmerdler): Debugging
	fmt.Println("checkLeft()")

	if startPos <= 0 {
		return false
	}

	leftChar := rune(line[startPos-1])
	return isSymbol(leftChar)
}

func checkAt(line string, pos int) bool {
	if pos >= len(line) {
		return false
	}

	charAt := rune(line[pos])
	return isSymbol(charAt)
}

func checkLine(line string, startPos, endPos int) bool {
	if checkRight(line, startPos, endPos) {
		return true
	}

	if checkLeft(line, startPos, endPos) {
		return true
	}

	for i := startPos; i == endPos; i++ {
		if checkAt(line, i) {
			return true
		}
	}

	return false
}

func partOne(lines []string) int {
	found := []string{}

	// FIXME(alecmerdler): Debugging - seeing the total of all the numbers to see the max and compare the total...
	max := 0

	for row, line := range lines {
		skipTo := 0

		for pos, char := range line {
			if pos < skipTo {
				continue
			}

			// FIXME(alecmerdler): Implement new approach: find every number and check if it is touching a symbol...
			if isDigit(char) {
				number := findAllDigits(line, pos)
				startPos := pos
				endPos := startPos + len(number) - 1

				// FIXME(alecmerdler): Debugging...
				fmt.Printf("L%d: %s\n", row, line)
				fmt.Printf("Number: %s, End Pos: %d\n", string(number), endPos)

				// FIXME(alecmerdler): Debugging - seeing the total of all the numbers to see the max and compare the total...
				val, _ := strconv.Atoi(number)
				max += val

				checkAbove := func(line string, startPos, endPos int) bool {
					return row > 0 && checkLine(lines[row-1], startPos, endPos)
				}
				checkBelow := func(line string, startPos, endPos int) bool {
					return row < len(lines)-1 && checkLine(lines[row+1], startPos, endPos)
				}

				checkFuncs := []checkFunc{checkRight, checkLeft, checkAbove, checkBelow}
				for _, check := range checkFuncs {
					if check(line, startPos, endPos) {
						found = append(found, number)
						continue
					}
				}

				skipTo = endPos + 1
			}
		}

		// FIXME(alecmerdler): Implement new approach: find every number and check if it is touching a symbol...
		fmt.Printf("\nMax: %d", max)
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
