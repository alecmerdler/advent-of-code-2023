package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	blue   = "\033[34m"
	yellow = "\033[33m"
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
	if endPos >= len(line)-1 {
		return false
	}

	rightChar := rune(line[endPos+1])
	return isSymbol(rightChar)
}

func checkLeft(line string, startPos, endPos int) bool {
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

	for i := startPos; i <= endPos; i++ {
		if checkAt(line, i) {
			return true
		}
	}

	return false
}

// FIXME(alecmerdler): Debugging function...
func printLines(lines []string, row, startPos, endPos int) {
	fmt.Printf("\n\nLine: %d", row+1)

	if row > 0 {
		for pos, char := range lines[row-1] {
			if pos == 0 {
				fmt.Print("\n")
			}

			if isSymbol(char) {
				fmt.Printf("%s%s%s", blue, string(char), reset)
			} else {
				fmt.Print(string(char))
			}
		}
	}

	for pos, char := range lines[row] {
		if pos == 0 {
			fmt.Print("\n")
		}

		if isSymbol(char) {
			fmt.Printf("%s%s%s", blue, string(char), reset)
		} else if pos == startPos && pos == endPos {
			fmt.Printf("%s%s%s", red, string(char), reset)
		} else if pos == startPos {
			fmt.Printf("%s%s", red, string(char))
		} else if pos == endPos {
			fmt.Printf("%s%s", string(char), reset)
		} else {
			fmt.Print(string(char))
		}
	}

	if row < len(lines)-1 {
		for pos, char := range lines[row+1] {
			if pos == 0 {
				fmt.Print("\n")
			}

			if isSymbol(char) {
				fmt.Printf("%s%s%s", blue, string(char), reset)
			} else {
				fmt.Print(string(char))
			}
		}
	}
}

func printChar(char rune, found bool) {
	if found {
		fmt.Printf("%s%s%s", yellow, string(char), reset)
	} else if isSymbol(char) {
		fmt.Printf("%s%s%s", blue, string(char), reset)
	} else if isDigit(char) {
		fmt.Printf("%s%s%s", red, string(char), reset)
	} else {
		fmt.Print(string(char))
	}
}

func partOne(lines []string) int {
	found := []string{}

	for row, line := range lines {
		skipTo := 0

		for pos, char := range line {
			foundWith := ""
			// FIXME(alecmerdler): Debugging... Looks like we aren't counting number touching digit directly above or below...
			if pos == 0 {
				fmt.Println("")
			}

			if pos < skipTo {
				continue
			}

			if isDigit(char) {
				number := findAllDigits(line, pos)
				startPos := pos
				endPos := startPos + len(number) - 1

				checkAbove := func(line string, startPos, endPos int) bool {
					return row > 0 && checkLine(lines[row-1], startPos, endPos)
				}
				checkBelow := func(line string, startPos, endPos int) bool {
					return row < len(lines)-1 && checkLine(lines[row+1], startPos, endPos)
				}

				checkFuncs := map[string]checkFunc{"checkRight": checkRight, "checkLeft": checkLeft, "checkAbove": checkAbove, "checkBelow": checkBelow}
				for checkName, check := range checkFuncs {
					if check(line, startPos, endPos) {
						foundWith = checkName
						found = append(found, number)
						break
					}
				}

				skipTo = endPos + 1
			}

			// FIXME(alecmerdler): Debugging...
			printChar(char, foundWith != "")
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
