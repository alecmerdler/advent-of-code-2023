package main

import (
	"fmt"
	"os"
	"strings"
)

func firstDigit(s string) (int, error) {
	for i := 0; i < len(s); i++ {
		value := rune(s[i]) - 48
		if value >= 0 && value < 10 {
			return int(value), nil
		}
	}

	return 0, fmt.Errorf("no digits in string: " + s)
}

func lastDigit(s string) (int, error) {
	for i := len(s) - 1; i >= 0; i-- {
		value := rune(s[i]) - 48
		if value >= 0 && value < 10 {
			return int(value), nil
		}
	}

	return 0, fmt.Errorf("no digits in string: " + s)
}

func main() {
	inputRaw, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	file := string(inputRaw)
	lines := strings.Split(file, "\n")

	total := 0
	for n, line := range lines {
		first, err := firstDigit(line)
		if err != nil {
			panic(err)
		}

		last, err := lastDigit(line)
		if err != nil {
			panic(err)
		}

		combined := (first * 10) + last
		total += combined

		// FIXME(alecmerdler): Debugging
		fmt.Printf("Line %d - found %d and %d = %d in %s. Total: %d\n\n", n, first, last, combined, line, total)
	}

	fmt.Printf("Answer: %d\n", total)
}
