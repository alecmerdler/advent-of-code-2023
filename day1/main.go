package main

import (
	"fmt"
	"os"
	"strings"
)

var wordDigits = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func firstDigit(s string, includeWordDigits bool) (int, error) {
	for i := 0; i < len(s); i++ {
		char := rune(s[i])

		// Check if 0-9
		value := char - 48
		if value >= 0 && value < 10 {
			return int(value), nil
		}

		if includeWordDigits {
			// Check if the first letter of zero-nine
			for intValue, wordDigit := range wordDigits {
				startIndex := i
				endIndex := i + len(wordDigit)
				if endIndex > len(s) {
					continue
				}

				substr := s[startIndex:endIndex]
				if substr == wordDigit {
					return intValue, nil
				}
			}
		}
	}

	return 0, fmt.Errorf("no digits in string: " + s)
}

func lastDigit(s string, includeWordDigits bool) (int, error) {
	for i := len(s) - 1; i >= 0; i-- {
		char := rune(s[i])

		// Check if 0-9
		value := char - 48
		if value >= 0 && value < 10 {
			return int(value), nil
		}

		if includeWordDigits {
			// Check if the first letter of zero-nine
			for intValue, wordDigit := range wordDigits {
				startIndex := i - len(wordDigit) + 1
				endIndex := i + 1
				if startIndex < 0 {
					continue
				}

				substr := s[startIndex:endIndex]
				if substr == wordDigit {
					return intValue, nil
				}
			}
		}
	}

	return 0, fmt.Errorf("no digits in string: " + s)
}

func partOne(lines []string) int {
	total := 0
	for n, line := range lines {
		first, err := firstDigit(line, false)
		if err != nil {
			panic(err)
		}

		last, err := lastDigit(line, false)
		if err != nil {
			panic(err)
		}

		combined := (first * 10) + last
		total += combined

		fmt.Printf("Line %d - found %d and %d = %d in %s. Total: %d\n\n", n+1, first, last, combined, line, total)
	}

	return total
}

func partTwo(lines []string) int {
	total := 0
	for n, line := range lines {
		first, err := firstDigit(line, true)
		if err != nil {
			panic(err)
		}

		last, err := lastDigit(line, true)
		if err != nil {
			panic(err)
		}

		combined := (first * 10) + last
		total += combined

		fmt.Printf("Line %d - found %d and %d = %d in %s. Total: %d\n\n", n+1, first, last, combined, line, total)
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
	fmt.Printf("Part one answer: %d\n", partOneAnswer)

	partTwoAnswer := partTwo(lines)
	fmt.Printf("Part two answer: %d\n", partTwoAnswer)
}
